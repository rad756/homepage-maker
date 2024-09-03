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
	settingsBtn := widget.NewButton("Settings", nil)

	mainContent := LoadMainContent(*MyApp)

	allContent := container.NewBorder(nil, settingsBtn, nil, nil, mainContent)

	MyApp.Win.SetContent(allContent)
}

func LoadSetupMenu(MyApp *logic.MyApp) {

}

func MakeWebsiteButton(Website logic.Website, MyApp *logic.MyApp) *fyne.Container {
	upBtn := widget.NewButtonWithIcon("", theme.MoveUpIcon(), nil)
	downBtn := widget.NewButtonWithIcon("", theme.MoveDownIcon(), nil)
	leftBtn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), nil)
	rightBtn := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), nil)
	mainBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), nil)
	lbl := widget.NewLabel(Website.Name)

	insideBorder := container.NewBorder(nil, lbl, nil, nil, mainBtn)

	if Website.Selected {
		return container.NewBorder(upBtn, downBtn, leftBtn, rightBtn, insideBorder)
	} else {
		return container.NewBorder(nil, nil, nil, nil, insideBorder)
	}
}

func MakeBlankButton(MyApp *logic.MyApp) *fyne.Container {
	mainBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), nil)

	return container.NewBorder(nil, nil, nil, nil, mainBtn)
}

func LoadMainContent(MyApp logic.MyApp) *fyne.Container {
	if len(MyApp.Content) == 0 {
		return MakeBlankButton(&MyApp)
	} else {
		return container.NewVBox(nil)
	}
}
