package tuiclient

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/client/models"
)

// formRegister форма создания секрета типа Логин/пароль
func (a *App) formAddLogPass(ctx context.Context) {
	a.tvievApp.form.Clear(true)
	var secretName, login, password string

	a.tvievApp.form.AddInputField("Имя секрета", "", 20, nil, func(name string) {
		secretName = name
	})

	a.tvievApp.form.AddInputField("login", "", 20, nil, func(username string) {
		login = username
	})

	a.tvievApp.form.AddPasswordField("password", "", 20, '*', func(pass string) {
		password = pass
	})

	a.tvievApp.form.AddButton("Сохранить", func() {
		if err := a.ucSecrets.CreateSecret(ctx,
			models.SecretTypeIDLogpass, models.SecretName(secretName),
			models.SecretLogPassValue{
				Login:    login,
				Password: password,
			},
		); err != nil {
			a.showModal(err.Error())
			return
		}
		a.listSecrets(ctx, models.SecretTypeIDLogpass)
		a.tvievApp.pages.SwitchToPage(pageAnyList)
	})

	a.tvievApp.form.AddButton("Вернуться", func() {
		a.listSecrets(ctx, models.SecretTypeIDLogpass)
		a.tvievApp.pages.SwitchToPage(pageAnyList)
	})
}
