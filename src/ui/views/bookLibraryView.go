package views

import (
	"AletheiaDesktop/src/models"
	"AletheiaDesktop/src/ui/components"
	"AletheiaDesktop/src/util/email"
	"AletheiaDesktop/src/util/shared"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func CreateBookLibraryContainer(book models.Book, appWindow fyne.Window, tabs *container.AppTabs) *fyne.Container {
	bookDetailsString := fmt.Sprintf(
		"Title: %s\nAuthor: %s\nFiletype: %s\nFilesize: %s\nLanguage: %s\nPages: %s\nPublisher: %s",
		book.Title, book.Author, book.Extension, book.Size, book.Language, book.Pages, book.Publisher,
	)

	bookDetailsLabel := widget.NewLabelWithStyle(bookDetailsString, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	bookDetailsLabel.Wrapping = fyne.TextWrapWord

	openButton := widget.NewButtonWithIcon("", theme.FileIcon(), func() {
		go func() {
			err := shared.OpenWithDefaultApp(book.Filepath)
			if err != nil {
				log.Println("Could not open book with default application.")
				shared.SendNotification("Error", "Aletheia cannot find an application that can open your book.")
			}
		}()
	})

	convertButton := widget.NewButtonWithIcon("", theme.ContentRedoIcon(), func() {
		ShowConversionPopup(appWindow, book, tabs)
	})

	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		confirmDialog := dialog.NewConfirm("Are you sure?", fmt.Sprintf("Do you want to delete %s?", book.Title), func(b bool) {
			if b {
				shared.DeleteBook(book)
				RefreshLibraryTab(appWindow, tabs)
			}
		}, appWindow)
		confirmDialog.Show()
	})

	openLibraryFolderButton := widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
		shared.OpenLibraryFolder()
	})

	var emailBookButton *widget.Button
	emailBookButton = widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {
		go func() {
			emailSent := email.SendBookEmail(book)
			if emailSent {
				shared.SendNotification("Success", "Your book is emailed successfully.")
				emailBookButton.SetIcon(theme.ConfirmIcon())
			} else {
				shared.SendNotification("Failed", "Your book could not be emailed. Check your credentials.")
				emailBookButton.SetIcon(theme.ErrorIcon())
			}
		}()
	})

	buttonContainer := container.NewHBox(openButton, openLibraryFolderButton, emailBookButton, convertButton, deleteButton, layout.NewSpacer())

	border := components.CreateBorderBox()

	bookCover := createBookDetailsTopView(book)

	borderedContainer := container.NewStack(border, container.NewVBox(bookDetailsLabel, buttonContainer))
	borderedContainerWithCover := container.NewHSplit(bookCover, borderedContainer)
	borderedContainerWithCover.SetOffset(0.10)

	return container.NewVBox(borderedContainerWithCover)
}
