package components

import (
	"AletheiaDesktop/src/models"
	"AletheiaDesktop/src/util/cache"
	"AletheiaDesktop/src/util/shared"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

func CreateBookCover(book models.Book) *fyne.Container {
	coverImageSize := fyne.NewSize(120, 200)
	var bookCover *canvas.Image

	var uri fyne.URI
	if book.CoverLink != "Unknown" {
		uri, _ = storage.ParseURI(book.CoverLink)
	} else {
		uri, _ = storage.ParseURI("https://cdn.pixabay.com/photo/2013/07/13/13/34/book-161117_960_720.png")
	}

	cachedImageExists, err := shared.Exists(book.CoverPath)
	if err != nil {
		cache.SaveCoverImage(book.CoverLink, book.CoverPath)
	}
	if cachedImageExists {
		bookCover = canvas.NewImageFromFile(book.CoverPath)
	} else {
		bookCover = canvas.NewImageFromURI(uri)

	}

	bookCover.FillMode = canvas.ImageFillContain
	bookCover.SetMinSize(coverImageSize)
	bookCoverContainer := container.NewPadded(container.NewCenter(bookCover))

	return bookCoverContainer
}
