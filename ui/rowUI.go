package ui

import (
	"errors"
	"homepage-maker/logic"
	"reflect"
	"slices"

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
	var linkEnt *widget.Entry
	var content *fyne.Container
	var radio *widget.RadioGroup

	nameEnt := widget.NewEntry()
	nameEnt.SetPlaceHolder("Enter Name of Label")
	linkEnt = widget.NewEntry()
	linkEnt.SetPlaceHolder("Enter Link of Hyperlink")

	labelBtn := widget.NewButton("Create Label Row", func() {
		if nameEnt.Text == "" {
			return
		}

		row := &logic.Row{Mode: "Label", Name: nameEnt.Text, Number: len(MyApp.Rows)}

		if radio.Selected == "Hyperlink" && linkEnt.Text == "" {
			return
		} else if radio.Selected == "Hyperlink" {
			row.Link = linkEnt.Text
		}

		if radio.Selected == "Sublink" {
			//location := MyApp.CurrentPage + nameEnt.Text + "/"
			logic.AddPage(nameEnt.Text, MyApp)

			row.Sublink = true
			logic.GetPages(MyApp)
		}

		MyApp.Rows = append(MyApp.Rows, *row)
		logic.CreateRowFile(MyApp)

		CreateRowPopUp.Hide()
		LoadGUI(MyApp)
	})

	orLbl := widget.NewLabel("OR")
	orLblCentered := container.NewCenter(orLbl)

	websiteRowBtn := widget.NewButton("Create Website Row", func() {
		row := &logic.Row{Mode: "Website", Number: len(MyApp.Rows)}

		MyApp.Rows = append(MyApp.Rows, *row)
		logic.CreateRowFile(MyApp)

		CreateRowPopUp.Hide()
		LoadGUI(MyApp)
	})

	exitBtn := widget.NewButton("Exit", func() { CreateRowPopUp.Hide() })

	radio = widget.NewRadioGroup([]string{"Label", "Hyperlink", "Sublink"}, func(s string) {
		if s == "Label" || s == "Sublink" {
			linkEnt.SetText("")
			content.Objects = []fyne.CanvasObject{radio, nameEnt, labelBtn, orLblCentered, websiteRowBtn, layout.NewSpacer(), exitBtn}
			content.Refresh()
		} else {
			content.Objects = []fyne.CanvasObject{radio, nameEnt, linkEnt, labelBtn, orLblCentered, websiteRowBtn, layout.NewSpacer(), exitBtn}
			content.Refresh()
		}
		nameEnt.Validate()
	})

	radio.Horizontal = true

	//Validate:
	nameEnt.Validator = func(in string) error {
		if in == "" {
			labelBtn.Disable()
			return errors.New("name cannot be empty")
		}

		if radio.Selected == "Sublink" && logic.SubpageContainsNameCheck(in, MyApp) {
			labelBtn.Disable()
			return errors.New("name cannot be contains this")
		}

		if radio.Selected == "Hyperlink" && (in == "" || linkEnt.Text == "") {
			labelBtn.Disable()
			return errors.New("hyperlink name cannot be empty")
		}
		labelBtn.Enable()
		return nil
	}

	linkEnt.Validator = func(in string) error {
		if radio.Selected != "Hyperlink" {
			return nil
		}
		if in == "" {
			labelBtn.Disable()
			return errors.New("hyperlink's link cannot be empty")
		}

		labelBtn.Enable()
		return nil
	}

	content = container.NewVBox(radio, nameEnt, labelBtn, orLblCentered, websiteRowBtn, layout.NewSpacer(), exitBtn)
	radio.SetSelected("Label")
	radio.Required = true

	CreateRowPopUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	CreateRowPopUp.Show()
}

func EditLabelPopUp(row int, MyApp *logic.MyApp) {
	var CreateRowPopUp *widget.PopUp

	nameEnt := widget.NewEntry()
	nameEnt.SetText(MyApp.Rows[row].Name)

	editBtn := widget.NewButton("Edit Label", func() {
		if nameEnt.Text == "" {
			return
		}
		row := &logic.Row{Mode: "Label", Name: nameEnt.Text, Number: row}

		MyApp.Rows[row.Number] = *row
		logic.CreateRowFile(MyApp)

		CreateRowPopUp.Hide()
		LoadGUI(MyApp)
	})

	deleteBtn := widget.NewButton("Delete Row", func() {
		ConfirmDeleteLabelRowPopUp(row, CreateRowPopUp, MyApp)
	})

	exitBtn := widget.NewButton("Exit", func() { CreateRowPopUp.Hide() })

	content := container.NewVBox(nameEnt, editBtn, layout.NewSpacer(), deleteBtn, layout.NewSpacer(), exitBtn)

	CreateRowPopUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	CreateRowPopUp.Show()
}

func EditSublinkPopUp(row int, MyApp *logic.MyApp) {
	var popUp *widget.PopUp
	var nameEnt *widget.Entry

	nameEnt = widget.NewEntry()
	nameEnt.SetText(MyApp.Rows[row].Name)

	editBtn := widget.NewButton("Edit Sublink", func() {
		row := &logic.Row{Mode: "Label", Name: nameEnt.Text, Number: row, Sublink: true}

		MyApp.Rows[row.Number] = *row
		logic.CreateRowFile(MyApp)

		popUp.Hide()
		LoadGUI(MyApp)
	})

	deleteBtn := widget.NewButton("Delete Sublink", func() {
		ConfirmDeleteLabelRowPopUp(row, popUp, MyApp)
	})

	exitBtn := widget.NewButton("Exit", func() { popUp.Hide() })

	//validate
	nameEnt.Validator = func(s string) error {
		if s == "" {
			editBtn.Disable()
			return errors.New("name cannot be empty")
		}

		if logic.SubpageContainsNameCheck(s, MyApp) {
			editBtn.Disable()
			return errors.New("already contains sublink/subpage")
		}

		editBtn.Enable()
		return nil
	}

	content := container.NewVBox(nameEnt, editBtn, layout.NewSpacer(), deleteBtn, layout.NewSpacer(), exitBtn)

	popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	popUp.Show()
}

func EditHyperlinkPopUp(row int, MyApp *logic.MyApp) {
	var popUp *widget.PopUp
	var nameEnt *widget.Entry
	var linkEnt *widget.Entry

	nameEnt = widget.NewEntry()
	nameEnt.SetText(MyApp.Rows[row].Name)

	linkEnt = widget.NewEntry()
	linkEnt.SetText(MyApp.Rows[row].Link)

	editBtn := widget.NewButton("Edit Hyperlink", func() {
		row := &logic.Row{Mode: "Label", Name: nameEnt.Text, Link: linkEnt.Text, Number: row}

		MyApp.Rows[row.Number] = *row
		logic.CreateRowFile(MyApp)

		popUp.Hide()
		LoadGUI(MyApp)
	})

	deleteRow := widget.NewButton("Delete Row", func() {
		ConfirmDeleteLabelRowPopUp(row, popUp, MyApp)
	})

	exitBtn := widget.NewButton("Dismiss", func() { popUp.Hide() })

	//validate
	nameEnt.Validator = func(s string) error {
		if s == "" {
			editBtn.Disable()
			return errors.New("name cannot be empty")
		}

		if linkEnt.Text != "" {
			editBtn.Enable()
		}

		return nil
	}

	linkEnt.Validator = func(s string) error {
		if s == "" {
			editBtn.Disable()
			return errors.New("link cannot be empty")
		}

		if nameEnt.Text != "" {
			editBtn.Enable()
		}

		return nil
	}

	content := container.NewVBox(nameEnt, linkEnt, editBtn, layout.NewSpacer(), deleteRow, layout.NewSpacer(), exitBtn)

	popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	popUp.Show()
}

func ConfirmDeleteLabelRowPopUp(row int, previousPopUp *widget.PopUp, MyApp *logic.MyApp) {
	var popUp *widget.PopUp
	var content *fyne.Container

	lbl := widget.NewLabel("Are you sure you want to delete the below row?")
	rowContent := LoadDummyLabelRow(MyApp.Rows[row], MyApp)

	subpagesLbl := widget.NewLabel("")

	if MyApp.Rows[row].Sublink {
		subpagesLbl.SetText("AND THE SUBLINKS/SUBPAGES LISTED BELOW\n" + *logic.GetSubpages(row, MyApp))
	}

	yesBtn := widget.NewButton("Yes", func() {
		if MyApp.Rows[row].Sublink {
			logic.DeletePageFolder(row, MyApp)
			logic.GetPages(MyApp)
		}
		MyApp.Rows = slices.Delete(MyApp.Rows, row, row+1)
		logic.OrderRows(MyApp)
		logic.CreateRowFile(MyApp)
		previousPopUp.Hide()
		popUp.Hide()
		LoadGUI(MyApp)
	})

	noBtn := widget.NewButton("No", func() {
		popUp.Hide()
	})

	if MyApp.Rows[row].Sublink {
		content = container.NewVBox(lbl, rowContent, subpagesLbl, yesBtn, noBtn)
	} else {
		content = container.NewVBox(lbl, rowContent, yesBtn, noBtn)
	}

	popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	popUp.Show()
}

func MakeBottomRowButton(MyApp *logic.MyApp) *fyne.Container {
	mainBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		MakeCreateRowPopUp(MyApp)
	})

	// If last row is a website row and it is empty, disable the button
	if len(MyApp.Rows) != 0 && MyApp.Rows[len(MyApp.Rows)-1].Mode == "Website" && len(MyApp.Rows[len(MyApp.Rows)-1].Websites) == 0 {
		mainBtn.Disable()
	}

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

func LoadDummyWebsiteRowItems(Row logic.Row, MyApp *logic.MyApp) *fyne.Container {
	var content []fyne.CanvasObject

	for i, v := range Row.Websites {

		content = append(content, MakeDummyWebsiteButton(Row.Number, i, &v, MyApp))
	}

	return container.NewGridWrap(fyne.NewSize(64, 108), content...)
}

func LoadLabelRow(Row logic.Row, MyApp *logic.MyApp) *fyne.Container {
	var lbl *widget.Button
	lbl = widget.NewButton("", func() {
		if MyApp.Reorder {
			if MyApp.Selected.Mode == "Label" && MyApp.Selected.Row == Row.Number {
				// If selected was current label
				MyApp.Selected.Mode = ""
				MyApp.Selected.Row = 0
				MyApp.Selected.Column = 0

				lbl.Importance = 0
				lbl.Refresh()
				LoadGUI(MyApp)
			} else {
				MyApp.Selected.Mode = "Label"
				MyApp.Selected.Row = Row.Number

				ClearButtonSelection(MyApp)

				lbl.Importance = 1
				lbl.Refresh()
				LoadGUI(MyApp)
			}
			SetReorderButtons(*MyApp)
		} else {
			if Row.Sublink {
				EditSublinkPopUp(Row.Number, MyApp)
				return
			}

			if Row.Link != "" {
				EditHyperlinkPopUp(Row.Number, MyApp)
				return
			}

			EditLabelPopUp(Row.Number, MyApp)
		}

		MyApp.Buttons = append(MyApp.Buttons, lbl)
	})

	// Set name of label to indicate if hyperlink, sublink or standard
	if Row.Link != "" {
		lbl.SetText("[H] " + Row.Name)
	} else if Row.Sublink {
		lbl.SetText("[S] " + Row.Name)
	} else {
		lbl.SetText(Row.Name)
	}

	if MyApp.Selected.Mode == "Label" && MyApp.Selected.Row == Row.Number {
		lbl.Importance = 1
	} else {
		lbl.Importance = 0
	}

	return container.NewHBox(lbl)
}

func LoadDummyLabelRow(Row logic.Row, MyApp *logic.MyApp) *fyne.Container {
	var lbl *widget.Button
	lbl = widget.NewButton(Row.Name, nil)

	return container.NewHBox(lbl)
}

func MoveLeft(row int, column int, MyApp *logic.MyApp) {
	currentRow := MyApp.Rows[row].Websites

	swapper := reflect.Swapper(currentRow)
	swapper(column, column-1)

	MyApp.Rows[row].Websites = currentRow

	logic.CurrentlySelected(row, column-1, MyApp)
	//MyApp.OldSelectedColumn = MyApp.Selected.Column
}

func MoveRight(row int, column int, MyApp *logic.MyApp) {
	currentRow := MyApp.Rows[row].Websites

	swapper := reflect.Swapper(currentRow)
	swapper(column+1, column)

	MyApp.Rows[row].Websites = currentRow

	logic.CurrentlySelected(row, column+1, MyApp)
	//MyApp.OldSelectedColumn = MyApp.Selected.Column
}

func MoveUp(row int, column int, MyApp *logic.MyApp) {
	if MyApp.Selected.Mode == "Label" || MyApp.Selected.Mode == "Website-Row" {
		rows := MyApp.Rows

		swapper := reflect.Swapper(rows)
		swapper(row, row-1)

		MyApp.Rows = rows

		logic.CurrentlySelected(row-1, column, MyApp)
		return
	}

	if MyApp.Selected.Mode == "Website" {
		website := MyApp.Rows[row].Websites[column]

		// If first row selected
		if MyApp.Selected.Row == 0 {
			newRow := logic.Row{Mode: "Website", Websites: []logic.Website{website}}
			logic.DeleteWebsite(row, column, MyApp)
			MyApp.Rows = append([]logic.Row{newRow}, MyApp.Rows...)
			logic.CurrentlySelected(0, 0, MyApp)
			return
		}

		// If row above is not website row and only item in the row (Basically row swap)
		if MyApp.Rows[row-1].Mode != "Website" && len(MyApp.Rows[row].Websites) == 1 {
			newRow := logic.Row{Mode: "Website", Websites: []logic.Website{website}}
			logic.DeleteWebsite(row, column, MyApp)
			MyApp.Rows = slices.Insert(MyApp.Rows, row-1, newRow)
			logic.CurrentlySelected(row-1, 0, MyApp)
			return
		}

		// If row above is not website row
		if MyApp.Rows[row-1].Mode != "Website" {
			newRow := logic.Row{Mode: "Website", Websites: []logic.Website{website}}
			logic.DeleteWebsite(row, column, MyApp)
			MyApp.Rows = slices.Insert(MyApp.Rows, row, newRow)
			logic.CurrentlySelected(row, 0, MyApp)
			return
		}

		// If current selected website is at column larger than row above
		if MyApp.Selected.Column >= len(MyApp.Rows[row-1].Websites) {
			MyApp.Rows[row-1].Websites = append(MyApp.Rows[row-1].Websites, website)
			logic.DeleteWebsite(row, column, MyApp)
			logic.CurrentlySelected(row-1, len(MyApp.Rows[row-1].Websites)-1, MyApp)
			return
		}

		// Insert website into row above at them same column position
		MyApp.Rows[row-1].Websites = slices.Insert(MyApp.Rows[row-1].Websites, column, website)
		logic.DeleteWebsite(row, column, MyApp)
		logic.CurrentlySelected(row-1, column, MyApp)
	}
}

func MoveDown(row int, column int, MyApp *logic.MyApp) {
	if MyApp.Selected.Mode == "Label" || MyApp.Selected.Mode == "Website-Row" {

		rows := MyApp.Rows

		swapper := reflect.Swapper(rows)
		swapper(row+1, row)

		MyApp.Rows = rows

		logic.CurrentlySelected(row+1, column, MyApp)
		return
	}

	if MyApp.Selected.Mode == "Website" {

		website := MyApp.Rows[row].Websites[column]

		// If last row selected
		if MyApp.Selected.Row == len(MyApp.Rows)-1 {
			newRow := logic.Row{Mode: "Website", Websites: []logic.Website{website}}
			logic.DeleteWebsite(row, column, MyApp)
			MyApp.Rows = append(MyApp.Rows, newRow)
			logic.CurrentlySelected(len(MyApp.Rows)-1, 0, MyApp)
			return
		}

		// If row below is not website row and only item in the row (Basically row swap)
		if MyApp.Rows[row+1].Mode != "Website" && len(MyApp.Rows[row].Websites) == 1 {
			newRow := logic.Row{Mode: "Website", Websites: []logic.Website{website}}
			logic.DeleteWebsite(row, column, MyApp)
			MyApp.Rows = slices.Insert(MyApp.Rows, row+1, newRow)
			logic.CurrentlySelected(row+1, 0, MyApp)
			return
		}

		// If row below is not website row
		if MyApp.Rows[row+1].Mode != "Website" {
			newRow := logic.Row{Mode: "Website", Websites: []logic.Website{website}}
			logic.DeleteWebsite(row, column, MyApp)
			MyApp.Rows = slices.Insert(MyApp.Rows, row+1, newRow)
			logic.CurrentlySelected(row+1, 0, MyApp)
			return
		}

		// If current selected website is at column larger than row below
		if MyApp.Selected.Column >= len(MyApp.Rows[row+1].Websites) {
			MyApp.Rows[row+1].Websites = append(MyApp.Rows[row+1].Websites, website)
			logic.DeleteWebsite(row, column, MyApp)
			logic.CurrentlySelected(row+1, len(MyApp.Rows[row+1].Websites)-1, MyApp)
			return
		}

		// Insert website into row below at the same column position WHILE selected website is the only in its row
		if MyApp.Selected.Column <= len(MyApp.Rows[row+1].Websites) && len(MyApp.Rows[row].Websites) == 1 {
			MyApp.Rows[row+1].Websites = slices.Insert(MyApp.Rows[row+1].Websites, column, website)
			logic.DeleteWebsite(row, column, MyApp)
			logic.CurrentlySelected(row, column, MyApp)
			return
		}

		// Insert website into row below at them same column position
		MyApp.Rows[row+1].Websites = slices.Insert(MyApp.Rows[row+1].Websites, column, website)
		logic.DeleteWebsite(row, column, MyApp)
		logic.CurrentlySelected(row+1, column, MyApp)

	}
}
