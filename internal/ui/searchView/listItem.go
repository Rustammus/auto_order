package searchView

import (
	"auto_order/internal/models"
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
		widget.NewLabel("Артикул:"),
		articleLabel,
		//layout.NewSpacer(),
		//widget.NewSeparator(),
		widget.NewLabel("Поставщик:"),
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
func updateItemTemplate(o fyne.CanvasObject, product models.SearchItem) {
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
	nameLabel.SetText(product.Title)
	priceLabel.SetText(fmt.Sprintf("%.2f ₽", product.Price))
	articleLabel.SetText(product.Article)
	//supplierLabel.SetText(supplierSlice[product.SupplierID])
	supplierLabel.SetText(supplierSlice[1])

	// Оформление количества в зависимости от наличия
	// if product.InStock && product.Quantity > 0
	if product.Count > 0 {
		quantityLabel.SetText(strconv.FormatInt(product.Count, 10) + " шт.")
		quantityLabel.TextStyle = fyne.TextStyle{Bold: true}
		//quantityLabel.Color = theme.PrimaryColor()
	} else {
		quantityLabel.SetText("нет в наличии")
		quantityLabel.TextStyle = fyne.TextStyle{Bold: false}
		//quantityLabel.Color = theme.ErrorColor()
	}

	// Обновляем стиль цены
	//if product.Price > 50000 {
	//	priceLabel.Color = theme.ErrorColor()
	//} else if product.Price > 20000 {
	//	priceLabel.Color = theme.WarningColor()
	//} else {
	//	priceLabel.Color = theme.PrimaryColor()
	//}
}
