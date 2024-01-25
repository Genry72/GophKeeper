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

	for index, secret := range secrets {
		switch s := secret.(type) {
		case models.SecretLogPass:
			a.tvievApp.list.AddItem(string(s.Name), "", rune(49+index), func() {
				a.formAddLogPass(ctx)
				a.tvievApp.pages.SwitchToPage(pageAny)
			})
		case models.SecretBankCard:
		case models.SecretBinary:
		case models.SecretText:
		default:
			a.log.Fatal(models.ErrUnckowType.Error())
		}
	}

	a.tvievApp.list.AddItem("Назад", "", 'p', func() {
		a.listSecretTypes(ctx)
		a.tvievApp.pages.SwitchToPage(pageAnyList)
	})
}
