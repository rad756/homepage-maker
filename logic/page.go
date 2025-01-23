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
	// if isHomepageEmpty(MyApp) {
	// 	initilizeHomepage(MyApp)
	// }
	initilizeHomepage(MyApp)

	path, _ := storage.Child(MyApp.App.Storage().RootURI(), "Homepage.json")

	file, _ := storage.Writer(path)

	mar, _ := json.Marshal(MyApp.Homepage)

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

func ReadHomepageFile(MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), "Homepage.json")

	file, _ := storage.LoadResourceFromURI(path)

	json.Unmarshal(file.Content(), &MyApp.Homepage)
	MyApp.CurrentPage = MyApp.Homepage

	fmt.Println(path)
	fmt.Println(MyApp.CurrentPage)
}

func ReadPagesFile(MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), "Homepage.json")

	file, _ := storage.LoadResourceFromURI(path)

	json.Unmarshal(file.Content(), &MyApp.Homepage)
}

func AddPage(newPage Page, MyApp *MyApp) {
	insertPage(newPage, MyApp)

	CreatePageFolder(newPage, MyApp)
	CreateHomepageFile(MyApp)
	CreateHTMLFile(MyApp)
}

func DeletePageFolder(Page Page, MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), Page.Location)
	_ = storage.Delete(path)
}

func insertPage(newPage Page, MyApp *MyApp) {
	// If its a subpage/sublink of the homepage, it adds it to first layer of its subpage
	if newPage.Depth == 2 {
		MyApp.Homepage.SubPages = append(MyApp.Homepage.SubPages, newPage)
		MyApp.CurrentPage = MyApp.Homepage
		return
	}

	// for _, v := range MyApp.Homepage.SubPages {
	// 	if MyApp.CurrentPage.Depth != newPage.Depth -1 {

	// 	}
	// }
}

// func recursiveInsert(newPage Page, parentPage Page, MyApp *MyApp) bool {
// 	for _, v := range parentPage.SubPages{
// 		if
// 	}
// }

func SubpageContainsNameCheck(name string, MyApp *MyApp) bool {
	for _, v := range MyApp.CurrentPage.SubPages {
		if v.Name == name {
			return true
		}
	}

	return false
}
