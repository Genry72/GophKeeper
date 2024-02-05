package tuiclient

import (
	"context"
	"fmt"
	"github.com/rivo/tview"
)

// listLogon Начальный список при запуске приложения.
func (a *App) listLogon(ctx context.Context, buildVersion, buildDate string) *tview.Flex {
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

	flex := tview.NewFlex()

	flex.SetDirection(tview.FlexRow).
		AddItem(list, 0, 10, true).
		AddItem(tview.NewTextView().SetText(fmt.Sprintf("Build version: %s\nBuld date: %s",
			buildVersion, buildDate)), 0, 1, false)

	return flex
}
