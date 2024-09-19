package views

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/internal/ui/components"
	"AletheiaDesktop/pkg/util/database"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateBookBookmarksContainer(book models.Book, appWindow fyne.Window, tabs *container.AppTabs) *fyne.Container {
	bookDetailsContainer := components.CreateBookDetails(book, true)

	unfavoriteButton := widget.NewButtonWithIcon("Unfavorite", theme.ContentRemoveIcon(), func() {
		database.UpdateDatabase(book, false, "favorited")
		refreshBookmarksTab(appWindow, tabs)
	})
	downloadButton := components.CreateDownloadButton(book)
	buttonContainer := container.NewHBox(unfavoriteButton, downloadButton, layout.NewSpacer())

	border := components.CreateBorderBox()

	bookCover := components.CreateBookCover(book)

	borderedContainer := container.NewStack(border, container.NewVBox(bookDetailsContainer, buttonContainer))
	borderedContainerWithCover := container.NewHSplit(bookCover, borderedContainer)
	borderedContainerWithCover.SetOffset(0.10)

	return container.NewVBox(borderedContainerWithCover)
}
