package ui

import (
	"homepage-maker/logic"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoadGUI(MyApp *logic.MyApp) {
	LoadMainMenu(MyApp)
}

func LoadMainMenu(MyApp *logic.MyApp) {
	settingsBtn := widget.NewButton("Settings", nil)

	content := container.NewBorder(nil, settingsBtn, nil, nil, nil)

	MyApp.Win.SetContent(content)
}

func LoadSetupMenu(MyApp *logic.MyApp) {

}
