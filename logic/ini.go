package logic

import "fyne.io/fyne/v2"

type MyApp struct {
	App      fyne.App
	Win      fyne.Window
	Rows     []Row
	GridSize fyne.Size
}

func Ini(MyApp *MyApp) {
	if MyApp.App.Preferences().BoolWithFallback("FirstRun", true) {
		MyApp.App.Preferences().SetBool("FirstRun", true)
	} else {
		MyApp.App.Preferences().SetBool("FirstRun", false)
	}

	MyApp.App.Preferences().SetString("RowFileName", "Rows.json")

	if MyApp.App.Preferences().Bool("FirstRun") {
		CreateRowFile(MyApp)
		CreateImgFolder(MyApp)
	} else {
		ReadRowFile(MyApp)
	}

	size := 64
	MyApp.GridSize.Height = float32(size)
	MyApp.GridSize.Width = float32(size)
}
