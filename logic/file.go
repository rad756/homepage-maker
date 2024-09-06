package logic

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"fyne.io/fyne/v2/storage"
)

func PathExists(s string, MyApp *MyApp) bool {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), s)
	exists, _ := storage.Exists(path)

	return exists
}

func DownloadIconToMemory(link string) []byte {
	uri := fmt.Sprintf("https://www.google.com/s2/favicons?domain=%s&sz=64", link)

	resp, err := http.Get(uri)

	defer resp.Body.Close()

	body := resp.Body

	buf := new(bytes.Buffer)

	_, err = io.Copy(buf, body)

	if err != nil {
		return []byte{}
	}

	return buf.Bytes()
}

func DownloadIcon(Website *Website, MyApp *MyApp) {
	uri := fmt.Sprintf("https://www.google.com/s2/favicons?domain=%s&sz=64", Website.Link)

	resp, _ := http.Get(uri)

	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	path, _ := storage.Child(MyApp.App.Storage().RootURI(), Website.IconLocation)

	file, _ := storage.Writer(path)

	file.Write(data)
}

func CreateImgFolder(MyApp *MyApp) {
	path, _ := storage.Child(MyApp.App.Storage().RootURI(), "Img")
	err := storage.CreateListable(path)

	if err != nil {
		log.Fatal(err)
	}
}
