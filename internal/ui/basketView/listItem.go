package basketView

import (
	"auto_order/internal/models"
	"auto_order/internal/sups"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var supplierSlice []string = []string{"КНЯГИНЯ", "АВТО-ТРАСТ", "ПРОГРЕСС", "ЭССКО"}

func NewListItem() {

}

// createItemTemplate создает шаблон для одного элемента списка
func createItemTemplate() fyne.CanvasObject {
	// Основные данные
	nameLabel := widget.NewLabel("")
	nameLabel.TextStyle.Bold = true
	nameLabel.Wrapping = fyne.TextWrapOff

	articleLabel := widget.NewLabel("")
	supplierLabel := widget.NewLabel("")

	// Количество и цена
	quantityLabel := widget.NewLabel("")
	priceLabel := widget.NewLabel("")
	priceLabel.TextStyle.Bold = true
	priceLabel.Alignment = fyne.TextAlignTrailing

	// Размещаем элементы в контейнерах
	topRow := container.NewHBox(
		nameLabel,
		layout.NewSpacer(),
		priceLabel,
	)

	bottomRow := container.NewHBox(
		widget.NewLabel("Продуктов:"),
		articleLabel,
		//layout.NewSpacer(),
		//widget.NewSeparator(),
		widget.NewLabel("Доступно:"),
		supplierLabel,
		layout.NewSpacer(),
		//widget.NewSeparator(),
		quantityLabel,
	)

	// Главный контейнер элемента
	return container.NewVBox(
		topRow,
		bottomRow,
		widget.NewSeparator(),
	)
}

// updateItemTemplate обновляет данные в шаблоне элемента
func updateItemTemplate(o fyne.CanvasObject, b *models.BasketList) {
	vbox := o.(*fyne.Container)

	// Получаем элементы из шаблона
	topRow := vbox.Objects[0].(*fyne.Container)
	bottomRow := vbox.Objects[1].(*fyne.Container)

	// Обновляем данные
	nameLabel := topRow.Objects[0].(*widget.Label)
	priceLabel := topRow.Objects[2].(*widget.Label)

	articleLabel := bottomRow.Objects[1].(*widget.Label)
	supplierLabel := bottomRow.Objects[3].(*widget.Label)
	quantityLabel := bottomRow.Objects[5].(*widget.Label)

	// Устанавливаем значения
	nameLabel.SetText(sups.SupNames[b.SupplierID])
	priceLabel.SetText(fmt.Sprintf("%.2f ₽", b.TotalSum))
	articleLabel.SetText(strconv.Itoa(len(b.Items)))
	//supplierLabel.SetText(supplierSlice[b.SupplierID])
	supplierLabel.SetText(strconv.FormatInt(b.TotalCount, 10))

	quantityLabel.SetText(strconv.FormatInt(b.TotalCount, 10) + " шт.")
}
