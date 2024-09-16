package views

import (
	"AletheiaDesktop/src/search"
	"AletheiaDesktop/src/ui/components"
	"AletheiaDesktop/src/util/database"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"log"
	"strings"
	"time"
)

func loadFavoriteBooks() (map[string]*search.Book, error) {
	userData, err := database.ReadDatabaseFile()

	if len(userData) == 0 {
		userData, err = database.ReadDatabaseFile()
	}

	if err != nil {
		return nil, err
	}

	if favoriteBooks, ok := userData["favoriteBooks"].(map[string]*search.Book); ok {
		return favoriteBooks, nil
	}
	return nil, nil
}

func updateBookmarksGrid(grid *fyne.Container, books map[string]*search.Book, filter string, appWindow fyne.Window, tabs *container.AppTabs) {
	grid.RemoveAll()

	filteredBooks := filterBooks(books, filter)

	for _, book := range filteredBooks {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(filter)) {

			bookLibraryContainer := CreateBookBookmarksContainer(*book, appWindow, tabs)
			grid.Add(bookLibraryContainer)
		}
	}
	grid.Refresh()
}

func refreshBookmarksTab(appWindow fyne.Window, tabs *container.AppTabs) {
	newBookmarksView := CreateBookmarksView(appWindow, tabs)
	tabs.Items[2] = newBookmarksView
	tabs.SelectIndex(2)

	tabs.Refresh()
}

func CreateBookmarksView(appWindow fyne.Window, tabs *container.AppTabs) *container.TabItem {
	favoriteBooks, err := loadFavoriteBooks()
	if err != nil {
		log.Printf("Could not read savedBooks %s", err)
	}

	filterInput := components.CreateFilterInput()
	topWidgets := container.NewWithoutLayout(filterInput)
	bookmarksViewGrid := container.NewVBox()
	bookmarksViewGridScrollable := container.NewVScroll(bookmarksViewGrid)

	var typingTimer *time.Timer
	filterInput.OnChanged = func(filter string) {
		if typingTimer != nil {
			typingTimer.Stop() // stop timer
		}

		typingTimer = time.AfterFunc(300*time.Millisecond, func() { // 300ms delay so filtering does not get laggy
			if favoriteBooks != nil {
				updateBookmarksGrid(bookmarksViewGrid, favoriteBooks, filter, appWindow, tabs)
			}
		})
	}

	if favoriteBooks != nil {
		updateBookmarksGrid(bookmarksViewGrid, favoriteBooks, "", appWindow, tabs)
	}

	bookmarksViewLayout := container.NewBorder(topWidgets, nil, nil, nil, bookmarksViewGridScrollable)
	bookmarksView := container.NewTabItemWithIcon("Bookmarks", theme.ContentAddIcon(), bookmarksViewLayout)
	return bookmarksView
}
