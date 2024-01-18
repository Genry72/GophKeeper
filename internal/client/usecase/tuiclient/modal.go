package tuiclient

func (a *App) showModal(message string, switchPage string) {
	a.modal.ClearButtons()
	a.modal.
		SetText(message).
		AddButtons([]string{"Понятно"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			a.pages.SwitchToPage(switchPage)
		})

	a.pages.SwitchToPage(pageModal)
}
