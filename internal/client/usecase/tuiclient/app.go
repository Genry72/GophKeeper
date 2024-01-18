package tuiclient

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"github.com/Genry72/GophKeeper/internal/client/usecase"
	"github.com/rivo/tview"
	"go.uber.org/zap"
)

const (
	textExit     = "Выход из программы"
	textRegister = "Регистрация"
	textLogon    = "Вход"
)

// Имена страниц
const (
	// pageLogon Первая страница
	pageLogon = "logon"
	// pageRegister Регистрация пользователя
	pageRegister = "register"
	// pageAuth Вход пользователя
	pageAuth = "auth"
	// Модалка
	pageModal = "modal"
)

type App struct {
	tvievApp  tvievApp
	userInfo  models.UserInfo
	netClient usecase.INetClient
	token     *string
	log       *zap.Logger
}

type tvievApp struct {
	app   *tview.Application
	form  *tview.Form
	pages *tview.Pages
	modal *tview.Modal
}

func NewApp(netclient usecase.INetClient, log *zap.Logger) *App {
	tvievApp := tvievApp{
		app:   tview.NewApplication(),
		form:  tview.NewForm(),
		pages: tview.NewPages(),
		modal: tview.NewModal(),
	}
	return &App{
		tvievApp:  tvievApp,
		netClient: netclient,
		log:       log,
	}
}

func (a *App) Run(ctx context.Context) error {
	firstForm := a.formLogon(ctx)

	a.tvievApp.pages.AddPage(pageLogon, firstForm, true, true)
	a.tvievApp.pages.AddPage(pageRegister, a.tvievApp.form, true, false)
	a.tvievApp.pages.AddPage(pageAuth, a.tvievApp.form, true, false)
	a.tvievApp.pages.AddPage(pageModal, a.tvievApp.modal, true, false)

	if err := a.tvievApp.app.SetRoot(a.tvievApp.pages, true).EnableMouse(true).Run(); err != nil {
		return fmt.Errorf("a.tvievApp.SetRoot: %w", err)
	}
	return nil
}

func (a *App) Stop() error {
	// todo синхронизация с сервером
	a.tvievApp.app.Stop()
	return nil
}

// GetUserInfo Возвращает логин и пароль пользователя
func (a *App) GetUserInfo() (string, string) {
	return a.userInfo.Username, a.userInfo.Password
}
