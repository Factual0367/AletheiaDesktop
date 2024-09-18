package views

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/internal/ui/components"
	"AletheiaDesktop/pkg/util/cache"
	"AletheiaDesktop/pkg/util/database"
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
		go func() {
			bookDetailsPopup := BookDetailsPopup(appWindow, book)
			bookDetailsPopup.Show()
		}()
	})

	var favoriteButton *widget.Button
	favoriteButton = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		database.UpdateDatabase(book, true, "favorited")
		// download covers before the user opens the
		// bookmarks view to prevent lag
		go func() {
			cache.SaveCoverImage(book.CoverLink, book.CoverPath)
		}()
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
