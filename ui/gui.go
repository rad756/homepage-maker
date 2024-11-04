package ui

import (
	"homepage-maker/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func LoadGUI(MyApp *logic.MyApp) {
	LoadMainMenu(MyApp)
}

func LoadMainMenu(MyApp *logic.MyApp) {
	var allContent *fyne.Container

	upBtn := widget.NewButtonWithIcon("", theme.MoveUpIcon(), nil)
	downBtn := widget.NewButtonWithIcon("", theme.MoveDownIcon(), nil)
	leftBtn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), nil)
	rightBtn := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), nil)

	topContent := container.NewGridWithColumns(4, upBtn, downBtn, leftBtn, rightBtn)

	mainContent := LoadRows(MyApp)

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

	bottomContent := container.NewGridWithColumns(2, reorderBtn, settingsBtn)

	if MyApp.Reorder {
		allContent = container.NewBorder(topContent, bottomContent, nil, nil, mainContent)
	} else {
		allContent = container.NewBorder(nil, bottomContent, nil, nil, mainContent)
	}

	MyApp.Win.SetContent(allContent)
}

func LoadSetupMenu(MyApp *logic.MyApp) {

}
