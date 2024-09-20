package components

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/pkg/util/database"
	"AletheiaDesktop/pkg/util/downloads"
	"AletheiaDesktop/pkg/util/shared"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateDownloadButton(myApp fyne.App, book models.Book) *widget.Button {
	var downloadButton *widget.Button
	downloadButton = widget.NewButtonWithIcon("Download", theme.DownloadIcon(), func() {
		downloadButton.Text = "Downloading"
		go func() {
			if !downloads.AddInProgressDownloads(&book) {
				shared.SendNotification(myApp, book.Title, "Downloading")
				success := book.Download()
				if success {
					shared.SendNotification(myApp, book.Title, "Downloaded successfully")
					downloadButton = widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {})
					database.UpdateDatabase(book, true, "downloaded") // true to add a book, false to remove
				} else {
					shared.SendNotification(myApp, book.Title, "Download failed. Is Libgen down?")
					log.Println(fmt.Sprintf("Download failed for book %s", book.Title))
				}
			}
		}()
	})

	return downloadButton
}
