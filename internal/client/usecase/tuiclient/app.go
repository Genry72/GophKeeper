package tuiclient

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/client/usecase"
	"github.com/rivo/tview"
	"go.uber.org/zap"
)

const (
	textExit     = "Выход из программы"
	textRegister = "Регистрация"
	textLogon    = "Вход"
)

// Имена страниц.
const (
	// pageLogon Первая страница.
	pageLogon = "logon"
	// pageAny Обновляемая страница.
	pageAny = "toform"
	// Обновляемый список.
	pageAnyList = "toList"
	// Модалка.
	pageModal = "modal"
)

type App struct {
	tvievApp  tvievApp
	ucUsers   usecase.Iusers
	ucSecrets usecase.ISecrets
	log       *zap.Logger
}

type tvievApp struct {
	app   *tview.Application
	form  *tview.Form
	pages *tview.Pages
	modal *tview.Modal
	list  *tview.List
}

func NewApp(ucUser usecase.Iusers, ucSecrets usecase.ISecrets, log *zap.Logger) *App {
	tvievApp := tvievApp{
		app:   tview.NewApplication(),
		form:  tview.NewForm(),
		pages: tview.NewPages(),
		modal: tview.NewModal(),
		list:  tview.NewList().ShowSecondaryText(false),
	}
	return &App{
		tvievApp:  tvievApp,
		ucUsers:   ucUser,
		ucSecrets: ucSecrets,
		log:       log,
	}
}

func (a *App) Run(ctx context.Context) error {
	a.tvievApp.pages.AddPage(pageLogon, a.listLogon(ctx), true, true)
	a.tvievApp.pages.AddPage(pageAny, a.tvievApp.form, true, false)
	a.tvievApp.pages.AddPage(pageModal, a.tvievApp.modal, true, false)
	a.tvievApp.pages.AddPage(pageAnyList, a.tvievApp.list, true, false)

	if err := a.tvievApp.app.SetRoot(a.tvievApp.pages, true).EnableMouse(true).Run(); err != nil {
		return fmt.Errorf("a.tvievApp.SetRoot: %w", err)
	}

	return nil
}

func (a *App) Stop() error {
	a.tvievApp.app.Stop()
	return nil
}
