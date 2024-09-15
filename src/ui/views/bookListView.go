package views

import (
	"AletheiaDesktop/src/search"
	"AletheiaDesktop/src/util/database"
	"AletheiaDesktop/src/util/shared"
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

func CreateBookListContainer(book search.Book, DetailsContainer *fyne.Container) *fyne.Container {
	bookDetailsString := fmt.Sprintf(
		"%s\n%s\nFiletype: %s\nFilesize: %s",
		book.Title, book.Author, book.Extension, book.Size)

	bookDetailsLabelContainer := container.NewVBox()
	bookDetailsLabel := widget.NewLabelWithStyle(bookDetailsString, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	bookDetailsLabel.Wrapping = fyne.TextWrapWord
	bookDetailsLabelContainer.Add(bookDetailsLabel)

	var downloadButton *widget.Button

	downloadButton = widget.NewButtonWithIcon("", theme.DownloadIcon(), func() {
		go func() {
			shared.SendNotification(book.Title, "Downloading")
			success := book.Download()
			if success {
				shared.SendNotification(book.Title, "Downloaded successfully")
				downloadButton = widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {})
				database.UpdateDatabase(book, true, "downloaded") // true to add a book, false to remove
				downloadButton.SetIcon(theme.ConfirmIcon())
			} else {
				shared.SendNotification(book.Title, "Download failed")
				log.Println(fmt.Sprintf("Download failed: %s"))
				downloadButton.SetIcon(theme.ErrorIcon())
			}
		}()
	})

	moreInformationButton := widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
		DetailsContainer.Objects = nil
		// new content for the selected book
		newDetailsView := CreateBookDetailsView(book, false)
		DetailsContainer.Add(newDetailsView)

		DetailsContainer.Refresh()
	})

	var favoriteButton *widget.Button
	favoriteButton = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		database.UpdateDatabase(book, true, "favorited")
		favoriteButton.SetIcon(theme.ContentRemoveIcon())
	})

	buttonContainer := container.NewHBox(
		moreInformationButton,
		favoriteButton,
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
