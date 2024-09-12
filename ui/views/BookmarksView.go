package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
)

func CreateBookmarksView(appWindow fyne.Window) *container.TabItem {
	filterInput := widget.NewEntry()
	filterInput.PlaceHolder = "Filter"
	libraryViewGrid := container.NewVBox()

	savedBooks, err := loadSavedBooks()
	if err != nil {
		log.Printf("Could not read savedBooks %s", err)
	}

	if savedBooks != nil {
		updateLibraryGrid(libraryViewGrid, savedBooks, "", appWindow)
	}

	var typingTimer *time.Timer

	filterInput.OnChanged = func(filter string) {
		if typingTimer != nil {
			typingTimer.Stop() // stop timer
		}

		typingTimer = time.AfterFunc(500*time.Millisecond, func() { // 500ms delay so filtering does not get laggy
			if savedBooks != nil {
				updateLibraryGrid(libraryViewGrid, savedBooks, filter, appWindow)
			}
		})
	}

	libraryViewLayout := container.NewBorder(filterInput, nil, nil, nil, libraryViewGrid)

	return container.NewTabItemWithIcon("Bookmarks", theme.ContentAddIcon(), libraryViewLayout)
}
