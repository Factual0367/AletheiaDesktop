package views

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/internal/ui/components"
	conversion2 "AletheiaDesktop/pkg/util/conversion"
	"AletheiaDesktop/pkg/util/shared"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func BookDetailsPopup(appWindow fyne.Window, book models.Book) *widget.PopUp {
	bookCover := components.CreateBookCover(book)
	bookDetails := components.CreateBookDetails(book, false)
	var bookDetailsPopup *widget.PopUp

	bookDetails.MinSize()

	bookDetailsWithButtons := container.NewVBox(
		bookDetails,
		container.NewHBox(
			widget.NewButton("Close", func() {
				bookDetailsPopup.Hide()
			}),
			components.CreateDownloadButton(book),
		),
	)

	bookDetailsPopupContainer := container.NewHSplit(
		bookCover,
		bookDetailsWithButtons,
	)

	bookDetailsPopup = widget.NewModalPopUp(bookDetailsPopupContainer, appWindow.Canvas())
	return bookDetailsPopup
}

func ConversionPopup(appWindow fyne.Window, book models.Book, tabs *container.AppTabs) *widget.PopUp {
	var targetFormat string
	var modal *widget.PopUp

	conversionContainer := container.NewVBox(
		widget.NewLabel("Which format do you want to convert to?"),
		widget.NewSelect([]string{"EPUB", "PDF", "MOBI"}, func(s string) { targetFormat = s }),
		widget.NewButtonWithIcon("Convert", theme.ContentRedoIcon(), func() {
			go func() {
				modal.Hide()
				if conversion2.ConvertToFormat(targetFormat, book) {
					shared.SendNotification("Success", "Your book is converted successfully.")
					RefreshLibraryTab(appWindow, tabs)
				} else {
					shared.SendNotification("Error", "Cannot convert book. Did you select the right format?")
				}
			}()
		}),
		widget.NewButton("Close", func() {
			modal.Hide()
		}),
	)

	modal = widget.NewModalPopUp(conversionContainer, appWindow.Canvas())
	return modal
}

func InstallCalibrePopup(appWindow fyne.Window) *widget.PopUp {
	var modal *widget.PopUp
	installCalibreContainer := container.NewVBox(
		widget.NewLabel("For this feature to work, you need to have Calibre installed on your system."),
		widget.NewLabel("Please visit: https://calibre-ebook.com/download and install it."),
		widget.NewButton("Close", func() {
			modal.Hide()
		}),
	)

	modal = widget.NewModalPopUp(installCalibreContainer, appWindow.Canvas())
	return modal
}

func ShowConversionPopup(appWindow fyne.Window, book models.Book, tabs *container.AppTabs) {
	if conversion2.CheckCalibreInstalled() {
		ConversionPopup(appWindow, book, tabs).Show()
	} else {
		InstallCalibrePopup(appWindow).Show()
	}
}
