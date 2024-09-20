package views

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/internal/ui/components"
	"AletheiaDesktop/pkg/util/email"
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

func CreateBookLibraryContainer(myApp fyne.App, book models.Book, appWindow fyne.Window, tabs *container.AppTabs) *fyne.Container {
	bookDetailsContainer := components.CreateBookDetails(book, true)

	openButton := widget.NewButtonWithIcon("Open", theme.FileIcon(), func() {
		go func() {
			err := shared.OpenWithDefaultApp(book.Filepath)
			if err != nil {
				log.Println("Could not open book with default application.")
				shared.SendNotification(myApp, "Error", "Aletheia cannot find an application that can open your book.")
			}
		}()
	})

	convertButton := widget.NewButtonWithIcon("Convert", theme.ContentRedoIcon(), func() {
		ShowConversionPopup(myApp, appWindow, book, tabs)
	})

	deleteButton := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
		confirmDialog := dialog.NewConfirm("Are you sure?", fmt.Sprintf("Do you want to delete %s?", book.Title), func(b bool) {
			if b {
				shared.DeleteBook(book)
				RefreshLibraryTab(myApp, appWindow, tabs)
			}
		}, appWindow)
		confirmDialog.Show()
	})

	openLibraryFolderButton := widget.NewButtonWithIcon("Open Location", theme.FolderOpenIcon(), func() {
		shared.OpenLibraryFolder()
	})

	var emailBookButton *widget.Button
	emailBookButton = widget.NewButtonWithIcon("Email", theme.MailSendIcon(), func() {
		go func() {
			emailSent := email.SendBookEmail(book)
			if emailSent {
				shared.SendNotification(myApp, "Success", "Your book is emailed successfully.")
				emailBookButton.SetIcon(theme.ConfirmIcon())
			} else {
				shared.SendNotification(myApp, "Failed", "Your book could not be emailed. Check your credentials.")
				emailBookButton.SetIcon(theme.ErrorIcon())
			}
		}()
	})

	buttonContainer := container.NewHBox(openButton, openLibraryFolderButton, emailBookButton, convertButton, deleteButton, layout.NewSpacer())

	border := components.CreateBorderBox()

	bookCover := components.CreateBookCover(book)

	borderedContainer := container.NewStack(border, container.NewVBox(bookDetailsContainer, buttonContainer))
	borderedContainerWithCover := container.NewHSplit(bookCover, borderedContainer)
	borderedContainerWithCover.SetOffset(0.10)

	return container.NewVBox(borderedContainerWithCover)
}
