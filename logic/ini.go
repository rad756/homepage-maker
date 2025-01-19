package logic

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MyApp struct {
	App      fyne.App
	Win      fyne.Window
	Rows     []Row
	GridSize fyne.Size
	Reorder  bool
	Selected Selected
	//OldSelectedColumn int
	Buttons     []*widget.Button
	Websites    []*Website
	UpBtn       *widget.Button
	DownBtn     *widget.Button
	LeftBtn     *widget.Button
	RightBtn    *widget.Button
	Pages       []Page
	CurrentPage Page
}

func Ini(MyApp *MyApp) {
	if MyApp.App.Preferences().BoolWithFallback("FirstRun", true) {
		MyApp.App.Preferences().SetBool("FirstRun", true)
	} else {
		MyApp.App.Preferences().SetBool("FirstRun", false)
	}

	MyApp.App.Preferences().SetString("RowFileName", "Rows.json")
	MyApp.App.Preferences().SetString("PageFileName", "Page.html")

	if MyApp.App.Preferences().Bool("FirstRun") {
		CreateImgFolder(MyApp)
		CreatePagesFolder(MyApp)
		CreatePageFile(MyApp)
		CreateRowFile(MyApp)
	} else {
		ReadPageFile(MyApp)
		ReadRowFile(MyApp)
	}

	size := 64
	MyApp.GridSize.Height = float32(size)
	MyApp.GridSize.Width = float32(size)
}

func CurrentlySelected(row int, column int, MyApp *MyApp) {
	MyApp.Selected.Row = row
	MyApp.Selected.Column = column
}
