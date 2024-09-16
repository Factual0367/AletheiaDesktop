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

func CreateBookListContainer(book search.Book, DetailsContainer *fyne.Container) *fyne.Container {
	bookDetailsString := fmt.Sprintf(
		"%s\n%s\nFiletype: %s\nFilesize: %s",
		book.Title, book.Author, book.Extension, book.Size)

	bookDetailsLabelContainer := container.NewVBox()
	bookDetailsLabel := widget.NewLabelWithStyle(bookDetailsString, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	bookDetailsLabel.Wrapping = fyne.TextWrapWord
	bookDetailsLabelContainer.Add(bookDetailsLabel)

	downloadButton := components.CreateDownloadButton(book)

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
