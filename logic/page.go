package logic

import (
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

func CreatePageFolder(location string, MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), GetCurrentPageName(MyApp)+location)
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
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), GetCurrentPageName(MyApp)+name)
	_ = storage.Delete(path)
}

func SubpageContainsNameCheck(name string, MyApp *MyApp) bool {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), GetCurrentPageName(MyApp))

	list, _ := storage.List(path)

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

func GetCurrentPageName(MyApp *MyApp) string {
	return lastDirectory(MyApp.Pages[MyApp.CurrentPage])
}
