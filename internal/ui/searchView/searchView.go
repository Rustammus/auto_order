package searchView

import (
	"auto_order/internal/repo"
	"auto_order/internal/service"
	"auto_order/internal/sups"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SearchView struct {
	*widget.List
	sb       *SearchBar
	searcher *service.Searcher
}

func NewSearchView(r *repo.ProductRepo) fyne.CanvasObject {
	var view SearchView
	// Создаем список с кастомным шаблоном
	list := widget.NewList(
		func() int {
			return view.searcher.Size()
		},
		func() fyne.CanvasObject {
			// Здесь определяем шаблон для одного элемента списка
			return createItemTemplate()
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			// Здесь обновляем данные для конкретного элемента
			updateItemTemplate(o, view.searcher.Item(i))
		},
	)

	view.sb = NewSearchWidget(func(text string, _ string) {
		view.searcher.Search(text, 0, sups.PlanetMask)
		view.Refresh()
	})

	view.List = list
	view.searcher = service.NewSearcher(r)

	border := container.NewBorder(view.sb, nil, nil, nil, view.List)
	border.Resize(fyne.NewSize(150, 400))
	return border

}
