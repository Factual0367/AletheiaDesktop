package views

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/internal/ui/components"
	"AletheiaDesktop/pkg/util/database"
	"AletheiaDesktop/pkg/util/shared"
	"log"
	"sort"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func updateLibraryGrid(grid *fyne.Container, books map[string]*models.Book, filter string, appWindow fyne.Window, tabs *container.AppTabs) {
	grid.RemoveAll()

	filteredBooks := shared.FilterBooks(books, filter)

	bookSlice := make(models.BookSlice, 0, len(filteredBooks))
	for _, book := range filteredBooks {
		bookSlice = append(bookSlice, book)
	}

	// sorting is necessasry here for the time being
	// to keep order in the gui
	sort.Sort(bookSlice)

	for _, book := range bookSlice {
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
	savedBooks, err := database.LoadSavedBooks()
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
