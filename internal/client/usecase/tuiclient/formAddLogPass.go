package tuiclient

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"sync"
)

// formRegister форма создания секрета типа Логин/пароль
func (a *App) formAddLogPass(ctx context.Context, secret *models.SecretLogPass) {
	a.tvievApp.form.Clear(true)

	var (
		existSecret models.SecretLogPass
		safeFn      func() error
	)

	// Изменение секрета
	if secret != nil {
		existSecret = *secret
		safeFn = func() error {
			return a.ucSecrets.EditSecret(ctx,
				existSecret.ID, existSecret.Name,
				models.SecretLogPassValue{
					Login:    existSecret.Value.Login,
					Password: existSecret.Value.Password,
				})
		}
	} else { // Добавление нового секрета
		safeFn = func() error {
			return a.ucSecrets.CreateSecret(ctx,
				models.SecretTypeIDLogpass, existSecret.Name,
				models.SecretLogPassValue{
					Login:    existSecret.Value.Login,
					Password: existSecret.Value.Password,
				})
		}
	}

	var once sync.Once

	// Кнопка сохранить появляется в зависимости от условий
	safeOnce := func() {
		a.tvievApp.form.AddButton("Сохранить", func() {
			if err := safeFn(); err != nil {
				a.showModal(err.Error())
				return
			}
			a.listSecrets(ctx, models.SecretTypeIDLogpass)
			a.tvievApp.pages.SwitchToPage(pageAnyList)
		})
	}

	a.tvievApp.form.AddInputField("Имя секрета", string(existSecret.Name), 20, nil, func(name string) {
		existSecret.Name = models.SecretName(name)
		once.Do(safeOnce)
	})

	a.tvievApp.form.AddInputField("login", existSecret.Value.Login, 20, nil, func(username string) {
		existSecret.Value.Login = username
		once.Do(safeOnce)
	})

	a.tvievApp.form.AddInputField("password", existSecret.Value.Password, 20, nil, func(username string) {
		existSecret.Value.Password = username
		once.Do(safeOnce)
	})

	// Показываем кнопку сохранить, если это отображение текущего секрета
	if secret == nil {
		once.Do(safeOnce)
	}

	a.tvievApp.form.AddButton("Вернуться", func() {
		a.listSecrets(ctx, models.SecretTypeIDLogpass)
		a.tvievApp.pages.SwitchToPage(pageAnyList)
	})

	if secret != nil {
		a.tvievApp.form.AddButton("Удалить", func() {
			if err := a.ucSecrets.DeleteSecret(ctx, existSecret.ID, existSecret.SecretTypeID); err != nil {
				a.showModal(err.Error())
				return
			}
			a.listSecrets(ctx, models.SecretTypeIDLogpass)
			a.tvievApp.pages.SwitchToPage(pageAnyList)
		})
	}
}
