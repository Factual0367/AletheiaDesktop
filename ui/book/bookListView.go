package book

import (
	"AletheiaDesktop/search"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gen2brain/beeep"
	"log"
)

func CreateBookListContainer(book *search.Book) *fyne.Container {
	bookDetailsString := fmt.Sprintf(
		"%s by %s",
		book.Title, book.Author)

	bookDetailsLabelContainer := container.NewVBox()
	bookDetailsLabel := widget.NewLabelWithStyle(bookDetailsString, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	bookDetailsLabel.Wrapping = fyne.TextWrapWord
	bookDetailsLabelContainer.Add(bookDetailsLabel)

	var downloadButton *widget.Button

	downloadButton = widget.NewButtonWithIcon("", theme.DownloadIcon(), func() {
		go func() {
			success := book.Download()
			if success {
				ok := beeep.Notify(book.Title, "Downloaded successfully", "")
				if ok != nil {
					log.Println("Could not send notification.")
				}
				downloadButton = widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {})

			} else {
				ok := beeep.Notify(book.Title, "Download failed", "")
				if ok != nil {
					log.Println("Could not send notification.")
				}
				log.Println(fmt.Sprintf("Download failed: %s"))
				downloadButton.SetIcon(theme.ErrorIcon())
			}
		}()
	})

	moreInformationButton := widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
		// placeholder
	})

	buttonContainer := container.NewHBox(
		moreInformationButton,
		downloadButton,
		layout.NewSpacer(),
	)

	bookDetailsLabelContainer.Add(buttonContainer)

	return bookDetailsLabelContainer
}
