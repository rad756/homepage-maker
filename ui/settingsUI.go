package ui

import (
	"homepage-maker/logic"
	"image/color"
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

	viewOrDeleteIconsBtn := widget.NewButton("View or Delete Downloaded Icons", func() { ShowDownloadedIcons(false, MyApp) })
	downloadIconBtn := widget.NewButton("Download Icon", func() { DownloadIconPopUp(MyApp) })

	aboutBtn := widget.NewButton("About", func() {})

	dismissBtn := widget.NewButton("Dismiss", func() { popUp.Hide() })

	content := container.NewVBox(viewOrDeleteIconsBtn, downloadIconBtn, layout.NewSpacer(), aboutBtn, layout.NewSpacer(), dismissBtn)
	popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	popUp.Resize(fyne.NewSize(200, 200))
	popUp.Show()
}

func ShowDownloadedIcons(whiteBackground bool, MyApp *logic.MyApp) {
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
			buttons = append(buttons, MakeDummyIconButton(v, hide, whiteBackground, MyApp))
		}

		center := container.NewGridWrap(fyne.NewSize(64, 108), buttons...)
		scrollCenter := container.NewVScroll(center)

		cck := widget.NewCheck("Preview with White Background", func(b bool) {
			whiteBackground = !whiteBackground
			popUp.Hide()
			ShowDownloadedIcons(whiteBackground, MyApp)
		})

		cck.Checked = whiteBackground

		centeredCck := container.NewCenter(cck)

		topContent := container.NewVBox(centeredLbl, centeredCck)

		content = container.NewBorder(topContent, btn, nil, nil, scrollCenter)

		popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
		popUp.Resize(fyne.NewSize(MyApp.Win.Canvas().Size().Width*0.9, MyApp.Win.Canvas().Size().Height*0.75))
	}
	popUp.Show()
}

func MakeDummyIconButton(iconLocation fyne.URI, hide func(), whiteBackground bool, MyApp *logic.MyApp) fyne.CanvasObject {
	var stack *fyne.Container
	file, _ := storage.LoadResourceFromURI(iconLocation)
	whiteBg := canvas.NewRectangle(color.White)
	whiteBgPadded := container.NewPadded(whiteBg)
	img := canvas.NewImageFromResource(file)
	imgPadded := container.NewPadded(img)
	lbl := widget.NewLabel(file.Name())

	btn := widget.NewButton("", func() {
		logic.DeleteFile(iconLocation, MyApp)
		hide()
		ShowDownloadedIcons(whiteBackground, MyApp)
	})

	if whiteBackground {
		stack = container.NewStack(btn, whiteBgPadded, imgPadded)
	} else {
		stack = container.NewStack(btn, imgPadded)
	}

	content := container.NewBorder(nil, lbl, nil, nil, stack)

	return container.NewGridWrap(fyne.NewSize(64, 108), content)
}

func DownloadIconPopUp(MyApp *logic.MyApp) {
	var popUp *widget.PopUp

	nameEnt := widget.NewEntry()
	nameEnt.SetPlaceHolder("Enter Name of Icon")

	linkEnt := widget.NewEntry()
	linkEnt.SetPlaceHolder("Enter Link")

	downloadFaviconBtn := widget.NewButton("Download Favicon of Link", func() {
		if nameEnt.Text == "" || linkEnt.Text == "" {
			return
		}

		icon16 := logic.DownloadIconToMemory(linkEnt.Text, "16")
		icon32 := logic.DownloadIconToMemory(linkEnt.Text, "32")
		icon64 := logic.DownloadIconToMemory(linkEnt.Text, "64")
		icon128 := logic.DownloadIconToMemory(linkEnt.Text, "128")

		DownloadFaviconDirectPopUP(nameEnt.Text, icon16, icon32, icon64, icon128, MyApp)
	})
	downloadIconBtn := widget.NewButton("Direct Download of Icon from Link", func() {
		if nameEnt.Text == "" || linkEnt.Text == "" {
			return
		}

		DownloadDirectIconPopUP(nameEnt.Text, linkEnt.Text, MyApp)
	})

	dismissBtn := widget.NewButton("Dismiss", func() { popUp.Hide() })

	content := container.NewVBox(nameEnt, linkEnt, layout.NewSpacer(), downloadFaviconBtn, downloadIconBtn, layout.NewSpacer(), dismissBtn)

	popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	popUp.Resize(fyne.NewSize(200, 250))
	popUp.Show()
}

func DownloadFaviconDirectPopUP(name string, icon16 []byte, icon32 []byte, icon64 []byte, icon128 []byte, MyApp *logic.MyApp) {
	var popUp *widget.PopUp
	var content *fyne.Container
	var stack16 *fyne.Container
	var stack32 *fyne.Container
	var stack64 *fyne.Container
	var stack128 *fyne.Container

	whiteBackground := false
	whiteBackgroundPadded := container.NewPadded(canvas.NewRectangle(color.White))

	website := &logic.Website{Name: name, IconLocation: "Img/" + name}
	lbl := widget.NewLabel("Please select an icon for the website")
	lblCentered := container.NewCenter(lbl)

	btn16 := widget.NewButton("", func() {
		website.Size = "16"
		logic.SaveIconFromMemory(website, icon16, MyApp)
		popUp.Hide()
	})
	icon16Padded := container.NewPadded(canvas.NewImageFromResource(fyne.NewStaticResource("temp-icon", icon16)))
	stack16 = container.NewStack(btn16, icon16Padded)
	lbl16 := widget.NewLabel("16px")
	lbl16Centered := container.NewCenter(lbl16)
	border16 := container.NewBorder(nil, lbl16Centered, nil, nil, stack16)

	btn32 := widget.NewButton("", func() {
		website.Size = "32"
		logic.SaveIconFromMemory(website, icon32, MyApp)
		popUp.Hide()
	})
	icon32Padded := container.NewPadded(canvas.NewImageFromResource(fyne.NewStaticResource("temp-icon", icon32)))
	stack32 = container.NewStack(btn32, icon32Padded)
	lbl32 := widget.NewLabel("32px")
	lbl32Centered := container.NewCenter(lbl32)
	border32 := container.NewBorder(nil, lbl32Centered, nil, nil, stack32)

	btn64 := widget.NewButton("", func() {
		website.Size = "64"
		logic.SaveIconFromMemory(website, icon64, MyApp)
		popUp.Hide()
	})
	icon64Padded := container.NewPadded(canvas.NewImageFromResource(fyne.NewStaticResource("temp-icon", icon64)))
	stack64 = container.NewStack(btn64, icon64Padded)
	lbl64 := widget.NewLabel("64px")
	lbl64Centered := container.NewCenter(lbl64)
	border64 := container.NewBorder(nil, lbl64Centered, nil, nil, stack64)

	btn128 := widget.NewButton("", func() {
		website.Size = "128"
		logic.SaveIconFromMemory(website, icon128, MyApp)
		popUp.Hide()
	})
	icon128Padded := container.NewPadded(canvas.NewImageFromResource(fyne.NewStaticResource("temp-icon", icon128)))
	stack128 = container.NewStack(btn128, icon128Padded)
	lbl128 := widget.NewLabel("128px")
	lbl128Centered := container.NewCenter(lbl128)
	border128 := container.NewBorder(nil, lbl128Centered, nil, nil, stack128)

	grid := container.NewGridWrap(fyne.NewSize(64, 108), border16, border32, border64, border128)

	exitBtn := widget.NewButton("Discard", func() { popUp.Hide() })

	cck := widget.NewCheck("Preview with White Background", func(b bool) {
		whiteBackground = !whiteBackground

		if whiteBackground {
			stack16.Objects = []fyne.CanvasObject{btn16, whiteBackgroundPadded, icon16Padded}
			stack32.Objects = []fyne.CanvasObject{btn32, whiteBackgroundPadded, icon32Padded}
			stack64.Objects = []fyne.CanvasObject{btn64, whiteBackgroundPadded, icon64Padded}
			stack128.Objects = []fyne.CanvasObject{btn128, whiteBackgroundPadded, icon128Padded}
		} else {
			stack16.Objects = []fyne.CanvasObject{btn16, icon16Padded}
			stack32.Objects = []fyne.CanvasObject{btn32, icon32Padded}
			stack64.Objects = []fyne.CanvasObject{btn64, icon64Padded}
			stack128.Objects = []fyne.CanvasObject{btn128, icon128Padded}
		}
	})

	content = container.NewVBox(lblCentered, cck, grid, exitBtn)
	popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	popUp.Resize(fyne.NewSize(276, 0))
	popUp.Show()
}

func DownloadDirectIconPopUP(name string, link string, MyApp *logic.MyApp) {
	var popUp *widget.PopUp

	lbl := widget.NewLabel("Do you want to save this icon?")
	lblCentered := container.NewCenter(lbl)

	iconBtn := widget.NewButton("", func() {})
	iconBtn.Resize(MyApp.GridSize)
	//whiteBg := canvas.NewRectangle(color.White)
	//whiteBgPadded := container.NewPadded(whiteBg)
	icon := logic.DownloadDirectIconToMemory(link)
	img := canvas.NewImageFromResource(fyne.NewStaticResource("temp-icon", icon))
	imgPadded := container.NewPadded(img)
	stack := container.NewStack(iconBtn, imgPadded)
	iconContainer := container.NewGridWrap(MyApp.GridSize, stack)
	iconCentered := container.NewCenter(iconContainer)

	yesBtn := widget.NewButton("Yes", func() {
		website := &logic.Website{Name: name, IconLocation: "Img/" + name}
		logic.SaveIconFromMemory(website, icon, MyApp)
		popUp.Hide()
	})
	noBtn := widget.NewButton("No", func() { popUp.Hide() })

	content := container.NewVBox(lblCentered, iconCentered, layout.NewSpacer(), yesBtn, layout.NewSpacer(), noBtn)
	popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	popUp.Resize(fyne.NewSize(200, 200))
	popUp.Show()
}
