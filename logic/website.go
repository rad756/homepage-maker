package logic

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

type Website struct {
	Name         string // Name for website button label
	Link         string // Link to the website
	IconLocation string // Path to icon/favicon
	Selected     bool   // Show or not buttons in border to move/re-order button in grid
}

func SaveWebsite(row int, Website *Website, MyApp *MyApp) {
	CurrentRow := MyApp.Rows[row]

	CurrentRow.Websites = append(CurrentRow.Websites, *Website)

	MyApp.Rows[row] = CurrentRow

	DownloadIcon(Website, MyApp)

	CreateRowFile(MyApp)
}

// func SaveIcon(Website *Website, MyApp *MyApp) fyne.Resource {
// 	file, _ := storage.Writer(Website.IconLocation)

// 	file.Write(dow)
// }

func LoadIcon(Website *Website, MyApp *MyApp) fyne.Resource {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), Website.IconLocation)
	file, _ := storage.LoadResourceFromURI(path)

	return file
}
