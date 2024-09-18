package views

import (
	"AletheiaDesktop/src/models"
	"AletheiaDesktop/src/ui/components"
	"AletheiaDesktop/src/util/database"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"log"
	"time"
)

func loadFavoriteBooks() (map[string]*models.Book, error) {
	userData, err := database.ReadDatabaseFile()
	if err != nil || len(userData) == 0 {
		return nil, err
	}

	if favoriteBooks, ok := userData["favoriteBooks"].(map[string]*models.Book); ok {
		return favoriteBooks, nil
	}
	return nil, nil
}

func updateBookmarksGrid(grid *fyne.Container, books map[string]*models.Book, filter string, appWindow fyne.Window, tabs *container.AppTabs) {
	grid.RemoveAll()

	filteredBooks := filterBooks(books, filter)
	for _, book := range filteredBooks {
		grid.Add(CreateBookBookmarksContainer(*book, appWindow, tabs))
	}

	grid.Refresh()
}

func refreshBookmarksTab(appWindow fyne.Window, tabs *container.AppTabs) {
	tabs.Items[2] = CreateBookmarksView(appWindow, tabs)
	tabs.SelectIndex(2)
	tabs.Refresh()
}

func CreateBookmarksView(appWindow fyne.Window, tabs *container.AppTabs) *container.TabItem {
	favoriteBooks, err := loadFavoriteBooks()
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

	bookmarksViewLayout := container.NewBorder(container.NewMax(filterInput), nil, nil, nil, bookmarksViewGridScrollable)
	return container.NewTabItemWithIcon("Bookmarks", theme.ContentAddIcon(), bookmarksViewLayout)
}
