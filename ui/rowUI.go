package ui

import (
	"fmt"
	"homepage-maker/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func LoadRows(MyApp *logic.MyApp) *fyne.Container {
	var content []fyne.CanvasObject

	if len(MyApp.Rows) == 0 {
		return container.NewVBox(MakeBottomRowButton(MyApp))
	}

	for i, v := range MyApp.Rows {
		content = append(content, LoadWebsiteRowItems(v, MyApp))

		if i == len(MyApp.Rows) {
			content = append(content, MakeBottomRowButton(MyApp))
		}
	}

	return container.NewVBox(content...)
}

func MakeCreateRowPopUp(MyApp *logic.MyApp) {
	var CreateRowPopUp *widget.PopUp

	labelBtn := widget.NewButton("Create Label Row", nil)
	websiteRowBtn := widget.NewButton("Create Website Row", func() {
		row := &logic.Row{Mode: "Website"}

		MyApp.Rows = append(MyApp.Rows, *row)

		CreateRowPopUp.Hide()
		LoadGUI(MyApp)
	})

	exitBtn := widget.NewButton("Exit", func() { CreateRowPopUp.Hide() })

	content := container.NewVBox(labelBtn, websiteRowBtn, layout.NewSpacer(), exitBtn)

	CreateRowPopUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	CreateRowPopUp.Show()
}

func MakeBottomRowButton(MyApp *logic.MyApp) *fyne.Container {
	mainBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		MakeCreateRowPopUp(MyApp)
	})

	return container.NewVBox(mainBtn)
}

func LoadWebsiteRowItems(Row logic.Row, MyApp *logic.MyApp) *fyne.Container {
	if len(Row.Websites) == 0 {
		return container.NewGridWrap(fyne.NewSize(32, 32), MakeBlankWebsiteButton(MyApp))
	}

	var content []fyne.CanvasObject

	website := logic.Website{}

	for i, v := range Row.Websites {
		fmt.Println(i)
		fmt.Println(v)

		if i == len(MyApp.Rows) {
			content = append(content, MakeBlankWebsiteButton(MyApp))
			break
		}

		content = append(content, MakeWebsiteButton(website, MyApp))

	}
	return container.NewGridWrap(fyne.NewSize(32, 32), content...)
}
