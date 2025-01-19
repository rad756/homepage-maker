package logic

import (
	"encoding/json"
	"fmt"
	"log"

	"fyne.io/fyne/v2/storage"
)

type Page struct {
	Name     string
	Location string
	Depth    int
	SubPages []Page
}

func CreateInitialPageFile(MyApp *MyApp) {
	if MyApp.Pages == nil {
		MyApp.CurrentPage = &Page{Location: "/Pages/", Depth: 1}
		MyApp.Pages = append(MyApp.Pages, *MyApp.CurrentPage)
	}

	path, _ := storage.Child(MyApp.App.Storage().RootURI(), MyApp.CurrentPage.Location+MyApp.App.Preferences().String("PageFileName"))

	file, _ := storage.Writer(path)

	mar, _ := json.Marshal(MyApp.Rows)

	file.Write(mar)
}

func CreatePageFile(Page Page, MyApp *MyApp) {
	name := Page.Location + Page.Name + "/" + MyApp.App.Preferences().String("PageFileName")

	fmt.Println(name)

	path, _ := storage.Child(MyApp.App.Storage().RootURI(), name)

	file, _ := storage.Writer(path)

	mar, _ := json.Marshal(MyApp.Rows)

	file.Write(mar)
}

func CreatePageFolder(Page Page, MyApp *MyApp) {
	name := Page.Location + Page.Name + "/"
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), name)
	err := storage.CreateListable(path)

	if err != nil {
		log.Fatal(err)
	}
}

func ReadPageFile(MyApp *MyApp) {
	if MyApp.Pages == nil {
		MyApp.CurrentPage = &Page{Location: "/Pages/", Depth: 1}
		MyApp.Pages = append(MyApp.Pages, *MyApp.CurrentPage)
	}

	name := MyApp.CurrentPage.Location + MyApp.App.Preferences().String("PageFileName")
	if PathExists(name, MyApp) {
		path, _ := storage.Child(MyApp.App.Storage().RootURI(), name)

		file, _ := storage.LoadResourceFromURI(path)

		json.Unmarshal(file.Content(), &MyApp.Rows)
	}
}
