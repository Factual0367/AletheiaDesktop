package views

import (
	"AletheiaDesktop/src/search"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
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
	border := canvas.NewRectangle(&color.NRGBA{R: 97, G: 97, B: 97, A: 50})
	border.StrokeColor = color.NRGBA{R: 97, G: 97, B: 97, A: 50}
	border.StrokeWidth = 2
	border.CornerRadius = 10

	borderedContainer := container.NewStack(border, container.NewVBox(bookDetailsLabel, progressBar))
	borderedContainerWithCover := container.NewVBox(borderedContainer)
	//borderedContainerWithCover.SetOffset(0.10)

	return container.NewVBox(borderedContainerWithCover)
}
