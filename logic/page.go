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
	pathRow, _ := storage.Child(path, MyApp.App.Preferences().String("RowFileName"))

	file, _ := storage.Writer(pathRow)

	file.Write(nil)

	pathPage, _ := storage.Child(path, "Page.html")

	file, _ = storage.Writer(pathPage)

	file.Write(GetBlankPage(MyApp))

	CreateRowFile(MyApp)
}

func AddPage(name string, MyApp *MyApp) {
	CreatePageFolder(name, MyApp)

	//CreateHTMLFile(MyApp)
}

func DeletePageFolder(path fyne.URI) {
	list, _ := storage.List(path)

	for _, v := range list {
		listable, _ := storage.CanList(v)
		if !listable {
			_ = storage.Delete(v)
			//fmt.Println("Deleted: " + v.Path())
			continue
		} else {
			DeletePageFolder(v)
		}
	}

	err := storage.Delete(path)
	//fmt.Println("Deleted: " + path.Path())
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

func GetSubpages(path fyne.URI, MyApp *MyApp) *string {
	var s *string

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

func PageRename(name string, rowToBeModitied int, MyApp *MyApp) {
	path, err := storage.Child(MyApp.Pages[MyApp.CurrentPage], MyApp.Rows[rowToBeModitied].Name)
	fmt.Println(err)

	//fmt.Println(path)

	newPath, err := storage.Child(MyApp.Pages[MyApp.CurrentPage], name)
	//newPath, err := storage.ParseURI(MyApp.Pages[MyApp.CurrentPage].Path() + "/" + name)
	fmt.Println(err)
	//fmt.Println(newPath)

	err = storage.Move(path, newPath)
	fmt.Println(err)
}

func ContainsFolder(path fyne.URI) bool {
	list, _ := storage.List(path)

	for _, v := range list {
		listable, _ := storage.CanList(v)
		if !listable {
			continue
		} else {
			return true
		}
	}

	return false
}
