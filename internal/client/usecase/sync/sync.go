package sync

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"github.com/Genry72/GophKeeper/internal/client/repositories"
	"github.com/Genry72/GophKeeper/internal/client/usecase"
	"go.uber.org/zap"
	"time"
)

// Sync Синхронизация секретов с сервером.
type Sync struct {
	localRepo    repositories.IrepoSecretsSync
	serverClient usecase.InetClientSecrets
	log          *zap.Logger
}

func NewSync(localRepo repositories.IrepoSecretsSync, serverClient usecase.InetClientSecrets, log *zap.Logger) *Sync {
	return &Sync{
		localRepo:    localRepo,
		serverClient: serverClient,
		log:          log,
	}
}

func (s *Sync) StartSync(ctx context.Context) error {
	s.log.Info("sync started")

	// После получения токена прогружаем типы секретов
	secretTypes, err := s.syncTypeSecrets(ctx)
	if err != nil {
		return fmt.Errorf("s.SyncTypeSecrets: %w", err)
	}

	periodicSync := func() {
		if err := s.syncSecrets(ctx, secretTypes); err != nil {
			s.log.Error("s.syncSecrets", zap.Error(err))
			return
		}
	}

	// При старте прогружаем все секреты
	periodicSync()

	t := time.NewTicker(time.Minute)

	go func() {
		for {
			select {
			case <-t.C:
				periodicSync()
			case <-ctx.Done():
				t.Stop()
				s.log.Info("sync stopped")
				return
			}
		}
	}()

	return nil
}

func (s *Sync) syncTypeSecrets(ctx context.Context) ([]models.SecretType, error) {
	src, err := s.serverClient.GetSecretTypes(ctx)
	if err != nil {
		return nil, fmt.Errorf("s.serverClient.GetSecretTypes: %w", err)
	}

	s.localRepo.SetSecretTypes(src)

	return src, nil
}

func (s *Sync) syncSecrets(ctx context.Context, secretTypes []models.SecretType) error {
	for _, st := range secretTypes {
		src, err := s.serverClient.GetSecretsBySecretTypeID(ctx, st.SecretTypeID)
		if err != nil {
			return fmt.Errorf("s.serverClient.GetSecretTypes: %w", err)
		}

		if err := s.localRepo.SyncSecrets(ctx, src); err != nil {
			return fmt.Errorf("s.localRepo.SyncSecrets: %w", err)
		}
	}

	return nil
}
