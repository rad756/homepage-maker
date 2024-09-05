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

	for _, v := range MyApp.Rows {

		if v.Mode == "Label" {
			content = append(content, LoadLabelRow(v, MyApp))

		} else if v.Mode == "Website" {
			content = append(content, LoadWebsiteRowItems(v, MyApp))
		}

	}

	content = append(content, MakeBottomRowButton(MyApp))

	return container.NewVBox(content...)
}

func MakeCreateRowPopUp(MyApp *logic.MyApp) {
	var CreateRowPopUp *widget.PopUp

	nameEnt := widget.NewEntry()
	nameEnt.SetPlaceHolder("Enter Name of Label")

	labelBtn := widget.NewButton("Create Label Row", func() {
		if nameEnt.Text == "" {
			return
		}
		row := &logic.Row{Mode: "Label", Name: nameEnt.Text, Number: len(MyApp.Rows)}

		MyApp.Rows = append(MyApp.Rows, *row)
		logic.CreateRowFile(MyApp)

		CreateRowPopUp.Hide()
		LoadGUI(MyApp)
	})
	websiteRowBtn := widget.NewButton("Create Website Row", func() {
		row := &logic.Row{Mode: "Website", Number: len(MyApp.Rows)}

		MyApp.Rows = append(MyApp.Rows, *row)
		logic.CreateRowFile(MyApp)

		CreateRowPopUp.Hide()
		LoadGUI(MyApp)
	})

	exitBtn := widget.NewButton("Exit", func() { CreateRowPopUp.Hide() })

	content := container.NewVBox(nameEnt, labelBtn, websiteRowBtn, layout.NewSpacer(), exitBtn)

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
		return container.NewGridWrap(MyApp.GridSize, MakeBlankWebsiteButton(MyApp))
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

func LoadLabelRow(Row logic.Row, MyApp *logic.MyApp) *fyne.Container {
	lbl := widget.NewLabel(Row.Name)

	return container.NewHBox(lbl)
}
