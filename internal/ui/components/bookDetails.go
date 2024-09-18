package components

import (
	"AletheiaDesktop/internal/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateBookDetails(book models.Book, shouldWrap bool) *fyne.Container {
	var bookDetailsString string
	var bookDetailsContainer *fyne.Container
	var bookDetailsLabel *widget.Label

	bookDetailsString = fmt.Sprintf(
		"Title: %s\n"+
			"Author: %s\n"+
			"Filetype: %s\n"+
			"Filesize: %s\n"+
			"Language: %s\n"+
			"Pages: %s\n"+
			"Publisher: %s",
		book.Title,
		book.Author,
		book.Extension,
		book.Size,
		book.Language,
		book.Pages,
		book.Publisher,
	)

	bookDetailsLabel = widget.NewLabel(bookDetailsString)

	if shouldWrap {
		bookDetailsLabel.Wrapping = fyne.TextWrapWord
	}

	bookDetailsContainer = container.NewVBox(
		bookDetailsLabel,
	)

	return bookDetailsContainer
}
