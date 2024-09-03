package logic

import "fyne.io/fyne/v2"

type Website struct {
	Name     string   // Name for website button label
	URL      string   // Link to the website
	Icon     fyne.URI // Path to icon/favicon
	Selected bool     // Show or not buttons in border to move/re-order button in grid
}
