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

func conversionPopup(appWindow fyne.Window, book search.Book, tabs *container.AppTabs) *widget.PopUp {
	var targetFormat string
	var modal *widget.PopUp

	conversionContainer := container.NewVBox(
		widget.NewLabel("Which format do you want to convert to?"),
		widget.NewSelect([]string{"EPUB", "PDF", "MOBI"}, func(s string) { targetFormat = s }),
		widget.NewButtonWithIcon("Convert", theme.ContentRedoIcon(), func() {
			go func() {
				modal.Hide()
				if conversion.ConvertToFormat(targetFormat, book) {
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

func installCalibrePopup(appWindow fyne.Window) *widget.PopUp {
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

func ShowConversionPopup(appWindow fyne.Window, book search.Book, tabs *container.AppTabs) {
	if conversion.CheckCalibreInstalled() {
		conversionPopup(appWindow, book, tabs).Show()
	} else {
		installCalibrePopup(appWindow).Show()
	}
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
	filter = strings.ToLower(filter)
	var filteredBooks []*search.Book
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), filter) {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return filteredBooks
}

func updateLibraryGrid(grid *fyne.Container, books map[string]*search.Book, filter string, appWindow fyne.Window, tabs *container.AppTabs) {
	grid.RemoveAll()
	for _, book := range filterBooks(books, filter) {
		if exists, err := shared.Exists(book.Filepath); exists && err == nil {
			grid.Add(CreateBookLibraryContainer(*book, appWindow, tabs))
		} else {
			log.Printf("Book does not exist, removing book from database: %s", book.Title)
			database.UpdateDatabase(*book, false, "downloaded")
		}
	}
	grid.Refresh()
}

func RefreshLibraryTab(appWindow fyne.Window, tabs *container.AppTabs) {
	tabs.Items[1] = CreateLibraryView(appWindow, tabs)
	tabs.SelectIndex(1)
	tabs.Refresh()
}

func CreateLibraryView(appWindow fyne.Window, tabs *container.AppTabs) *container.TabItem {
	savedBooks, err := loadSavedBooks()
	if err != nil {
		log.Printf("Could not read savedBooks: %s", err)
		return nil
	}

	libraryViewGrid := container.NewVBox()
	libraryViewGridScrollable := container.NewVScroll(libraryViewGrid)

	filterInput := components.CreateFilterInput()
	var typingTimer *time.Timer
	filterInput.OnChanged = func(filter string) {
		if typingTimer != nil {
			typingTimer.Stop()
		}
		typingTimer = time.AfterFunc(300*time.Millisecond, func() {
			if savedBooks != nil {
				updateLibraryGrid(libraryViewGrid, savedBooks, filter, appWindow, tabs)
			}
		})
	}

	topWidgets := container.NewWithoutLayout(filterInput)
	if savedBooks != nil {
		updateLibraryGrid(libraryViewGrid, savedBooks, "", appWindow, tabs)
	}

	libraryViewLayout := container.NewBorder(topWidgets, nil, nil, nil, libraryViewGridScrollable)
	return container.NewTabItemWithIcon("Library", theme.StorageIcon(), libraryViewLayout)
}
