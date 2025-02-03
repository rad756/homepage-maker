package logic

import (
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

func CreatePageFolder(name string, MyApp *MyApp) {
	//path, _ := storage.Child(MyApp.App.Storage().RootURI(), GetCurrentPageName(MyApp)+location)
	//path, _ := storage.Child(MyApp.Pages[MyApp.CurrentPage], MyApp.App.Preferences().String("RowFileName"))
	path, _ := storage.Child(MyApp.Pages[MyApp.CurrentPage], name)
	err := storage.CreateListable(path)

	if err != nil {
		log.Fatal(err)
	}

	// Creates empty
	path, _ = storage.Child(path, MyApp.App.Preferences().String("RowFileName"))

	file, _ := storage.Writer(path)

	file.Write(nil)

	CreateRowFile(MyApp)
}

func AddPage(name string, MyApp *MyApp) {
	CreatePageFolder(name, MyApp)
	CreateHTMLFile(MyApp)
}

func DeletePageFolder(row int, MyApp *MyApp) {
	path, err := storage.Child(MyApp.Pages[MyApp.CurrentPage], MyApp.Rows[row].Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)

	err = storage.Delete(path)
	if err != nil {
		fmt.Println(err)
	}
}

func SubpageContainsNameCheck(name string, MyApp *MyApp) bool {
	list, _ := storage.List(MyApp.Pages[MyApp.CurrentPage])

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
			continue
		}

		MyApp.Pages = append(MyApp.Pages, v)
		getDirectories(v, MyApp)
	}
}

func GetCurrentPageName(MyApp *MyApp) string {
	return lastDirectory(MyApp.Pages[MyApp.CurrentPage])
}

func GetSubpages(row int, MyApp *MyApp) *string {
	var s *string

	path, _ := storage.Child(MyApp.Pages[MyApp.CurrentPage], MyApp.Rows[row].Name)

	x := ""

	s = &x

	s = getSubDirectories(path, s, MyApp)

	return s
}

func getSubDirectories(path fyne.URI, s *string, MyApp *MyApp) *string {
	list, _ := storage.List(path)

	for _, v := range list {
		listable, _ := storage.CanList(v)
		if !listable {
			continue
		}

		x := strings.Split(v.Path(), "/")

		z := *s + x[len(x)-1] + "\n"
		s = &z

		s = getSubDirectories(v, s, MyApp)
	}

	return s
}
