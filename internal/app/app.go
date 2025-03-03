package app

import "auto_order/internal/ui"

func Run() {
	appUI := ui.NewAppUI()
	appUI.Run()
}
