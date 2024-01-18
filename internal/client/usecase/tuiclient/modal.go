package tuiclient

func (a *App) showModal(message string, switchPage string) {
	a.tvievApp.modal.ClearButtons()
	a.tvievApp.modal.
		SetText(message).
		AddButtons([]string{"Понятно"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			a.tvievApp.pages.SwitchToPage(switchPage)
		})

	a.tvievApp.pages.SwitchToPage(pageModal)
}
