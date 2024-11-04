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
	reorderText := ""

	if MyApp.Reorder {
		reorderText = "Reorder: ON"
	} else {
		reorderText = "Reorder: OFF"
	}
	reorderBtn := widget.NewButton(reorderText, func() {
		if MyApp.Reorder {
			MyApp.Reorder = false
			LoadGUI(MyApp)
		} else {
			MyApp.Reorder = true
			LoadGUI(MyApp)
		}
	})

	if MyApp.Reorder {
		reorderBtn.Importance = 1
	}
	settingsBtn := widget.NewButton("Settings", nil)

	mainContent := LoadRows(MyApp)

	bottomContent := container.NewGridWithColumns(2, reorderBtn, settingsBtn)

	allContent := container.NewBorder(nil, bottomContent, nil, nil, mainContent)

	MyApp.Win.SetContent(allContent)
}

func LoadSetupMenu(MyApp *logic.MyApp) {

}
