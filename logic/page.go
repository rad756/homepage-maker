package logic

import (
	"encoding/json"
	"log"

	"fyne.io/fyne/v2/storage"
)

type Page struct {
	Name     string
	Location string
	Depth    int
	SubPages []Page
}

func isHomepageEmpty(MyApp *MyApp) bool {
	if MyApp.Homepage.Name == "" && MyApp.Homepage.Location == "" && MyApp.Homepage.Depth == 0 && MyApp.Homepage.SubPages == nil {
		return true
	}
	return false
}

func initilizeHomepage(MyApp *MyApp) {
	MyApp.CurrentPage = Page{Name: "Homepage", Location: "/Homepage/", Depth: 1}
	MyApp.Homepage = MyApp.CurrentPage
}

func CreateInitialHomepageFile(MyApp *MyApp) {
	if isHomepageEmpty(MyApp) {
		initilizeHomepage(MyApp)
	}

	path, _ := storage.Child(MyApp.App.Storage().RootURI(), MyApp.CurrentPage.Location+MyApp.App.Preferences().String("PageFileName"))

	file, _ := storage.Writer(path)

	mar, _ := json.Marshal(MyApp.CurrentPage)

	file.Write(mar)
}

//// This will go into HTML.go
// func CreateHTMLFile(Page Page, MyApp *MyApp) {
// 	name := Page.Location + Page.Name + "/" + MyApp.App.Preferences().String("PageFileName")

// 	path, _ := storage.Child(MyApp.App.Storage().RootURI(), name)

// 	file, _ := storage.Writer(path)

// 	mar, _ := json.Marshal(MyApp.CurrentPage)

// 	file.Write(mar)
// }

func CreateHomepageFile(MyApp *MyApp) {
	name := "Homepage.json"

	path, _ := storage.Child(MyApp.App.Storage().RootURI(), name)

	file, _ := storage.Writer(path)

	mar, _ := json.Marshal(MyApp.Homepage)

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
	if isHomepageEmpty(MyApp) {
		initilizeHomepage(MyApp)
	}

	name := MyApp.CurrentPage.Location + MyApp.App.Preferences().String("PageFileName")
	if PathExists(name, MyApp) {
		path, _ := storage.Child(MyApp.App.Storage().RootURI(), name)

		file, _ := storage.LoadResourceFromURI(path)

		json.Unmarshal(file.Content(), &MyApp.CurrentPage)
	}
}

func ReadPagesFile(MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), "Homepage.json")

	file, _ := storage.LoadResourceFromURI(path)

	json.Unmarshal(file.Content(), MyApp.Homepage)
}

func AddPage(newPage Page, MyApp *MyApp) {

	CreatePageFolder(newPage, MyApp)
	CreateHomepageFile(MyApp)
	CreateHTMLFile(MyApp)
}

func DeletePageFolder(Page Page, MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), Page.Location)
	_ = storage.Delete(path)
}
