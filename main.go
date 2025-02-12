package main

import (
	"hometab-builder/logic"
	"hometab-builder/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	MyApp := &logic.MyApp{App: app.New()}

	if MyApp.App.Metadata().Release {
		MyApp.App = app.NewWithID("com.github.rad756.hometab-builder")
	} else {
		MyApp.App = app.NewWithID("com.github.rad756.hometab-builder.testing")
	}

	MyApp.Win = MyApp.App.NewWindow("HomeTab-Builder")
	MyApp.Win.Resize(fyne.NewSize(1200, 800))

	logic.Ini(MyApp)
	MyApp.App.Preferences().SetBool("FirstRun", false)

	ui.LoadGUI(MyApp)

	MyApp.Win.ShowAndRun()
}
