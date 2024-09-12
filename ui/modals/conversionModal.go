package modals

import (
	"AletheiaDesktop/search"
	"AletheiaDesktop/util/conversion"
	"AletheiaDesktop/util/shared"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func conversionPopup(appWindow fyne.Window, book search.Book, modal *widget.PopUp) *widget.PopUp {

	var targetFormat string
	conversionContainer := container.NewVBox()
	conversionLabel := widget.NewLabel("Which format do you want to convert to?")
	conversionSelector := widget.NewSelect([]string{"EPUB", "PDF", "KEPUB", "DJVU"}, func(s string) {
		targetFormat = s
		fmt.Println(s)
	})

	convertButton := widget.NewButtonWithIcon("Convert", theme.ContentRedoIcon(), func() {
		converted := conversion.ConvertToFormat(targetFormat, book)
		if !converted {
			shared.SendNotification("Error", "Cannot convert book, do you have Calibre installed?")
		} else {
			shared.SendNotification("Success", "Your book is converted successfully.")
		}
	})

	conversionContainer.Add(conversionLabel)
	conversionContainer.Add(conversionSelector)
	conversionContainer.Add(convertButton)
	conversionContainer.Add(widget.NewButton("Close", func() { modal.Hide() }))
	modal = widget.NewModalPopUp(
		conversionContainer,
		appWindow.Canvas(),
	)
	return modal
}

func installCalibrePopup(appWindow fyne.Window, modal *widget.PopUp) *widget.PopUp {
	installCalibreContainer := container.NewVBox()
	directions := widget.NewLabel(
		"For this feature to work you need to have Calibre installed on your system.")
	websiteDirections := widget.NewLabel(
		"Please visit: https://calibre-ebook.com/download and install it.")
	installCalibreContainer.Add(directions)
	installCalibreContainer.Add(websiteDirections)
	installCalibreContainer.Add(widget.NewButton("Close", func() { modal.Hide() }))
	modal = widget.NewModalPopUp(
		installCalibreContainer,
		appWindow.Canvas(),
	)
	return modal
}

func ShowConversionPopup(appWindow fyne.Window, book search.Book) *widget.PopUp {
	var modal *widget.PopUp

	calibreExists := conversion.CheckCalibreInstalled()
	fmt.Println(calibreExists)

	if calibreExists {
		modal = conversionPopup(appWindow, book, modal)
		modal.Show()
	} else {
		modal = installCalibrePopup(appWindow, modal)
		modal.Show()
	}

	return modal
}
