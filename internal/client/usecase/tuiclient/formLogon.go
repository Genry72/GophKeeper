package tuiclient

import (
	"context"
	"github.com/rivo/tview"
)

// formLogon Первичная форма для выбора: войти или зарегистрироваться
func (a *App) formLogon(ctx context.Context) *tview.Form {
	mainForm := tview.NewForm()
	mainForm.AddButton(textRegister, func() {
		a.tvievApp.pages.SwitchToPage(pageRegister)
		a.formRegister(ctx)
	})

	mainForm.AddButton(textLogon, func() {
		a.tvievApp.pages.SwitchToPage(pageAuth)
		a.formAuth(ctx)
	})

	mainForm.AddButton(textExit, func() {
		a.tvievApp.app.Stop()
	})

	return mainForm
}
