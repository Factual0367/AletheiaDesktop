package views

import (
	"AletheiaDesktop/src/models"
	"AletheiaDesktop/src/ui/components"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateBookDownloadsContainer(book *models.Book) *fyne.Container {

	bookDetailsString := fmt.Sprintf(
		"Title: %s\nAuthor: %s",
		book.Title, book.Author,
	)

	bookDetailsLabel := widget.NewLabelWithStyle(bookDetailsString, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	bookDetailsLabel.Wrapping = fyne.TextWrapWord

	progressBar := widget.NewProgressBar()
	progressBar.SetValue(book.DownloadProgress)

	border := components.CreateBorderBox()

	borderedContainer := container.NewStack(border, container.NewVBox(bookDetailsLabel, progressBar))

	return container.NewVBox(borderedContainer)
}
