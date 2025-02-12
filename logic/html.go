package logic

import (
	"encoding/json"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

func CreateHTMLFile(MyApp *MyApp) {
	var page string

	path, _ := storage.Child(MyApp.Pages[MyApp.CurrentPage], "Rows.json")

	page = `<!DOCTYPE html>
<html lang="en">
`
	page = appendHead(page, MyApp)
	page = appendBody(page, path, MyApp)

	path, _ = storage.Child(MyApp.Pages[MyApp.CurrentPage], "Page.html")

	file, _ := storage.Writer(path)

	file.Write([]byte(page))
}

func GetBlankPage(MyApp *MyApp) []byte {
	page := ""
	page = appendHead(page, MyApp)
	page = page + `<body></body>`
	return []byte(page)
}

func appendHead(page string, MyApp *MyApp) string {
	head := `<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>` + GetCurrentPageName(MyApp) + `</title>
</head>
<style>
    body {
        background-color: rgb(119, 119, 119);
        padding: 0px;
        margin: 0px;
        margin-left: 10px;
        margin-top: 5px;
    }

    ul {
        list-style: none;
        padding-left: 20px;
        text-align: justify;
    }

    li {
        display: inline-block;
        text-align: center;
        padding-right: 20px;
    }

    h2{
        padding: 0px;
        margin: 0px;
    }

    img {
        width: 60px;
        height: 60px;
        background-color: rgb(100, 100, 100);
        border-radius: 20%;
    }

    name {
        display: block;
        color: white;
        font-size: 12px;
    }

    .white{
        background-color: aliceblue;
    }

</style>
`

	return page + head
}

func appendBody(page string, path fyne.URI, MyApp *MyApp) string {
	var row []Row

	depth := countDepth(path, MyApp)

	file, _ := storage.LoadResourceFromURI(path)

	json.Unmarshal(file.Content(), &row)

	page = page + `<body>
`

	for _, v := range row {
		if v.Mode == "Label" {
			page = appendLabel(page, v)
		}
		if v.Mode == "Website" {
			page = appendWebsite(page, v, depth)
		}
	}

	page = page + `</body>`
	return page
}

func appendLabel(page string, row Row) string {
	if row.Sublink {
		page = page + `<a href="` + row.Name + `/Page.html">` + row.Name + `</a><br>`
		return page
	}

	if row.Link != "" {
		page = page + `<a href="` + row.Link + `">` + row.Name + `</a><br>`
		return page
	}

	// If just standard label
	page = page + row.Name + `<br>`
	return page
}

func appendWebsite(page string, row Row, depth int) string {
	page = page + `<ul>`

	for _, v := range row.Websites {
		var dots string
		var link string

		for i := 0; i < depth; i++ {
			dots = dots + "../"
		}

		if v.Subsite {
			link = v.Name + `/Page.html`
		} else {
			link = v.Link
		}
		page = page + `<li>`

		page = page + `<a href="` + link + `"><img src="` + dots + v.IconLocation + `" `

		if v.WhiteBg {
			page = page + `class="white"`
		}

		page = page + `/></a><Name>` + v.Name + `</Name>`

		page = page + `</li>`
	}

	page = page + `</ul>`
	return page
}

func countDepth(path fyne.URI, MyApp *MyApp) int {
	pathWithoutRoot := strings.Replace(path.Path(), MyApp.App.Storage().RootURI().Path(), "", -1)
	depth := strings.Count(pathWithoutRoot, "/")

	return depth - 1
}

func RegenerateHTML(MyApp *MyApp) {
	for _, v := range MyApp.Pages {
		var page string

		path, _ := storage.Child(v, "Rows.json")

		page = `<!DOCTYPE html>
	<html lang="en">
	`
		page = appendHead(page, MyApp)
		page = appendBody(page, path, MyApp)

		path, _ = storage.Child(v, "Page.html")

		file, _ := storage.Writer(path)

		file.Write([]byte(page))
	}
}
