package tuiclient

import (
	"context"
	"fmt"
)

// formRegister Список доступных типов секретов.
func (a *App) listSecretTypes(ctx context.Context) {
	a.tvievApp.list.Clear()

	types, err := a.ucSecrets.GetSecretTypes(ctx)
	if err != nil {
		a.showModal(err.Error())
		return
	}

	for index, t := range types {
		secretID := t.SecretTypeID
		secretName := t.SecretTypeName
		a.log.Info(fmt.Sprintf("%+v", t))
		a.tvievApp.list.AddItem(string(secretName), "", rune(49+index), func() {
			a.listSecrets(ctx, secretID)
			a.tvievApp.pages.SwitchToPage(pageAnyList)
		})
	}

	a.tvievApp.list.AddItem(textExit, "", 'q', func() {
		a.tvievApp.app.Stop()
	})
}
