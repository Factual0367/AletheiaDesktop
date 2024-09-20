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
	previousBookmarksSize           int = 0
	previousFilteredBookmarkedBooks models.BookSlice
)

func updateBookmarksGrid(myApp fyne.App, grid *fyne.Container, books map[string]*models.Book, filter string, appWindow fyne.Window, tabs *container.AppTabs) {
	grid.RemoveAll()

	filteredBooks := shared.FilterBooks(books, filter)
	currentBookmarksSize := len(filteredBooks)

	if currentBookmarksSize != previousBookmarksSize {
		log.Println("Bookmarks size changed, rebuilding book slice.")
		previousBookmarksSize = currentBookmarksSize

		previousFilteredBookmarkedBooks = make(models.BookSlice, 0, currentBookmarksSize)
		for _, book := range filteredBooks {
			previousFilteredBookmarkedBooks = append(previousFilteredBookmarkedBooks, book)
		}

		sort.Sort(previousFilteredBookmarkedBooks)
	}

	for _, book := range previousFilteredBookmarkedBooks {
		bookContainer := CreateBookBookmarksContainer(myApp, *book, appWindow, tabs)
		grid.Add(bookContainer)
	}

	grid.Refresh()
}

func refreshBookmarksTab(myApp fyne.App, appWindow fyne.Window, tabs *container.AppTabs) {
	tabs.Items[2] = CreateBookmarksView(myApp, appWindow, tabs)
	tabs.SelectIndex(2)
	tabs.Refresh()
}

func CreateBookmarksView(myApp fyne.App, appWindow fyne.Window, tabs *container.AppTabs) *container.TabItem {
	favoriteBooks, err := database.LoadFavoriteBooks()
	if err != nil {
		log.Printf("Could not read favoriteBooks: %s", err)
	}

	filterInput := components.CreateFilterInput()
	bookmarksViewGrid := container.NewVBox()
	bookmarksViewGridScrollable := container.NewVScroll(bookmarksViewGrid)

	var typingTimer *time.Timer
	filterInput.OnChanged = func(filter string) {
		if typingTimer != nil {
			typingTimer.Stop()
		}
		typingTimer = time.AfterFunc(300*time.Millisecond, func() {
			if favoriteBooks != nil {
				updateBookmarksGrid(myApp, bookmarksViewGrid, favoriteBooks, filter, appWindow, tabs)
			}
		})
	}

	if favoriteBooks != nil {
		updateBookmarksGrid(myApp, bookmarksViewGrid, favoriteBooks, "", appWindow, tabs)
	}

	bookmarksViewLayout := container.NewBorder(container.NewStack(filterInput), nil, nil, nil, bookmarksViewGridScrollable)
	return container.NewTabItemWithIcon("Bookmarks", theme.ContentAddIcon(), bookmarksViewLayout)
}
