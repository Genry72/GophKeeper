package tuiclient

func (a *App) showModal(message string) {
	a.tvievApp.modal.ClearButtons()
	a.tvievApp.modal.
		SetText(message).
		AddButtons([]string{"Понятно"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			a.tvievApp.pages.SwitchToPage(pageAny)
		})

	a.tvievApp.pages.SwitchToPage(pageModal)
}
