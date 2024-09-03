package logic

import "fyne.io/fyne/v2"

type MyApp struct {
	App  fyne.App
	Win  fyne.Window
	Rows []Row
}

func Ini(MyApp *MyApp) {
	if MyApp.App.Preferences().BoolWithFallback("FirstRun", true) {
		MyApp.App.Preferences().SetBool("FirstRun", true)
	} else {
		MyApp.App.Preferences().SetBool("FirstRun", false)
	}

	MyApp.App.Preferences().SetString("ContentFileName", "Content.json")

	if MyApp.App.Preferences().Bool("FirstRun") {
		CreateRowFile(MyApp)
	}
}
