package tuiclient

import (
	"context"
	"github.com/rivo/tview"
)

// listLogon Начальный список при запуске приложения.
func (a *App) listLogon(ctx context.Context) *tview.List {
	list := tview.NewList().ShowSecondaryText(false)
	list.AddItem(textLogon, "", 'l', func() {
		a.formAuth(ctx)
		a.tvievApp.pages.SwitchToPage(pageAny)

	})

	list.AddItem(textRegister, "", 'r', func() {
		a.tvievApp.pages.SwitchToPage(pageAny)
		a.formRegister(ctx)
	})

	list.AddItem(textExit, "", 'q', func() {
		a.tvievApp.app.Stop()
	})

	return list
}
