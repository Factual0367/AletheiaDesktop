package views

import (
	"AletheiaDesktop/src/search"
	"AletheiaDesktop/src/ui/components"
	"AletheiaDesktop/src/util/database"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func CreateBookBookmarksContainer(book search.Book, appWindow fyne.Window, tabs *container.AppTabs) *fyne.Container {
	bookDetailsString := fmt.Sprintf(
		"Title: %s\nAuthor: %s\nFiletype: %s\nFilesize: %s\nLanguage: %s\nPages: %s\nPublisher: %s",
		book.Title, book.Author, book.Extension, book.Size, book.Language, book.Pages, book.Publisher,
	)

	bookDetailsLabel := widget.NewLabelWithStyle(bookDetailsString, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	bookDetailsLabel.Wrapping = fyne.TextWrapWord

	unfavoriteButton := widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), func() {
		refreshBookmarksTab(appWindow, tabs)
		database.UpdateDatabase(book, false, "favorited")
	})

	downloadButton := components.CreateDownloadButton(book)

	buttonContainer := container.NewHBox(unfavoriteButton, downloadButton, layout.NewSpacer())

	border := canvas.NewRectangle(&color.NRGBA{R: 97, G: 97, B: 97, A: 50})
	border.StrokeColor = color.NRGBA{R: 97, G: 97, B: 97, A: 50}
	border.StrokeWidth = 2
	border.CornerRadius = 10

	bookCover := createBookDetailsTopView(book)

	borderedContainer := container.NewStack(border, container.NewVBox(bookDetailsLabel, buttonContainer))
	borderedContainerWithCover := container.NewHSplit(bookCover, borderedContainer)
	borderedContainerWithCover.SetOffset(0.10)

	return container.NewVBox(borderedContainerWithCover)
}
