package logic

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

type Website struct {
	Name         string // Name for website button label
	Link         string // Link to the website
	IconLocation string // Path to icon/favicon
	Size         string // Size of icon
	WhiteBg      bool   // Toggles white background behind the icon
}

func SaveWebsite(row int, Website *Website, MyApp *MyApp) {
	CurrentRow := MyApp.Rows[row]

	CurrentRow.Websites = append(CurrentRow.Websites, *Website)

	MyApp.Rows[row] = CurrentRow

	//DownloadIcon(Website, MyApp)

	CreateRowFile(MyApp)
}

func EditWebsite(row int, column int, Website *Website, MyApp *MyApp) {
	CurrentRow := MyApp.Rows[row]

	CurrentRow.Websites[column] = *Website

	MyApp.Rows[row] = CurrentRow

	//DownloadIcon(Website, MyApp)

	CreateRowFile(MyApp)
}

func DeleteWebsite(row int, column int, MyApp *MyApp) {
	if len(MyApp.Rows[row].Websites) == 1 {
		MyApp.Rows = slices.Delete(MyApp.Rows, row, row+1)
	} else {
		MyApp.Rows[row].Websites = slices.Delete(MyApp.Rows[row].Websites, column, column+1)
	}
}

func LoadIcon(Website *Website, MyApp *MyApp) fyne.Resource {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), Website.IconLocation)
	file, _ := storage.LoadResourceFromURI(path)

	return file
}
