package tuiclient

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/client/models"
)

// formRegister форма регистрации пользователя
func (a *App) formAuth(ctx context.Context) {
	a.form.Clear(true)
	user := a.userInfo

	a.form.AddInputField("Имя пользователя", "", 20, nil, func(username string) {
		user.Username = username
	})

	a.form.AddPasswordField("Пароль", "", 20, '*', func(pass string) {
		user.Password = pass
	})

	a.form.AddButton("Войти", func() {
		if len(user.Username) < 4 || len(user.Password) < 4 {
			a.showModal(models.ErrLenLogPass.Error(), pageAuth)
			return
		}

		token, err := a.netClient.Auth(ctx, user.Username, user.Password)
		if err != nil {
			a.showModal(err.Error(), pageAuth)
			return
		}

		a.userInfo = user
		a.token = token
		a.pages.SwitchToPage(pageLogon) // todo нужна норм страничка
	})

	a.form.AddButton("Вернуться", func() {
		a.pages.SwitchToPage(pageLogon)
	})
}
