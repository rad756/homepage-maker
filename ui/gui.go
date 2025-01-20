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
	//var pages []string

	// for i, v := range MyApp.Pages {
	// 	//pages = append(pages, )
	// 	fmt.Println(i)
	// 	fmt.Println(v)
	// }

	//pageSel := widget.NewSelect()

	upBtn = widget.NewButtonWithIcon("", theme.MoveUpIcon(), func() {
		MoveUp(MyApp.Selected.Row, MyApp.Selected.Column, MyApp)
		logic.OrderRows(MyApp)
		logic.CreateRowFile(MyApp)
		LoadGUI(MyApp)
	})
	downBtn = widget.NewButtonWithIcon("", theme.MoveDownIcon(), func() {
		MoveDown(MyApp.Selected.Row, MyApp.Selected.Column, MyApp)
		logic.OrderRows(MyApp)
		logic.CreateRowFile(MyApp)
		LoadGUI(MyApp)
	})
	leftBtn = widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		MoveLeft(MyApp.Selected.Row, MyApp.Selected.Column, MyApp)
		logic.CreateRowFile(MyApp)
		LoadGUI(MyApp)
	})
	rightBtn = widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		MoveRight(MyApp.Selected.Row, MyApp.Selected.Column, MyApp)
		logic.CreateRowFile(MyApp)
		LoadGUI(MyApp)
	})

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

	settingsBtn := widget.NewButton("Settings", func() { CreateSettingsPopUp(MyApp) })

	bottomContent := container.NewGridWithColumns(2, reorderBtn, settingsBtn)

	centerScroll := container.NewVScroll(mainContent)

	if MyApp.Reorder {
		allContent = container.NewBorder(topContent, bottomContent, nil, nil, centerScroll)
	} else {
		allContent = container.NewBorder(nil, bottomContent, nil, nil, centerScroll)
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
		if MyApp.Selected.Row == 0 && len(MyApp.Rows[MyApp.Selected.Row].Websites) == 1 {
			MyApp.UpBtn.Disable()
		}

		// If in bottom row and only item in its row
		if MyApp.Selected.Row == len(MyApp.Rows)-1 && len(MyApp.Rows[len(MyApp.Rows)-1].Websites) == 1 {
			MyApp.DownBtn.Disable()
		}

		return
	}

	// If label or website-row selected
	if MyApp.Selected.Mode == "Label" || MyApp.Selected.Mode == "Website-Row" {
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
