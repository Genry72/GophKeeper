package tuiclient

import (
	"context"
	"github.com/rivo/tview"
)

// formLogon Первичная форма для выбора: войти или зарегистрироваться
func (a *App) formLogon(ctx context.Context) *tview.Form {
	mainForm := tview.NewForm()
	mainForm.AddButton(textRegister, func() {
		a.pages.SwitchToPage(pageRegister)
		a.formRegister(ctx)
	})

	mainForm.AddButton(textLogon, func() {
		a.pages.SwitchToPage(pageAuth)
		a.formAuth(ctx)
	})

	mainForm.AddButton(textExit, func() {
		a.tvievApp.Stop()
	})

	return mainForm
}
