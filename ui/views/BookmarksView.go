package views

import (
	"AletheiaDesktop/search"
	book2 "AletheiaDesktop/ui/book"
	"AletheiaDesktop/util/database"
	"fmt"
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
		fmt.Println(len(userData))
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

func updateBookmarksGrid(grid *fyne.Container, books map[string]*search.Book, filter string, appWindow fyne.Window) {
	grid.Objects = nil

	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(filter)) {

			bookLibraryContainer := book2.CreateBookBookmarksContainer(*book, appWindow)
			grid.Add(bookLibraryContainer)
		}
	}
	grid.Refresh()
}

func CreateBookmarksView(appWindow fyne.Window) *container.TabItem {
	filterInput := widget.NewEntry()
	filterInput.PlaceHolder = "Filter"
	bookmarksViewGrid := container.NewVBox()

	favoriteBooks, err := loadFavoriteBooks()
	if err != nil {
		log.Printf("Could not read savedBooks %s", err)
	}

	if favoriteBooks != nil {
		updateBookmarksGrid(bookmarksViewGrid, favoriteBooks, "", appWindow)
	}

	var typingTimer *time.Timer

	filterInput.OnChanged = func(filter string) {
		if typingTimer != nil {
			typingTimer.Stop() // stop timer
		}

		typingTimer = time.AfterFunc(500*time.Millisecond, func() { // 500ms delay so filtering does not get laggy
			if favoriteBooks != nil {
				updateBookmarksGrid(bookmarksViewGrid, favoriteBooks, filter, appWindow)
			}
		})
	}

	bookmarksViewLayout := container.NewBorder(filterInput, nil, nil, nil, bookmarksViewGrid)

	return container.NewTabItemWithIcon("Bookmarks", theme.ContentAddIcon(), bookmarksViewLayout)
}