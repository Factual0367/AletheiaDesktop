package views

import (
	"AletheiaDesktop/src/models"
	"AletheiaDesktop/src/ui/components"
	"AletheiaDesktop/src/util/database"
	"AletheiaDesktop/src/util/shared"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"log"
	"strings"
	"time"
)

func loadSavedBooks() (map[string]*models.Book, error) {
	userData, err := database.ReadDatabaseFile()
	if err != nil {
		return nil, err
	}

	if savedBooks, ok := userData["savedBooks"].(map[string]*models.Book); ok {
		return savedBooks, nil
	}
	return nil, nil
}

func filterBooks(books map[string]*models.Book, filter string) []*models.Book {
	filter = strings.ToLower(filter)
	var filteredBooks []*models.Book
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), filter) {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return filteredBooks
}

func updateLibraryGrid(grid *fyne.Container, books map[string]*models.Book, filter string, appWindow fyne.Window, tabs *container.AppTabs) {
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

	topWidgets := container.NewMax(filterInput)
	if savedBooks != nil {
		updateLibraryGrid(libraryViewGrid, savedBooks, "", appWindow, tabs)
	}

	libraryViewLayout := container.NewBorder(topWidgets, nil, nil, nil, libraryViewGridScrollable)
	return container.NewTabItemWithIcon("Library", theme.StorageIcon(), libraryViewLayout)
}
