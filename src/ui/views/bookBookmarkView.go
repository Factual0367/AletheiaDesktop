package views

import (
	"AletheiaDesktop/src/search"
	"AletheiaDesktop/src/ui/components"
	"AletheiaDesktop/src/util/database"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateBookBookmarksContainer(book search.Book, appWindow fyne.Window, tabs *container.AppTabs) *fyne.Container {
	bookDetailsString := fmt.Sprintf(
		"Title: %s\nAuthor: %s\nFiletype: %s\nFilesize: %s\nLanguage: %s\nPages: %s\nPublisher: %s",
		book.Title, book.Author, book.Extension, book.Size, book.Language, book.Pages, book.Publisher,
	)

	bookDetailsLabel := widget.NewLabelWithStyle(bookDetailsString, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	bookDetailsLabel.Wrapping = fyne.TextWrapWord

	unfavoriteButton := widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), func() {
		database.UpdateDatabase(book, false, "favorited")
		refreshBookmarksTab(appWindow, tabs)
	})

	downloadButton := components.CreateDownloadButton(book)

	buttonContainer := container.NewHBox(unfavoriteButton, downloadButton, layout.NewSpacer())

	border := components.CreateBorderBox()

	bookCover := createBookDetailsTopView(book)

	borderedContainer := container.NewStack(border, container.NewVBox(bookDetailsLabel, buttonContainer))
	borderedContainerWithCover := container.NewHSplit(bookCover, borderedContainer)
	borderedContainerWithCover.SetOffset(0.10)

	return container.NewVBox(borderedContainerWithCover)
}
