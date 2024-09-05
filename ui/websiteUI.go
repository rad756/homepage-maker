package ui

import (
	"homepage-maker/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeWebsiteButton(Website logic.Website, MyApp *logic.MyApp) *fyne.Container {
	upBtn := widget.NewButtonWithIcon("", theme.MoveUpIcon(), nil)
	downBtn := widget.NewButtonWithIcon("", theme.MoveDownIcon(), nil)
	leftBtn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), nil)
	rightBtn := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), nil)
	mainBtn := widget.NewButtonWithIcon("", theme.HelpIcon(), nil)
	lbl := widget.NewLabel(Website.Name)

	insideBorder := container.NewBorder(nil, lbl, nil, nil, mainBtn)

	if Website.Selected {
		return container.NewBorder(upBtn, downBtn, leftBtn, rightBtn, insideBorder)
	} else {
		return container.NewBorder(nil, nil, nil, nil, insideBorder)
	}
}

func MakeBlankWebsiteButton(MyApp *logic.MyApp) *fyne.Container {
	mainBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		MakeCreateWebsiteButtonPopUp(MyApp)
	})

	return container.NewBorder(nil, nil, nil, nil, mainBtn)
}

func MakeCreateWebsiteButtonPopUp(MyApp *logic.MyApp) {
	var createWebsiteButtonPopUp *widget.PopUp
	var iconBtn *widget.Button
	var nameEnt *widget.Entry
	var linkEnt *widget.Entry

	iconBtn = widget.NewButtonWithIcon("", theme.DownloadIcon(), func() {
		icon := logic.DownloadImageToMemory(linkEnt.Text)
		iconBtn.Icon = fyne.NewStaticResource("temp-icon", icon)
	})
	iconContainer := container.NewGridWrap(MyApp.GridSize, iconBtn)
	iconCentered := container.NewCenter(iconContainer)

	nameEnt = widget.NewEntry()
	nameEnt.SetPlaceHolder("Enter Name of Website")

	linkEnt = widget.NewEntry()
	linkEnt.SetPlaceHolder("Enter Link to Website")

	saveBtn := widget.NewButton("Save Website", func() {
		iconLocation, _ := storage.Child(MyApp.App.Storage().RootURI(), "Img/"+nameEnt.Text)
		website := &logic.Website{Name: nameEnt.Text, Link: linkEnt.Text, IconLocation: iconLocation}
		logic.SaveWebsite(website, MyApp)
	})
	exitBtn := widget.NewButton("Discard", func() { createWebsiteButtonPopUp.Hide() })

	content := container.NewVBox(iconCentered, nameEnt, linkEnt, saveBtn, exitBtn)

	createWebsiteButtonPopUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	createWebsiteButtonPopUp.Resize(fyne.NewSize(200, 0))
	createWebsiteButtonPopUp.Show()
}
