package logic

import (
	"encoding/json"

	"fyne.io/fyne/v2/storage"
)

type Row struct {
	Mode     string    // Label or Website
	Name     string    // Optional - needed for label
	Websites []Website // Optional - needed for button
	Number   int       // Needed
	Sublink  bool      // Optional - Only needed for sublink
	Link     string    // Optional - only for Label (hyperlink)
}

type Selected struct {
	Mode   string //Label or Website or "" for no selection
	Row    int
	Column int
}

// path, _ := storage.Child(MyApp.Pages[MyApp.CurrentPage], MyApp.App.Preferences().String("RowFileName"))
func CreateRowFile(MyApp *MyApp) {
	path, _ := storage.Child(MyApp.Pages[MyApp.CurrentPage], MyApp.App.Preferences().String("RowFileName"))

	file, _ := storage.Writer(path)

	mar, _ := json.Marshal(MyApp.Rows)

	file.Write(mar)

	CreateHTMLFile(MyApp)
}

func ReadRowFile(MyApp *MyApp) {
	path, _ := storage.Child(MyApp.Pages[MyApp.CurrentPage], MyApp.App.Preferences().String("RowFileName"))

	file, _ := storage.LoadResourceFromURI(path)

	//Clear Rows
	MyApp.Rows = nil

	//Fill Rows
	json.Unmarshal(file.Content(), &MyApp.Rows)
}

func OrderRows(MyApp *MyApp) {
	for i, _ := range MyApp.Rows {
		MyApp.Rows[i].Number = i
	}
}

func ContainsSubsite(row int, MyApp *MyApp) bool {

	websites := MyApp.Rows[row].Websites

	for _, v := range websites {
		if v.Subsite {
			return true
		}
	}

	return false
}
