package book

import (
	"AletheiaDesktop/search"
	"AletheiaDesktop/util/database"
	"AletheiaDesktop/util/shared"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

func CreateBookLibraryContainer(book search.Book) *fyne.Container {
	bookDetailsString := fmt.Sprintf(
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

	bookDetailsLabelContainer := container.NewVBox()
	bookDetailsLabel := widget.NewLabelWithStyle(bookDetailsString, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	bookDetailsLabel.Wrapping = fyne.TextWrapWord
	bookDetailsLabelContainer.Add(bookDetailsLabel)

	var downloadButton *widget.Button

	downloadButton = widget.NewButtonWithIcon("", theme.DownloadIcon(), func() {
		go func() {
			success := book.Download()
			if success {
				shared.SendNotification(book.Title, "Downloaded successfully")
				downloadButton = widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {})
				database.UpdateDatabase(book)
			} else {
				shared.SendNotification(book.Title, "Download failed")
				log.Println(fmt.Sprintf("Download failed: %s"))
				downloadButton.SetIcon(theme.ErrorIcon())
			}
		}()
	})

	convertButton := widget.NewButtonWithIcon("", theme.ContentRedoIcon(), func() {})

	buttonContainer := container.NewHBox(
		convertButton,
		downloadButton,
		layout.NewSpacer(),
	)

	// add some boxing
	border := canvas.NewRectangle(&color.NRGBA{R: 97, G: 97, B: 97, A: 50})
	border.StrokeColor = color.NRGBA{R: 97, G: 97, B: 97, A: 50}
	border.StrokeWidth = 2
	border.CornerRadius = 10

	bookDetailsLabelContainer.Add(buttonContainer)
	borderedContainer := container.NewStack(border, bookDetailsLabelContainer)

	return borderedContainer
}
