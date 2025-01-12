package ui

import (
	"fmt"
	"homepage-maker/logic"
	"regexp"
	"slices"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeWebsiteButton(row int, column int, Website *logic.Website, MyApp *logic.MyApp) *fyne.Container {
	var mainBtn *widget.Button
	mainBtn = widget.NewButtonWithIcon("", nil, func() {
		if MyApp.Reorder {
			if MyApp.Selected.Mode == "Website" && MyApp.Selected.Row == row && MyApp.Selected.Column == column {
				// If selected was current website
				MyApp.Selected.Mode = ""
				MyApp.Selected.Row = 0
				MyApp.Selected.Column = 0

				mainBtn.Importance = 0
				mainBtn.Refresh()
				LoadGUI(MyApp)
			} else {
				MyApp.Selected.Mode = "Website"
				MyApp.Selected.Row = row
				MyApp.Selected.Column = column
				//MyApp.OldSelectedColumn = column

				ClearButtonSelection(MyApp)

				mainBtn.Importance = 1
				mainBtn.Refresh()
				LoadGUI(MyApp)
			}
			SetReorderButtons(*MyApp)
		} else {
			EditWebsitePopUp(row, column, Website, MyApp)
		}
	})

	if MyApp.Selected.Mode == "Website" && MyApp.Selected.Row == row && MyApp.Selected.Column == column {
		mainBtn.Importance = 1
	} else {
		mainBtn.Importance = 0
	}

	img := canvas.NewImageFromResource(logic.LoadIcon(Website, MyApp))
	imgPadded := container.NewPadded(img)
	lbl := widget.NewLabel(Website.Name)

	stack := container.NewStack(mainBtn, imgPadded)

	MyApp.Buttons = append(MyApp.Buttons, mainBtn)
	MyApp.Websites = append(MyApp.Websites, Website)

	return container.NewBorder(nil, lbl, nil, nil, stack)
}

func MakeDummyWebsiteButton(row int, column int, Website *logic.Website, MyApp *logic.MyApp) *fyne.Container {
	var mainBtn *widget.Button
	mainBtn = widget.NewButtonWithIcon("", nil, nil)

	img := canvas.NewImageFromResource(logic.LoadIcon(Website, MyApp))
	imgPadded := container.NewPadded(img)
	lbl := widget.NewLabel(Website.Name)

	stack := container.NewStack(mainBtn, imgPadded)

	return container.NewBorder(nil, lbl, nil, nil, stack)
}

func MakeBlankWebsiteButton(row int, MyApp *logic.MyApp) *fyne.Container {
	mainBtn := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), func() {
		MakeCreateWebsiteButtonPopUp(row, MyApp)
	})

	blankLbl := widget.NewLabel("")

	return container.NewBorder(nil, blankLbl, nil, nil, mainBtn)
}

func MakeMoveRowButton(row int, MyApp *logic.MyApp) *fyne.Container {
	var mainBtn *widget.Button

	mainBtn = widget.NewButtonWithIcon("", theme.MoreHorizontalIcon(), func() {
		if MyApp.Selected.Mode == "Website-Row" && MyApp.Selected.Row == row {
			// If selected was current website
			MyApp.Selected.Mode = ""
			MyApp.Selected.Row = 0
			MyApp.Selected.Column = 0

			mainBtn.Importance = 0
			mainBtn.Refresh()
			SetReorderButtons(*MyApp)
			LoadGUI(MyApp)
		} else {
			MyApp.Selected.Mode = "Website-Row"
			MyApp.Selected.Row = row

			ClearButtonSelection(MyApp)

			mainBtn.Importance = 1
			mainBtn.Refresh()
			SetReorderButtons(*MyApp)
			LoadGUI(MyApp)
		}
	})

	MyApp.Buttons = append(MyApp.Buttons, mainBtn)

	if MyApp.Selected.Mode == "Website-Row" && MyApp.Selected.Row == row {
		mainBtn.Importance = 1
	} else {
		mainBtn.Importance = 0
	}

	blankLbl := widget.NewLabel("")

	return container.NewBorder(nil, blankLbl, nil, nil, mainBtn)
}

func MakeCreateWebsiteButtonPopUp(row int, MyApp *logic.MyApp) {
	var createWebsiteButtonPopUp *widget.PopUp
	var iconBtn *widget.Button
	var nameEnt *widget.Entry
	var linkEnt *widget.Entry
	var img *canvas.Image
	var size string

	website := &logic.Website{}
	lbl := widget.NewLabel("Add website to row or delete row")

	iconBtn = widget.NewButton("", func() {})
	iconBtn.Resize(MyApp.GridSize)
	//whiteBg := canvas.NewRectangle(color.White)
	//whiteBgPadded := container.NewPadded(whiteBg)
	img = canvas.NewImageFromResource(nil)
	imgPadded := container.NewPadded(img)
	stack := container.NewStack(iconBtn, imgPadded)
	iconContainer := container.NewGridWrap(MyApp.GridSize, stack)
	iconCentered := container.NewCenter(iconContainer)

	whiteBgCck := widget.NewCheck("White Background", func(b bool) {
		fmt.Println(b)
	})

	nameEnt = widget.NewEntry()
	nameEnt.SetPlaceHolder("Enter Name of Website")

	linkEnt = widget.NewEntry()
	linkEnt.SetPlaceHolder("Enter Link to Website")

	faviconDownloadBtn := widget.NewButton("Download Website's Icon", func() {
		icon16 := logic.DownloadIconToMemory(linkEnt.Text, "16")
		icon32 := logic.DownloadIconToMemory(linkEnt.Text, "32")
		icon64 := logic.DownloadIconToMemory(linkEnt.Text, "64")
		icon128 := logic.DownloadIconToMemory(linkEnt.Text, "128")

		website.Name = nameEnt.Text
		website.Link = linkEnt.Text
		website.IconLocation = "Img/" + nameEnt.Text

		DownloadFaviconPopUP(linkEnt.Text, icon16, icon32, icon64, icon128, &size, img, website, MyApp)
	})

	chooseSavedIconBtn := widget.NewButton("Choose Downloaded Icon", func() {
		chooseSavedIconPopUp(img, website, MyApp)
	})

	saveBtn := widget.NewButton("Save Website", func() {
		website.Name = nameEnt.Text
		website.Link = linkEnt.Text
		website.Size = size
		logic.SaveWebsite(row, website, MyApp)
		createWebsiteButtonPopUp.Hide()
		LoadGUI(MyApp)
	})

	deleteRowBtn := widget.NewButton("Delete Row", func() {
		ConfirmDeleteWebsiteRowPopUp(row, createWebsiteButtonPopUp, MyApp)
	})
	exitBtn := widget.NewButton("Discard", func() { createWebsiteButtonPopUp.Hide() })

	content := container.NewVBox(lbl, iconCentered, whiteBgCck, nameEnt, linkEnt, faviconDownloadBtn, chooseSavedIconBtn, saveBtn, deleteRowBtn, exitBtn)

	createWebsiteButtonPopUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	createWebsiteButtonPopUp.Resize(fyne.NewSize(200, 0))
	createWebsiteButtonPopUp.Show()
}

func DownloadFaviconPopUP(name string, icon16 []byte, icon32 []byte, icon64 []byte, icon128 []byte, size *string, img *canvas.Image, website *logic.Website, MyApp *logic.MyApp) {
	var popUp *widget.PopUp

	lbl := widget.NewLabel("Please select an icon for the website")
	lblCentered := container.NewCenter(lbl)

	btn16 := widget.NewButton("", func() {
		*size = "16"
		website.Size = *size
		logic.SaveIconFromMemory(website, icon16, MyApp)
		img.Resource = fyne.NewStaticResource("temp-icon", icon16)
		img.Refresh()
		popUp.Hide()
	})
	icon16Padded := container.NewPadded(canvas.NewImageFromResource(fyne.NewStaticResource("temp-icon", icon16)))
	stack16 := container.NewStack(btn16, icon16Padded)
	lbl16 := widget.NewLabel("16px")
	lbl16Centered := container.NewCenter(lbl16)
	border16 := container.NewBorder(nil, lbl16Centered, nil, nil, stack16)

	btn32 := widget.NewButton("", func() {
		*size = "32"
		website.Size = *size
		logic.SaveIconFromMemory(website, icon32, MyApp)
		img.Resource = fyne.NewStaticResource("temp-icon", icon32)
		img.Refresh()
		popUp.Hide()
	})
	icon32Padded := container.NewPadded(canvas.NewImageFromResource(fyne.NewStaticResource("temp-icon", icon32)))
	stack32 := container.NewStack(btn32, icon32Padded)
	lbl32 := widget.NewLabel("32px")
	lbl32Centered := container.NewCenter(lbl32)
	border32 := container.NewBorder(nil, lbl32Centered, nil, nil, stack32)

	btn64 := widget.NewButton("", func() {
		*size = "64"
		website.Size = *size
		logic.SaveIconFromMemory(website, icon64, MyApp)
		img.Resource = fyne.NewStaticResource("temp-icon", icon64)
		img.Refresh()
		popUp.Hide()
	})
	icon64Padded := container.NewPadded(canvas.NewImageFromResource(fyne.NewStaticResource("temp-icon", icon64)))
	stack64 := container.NewStack(btn64, icon64Padded)
	lbl64 := widget.NewLabel("64px")
	lbl64Centered := container.NewCenter(lbl64)
	border64 := container.NewBorder(nil, lbl64Centered, nil, nil, stack64)

	btn128 := widget.NewButton("", func() {
		*size = "128"
		website.Size = *size
		logic.SaveIconFromMemory(website, icon128, MyApp)
		img.Resource = fyne.NewStaticResource("temp-icon", icon128)
		img.Refresh()
		popUp.Hide()
	})
	icon128Padded := container.NewPadded(canvas.NewImageFromResource(fyne.NewStaticResource("temp-icon", icon128)))
	stack128 := container.NewStack(btn128, icon128Padded)
	lbl128 := widget.NewLabel("128px")
	lbl128Centered := container.NewCenter(lbl128)
	border128 := container.NewBorder(nil, lbl128Centered, nil, nil, stack128)

	grid := container.NewGridWrap(fyne.NewSize(64, 108), border16, border32, border64, border128)

	exitBtn := widget.NewButton("Discard", func() { popUp.Hide() })

	content := container.NewVBox(lblCentered, grid, exitBtn)
	popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	popUp.Resize(fyne.NewSize(276, 0))
	popUp.Show()
}

func chooseSavedIconPopUp(image *canvas.Image, website *logic.Website, MyApp *logic.MyApp) {
	path, err := storage.Child(MyApp.App.Storage().RootURI(), "Img")

	if err != nil {
		return
	}

	list, _ := storage.List(path)

	slices.SortFunc(list, func(a, b fyne.URI) int {
		return strings.Compare(strings.ToLower(a.Name()), strings.ToLower(b.Name()))
	})

	fmt.Println(list)

	var popUp *widget.PopUp
	hide := func() { popUp.Hide() }
	var content *fyne.Container
	btn := widget.NewButton("Dismiss", func() { popUp.Hide() })

	if len(list) == 0 {
		lbl := widget.NewLabel("You have to download icons first in prior popup!")
		centeredLbl := container.NewCenter(lbl)

		content = container.NewVBox(centeredLbl, btn)

		popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	} else {
		lbl := widget.NewLabel("Click to choose an icon for website")
		centeredLbl := container.NewCenter(lbl)
		var buttons []fyne.CanvasObject

		for _, v := range list {
			buttons = append(buttons, MakeIconSelectButton(v, hide, image, website, MyApp))
		}

		center := container.NewGridWrap(fyne.NewSize(64, 108), buttons...)
		scrollCenter := container.NewVScroll(center)

		content = container.NewBorder(centeredLbl, btn, nil, nil, scrollCenter)

		popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
		popUp.Resize(fyne.NewSize(MyApp.Win.Canvas().Size().Width*0.9, MyApp.Win.Canvas().Size().Height*0.75))
	}
	popUp.Show()
}

func MakeIconSelectButton(iconLocation fyne.URI, hidePopUp func(), image *canvas.Image, website *logic.Website, MyApp *logic.MyApp) fyne.CanvasObject {
	regex := regexp.MustCompile(MyApp.App.Storage().RootURI().Path())
	iconRelativePath := regex.ReplaceAllLiteralString(iconLocation.Path(), "")

	file, _ := storage.LoadResourceFromURI(iconLocation)
	img := canvas.NewImageFromResource(file)
	imgPadded := container.NewPadded(img)
	lbl := widget.NewLabel(file.Name())

	btn := widget.NewButton("", func() {
		website.IconLocation = iconRelativePath
		image.Resource = file
		image.Refresh()
		hidePopUp()
	})

	stack := container.NewStack(btn, imgPadded)

	content := container.NewBorder(nil, lbl, nil, nil, stack)

	return container.NewGridWrap(fyne.NewSize(64, 108), content)
}

func EditWebsitePopUp(row int, column int, Website *logic.Website, MyApp *logic.MyApp) {
	var createWebsiteButtonPopUp *widget.PopUp
	var iconBtn *widget.Button
	var nameEnt *widget.Entry
	var linkEnt *widget.Entry
	var img *canvas.Image
	var size string

	website := MyApp.Rows[row].Websites[column]

	img = canvas.NewImageFromResource(logic.LoadIcon(&website, MyApp))
	imgPadded := container.NewPadded(img)

	iconBtn = widget.NewButton("", func() {})

	stack := container.NewStack(iconBtn, imgPadded)
	stack.Resize(MyApp.GridSize)
	iconContainer := container.NewGridWrap(MyApp.GridSize, stack)
	iconCentered := container.NewCenter(iconContainer)

	nameEnt = widget.NewEntry()
	nameEnt.SetText(website.Name)

	linkEnt = widget.NewEntry()
	linkEnt.SetText(website.Link)

	faviconDownloadBtn := widget.NewButton("Download Website's Icon", func() {
		icon16 := logic.DownloadIconToMemory(linkEnt.Text, "16")
		icon32 := logic.DownloadIconToMemory(linkEnt.Text, "32")
		icon64 := logic.DownloadIconToMemory(linkEnt.Text, "64")
		icon128 := logic.DownloadIconToMemory(linkEnt.Text, "128")

		website.Name = nameEnt.Text
		website.Link = linkEnt.Text
		website.IconLocation = "Img/" + nameEnt.Text

		DownloadFaviconPopUP(linkEnt.Text, icon16, icon32, icon64, icon128, &size, img, &website, MyApp)
	})

	chooseSavedIconBtn := widget.NewButton("Choose Downloaded Icon", func() {
		chooseSavedIconPopUp(img, &website, MyApp)
	})

	editBtn := widget.NewButton("Edit Website", func() {
		website.Name = nameEnt.Text
		website.Link = linkEnt.Text
		website.Size = size
		logic.EditWebsite(row, column, &website, MyApp)
		createWebsiteButtonPopUp.Hide()
		LoadGUI(MyApp)
	})
	deleteBtn := widget.NewButton("Delete Website", func() {
		ConfirmDeleteWebsitePopUp(row, column, &website, createWebsiteButtonPopUp, MyApp)
	})
	exitBtn := widget.NewButton("Discard", func() { createWebsiteButtonPopUp.Hide() })

	content := container.NewVBox(iconCentered, nameEnt, linkEnt, faviconDownloadBtn, chooseSavedIconBtn, editBtn, deleteBtn, exitBtn)

	createWebsiteButtonPopUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	createWebsiteButtonPopUp.Resize(fyne.NewSize(200, 0))
	createWebsiteButtonPopUp.Show()
}

func ConfirmDeleteWebsitePopUp(row int, column int, website *logic.Website, previousPopUp *widget.PopUp, MyApp *logic.MyApp) {
	var popUp *widget.PopUp

	lbl := widget.NewLabel("Are you sure you want to delete the website below?")

	toBeDeleted := container.NewGridWrap(fyne.NewSize(64, 108), MakeDummyWebsiteButton(row, column, website, MyApp))
	centeredToBeDeleted := container.NewCenter(toBeDeleted)

	yesBtn := widget.NewButton("Yes", func() {
		logic.DeleteWebsite(row, column, MyApp)
		logic.OrderRows(MyApp)
		logic.CreateRowFile(MyApp)
		previousPopUp.Hide()
		popUp.Hide()
		LoadGUI(MyApp)
	})
	noBtn := widget.NewButton("No", func() {
		popUp.Hide()
	})

	content := container.NewVBox(lbl, centeredToBeDeleted, yesBtn, noBtn)
	popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	popUp.Show()
}

func ConfirmDeleteWebsiteRowPopUp(row int, previousPopUp *widget.PopUp, MyApp *logic.MyApp) {
	var popUp *widget.PopUp

	lbl := widget.NewLabel("Are you sure you want to delete the below row?")
	rowContent := LoadDummyWebsiteRowItems(MyApp.Rows[row], MyApp)

	yesBtn := widget.NewButton("Yes", func() {
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

	content := container.NewVBox(lbl, rowContent, yesBtn, noBtn)
	popUp = widget.NewModalPopUp(content, MyApp.Win.Canvas())
	popUp.Show()
}
