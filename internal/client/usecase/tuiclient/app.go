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
	tvievApp  *tview.Application
	form      *tview.Form
	pages     *tview.Pages
	modal     *tview.Modal
	userInfo  models.UserInfo
	netClient usecase.INetClient
	token     string
	log       *zap.Logger
}

func NewApp(netclient usecase.INetClient, log *zap.Logger) *App {
	return &App{
		tvievApp:  tview.NewApplication(),
		form:      tview.NewForm(),
		pages:     tview.NewPages(),
		modal:     tview.NewModal(),
		netClient: netclient,
		log:       log,
	}
}

func (a *App) Run(ctx context.Context) error {
	firstForm := a.formLogon(ctx)

	a.pages.AddPage(pageLogon, firstForm, true, true)
	a.pages.AddPage(pageRegister, a.form, true, false)
	a.pages.AddPage(pageAuth, a.form, true, false)
	a.pages.AddPage(pageModal, a.modal, true, false)

	if err := a.tvievApp.SetRoot(a.pages, true).EnableMouse(true).Run(); err != nil {
		return fmt.Errorf("a.tvievApp.SetRoot: %w", err)
	}
	return nil
}

func (a *App) Stop() error {
	// todo синхронизация с сервером
	a.tvievApp.Stop()
	return nil
}

// GetUserInfo Возвращает логин и пароль пользователя
func (a *App) GetUserInfo() (string, string) {
	return a.userInfo.Username, a.userInfo.Password
}
