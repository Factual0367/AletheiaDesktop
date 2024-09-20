package views

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/internal/ui/components"
	"AletheiaDesktop/pkg/util/shared"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createButton(label string, icon fyne.Resource, tapped func()) *widget.Button {
	return widget.NewButtonWithIcon(label, icon, tapped)
}

func handleOpenBook(myApp fyne.App, filepath string) {
	go func() {
		err := shared.OpenWithDefaultApp(filepath)
		if err != nil {
			log.Println("Could not open book with default application.")
			shared.SendNotification(myApp, "Error", "Aletheia cannot find an application that can open your book.")
		}
	}()
}

func showDeleteConfirmationDialog(appWindow fyne.Window, book models.Book, myApp fyne.App, tabs *container.AppTabs) {
	confirmDialog := dialog.NewConfirm("Are you sure?", fmt.Sprintf("Do you want to delete %s?", book.Title), func(b bool) {
		if b {
			shared.DeleteBook(book)
			RefreshLibraryTab(myApp, appWindow, tabs)
		}
	}, appWindow)
	confirmDialog.Show()
}

func CreateBookLibraryContainer(myApp fyne.App, book models.Book, appWindow fyne.Window, tabs *container.AppTabs) *fyne.Container {
	bookDetailsContainer := components.CreateBookDetails(book, true)
	bookCover := components.CreateBookCover(book)
	border := components.CreateBorderBox()

	openButton := createButton("Open", theme.FileIcon(), func() {
		handleOpenBook(myApp, book.Filepath)
	})

	convertButton := createButton("Convert", theme.ContentRedoIcon(), func() {
		ShowConversionPopup(myApp, appWindow, book, tabs)
	})

	deleteButton := createButton("Delete", theme.DeleteIcon(), func() {
		showDeleteConfirmationDialog(appWindow, book, myApp, tabs)
	})

	openLibraryFolderButton := createButton("Open Location", theme.FolderOpenIcon(), shared.OpenLibraryFolder)

	emailBookButton := components.CreateEmailButton(myApp, book)

	buttonContainer := container.NewHBox(openButton, openLibraryFolderButton, emailBookButton, convertButton, deleteButton, layout.NewSpacer())

	borderedContainer := container.NewStack(border, container.NewVBox(bookDetailsContainer, buttonContainer))
	borderedContainerWithCover := container.NewHSplit(bookCover, borderedContainer)
	borderedContainerWithCover.SetOffset(0.10)

	return container.NewVBox(borderedContainerWithCover)
}
