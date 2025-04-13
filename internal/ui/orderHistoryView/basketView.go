package orderHistoryView

import (
	"auto_order/internal/models"
	"auto_order/internal/service"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type OrderHistoryView struct {
	mainWindow     fyne.Window
	mainApp        fyne.App
	selectedBasket *models.BasketList
	toOrder        *service.ToOrder
	basketTable    *widget.Table
}

func NewOrdersView(basketService *service.BasketsService, mainWindow fyne.Window, mainApp fyne.App) fyne.CanvasObject {
	// create new OrderHistoryView
	b := &OrderHistoryView{
		mainWindow: mainWindow,
		mainApp:    mainApp,
	}

	list := b.newBasketsList()

	// Create main container

	// Create new base Table

	// Create info area
	// Область для дополнительной информации
	infoLabel := widget.NewLabel("Выберите товар для просмотра дополнительной информации")
	infoLabel.Wrapping = fyne.TextWrapWord
	infoArea := container.NewScroll(infoLabel)
	infoArea.SetMinSize(fyne.NewSize(500, 100))

	b.basketTable = b.newOrdersTable()
	// Создаем основной контейнер с разделением на две области
	split := container.NewVSplit(
		list,
		b.basketTable,
	)
	split.Offset = 0.5

	return split
}

func (t *OrderHistoryView) newOrdersTable() *widget.Table {
	baseTable := widget.NewTableWithHeaders(
		func() (int, int) {
			return t.selectedBasket.Size()
		},
		func() fyne.CanvasObject {
			l := canvas.NewText("Undefined", color.White)
			return l
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			obj := o.(*canvas.Text)
			obj.Text = t.selectedBasket.Cell(i.Row, i.Col)
			if i.Col == 1 || i.Col == 2 {
				obj.Alignment = fyne.TextAlignLeading
			} else {
				obj.Alignment = fyne.TextAlignCenter
			}
		})
	baseTable.ShowHeaderColumn = false
	baseTable.ShowHeaderRow = true
	baseTable.CreateHeader = func() fyne.CanvasObject {
		l := canvas.NewText("Undefined", color.White)
		l.Alignment = fyne.TextAlignCenter
		return l
	}
	baseTable.UpdateHeader = func(id widget.TableCellID, o fyne.CanvasObject) {

		text := ""
		switch id.Col {
		case 0:
			text = "Код"
		case 1:
			text = "Название"
		case 2:
			text = "Артикул"
		case 3:
			text = "Кол-во"
		case 4:
			text = "Стоимость"
		}

		o.(*canvas.Text).Text = text
	}

	return baseTable
}

func (t *OrderHistoryView) newBasketsList() *widget.List {
	list := widget.NewList(
		func() int {
			return t.toOrder.Size()
		},
		func() fyne.CanvasObject {
			// Здесь определяем шаблон для одного элемента списка
			return createItemTemplate()
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			// Здесь обновляем данные для конкретного элемента
			updateItemTemplate(o, t.toOrder.Basket(i))
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		t.selectedBasket = t.toOrder.Basket(id)
		t.basketTable.Refresh()
	}

	return list
}
