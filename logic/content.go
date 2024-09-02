package logic

import (
	"encoding/json"

	"fyne.io/fyne/v2/storage"
)

type Content struct {
	Mode     string    // Label or Button
	Name     string    // Optional - needed for label
	Websites []Website // Optional - needed for button
}

func CreateContentFile(MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), MyApp.App.Preferences().String("ContentFileName"))

	file, _ := storage.Writer(path)

	mar, _ := json.Marshal(MyApp.Content)

	file.Write(mar)
}
