package views

import (
	"AletheiaDesktop/src/search"
	conversion2 "AletheiaDesktop/src/util/conversion"
	database2 "AletheiaDesktop/src/util/database"
	shared2 "AletheiaDesktop/src/util/shared"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"strings"
	"time"
)

func conversionPopup(appWindow fyne.Window, book search.Book, modal *widget.PopUp, tabs *container.AppTabs) *widget.PopUp {

	var targetFormat string
	conversionContainer := container.NewVBox()
	conversionLabel := widget.NewLabel("Which format do you want to convert to?")
	conversionSelector := widget.NewSelect([]string{"EPUB", "PDF", "MOBI", "DJVU"}, func(s string) {
		targetFormat = s
		fmt.Println(s)
	})

	convertButton := widget.NewButtonWithIcon("Convert", theme.ContentRedoIcon(), func() {
		go func() {
			converted := conversion2.ConvertToFormat(targetFormat, book)
			if !converted {
				shared2.SendNotification("Error", "Cannot convert book. Did you select the right format?")
			} else {
				shared2.SendNotification("Success", "Your book is converted successfully.")
				RefreshLibraryTab(appWindow, tabs)

			}
		}()
		modal.Hide()

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

func ShowConversionPopup(appWindow fyne.Window, book search.Book, tabs *container.AppTabs) *widget.PopUp {
	var modal *widget.PopUp

	calibreExists := conversion2.CheckCalibreInstalled()
	fmt.Println(calibreExists)

	if calibreExists {
		modal = conversionPopup(appWindow, book, modal, tabs)
		modal.Show()
	} else {
		modal = installCalibrePopup(appWindow, modal)
		modal.Show()
	}

	return modal
}

func loadSavedBooks() (map[string]*search.Book, error) {
	userData, err := database2.ReadDatabaseFile()

	if len(userData) == 0 {
		fmt.Println(len(userData))
		userData, err = database2.ReadDatabaseFile()
	}

	if err != nil {
		return nil, err
	}

	if savedBooks, ok := userData["savedBooks"].(map[string]*search.Book); ok {
		return savedBooks, nil
	}
	return nil, nil
}

func updateLibraryGrid(grid *fyne.Container, books map[string]*search.Book, filter string, appWindow fyne.Window, tabs *container.AppTabs) {
	grid.Objects = nil

	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(filter)) {
			bookFileExists, err := shared2.Exists(book.Filepath)
			if err != nil {
				log.Printf("Book does not exist, removing book from database. Title: %s", book.Title)
				database2.UpdateDatabase(*book, false, "downloaded")
			}
			if bookFileExists {
				bookLibraryContainer := CreateBookLibraryContainer(*book, appWindow, tabs)
				grid.Add(bookLibraryContainer)
			}

		}
	}
	grid.Refresh()
}

func RefreshLibraryTab(appWindow fyne.Window, tabs *container.AppTabs) {

	newLibraryView := CreateLibraryView(appWindow, tabs)
	tabs.Items[1] = newLibraryView
	tabs.SelectIndex(1)
	tabs.Refresh()
}

func CreateLibraryView(appWindow fyne.Window, tabs *container.AppTabs) *container.TabItem {
	filterInput := widget.NewEntry()
	filterInput.PlaceHolder = "Filter"
	filterInput.Resize(fyne.NewSize(800, filterInput.MinSize().Height)) // Set the desired width

	topWidgets := container.NewWithoutLayout(filterInput)

	libraryViewGrid := container.NewVBox()
	libraryViewGridScrollable := container.NewVScroll(libraryViewGrid)

	savedBooks, err := loadSavedBooks()
	if err != nil {
		log.Printf("Could not read savedBooks %s", err)
	}

	if savedBooks != nil {
		updateLibraryGrid(libraryViewGrid, savedBooks, "", appWindow, tabs)
	}

	var typingTimer *time.Timer

	filterInput.OnChanged = func(filter string) {
		if typingTimer != nil {
			typingTimer.Stop() // stop timer
		}

		typingTimer = time.AfterFunc(500*time.Millisecond, func() { // 500ms delay so filtering does not get laggy
			if savedBooks != nil {
				updateLibraryGrid(libraryViewGrid, savedBooks, filter, appWindow, tabs)
			}
		})
	}

	libraryViewLayout := container.NewBorder(topWidgets, nil, nil, nil, libraryViewGridScrollable)

	return container.NewTabItemWithIcon("Library", theme.StorageIcon(), libraryViewLayout)
}
