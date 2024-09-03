package logic

import (
	"encoding/json"

	"fyne.io/fyne/v2/storage"
)

type Row struct {
	Mode     string    // Label or Website
	Name     string    // Optional - needed for label
	Websites []Website // Optional - needed for button
}

func CreateRowFile(MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), MyApp.App.Preferences().String("RowFileName"))

	file, _ := storage.Writer(path)

	mar, _ := json.Marshal(MyApp.Rows)

	file.Write(mar)
}
