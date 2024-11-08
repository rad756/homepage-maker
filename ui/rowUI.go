package ui

import (
	"homepage-maker/logic"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func LoadRows(MyApp *logic.MyApp) *fyne.Container {
	var content []fyne.CanvasObject

	for _, v := range MyApp.Rows {

		if v.Mode == "Label" {
			content = append(content, LoadLabelRow(v, MyApp))

		} else if v.Mode == "Website" {
			content = append(content, LoadWebsiteRowItems(v, MyApp))
		}

	}

	if !MyApp.Reorder {
		content = append(content, MakeBottomRowButton(MyApp))
	}

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

func EditLabelPopUp(row int, MyApp *logic.MyApp) {
	var CreateRowPopUp *widget.PopUp

	nameEnt := widget.NewEntry()
	nameEnt.SetText(MyApp.Rows[row].Name)

	labelBtn := widget.NewButton("Edit Label", func() {
		if nameEnt.Text == "" {
			return
		}
		row := &logic.Row{Mode: "Label", Name: nameEnt.Text, Number: row}

		MyApp.Rows[row.Number] = *row
		logic.CreateRowFile(MyApp)

		CreateRowPopUp.Hide()
		LoadGUI(MyApp)
	})

	exitBtn := widget.NewButton("Exit", func() { CreateRowPopUp.Hide() })

	content := container.NewVBox(nameEnt, labelBtn, layout.NewSpacer(), exitBtn)

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
	var content []fyne.CanvasObject

	for i, v := range Row.Websites {

		content = append(content, MakeWebsiteButton(Row.Number, i, &v, MyApp))
	}

	if !MyApp.Reorder {
		content = append(content, MakeBlankWebsiteButton(Row.Number, MyApp))
	} else {
		content = append(content, MakeMoveRowButton(Row.Number, MyApp))
	}

	return container.NewGridWrap(fyne.NewSize(64, 108), content...)
}

func LoadLabelRow(Row logic.Row, MyApp *logic.MyApp) *fyne.Container {
	var lbl *widget.Button
	lbl = widget.NewButton(Row.Name, func() {
		if MyApp.Reorder {
			if MyApp.Selected.Mode == "Label" && MyApp.Selected.Row == Row.Number {
				// If selected was current label
				MyApp.Selected.Mode = ""
				MyApp.Selected.Row = 0
				MyApp.Selected.Column = 0

				lbl.Importance = 0
				lbl.Refresh()
			} else {
				MyApp.Selected.Mode = "Label"
				MyApp.Selected.Row = Row.Number

				ClearButtonSelection(MyApp)

				lbl.Importance = 1
				lbl.Refresh()
			}
			SetReorderButtons(*MyApp)
		} else {
			EditLabelPopUp(Row.Number, MyApp)
		}

		MyApp.Buttons = append(MyApp.Buttons, lbl)
	})

	if MyApp.Selected.Mode == "Label" && MyApp.Selected.Row == Row.Number {
		lbl.Importance = 1
	}

	return container.NewHBox(lbl)
}

func MoveLeft(row int, column int, MyApp *logic.MyApp) {
	currentRow := MyApp.Rows[row].Websites

	swapper := reflect.Swapper(currentRow)
	swapper(column, column-1)

	MyApp.Rows[row].Websites = currentRow

	logic.CurrentlySelected(row, column-1, MyApp)
}

func MoveRight(row int, column int, MyApp *logic.MyApp) {
	currentRow := MyApp.Rows[row].Websites

	swapper := reflect.Swapper(currentRow)
	swapper(column+1, column)

	MyApp.Rows[row].Websites = currentRow

	logic.CurrentlySelected(row, column+1, MyApp)
}

func MoveUp(row int, column int, MyApp *logic.MyApp) {
	if MyApp.Selected.Mode == "Label" {
		rows := MyApp.Rows

		swapper := reflect.Swapper(rows)
		swapper(row, row-1)

		MyApp.Rows = rows

		logic.CurrentlySelected(row-1, column, MyApp)
	}

	for i, _ := range MyApp.Rows {
		MyApp.Rows[i].Number = i
	}
}

func MoveDown(row int, column int, MyApp *logic.MyApp) {
	if MyApp.Selected.Mode == "Label" {

		rows := MyApp.Rows

		swapper := reflect.Swapper(rows)
		swapper(row+1, row)

		MyApp.Rows = rows

		logic.CurrentlySelected(row+1, column, MyApp)
	}

	for i, _ := range MyApp.Rows {
		MyApp.Rows[i].Number = i
	}
}
