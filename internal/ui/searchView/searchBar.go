package searchView

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SearchBar struct {
	widget.BaseWidget
	searchEntry      *widget.Entry
	searchFlag       *widget.Select
	searchButt       *widget.Button
	onSearchCallback func(string, string)

	flagText string
}

func NewSearchWidget(onSearch func(string, string)) *SearchBar {
	var sw *SearchBar

	selectable := []string{"Все", "Княгиня", "Планета", "Прогресс", "Эсско"}

	sw = &SearchBar{
		// Input text widget
		searchEntry: widget.NewEntry(),

		searchFlag: widget.NewSelect(selectable, func(s string) {
			sw.flagText = s
		}),

		searchButt: widget.NewButtonWithIcon("", theme.Icon(theme.IconNameSearch), func() {
			if sw.onSearchCallback != nil {
				sw.onSearchCallback(sw.searchEntry.Text, sw.flagText)
			}
		}),

		onSearchCallback: onSearch,
	}

	sw.ExtendBaseWidget(sw)

	sw.searchFlag.SetSelectedIndex(0)
	return sw
}

func (sw *SearchBar) CreateRenderer() fyne.WidgetRenderer {

	boarder := container.NewBorder(
		nil,
		nil,
		widget.NewLabel("Поиск:"),
		container.NewHBox(
			sw.searchFlag,
			sw.searchButt,
		),
		sw.searchEntry)

	r := widget.NewSimpleRenderer(
		boarder,
		//container.NewVBox(
		//	container.NewBorder(nil, nil, widget.NewLabel("Поиск:"), nil, sb.searchEntry),
		//	sb.searchType,
		//),
	)

	return r
}
