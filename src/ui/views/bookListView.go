package views

import (
	"AletheiaDesktop/src/models"
	"AletheiaDesktop/src/ui/components"
	"AletheiaDesktop/src/util/database"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateBookListContainer(book models.Book, appWindow fyne.Window) *fyne.Container {
	bookDetailsString := fmt.Sprintf(
		"%s\n%s\nFiletype: %s\nFilesize: %s",
		book.Title, book.Author, book.Extension, book.Size)

	bookDetailsLabelContainer := container.NewVBox()
	bookDetailsLabel := widget.NewLabelWithStyle(bookDetailsString, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	bookDetailsLabel.Wrapping = fyne.TextWrapWord
	bookDetailsLabelContainer.Add(bookDetailsLabel)

	downloadButton := components.CreateDownloadButton(book)

	moreInformationButton := widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
		// DetailsContainer.Objects = nil
		// new content for the selected book
		// newDetailsView := CreateBookDetailsView(book, false)
		// DetailsContainer.Add(newDetailsView)
		// DetailsContainer.Refresh()
		bookDetailsPopup := components.BookDetailsPopup(appWindow, book)
		bookDetailsPopup.Show()
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
	bookDetailsLabelContainer.Add(buttonContainer)

	// add some boxing
	border := components.CreateBorderBox()
	borderedContainer := container.NewStack(border, bookDetailsLabelContainer)

	return borderedContainer
}
