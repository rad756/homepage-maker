package ui

import (
	"homepage-maker/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func LoadGUI(MyApp *logic.MyApp) {
	MyApp.Buttons = []*widget.Button{}
	MyApp.Websites = []*logic.Website{}
	LoadMainMenu(MyApp)
}

func LoadMainMenu(MyApp *logic.MyApp) {
	var allContent *fyne.Container
	var upBtn *widget.Button
	var downBtn *widget.Button
	var leftBtn *widget.Button
	var rightBtn *widget.Button

	upBtn = widget.NewButtonWithIcon("", theme.MoveUpIcon(), nil)
	downBtn = widget.NewButtonWithIcon("", theme.MoveDownIcon(), nil)
	leftBtn = widget.NewButtonWithIcon("", theme.NavigateBackIcon(), nil)
	rightBtn = widget.NewButtonWithIcon("", theme.NavigateNextIcon(), nil)

	MyApp.UpBtn = upBtn
	MyApp.DownBtn = downBtn
	MyApp.LeftBtn = leftBtn
	MyApp.RightBtn = rightBtn

	SetReorderButtons(*MyApp)

	topContent := container.NewGridWithColumns(4, upBtn, downBtn, leftBtn, rightBtn)

	mainContent := LoadRows(MyApp)

	reorderText := ""

	if MyApp.Reorder {
		reorderText = "Reorder: ON"
	} else {
		reorderText = "Reorder: OFF"
	}
	reorderBtn := widget.NewButton(reorderText, func() {
		MyApp.Selected.Mode = ""
		MyApp.Selected.Row = 0
		MyApp.Selected.Column = 0

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

func ClearButtonSelection(MyApp *logic.MyApp) {
	for _, v := range MyApp.Buttons {
		v.Importance = 0
		v.Refresh()
	}
}

func SetReorderButtons(MyApp logic.MyApp) {
	// Enable all buttons, then disable invalid choices
	MyApp.UpBtn.Enable()
	MyApp.DownBtn.Enable()
	MyApp.LeftBtn.Enable()
	MyApp.RightBtn.Enable()

	// If nothing selected
	if MyApp.Selected.Mode == "" {
		MyApp.UpBtn.Disable()
		MyApp.DownBtn.Disable()
		MyApp.LeftBtn.Disable()
		MyApp.RightBtn.Disable()
		return
	}

	// If website selected
	if MyApp.Selected.Mode == "Website" {
		// If selected is first column in row
		if MyApp.Selected.Column == 0 {
			MyApp.LeftBtn.Disable()
		}

		// If selected is last column in row
		if len(MyApp.Rows[MyApp.Selected.Row].Websites) == MyApp.Selected.Column+1 {
			MyApp.RightBtn.Disable()
		}

		// If in top row and only item in its row
		if MyApp.Selected.Row == 0 && MyApp.Selected.Column == 1 {
			MyApp.UpBtn.Disable()
		}

		// If in bottom row and only item in its row
		if len(MyApp.Rows) == MyApp.Selected.Row+1 && MyApp.Selected.Column == 1 {
			MyApp.DownBtn.Disable()
		}

		return
	}

	if MyApp.Selected.Mode == "Label" {
		MyApp.LeftBtn.Disable()
		MyApp.RightBtn.Disable()

		// If selected in first row
		if MyApp.Selected.Row == 0 {
			MyApp.UpBtn.Disable()
		}

		// If selected in last row
		if len(MyApp.Rows) == MyApp.Selected.Row+1 {
			MyApp.DownBtn.Disable()
		}
	}
}
