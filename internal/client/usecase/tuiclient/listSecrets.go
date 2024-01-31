package tuiclient

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/client/models"
)

// listSecrets Список секретов.
func (a *App) listSecrets(ctx context.Context, secretTypeID models.SecretTypeID) {
	a.tvievApp.list.Clear()

	secrets, err := a.ucSecrets.GetSecretBySecretTypeID(ctx, secretTypeID)
	if err != nil {
		a.showModal(err.Error())
		return
	}

	var makeSecretFn func()

	switch secretTypeID {
	case models.SecretTypeIDLogpass:
		makeSecretFn = func() {
			a.formAddLogPass(ctx, nil)
		}

		for _, secret := range secrets {
			s := secret.(models.SecretLogPass)
			a.tvievApp.list.AddItem(string(secret.(models.SecretLogPass).Name), "", rune('!'), func() {
				a.formAddLogPass(ctx, &s)
				a.tvievApp.pages.SwitchToPage(pageAny)
			})
		}

	case models.SecretTypeIDBankCard:
		makeSecretFn = func() {
			a.formBankCard(ctx, nil)
		}

		for _, secret := range secrets {
			s := secret.(models.SecretBankCard)
			a.tvievApp.list.AddItem(string(secret.(models.SecretBankCard).Name), "", rune('!'), func() {
				a.formBankCard(ctx, &s)
				a.tvievApp.pages.SwitchToPage(pageAny)
			})
		}
	case models.SecretTypeIDBinary:
		makeSecretFn = func() {
			a.formAddBinary(ctx, "", nil)
		}
		for _, secret := range secrets {
			s := secret.(models.SecretBinary)
			a.tvievApp.list.AddItem(string(secret.(models.SecretBinary).Name), "", rune('!'), func() {
				a.formAddBinary(ctx, "", &s)
				a.tvievApp.pages.SwitchToPage(pageAny)
			})
		}
	case models.SecretTypeIDText:
	default:
		a.log.Fatal(models.ErrUnckowType.Error())

	}

	a.tvievApp.list.AddItem("Создать секрет", "", 'a', func() {
		makeSecretFn()
		a.tvievApp.pages.SwitchToPage(pageAny)
	})

	a.tvievApp.list.AddItem("Назад", "", 'p', func() {
		a.listSecretTypes(ctx)
		a.tvievApp.pages.SwitchToPage(pageAnyList)
	})

}
