package tuiclient

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"os"
)

// formAddBinary форма создания секрета типа binary(добавление файла)
func (a *App) formAddBinary(ctx context.Context, filePath string, secret *models.SecretBinary) {
	a.tvievApp.form.Clear(true)

	var secretName string
	// Создание нового секрета
	if secret == nil {
		if filePath != "" {
			a.tvievApp.form.AddInputField("Имя секрета", "", 20, nil, func(name string) {
				secretName = name
			})

			a.tvievApp.form.AddInputField("Путь до файла", filePath, 20, nil, func(name string) {
				filePath = name
			})
		}

		a.tvievApp.form.AddButton("Выбрать файл", func() {
			a.tree(ctx)
			a.tvievApp.pages.SwitchToPage(pageTree)
		})

		a.tvievApp.form.AddButton("Сохранить", func() {
			f, err := os.ReadFile(filePath)
			if err != nil {
				a.showModal(err.Error())
				return
			}
			a.log.Info(string(f))
			if err := a.ucSecrets.CreateSecret(ctx,
				models.SecretTypeIDBinary, models.SecretName(secretName), models.SecretBinaryValue(f)); err != nil {
				a.showModal(err.Error())
				return
			}
			a.listSecrets(ctx, models.SecretTypeIDBinary)
			a.tvievApp.pages.SwitchToPage(pageAnyList)
		})
	} else {
		// Редактирование секрета
		a.tvievApp.form.AddInputField("Путь для сохранения файла", "", 20, nil, func(name string) {
			filePath = name
		})

		a.tvievApp.form.AddButton("Скачать", func() {
			f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
			if err != nil {
				a.showModal(err.Error())
				return
			}
			a.log.Info(string(secret.Value))
			if _, err := f.Write(secret.Value); err != nil {
				a.showModal(err.Error())
				return
			}
			a.listSecrets(ctx, models.SecretTypeIDBinary)
			a.tvievApp.pages.SwitchToPage(pageAnyList)
		})

		a.tvievApp.form.AddButton("Удалить секрет", func() {
			if err := a.ucSecrets.DeleteSecret(ctx, secret.ID, secret.SecretTypeID); err != nil {
				a.showModal(err.Error())
				return
			}
			a.listSecrets(ctx, models.SecretTypeIDBinary)
			a.tvievApp.pages.SwitchToPage(pageAnyList)
		})
	}

	a.tvievApp.form.AddButton("Вернуться", func() {
		a.listSecrets(ctx, models.SecretTypeIDBinary)
		a.tvievApp.pages.SwitchToPage(pageAnyList)
	})
}
