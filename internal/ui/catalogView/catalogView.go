package catalogView

import (
	"auto_order/internal/service"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"math/rand"
)

type CatalogTable struct {
	*widget.Table
	currMousePos fyne.Position
	popupmenu    *widget.PopUpMenu
	selectedCell widget.TableCellID
	catalog      *service.Catalog
	toOrder      *service.ToOrder
	mainWindow   fyne.Window
	mainApp      fyne.App
	il           *widget.Label
}

// MouseDown response to desktop mouse event
func (t *CatalogTable) MouseDown(e *desktop.MouseEvent) {

	if e.Button == desktop.MouseButtonSecondary {
		log.Printf("MouseDownSecondary AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
		t.Table.Tapped(&e.PointEvent)
		return
	} else if e.Button == desktop.MouseButtonPrimary {
		log.Printf("MouseDownPrimary   BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	}

	t.Table.MouseDown(e)
}

// MouseMoved response to desktop mouse event
func (t *CatalogTable) MouseMoved(ev *desktop.MouseEvent) {
	t.currMousePos = ev.Position
	t.Table.MouseMoved(ev)
}

func (t *CatalogTable) ShowPopupMenu() {
	if t.popupmenu == nil {
		menu := fyne.NewMenu("Right Menu",
			fyne.NewMenuItem("Копировать", func() {
				t.CopyCellToClipboard()
			}),
			fyne.NewMenuItem("К заказу", func() {
				t.AddToOrder()
			}),
			fyne.NewMenuItem("В корзину", func() {
				t.AddToOrder()
			}),
			fyne.NewMenuItem("Добавить заметку", func() {
				t.AddToOrder()
			}),
		)
		c := t.mainApp.Driver().CanvasForObject(t)
		t.popupmenu = widget.NewPopUpMenu(menu, c)
	}
	t.popupmenu.ShowAtRelativePosition(t.currMousePos, t)
}

func (t *CatalogTable) CopyCellToClipboard() {
	t.mainWindow.Clipboard().SetContent(t.catalog.Cell(t.selectedCell.Row, t.selectedCell.Col))
}

func (t *CatalogTable) AddToOrder() {
	//TODO

	p := t.catalog.GetProduct(t.selectedCell.Row)
	if p == nil {
		return
	}

	t.toOrder.AddToOrder(p)

}

func (t *CatalogTable) Tapped(e *fyne.PointEvent) {
	log.Printf("Tapped: %v", e)
	//t.Table.Tapped(e)
}

func (t *CatalogTable) UpdateData() {
	if t.il != nil {
		name := t.catalog.Cell(t.selectedCell.Row, 4)
		art := t.catalog.Cell(t.selectedCell.Row, 5)

		if art == "" {
			art = "Отсутствует"
		}

		supsCount := rand.Intn(3)

		s := fmt.Sprintf(
			`Наименование: %s
Артикул: %s
Комментарии: Отсутствуют

Кол-во поставщиков: %d`, name, art, supsCount,
		)

		t.il.SetText(s)
	}
	//t.Table.Tapped(e)
}

func NewCatalogView(catalog *service.Catalog, toOrder *service.ToOrder, mainWindow fyne.Window, mainApp fyne.App) fyne.CanvasObject {
	// create new CatalogTable
	table := &CatalogTable{
		catalog:    catalog,
		toOrder:    toOrder,
		mainWindow: mainWindow,
		mainApp:    mainApp,
	}

	// Create main container

	// Create new base Table
	baseTable := widget.NewTableWithHeaders(
		func() (int, int) {
			return table.catalog.Size()
		},
		func() fyne.CanvasObject {
			l := canvas.NewText("Undefined", color.White)
			l.Alignment = fyne.TextAlignCenter

			return l
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			obj := o.(*canvas.Text)
			obj.Text = table.catalog.Cell(i.Row, i.Col)
			if i.Col == 4 {
				obj.Alignment = fyne.TextAlignLeading
			} else {
				obj.Alignment = fyne.TextAlignCenter
			}
		})

	table.Table = baseTable

	table.ExtendBaseWidget(table)
	//

	table.OnSelected = func(id widget.TableCellID) {
		log.Printf("selected")
		table.selectedCell = id
		table.ShowPopupMenu()
		table.UpdateData()
	}

	table.ShowHeaderColumn = false

	// Заголовки таблицы
	table.CreateHeader = func() fyne.CanvasObject {
		l := canvas.NewText("Undefined", color.White)
		l.Alignment = fyne.TextAlignCenter
		return l
	}

	table.UpdateHeader = func(id widget.TableCellID, o fyne.CanvasObject) {
		o.(*canvas.Text).Text = catalog.ColumnName(id.Col)
	}

	table.SetColumnWidth(4, 500)

	// Create info area
	// Область для дополнительной информации
	infoLabel := widget.NewLabel("Выберите товар для просмотра дополнительной информации")
	infoLabel.Wrapping = fyne.TextWrapWord
	infoArea := container.NewScroll(infoLabel)
	infoArea.SetMinSize(fyne.NewSize(500, 100))

	table.il = infoLabel

	searchFunc := func(text, t string) {
		if text == "" {
			table.catalog.ListAll()
			table.Refresh()
		}
		switch t {
		case "Имя, Артикул":
			table.catalog.SearchByText(text)
			table.Refresh()
		case "По ID":
			//id, _ := strconv.ParseInt(text, 10, 64)
			table.catalog.SearchByCode(text)
			table.Refresh()
		}
	}

	// Создаем основной контейнер с разделением на две области
	searchBar := NewSearchWidget(searchFunc)
	split := container.NewVSplit(
		container.NewBorder(searchBar, nil, nil, nil, table),
		infoArea,
	)
	split.Offset = 0.7 // Начальное соотношение размеров (70% таблица, 30% информация)

	return split
}
