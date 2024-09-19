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

func updateBookmarksGrid(grid *fyne.Container, books map[string]*models.Book, filter string, appWindow fyne.Window, tabs *container.AppTabs) {
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
		bookContainer := CreateBookBookmarksContainer(*book, appWindow, tabs)
		grid.Add(bookContainer)
	}

	grid.Refresh()
}

func refreshBookmarksTab(appWindow fyne.Window, tabs *container.AppTabs) {
	tabs.Items[2] = CreateBookmarksView(appWindow, tabs)
	tabs.SelectIndex(2)
	tabs.Refresh()
}

func CreateBookmarksView(appWindow fyne.Window, tabs *container.AppTabs) *container.TabItem {
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
				updateBookmarksGrid(bookmarksViewGrid, favoriteBooks, filter, appWindow, tabs)
			}
		})
	}

	if favoriteBooks != nil {
		updateBookmarksGrid(bookmarksViewGrid, favoriteBooks, "", appWindow, tabs)
	}

	bookmarksViewLayout := container.NewBorder(container.NewStack(filterInput), nil, nil, nil, bookmarksViewGridScrollable)
	return container.NewTabItemWithIcon("Bookmarks", theme.ContentAddIcon(), bookmarksViewLayout)
}
