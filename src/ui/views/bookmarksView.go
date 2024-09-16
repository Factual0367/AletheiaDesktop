package views

import (
	"AletheiaDesktop/src/search"
	"AletheiaDesktop/src/util/database"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
	grid.Objects = nil

	for _, book := range books {
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
	filterInput := widget.NewEntry()
	filterInput.PlaceHolder = "Filter"
	filterInput.Resize(fyne.NewSize(800, filterInput.MinSize().Height)) // Set the desired width

	topWidgets := container.NewWithoutLayout(filterInput)

	bookmarksViewGrid := container.NewVBox()
	bookmarksViewGridScrollable := container.NewVScroll(bookmarksViewGrid)
	favoriteBooks, err := loadFavoriteBooks()
	if err != nil {
		log.Printf("Could not read savedBooks %s", err)
	}

	if favoriteBooks != nil {
		updateBookmarksGrid(bookmarksViewGrid, favoriteBooks, "", appWindow, tabs)
	}

	var typingTimer *time.Timer

	filterInput.OnChanged = func(filter string) {
		if typingTimer != nil {
			typingTimer.Stop() // stop timer
		}

		typingTimer = time.AfterFunc(500*time.Millisecond, func() { // 500ms delay so filtering does not get laggy
			if favoriteBooks != nil {
				updateBookmarksGrid(bookmarksViewGrid, favoriteBooks, filter, appWindow, tabs)
			}
		})
	}

	bookmarksViewLayout := container.NewBorder(topWidgets, nil, nil, nil, bookmarksViewGridScrollable)

	return container.NewTabItemWithIcon("Bookmarks", theme.ContentAddIcon(), bookmarksViewLayout)
}
