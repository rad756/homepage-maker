package main

import (
	"homepage-maker/logic"
	"homepage-maker/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	MyApp := &logic.MyApp{App: app.New()}

	if MyApp.App.Metadata().Release {
		MyApp.App = app.NewWithID("com.github.rad756.homepage-maker")
	} else {
		MyApp.App = app.NewWithID("com.github.rad756.homepage-maker.testing")
	}

	MyApp.Win = MyApp.App.NewWindow("HomePage-Maker")
	MyApp.Win.Resize(fyne.NewSize(1200, 800))

	logic.Ini(MyApp)

	ui.LoadGUI(MyApp)

	MyApp.Win.ShowAndRun()
}
