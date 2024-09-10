package book

import (
	"AletheiaDesktop/search"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gen2brain/beeep"
	"log"
)

func CreateBookDetailsView(book search.Book) *container.Split {
	coverImageSize := fyne.NewSize(120, 200)
	var bookCover *canvas.Image

	var uri fyne.URI
	if book.CoverLink != "Unknown" {
		uri, _ = storage.ParseURI(book.CoverLink)
	} else {
		uri, _ = storage.ParseURI("https://cdn.pixabay.com/photo/2013/07/13/13/34/book-161117_960_720.png")
	}

	bookCover = canvas.NewImageFromURI(uri)
	bookCover.FillMode = canvas.ImageFillContain
	bookCover.SetMinSize(coverImageSize)

	centeredBookCover := container.NewPadded(container.NewCenter(bookCover))

	var bookDetailsString string
	var bottomSplit *fyne.Container
	var bookDetailsLabel *widget.Label
	var downloadButton *widget.Button

	downloadButton = widget.NewButtonWithIcon("", theme.DownloadIcon(), func() {
		go func() {
			success := book.Download()
			if success {
				ok := beeep.Notify(book.Title, "Downloaded successfully", "")
				if ok != nil {
					log.Println("Could not send notification.")
				}
				downloadButton = widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {})

			} else {
				ok := beeep.Notify(book.Title, "Download failed", "")
				if ok != nil {
					log.Println("Could not send notification.")
				}
				log.Println(fmt.Sprintf("Download failed: %s"))
				downloadButton.SetIcon(theme.ErrorIcon())
			}
		}()
	})

	if len(book.Author) > 0 {
		bookDetailsString = fmt.Sprintf(
			"Title: %s\nAuthor: %s\nFiletype: %s\nFilesize: %s\nLanguage: %s\nPages: %s\nPublisher: %s",
			book.Title, book.Author, book.Extension, book.Size, book.Language, book.Pages, book.Publisher)
		bookDetailsLabel = widget.NewLabel(bookDetailsString)
		bottomSplit = container.NewVBox(
			bookDetailsLabel,
			downloadButton,
		)
	} else {
		bookDetailsString = fmt.Sprintf(
			"Select a book to view details",
		)
		bookDetailsLabel = widget.NewLabel(bookDetailsString)
		bottomSplit = container.NewVBox(
			bookDetailsLabel,
		)
	}

	detailsSplit := container.NewVSplit(centeredBookCover, bottomSplit)
	detailsSplit.SetOffset(0.25)

	return detailsSplit
}
