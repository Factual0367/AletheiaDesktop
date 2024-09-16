package views

import (
	"AletheiaDesktop/src/search"
	"AletheiaDesktop/src/ui/components"
	"AletheiaDesktop/src/util/conversion"
	"AletheiaDesktop/src/util/database"
	"AletheiaDesktop/src/util/shared"
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
	conversionSelector := widget.NewSelect([]string{"EPUB", "PDF", "MOBI"}, func(s string) {
		targetFormat = s
	})

	convertButton := widget.NewButtonWithIcon("Convert", theme.ContentRedoIcon(), func() {
		go func() {
			converted := conversion.ConvertToFormat(targetFormat, book)
			if !converted {
				shared.SendNotification("Error", "Cannot convert book. Did you select the right format?")
			} else {
				shared.SendNotification("Success", "Your book is converted successfully.")
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

	calibreExists := conversion.CheckCalibreInstalled()

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
	userData, err := database.ReadDatabaseFile()

	if err != nil {
		return nil, err
	}

	if savedBooks, ok := userData["savedBooks"].(map[string]*search.Book); ok {
		return savedBooks, nil
	}
	return nil, nil
}

func filterBooks(books map[string]*search.Book, filter string) []*search.Book {
	filteredBooks := []*search.Book{}
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(filter)) {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return filteredBooks
}

func updateLibraryGrid(grid *fyne.Container, books map[string]*search.Book, filter string, appWindow fyne.Window, tabs *container.AppTabs) {
	grid.RemoveAll()

	filteredBooks := filterBooks(books, filter)

	for _, book := range filteredBooks {
		bookFileExists, err := shared.Exists(book.Filepath)
		if err != nil {
			log.Printf("Book does not exist, removing book from database. Title: %s", book.Title)
			database.UpdateDatabase(*book, false, "downloaded")
			continue
		}
		if bookFileExists {
			bookLibraryContainer := CreateBookLibraryContainer(*book, appWindow, tabs)
			grid.Add(bookLibraryContainer)
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
	// load data to show
	savedBooks, err := loadSavedBooks()
	if err != nil {
		log.Printf("Could not read savedBooks %s", err)
	}

	libraryViewGrid := container.NewVBox()
	libraryViewGridScrollable := container.NewVScroll(libraryViewGrid)
	var typingTimer *time.Timer
	filterInput := components.CreateFilterInput()
	filterInput.OnChanged = func(filter string) {
		if typingTimer != nil {
			typingTimer.Stop() // stop timer
		}

		typingTimer = time.AfterFunc(300*time.Millisecond, func() { // 300ms delay so filtering does not get laggy
			if savedBooks != nil {
				updateLibraryGrid(libraryViewGrid, savedBooks, filter, appWindow, tabs)
			}
		})
	}
	topWidgets := container.NewWithoutLayout(filterInput)

	// add data to view
	if savedBooks != nil {
		updateLibraryGrid(libraryViewGrid, savedBooks, "", appWindow, tabs)
	}

	libraryViewLayout := container.NewBorder(topWidgets, nil, nil, nil, libraryViewGridScrollable)
	libraryView := container.NewTabItemWithIcon("Library", theme.StorageIcon(), libraryViewLayout)
	return libraryView
}
