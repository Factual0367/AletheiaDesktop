package views

import (
	"AletheiaDesktop/src/search"
	"AletheiaDesktop/src/util/cache"
	"AletheiaDesktop/src/util/database"
	"AletheiaDesktop/src/util/shared"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func createBookDetailsTopView(book search.Book) *fyne.Container {
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
	topView := container.NewPadded(container.NewCenter(bookCover))

	return topView
}

func createDownloadButton(book search.Book) *widget.Button {
	var downloadButton *widget.Button

	downloadButton = widget.NewButtonWithIcon("", theme.DownloadIcon(), func() {
		go func() {
			shared.SendNotification(book.Title, "Downloading")
			success := book.Download()
			if success {
				shared.SendNotification(book.Title, "Downloaded successfully")
				downloadButton = widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {})
				database.UpdateDatabase(book, true, "downloaded") // true to add, false to remove from database
			} else {
				shared.SendNotification(book.Title, "Download failed")
				log.Println(fmt.Sprintf("Download failed: %s"))
				downloadButton.SetIcon(theme.ErrorIcon())
			}
		}()
	})
	return downloadButton
}

func createBookDetailsBottomView(book search.Book) *fyne.Container {
	var bookDetailsString string
	var bottomView *fyne.Container
	var bookDetailsLabel *widget.Label

	downloadButton := createDownloadButton(book)

	if book.Title == "" {
		// this is the default view
		// user not having selected a book
		bookDetailsString = fmt.Sprintf(
			"Select a book to view details",
		)
		bookDetailsLabel = widget.NewLabel(bookDetailsString)
		bottomView = container.NewVBox(
			bookDetailsLabel,
		)
		return bottomView
	}

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
	bookDetailsLabel.Wrapping = fyne.TextWrapWord

	bottomView = container.NewVBox(
		bookDetailsLabel,
		downloadButton,
	)

	return bottomView
}

func CreateBookDetailsView(book search.Book, isDefaultBook bool) *fyne.Container {
	var bottomView *fyne.Container

	if isDefaultBook {
		// no details to show
		bottomView = container.NewVBox(
			widget.NewLabel("Select a book to view details"))
	} else {
		bottomView = createBookDetailsBottomView(book)

	}

	topView := createBookDetailsTopView(book)
	detailsSplit := container.NewVSplit(topView, bottomView)
	detailsSplit.SetOffset(0.25)
	detailsSplitContainer := container.NewStack(detailsSplit)
	return detailsSplitContainer
}
