package views

import (
	"AletheiaDesktop/src/search"
	"AletheiaDesktop/src/ui/components"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateBookDownloadsContainer(book *search.Book) *fyne.Container {

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
	borderedContainerWithCover := container.NewVBox(borderedContainer)
	//borderedContainerWithCover.SetOffset(0.10)

	return container.NewVBox(borderedContainerWithCover)
}
