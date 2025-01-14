package ui

import (
	"homepage-maker/logic"
	"slices"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func CreateSettingsPopUp(MyApp *logic.MyApp) {
	var popUp *widget.PopUp

	viewOrDeleteIconsBtn := widget.NewButton("View or Delete Downloaded Icons", func() { ShowDownloadedIcons(MyApp) })

	aboutBtn := widget.NewButton("About", func() {})

	dismissBtn := widget.NewButton("Dismiss", func() { popUp.Hide() })

	content := container.NewVBox(viewOrDeleteIconsBtn, layout.NewSpacer(), aboutBtn, layout.NewSpacer(), dismissBtn)
	popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	popUp.Resize(fyne.NewSize(200, 200))
	popUp.Show()
}

func ShowDownloadedIcons(MyApp *logic.MyApp) {
	path, err := storage.Child(MyApp.App.Storage().RootURI(), "Img")

	if err != nil {
		return
	}

	list, _ := storage.List(path)

	slices.SortFunc(list, func(a, b fyne.URI) int {
		return strings.Compare(strings.ToLower(a.Name()), strings.ToLower(b.Name()))
	})

	var popUp *widget.PopUp
	hide := func() { popUp.Hide() }
	var content *fyne.Container
	btn := widget.NewButton("Dismiss", func() { hide() })

	if len(list) == 0 {
		lbl := widget.NewLabel("You have to download icons first in prior popup!")
		centeredLbl := container.NewCenter(lbl)

		content = container.NewVBox(centeredLbl, btn)

		popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	} else {
		lbl := widget.NewLabel("Click icon to delete")
		centeredLbl := container.NewCenter(lbl)
		var buttons []fyne.CanvasObject

		for _, v := range list {
			buttons = append(buttons, MakeDummyIconButton(v, hide, MyApp))
		}

		center := container.NewGridWrap(fyne.NewSize(64, 108), buttons...)
		scrollCenter := container.NewVScroll(center)

		content = container.NewBorder(centeredLbl, btn, nil, nil, scrollCenter)

		popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
		popUp.Resize(fyne.NewSize(MyApp.Win.Canvas().Size().Width*0.9, MyApp.Win.Canvas().Size().Height*0.75))
	}
	popUp.Show()
}

func MakeDummyIconButton(iconLocation fyne.URI, hide func(), MyApp *logic.MyApp) fyne.CanvasObject {
	file, _ := storage.LoadResourceFromURI(iconLocation)
	img := canvas.NewImageFromResource(file)
	imgPadded := container.NewPadded(img)
	lbl := widget.NewLabel(file.Name())

	btn := widget.NewButton("", func() {
		logic.DeleteFile(iconLocation, MyApp)
		hide()
		ShowDownloadedIcons(MyApp)
	})

	stack := container.NewStack(btn, imgPadded)

	content := container.NewBorder(nil, lbl, nil, nil, stack)

	return container.NewGridWrap(fyne.NewSize(64, 108), content)
}
