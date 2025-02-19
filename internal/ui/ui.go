package ui

import (
	"auto_order/internal/repo"
	"auto_order/internal/service"
	"auto_order/internal/ui/basketView"
	"auto_order/internal/ui/catalogView"
	"auto_order/internal/ui/searchView"
	"auto_order/internal/ui/toOrderView"
	"auto_order/pkg/client/sqlitego"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var MainWindow fyne.Window
var MainApp fyne.App

type AppUI struct {
}

func NewAppUI() *AppUI {
	return &AppUI{}
}

func (ui *AppUI) Run() {
	a := app.New()
	MainApp = a
	w := a.NewWindow("Hello")
	MainWindow = w
	// Menu Header
	SetMainMenu(w, a)

	// Repo and DB init
	db, err := sqlitego.NewDB()
	if err != nil {
		panic(err)
	}

	prodRepo := repo.NewProductRepo(db)
	if prodRepo == nil {
		panic("repo is nil!!!")
	}

	catalogS := service.NewCatalog(prodRepo)
	toOrderS := service.NewToOrderService(repo.NewToOrderRepo(db))
	basketS := service.NewBaskets(prodRepo)

	tabs := container.NewAppTabs(
	//container.NewTabItem("Каталог", catalogView.NewCatalogView(catalogS, toOrderS, w, a)),
	//container.NewTabItem("Поиск", searchView.NewSearchView(prodRepo)),
	//container.NewTabItem("К заказу", toOrderView.NewToOrderView(toOrderS, w, a)),
	//container.NewTabItem("Корзины", basketView.NewBasketsView(basketS, w, a)),
	//container.NewTabItem("Заказы", widget.NewLabel("World!")),
	//container.NewTabItem("Пустая вкладка", container.NewCenter(widget.NewLabel("Это пустая вкладка!"))),
	)

	//hello := widget.NewLabel("Hello Fyne!")
	//w.SetContent(container.NewVBox(
	//	hello,
	//	widget.NewButton("Hi!", func() {
	//		hello.SetText("Welcome :)")
	//	}),
	//))

	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)

	w.Resize(fyne.NewSize(1280, 720))
	w.ShowAndRun()
}
