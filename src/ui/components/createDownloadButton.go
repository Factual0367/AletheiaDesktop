package components

import (
	"AletheiaDesktop/src/models"
	"AletheiaDesktop/src/util/database"
	"AletheiaDesktop/src/util/downloads"
	"AletheiaDesktop/src/util/shared"
	"fmt"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func CreateDownloadButton(book models.Book) *widget.Button {
	var downloadButton *widget.Button

	downloadButton = widget.NewButtonWithIcon("", theme.DownloadIcon(), func() {
		go func() {
			if !downloads.AddInProgressDownloads(&book) {
				shared.SendNotification(book.Title, "Downloading")
				success := book.Download()
				if success {
					shared.SendNotification(book.Title, "Downloaded successfully")
					downloadButton = widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {})
					database.UpdateDatabase(book, true, "downloaded") // true to add a book, false to remove
					downloadButton.SetIcon(theme.ConfirmIcon())
				} else {
					shared.SendNotification(book.Title, "Download failed. Is Libgen down?")
					log.Println(fmt.Sprintf("Download failed: %s"))
					downloadButton.SetIcon(theme.ErrorIcon())
				}
			}
		}()
	})

	return downloadButton
}
