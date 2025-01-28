package logic

import (
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

//// This will go into HTML.go
// func CreateHTMLFile(Page Page, MyApp *MyApp) {
// 	name := Page.Location + Page.Name + "/" + MyApp.App.Preferences().String("PageFileName")

// 	path, _ := storage.Child(MyApp.App.Storage().RootURI(), name)

// 	file, _ := storage.Writer(path)

// 	mar, _ := json.Marshal(MyApp.CurrentPage)

// 	file.Write(mar)
// }

func CreatePageFolder(location string, MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), MyApp.CurrentPage+location)
	err := storage.CreateListable(path)

	if err != nil {
		log.Fatal(err)
	}
}

func AddPage(name string, MyApp *MyApp) {
	CreatePageFolder(name, MyApp)
	CreateHTMLFile(MyApp)
}

func DeletePageFolder(name string, MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), MyApp.CurrentPage+name)
	_ = storage.Delete(path)
}

func SubpageContainsNameCheck(name string, MyApp *MyApp) bool {
	// for _, v := range MyApp.CurrentPage.SubPages {
	// 	if v.Name == name {
	// 		return true
	// 	}
	// }
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), MyApp.CurrentPage)

	//y := strings.Replace(x.Path(), MyApp.App.Storage().RootURI().Path(), "", -1)

	list, _ := storage.List(path)

	//fmt.Println(list)
	for _, v := range list {
		if strings.EqualFold(lastDirectory(v), name) {
			return true
		}
	}

	return false
}

func lastDirectory(location fyne.URI) string {
	directories := strings.Split(location.Path(), "/")
	return directories[len(directories)-1]
}

func GetPages(MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), "Homepage")
	MyApp.Pages = []fyne.URI{path}
	getDirectories(path, MyApp)
	for _, v := range MyApp.Pages {
		fmt.Println(v)
	}
}

func getDirectories(path fyne.URI, MyApp *MyApp) {
	list, _ := storage.List(path)

	for _, v := range list {
		listable, _ := storage.CanList(v)
		if !listable {
			return
		}

		MyApp.Pages = append(MyApp.Pages, v)
		getDirectories(v, MyApp)
	}

}

func containsDirectories(directory fyne.URI) bool {
	listable, _ := storage.CanList(directory)

	return listable
}
