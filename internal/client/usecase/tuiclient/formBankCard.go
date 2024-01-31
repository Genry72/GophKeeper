package tuiclient

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"strconv"
	"sync"
)

// formBankCard форма создания секрета банковской карты
func (a *App) formBankCard(ctx context.Context, secret *models.SecretBankCard) {
	a.tvievApp.form.Clear(true)

	var (
		existSecret models.SecretBankCard
		safeFn      func() error
	)

	// Изменение секрета
	if secret != nil {
		existSecret = *secret
		safeFn = func() error {
			return a.ucSecrets.EditSecret(ctx,
				existSecret.ID, existSecret.Name,
				models.SecretBankCardValue{
					CardNumber: existSecret.Value.CardNumber,
					CardDateTo: existSecret.Value.CardDateTo,
					Cvv:        existSecret.Value.Cvv,
				})
		}
	} else { // Добавление нового секрета
		safeFn = func() error {
			return a.ucSecrets.CreateSecret(ctx,
				models.SecretTypeIDBankCard, existSecret.Name,
				models.SecretBankCardValue{
					CardNumber: existSecret.Value.CardNumber,
					CardDateTo: existSecret.Value.CardDateTo,
					Cvv:        existSecret.Value.Cvv,
				})
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
			a.listSecrets(ctx, models.SecretTypeIDBankCard)
			a.tvievApp.pages.SwitchToPage(pageAnyList)
		})
	}

	a.tvievApp.form.AddInputField("Имя секрета", string(existSecret.Name), 20, nil, func(name string) {
		existSecret.Name = models.SecretName(name)
		once.Do(safeOnce)
	})

	a.tvievApp.form.AddInputField("card number", fmt.Sprint(existSecret.Value.CardNumber), 20,
		func(textToCheck string, lastChar rune) bool {
			if _, err := strconv.ParseInt(textToCheck, 10, 64); err != nil {
				return false
			}
			return true
		}, func(cardNumber string) {
			numCard, _ := strconv.ParseInt(cardNumber, 10, 64)
			existSecret.Value.CardNumber = numCard
			once.Do(safeOnce)
		})

	a.tvievApp.form.AddInputField("date to Year", fmt.Sprint(existSecret.Value.CardDateTo.Year), 20,
		func(textToCheck string, lastChar rune) bool {
			if _, err := strconv.ParseInt(textToCheck, 10, 64); err != nil {
				return false
			}
			return true
		}, func(year string) {
			yearCard, _ := strconv.ParseInt(year, 10, 64)
			existSecret.Value.CardDateTo.Year = int(yearCard)
			once.Do(safeOnce)
		})

	a.tvievApp.form.AddInputField("date to month", fmt.Sprint(existSecret.Value.CardDateTo.Month), 20,
		func(textToCheck string, lastChar rune) bool {
			if _, err := strconv.ParseInt(textToCheck, 10, 64); err != nil {
				return false
			}
			return true
		}, func(month string) {
			monthCard, _ := strconv.ParseInt(month, 10, 64)
			existSecret.Value.CardDateTo.Month = int(monthCard)
			once.Do(safeOnce)
		})

	a.tvievApp.form.AddInputField("csv", fmt.Sprint(existSecret.Value.Cvv), 20,
		func(textToCheck string, lastChar rune) bool {
			if _, err := strconv.ParseInt(textToCheck, 10, 64); err != nil {
				return false
			}
			return true
		}, func(csv string) {
			csvCard, _ := strconv.ParseInt(csv, 10, 64)
			existSecret.Value.Cvv = int(csvCard)
			once.Do(safeOnce)
		})

	// Показываем кнопку сохранить, если это отображение текущего секрета
	if secret == nil {
		once.Do(safeOnce)
	}

	a.tvievApp.form.AddButton("Вернуться", func() {
		a.listSecrets(ctx, models.SecretTypeIDBankCard)
		a.tvievApp.pages.SwitchToPage(pageAnyList)
	})

	if secret != nil {
		a.tvievApp.form.AddButton("Удалить", func() {
			if err := a.ucSecrets.DeleteSecret(ctx, existSecret.ID, existSecret.SecretTypeID); err != nil {
				a.showModal(err.Error())
				return
			}
			a.listSecrets(ctx, models.SecretTypeIDBankCard)
			a.tvievApp.pages.SwitchToPage(pageAnyList)
		})
	}
}
