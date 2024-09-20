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

var (
	previousLibrarySize             int = 0
	previousFilteredDownloadedBooks models.BookSlice
)

func updateLibraryGrid(myApp fyne.App, grid *fyne.Container, books map[string]*models.Book, filter string, appWindow fyne.Window, tabs *container.AppTabs) {
	grid.RemoveAll()

	filteredBooks := shared.FilterBooks(books, filter)
	currentLibrarySize := len(filteredBooks)

	// recreate bookSlice if the size of filteredBooks has changed
	if currentLibrarySize != previousLibrarySize {
		log.Println("Library size has changed, recreating book slice.")
		previousLibrarySize = currentLibrarySize

		previousFilteredDownloadedBooks = make(models.BookSlice, 0, currentLibrarySize)
		for _, book := range filteredBooks {
			previousFilteredDownloadedBooks = append(previousFilteredDownloadedBooks, book)
		}

		// sorting is necessary here for the time being to keep order in the GUI
		sort.Sort(previousFilteredDownloadedBooks)
	}

	for _, book := range previousFilteredDownloadedBooks {
		if exists, err := shared.Exists(book.Filepath); exists && err == nil {
			grid.Add(CreateBookLibraryContainer(myApp, *book, appWindow, tabs))
		} else {
			log.Printf("Book does not exist, removing book from database: %s", book.Title)
			database.UpdateDatabase(*book, false, "downloaded")
		}
	}

	grid.Refresh()
}

func RefreshLibraryTab(myApp fyne.App, appWindow fyne.Window, tabs *container.AppTabs) {
	tabs.Items[1] = CreateLibraryView(myApp, appWindow, tabs)
	tabs.SelectIndex(1)
	tabs.Refresh()
}

func CreateLibraryView(myApp fyne.App, appWindow fyne.Window, tabs *container.AppTabs) *container.TabItem {
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
				updateLibraryGrid(myApp, libraryViewGrid, savedBooks, filter, appWindow, tabs)
			}
		})
	}

	topWidgets := container.NewStack(filterInput)
	if savedBooks != nil {
		updateLibraryGrid(myApp, libraryViewGrid, savedBooks, "", appWindow, tabs)
	}

	libraryViewLayout := container.NewBorder(topWidgets, nil, nil, nil, libraryViewGridScrollable)
	return container.NewTabItemWithIcon("Library", theme.StorageIcon(), libraryViewLayout)
}
