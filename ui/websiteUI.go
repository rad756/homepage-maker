package ui

import (
	"homepage-maker/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeWebsiteButton(row int, column int, Website *logic.Website, MyApp *logic.MyApp) *fyne.Container {
	var mainBtn *widget.Button
	mainBtn = widget.NewButtonWithIcon("", theme.HelpIcon(), func() {
		if MyApp.Reorder {
			if MyApp.Selected.Mode == "Website" && MyApp.Selected.Row == row && MyApp.Selected.Column == column {
				// If selected was current website
				MyApp.Selected.Mode = ""
				MyApp.Selected.Row = 0
				MyApp.Selected.Column = 0

				mainBtn.Importance = 0
				mainBtn.Refresh()
			} else {
				MyApp.Selected.Mode = "Website"
				MyApp.Selected.Row = row
				MyApp.Selected.Column = column

				ClearButtonSelection(MyApp)

				mainBtn.Importance = 1
				mainBtn.Refresh()
			}
		} else {
			EditWebsitePopUp(row, column, MyApp)
		}
	})
	mainBtn.Icon = logic.LoadIcon(Website, MyApp)
	lbl := widget.NewLabel(Website.Name)

	MyApp.Buttons = append(MyApp.Buttons, mainBtn)
	MyApp.Websites = append(MyApp.Websites, Website)

	return container.NewBorder(nil, lbl, nil, nil, mainBtn)
}

func MakeBlankWebsiteButton(row int, MyApp *logic.MyApp) *fyne.Container {
	mainBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		MakeCreateWebsiteButtonPopUp(row, MyApp)
	})

	blankLbl := widget.NewLabel("")

	return container.NewBorder(nil, blankLbl, nil, nil, mainBtn)
}

func MakeCreateWebsiteButtonPopUp(row int, MyApp *logic.MyApp) {
	var createWebsiteButtonPopUp *widget.PopUp
	var iconBtn *widget.Button
	var nameEnt *widget.Entry
	var linkEnt *widget.Entry

	iconBtn = widget.NewButtonWithIcon("", theme.DownloadIcon(), func() {
		icon := logic.DownloadIconToMemory(linkEnt.Text)
		iconBtn.Icon = fyne.NewStaticResource("temp-icon", icon)
	})
	iconBtn.Resize(MyApp.GridSize)
	iconContainer := container.NewGridWrap(MyApp.GridSize, iconBtn)
	iconCentered := container.NewCenter(iconContainer)

	nameEnt = widget.NewEntry()
	nameEnt.SetPlaceHolder("Enter Name of Website")

	linkEnt = widget.NewEntry()
	linkEnt.SetPlaceHolder("Enter Link to Website")

	saveBtn := widget.NewButton("Save Website", func() {
		iconLocation := "Img/" + nameEnt.Text
		website := &logic.Website{Name: nameEnt.Text, Link: linkEnt.Text, IconLocation: iconLocation}
		logic.SaveWebsite(row, website, MyApp)
		createWebsiteButtonPopUp.Hide()
		LoadGUI(MyApp)
	})
	exitBtn := widget.NewButton("Discard", func() { createWebsiteButtonPopUp.Hide() })

	content := container.NewVBox(iconCentered, nameEnt, linkEnt, saveBtn, exitBtn)

	createWebsiteButtonPopUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	createWebsiteButtonPopUp.Resize(fyne.NewSize(200, 0))
	createWebsiteButtonPopUp.Show()
}

func EditWebsitePopUp(row int, column int, MyApp *logic.MyApp) {
	var createWebsiteButtonPopUp *widget.PopUp
	var iconBtn *widget.Button
	var nameEnt *widget.Entry
	var linkEnt *widget.Entry

	website := MyApp.Rows[row].Websites[column]

	iconBtn = widget.NewButtonWithIcon("", logic.LoadIcon(&website, MyApp), func() {
		icon := logic.DownloadIconToMemory(linkEnt.Text)
		iconBtn.Icon = fyne.NewStaticResource("temp-icon", icon)
	})
	iconBtn.Resize(MyApp.GridSize)
	iconContainer := container.NewGridWrap(MyApp.GridSize, iconBtn)
	iconCentered := container.NewCenter(iconContainer)

	nameEnt = widget.NewEntry()
	nameEnt.SetText(website.Name)

	linkEnt = widget.NewEntry()
	linkEnt.SetText(website.Link)

	editBtn := widget.NewButton("Edit Website", func() {
		iconLocation := "Img/" + nameEnt.Text
		website := &logic.Website{Name: nameEnt.Text, Link: linkEnt.Text, IconLocation: iconLocation}
		logic.EditWebsite(row, column, website, MyApp)
		createWebsiteButtonPopUp.Hide()
		LoadGUI(MyApp)
	})
	exitBtn := widget.NewButton("Discard", func() { createWebsiteButtonPopUp.Hide() })

	content := container.NewVBox(iconCentered, nameEnt, linkEnt, editBtn, exitBtn)

	createWebsiteButtonPopUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	createWebsiteButtonPopUp.Resize(fyne.NewSize(200, 0))
	createWebsiteButtonPopUp.Show()
}
