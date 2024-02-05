package tuiclient

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"sync"
)

// formAddText форма создания секрета с произвольным текстом
func (a *App) formAddText(ctx context.Context, secret *models.SecretText) {
	a.tvievApp.form.Clear(true)

	var (
		existSecret models.SecretText
		safeFn      func() error
	)

	// Изменение секрета
	if secret != nil {
		existSecret = *secret
		safeFn = func() error {
			return a.ucSecrets.EditSecret(ctx,
				existSecret.ID, existSecret.Name,
				existSecret.Value)
		}
	} else { // Добавление нового секрета
		safeFn = func() error {
			return a.ucSecrets.CreateSecret(ctx,
				models.SecretTypeIDText, existSecret.Name,
				existSecret.Value)
		}
	}

	var once sync.Once

	// Кнопка сохранить появляется в зависимости от условий
	safeOnce := func() {
		a.tvievApp.form.AddButton("Сохранить", func() {
			if err := safeFn(); err != nil {
				a.showModal(err.Error())
				return
			}
			a.listSecrets(ctx, models.SecretTypeIDText)
			a.tvievApp.pages.SwitchToPage(pageAnyList)
		})
	}

	a.tvievApp.form.AddInputField("Имя секрета", string(existSecret.Name), 20, nil, func(name string) {
		existSecret.Name = models.SecretName(name)
		once.Do(safeOnce)
	})

	a.tvievApp.form.AddTextArea("text", string(existSecret.Value), 0, 50, 0, func(text string) {
		existSecret.Value = models.SecretTextValue(text)
		once.Do(safeOnce)
	})

	// Показываем кнопку сохранить, если это отображение текущего секрета
	if secret == nil {
		once.Do(safeOnce)
	}

	a.tvievApp.form.AddButton("Вернуться", func() {
		a.listSecrets(ctx, models.SecretTypeIDText)
		a.tvievApp.pages.SwitchToPage(pageAnyList)
	})

	if secret != nil {
		a.tvievApp.form.AddButton("Удалить", func() {
			if err := a.ucSecrets.DeleteSecret(ctx, existSecret.ID, existSecret.SecretTypeID); err != nil {
				a.showModal(err.Error())
				return
			}
			a.listSecrets(ctx, models.SecretTypeIDText)
			a.tvievApp.pages.SwitchToPage(pageAnyList)
		})
	}
}
