package ui

import (
	"fyne.io/fyne/v2"
)

func SetMainMenu(w fyne.Window, a fyne.App) {
	// Создаем главное меню
	mainMenu := fyne.NewMainMenu(
		// Первое меню (обычно File)
		fyne.NewMenu("Файл",
			fyne.NewMenuItem("Импорт каталога", func() {
				//TODO
			}),
			fyne.NewMenuItem("Импорт товаров к заказу", func() {
				//TODO
			}),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Выход", func() {
				a.Quit()
			}),
		),

		// Меню настроек
		fyne.NewMenu("Настройки",
			fyne.NewMenuItem("Параметры отображения", func() {
				//TODO
			}),
			fyne.NewMenuItem("Экспорт данных", func() {
				//TODO
			}),
		),

		// Меню помощи (обычно последнее)
		fyne.NewMenu("Помощь",
			fyne.NewMenuItem("О программе", func() {
				//TODO
			}),
		),
	)

	// Устанавливаем меню
	w.SetMainMenu(mainMenu)
}
