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
	Sublink  bool      //Optional - Only needed for sublink
	Link     string    // Optional - only for Label (hyperlink)
}

type Selected struct {
	Mode   string //Label or Website or "" for no selection
	Row    int
	Column int
}

func CreateRowFile(MyApp *MyApp) {
	if MyApp.Pages == nil {
		MyApp.CurrentPage = Page{Location: "/Pages/", Depth: 1}
	}

	name := MyApp.CurrentPage.Location + MyApp.App.Preferences().String("RowFileName")

	path, _ := storage.Child(MyApp.App.Storage().RootURI(), name)

	file, _ := storage.Writer(path)

	mar, _ := json.Marshal(MyApp.Rows)

	file.Write(mar)
}

func ReadRowFile(MyApp *MyApp) {
	name := MyApp.CurrentPage.Location + MyApp.App.Preferences().String("RowFileName")
	if PathExists(name, MyApp) {
		path, _ := storage.Child(MyApp.App.Storage().RootURI(), name)

		file, _ := storage.LoadResourceFromURI(path)

		json.Unmarshal(file.Content(), &MyApp.Rows)
	}
}

func OrderRows(MyApp *MyApp) {
	for i, _ := range MyApp.Rows {
		MyApp.Rows[i].Number = i
	}
}
