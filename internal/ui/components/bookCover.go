package components

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/pkg/util/cache"
	"AletheiaDesktop/pkg/util/shared"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

var temporaryBookCache = make(map[string]*canvas.Image)

func CreateBookCover(book models.Book) *fyne.Container {
	coverImageSize := fyne.NewSize(120, 200)
	var bookCover *canvas.Image

	var uri fyne.URI

	if bookCover, ok := temporaryBookCache[book.ID]; ok {
		log.Println("Image exists in cache, returning it.")
		bookCover = temporaryBookCache[book.ID]
		bookCover.FillMode = canvas.ImageFillContain
		bookCover.SetMinSize(coverImageSize)
		bookCoverContainer := container.NewPadded(container.NewCenter(bookCover))
		return bookCoverContainer
	}

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

	temporaryBookCache[book.ID] = bookCover
	bookCover.FillMode = canvas.ImageFillContain
	bookCover.SetMinSize(coverImageSize)
	bookCoverContainer := container.NewPadded(container.NewCenter(bookCover))

	return bookCoverContainer
}
