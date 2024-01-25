package tuiclient

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/client/models"
)

// formRegister форма регистрации пользователя.
func (a *App) formRegister(ctx context.Context) {
	a.tvievApp.form.Clear(true)

	var login, password string

	a.tvievApp.form.AddInputField("Имя пользователя", "", 20, nil, func(username string) {
		login = username
	})

	a.tvievApp.form.AddPasswordField("Пароль", "", 20, '*', func(pass string) {
		password = pass
	})

	a.tvievApp.form.AddButton("Регистрация", func() {
		if len(login) < 4 || len(password) < 4 {
			a.showModal(models.ErrLenLogPass.Error())
			return
		}

		if err := a.ucUsers.Register(ctx, login, password); err != nil {
			a.showModal(err.Error())
			return
		}

		a.listSecretTypes(ctx)
		a.tvievApp.pages.SwitchToPage(pageAnyList)
	})

	a.tvievApp.form.AddButton("Вернуться", func() {
		a.tvievApp.pages.SwitchToPage(pageLogon)
	})
}
