package logic

import (
	"encoding/json"

	"fyne.io/fyne/v2/storage"
)

type Page struct {
	Location string
	Depth    int
}

func CreatePageFile(MyApp *MyApp) {
	if MyApp.Pages == nil {
		MyApp.CurrentPage = Page{Location: "/Pages/", Depth: 1}
		MyApp.Pages = append(MyApp.Pages, MyApp.CurrentPage)
	}

	name := MyApp.CurrentPage.Location + MyApp.App.Preferences().String("PageFileName")

	path, _ := storage.Child(MyApp.App.Storage().RootURI(), name)

	file, _ := storage.Writer(path)

	mar, _ := json.Marshal(MyApp.Rows)

	file.Write(mar)
}

func ReadPageFile(MyApp *MyApp) {
	if MyApp.Pages == nil {
		MyApp.CurrentPage = Page{Location: "/Pages/", Depth: 1}
		MyApp.Pages = append(MyApp.Pages, MyApp.CurrentPage)
	}

	name := MyApp.CurrentPage.Location + MyApp.App.Preferences().String("PageFileName")
	if PathExists(name, MyApp) {
		path, _ := storage.Child(MyApp.App.Storage().RootURI(), name)

		file, _ := storage.LoadResourceFromURI(path)

		json.Unmarshal(file.Content(), &MyApp.Rows)
	}
}
