package toOrderView

import (
	"auto_order/internal/service"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

type ToOrderTable struct {
	*widget.Table
	currMousePos fyne.Position
	popupmenu    *widget.PopUpMenu
	selectedCell widget.TableCellID
	toOrder      *service.ToOrder
	mainWindow   fyne.Window
	mainApp      fyne.App
}

// MouseDown response to desktop mouse event
func (t *ToOrderTable) MouseDown(e *desktop.MouseEvent) {

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
func (t *ToOrderTable) MouseMoved(ev *desktop.MouseEvent) {
	t.currMousePos = ev.Position
	t.Table.MouseMoved(ev)
}

func (t *ToOrderTable) ShowPopupMenu() {
	if t.popupmenu == nil {
		menu := fyne.NewMenu("Right Menu",
			fyne.NewMenuItem("Copy", func() {
				t.CopyCellToClipboard()
			}))
		c := t.mainApp.Driver().CanvasForObject(t)
		t.popupmenu = widget.NewPopUpMenu(menu, c)
	}
	t.popupmenu.ShowAtRelativePosition(t.currMousePos, t)
}

func (t *ToOrderTable) CopyCellToClipboard() {
	t.mainWindow.Clipboard().SetContent(t.toOrder.Cell(t.selectedCell.Row, t.selectedCell.Col))
}

func (t *ToOrderTable) Tapped(e *fyne.PointEvent) {
	//log.Printf("Tapped: %v", e)
	//t.Table.Tapped(e)
}

func NewToOrderView(toOrder *service.ToOrder, mainWindow fyne.Window, mainApp fyne.App) fyne.CanvasObject {
	// create new ToOrderTable
	table := &ToOrderTable{
		toOrder:    toOrder,
		mainWindow: mainWindow,
		mainApp:    mainApp,
	}

	// Create main container

	// Create new base Table
	baseTable := widget.NewTableWithHeaders(
		func() (int, int) {
			return table.toOrder.Size()
		},
		func() fyne.CanvasObject {
			l := canvas.NewText("Undefined", color.White)
			return l
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			obj := o.(*canvas.Text)
			obj.Text = table.toOrder.Cell(i.Row, i.Col)
			if i.Col == 2 || i.Col == 3 {
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
	}

	table.ShowHeaderColumn = false

	table.CreateHeader = func() fyne.CanvasObject {
		return canvas.NewText("Undefined", color.White)
	}

	table.UpdateHeader = func(id widget.TableCellID, o fyne.CanvasObject) {
		o.(*canvas.Text).Text = toOrder.ColumnName(id.Col)
	}

	// Create info area
	// Область для дополнительной информации
	infoLabel := widget.NewLabel("Выберите товар для просмотра дополнительной информации")
	infoLabel.Wrapping = fyne.TextWrapWord
	infoArea := container.NewScroll(infoLabel)
	infoArea.SetMinSize(fyne.NewSize(500, 100))

	searchFunc := func(text, t string) {
		if text == "" {
			table.toOrder.Search()
			table.Refresh()
		}
		//switch t {
		//case "Имя, Артикул":
		//	table.toOrder.SearchByText(text)
		//	table.Refresh()
		//case "По ID":
		//	//id, _ := strconv.ParseInt(text, 10, 64)
		//	table.toOrder.SearchByCode(text)
		//	table.Refresh()
		//}
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
