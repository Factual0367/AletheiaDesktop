package views

import (
	"AletheiaDesktop/search"
	book2 "AletheiaDesktop/ui/book"
	"AletheiaDesktop/util/database"
	"AletheiaDesktop/util/shared"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"strings"
	"time"
)

func loadSavedBooks() (map[string]*search.Book, error) {
	userData, err := database.ReadDatabaseFile()

	if len(userData) == 0 {
		fmt.Println(len(userData))
		userData, err = database.ReadDatabaseFile()
	}

	if err != nil {
		return nil, err
	}

	if savedBooks, ok := userData["savedBooks"].(map[string]*search.Book); ok {
		return savedBooks, nil
	}
	return nil, nil
}

func updateLibraryGrid(grid *fyne.Container, books map[string]*search.Book, filter string, appWindow fyne.Window) {
	grid.Objects = nil

	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(filter)) {
			bookFileExists, err := shared.Exists(book.Filepath)
			if err != nil {
				log.Printf("Book does not exist, removing book from database. Title: %s", book.Title)
				database.UpdateDatabase(*book, false, "downloaded")
			}
			if bookFileExists {
				bookLibraryContainer := book2.CreateBookLibraryContainer(*book, appWindow)
				grid.Add(bookLibraryContainer)
			}

		}
	}
	grid.Refresh()
}

func refreshLibraryViewButton(appWindow fyne.Window, tabs *container.AppTabs) *widget.Button {

	refreshButton := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		newLibraryView := CreateLibraryView(appWindow, tabs)
		tabs.Items[1] = newLibraryView
		tabs.Refresh()
	})

	return refreshButton
}

func CreateLibraryView(appWindow fyne.Window, tabs *container.AppTabs) *container.TabItem {
	filterInput := widget.NewEntry()
	filterInput.PlaceHolder = "Filter"
	filterInput.Resize(fyne.NewSize(800, filterInput.MinSize().Height)) // Set the desired width

	refreshButton := refreshLibraryViewButton(appWindow, tabs)
	refreshButton.Resize(fyne.NewSize(40, filterInput.MinSize().Height))
	refreshButton.Move(fyne.NewPos(805, 0))

	topWidgets := container.NewWithoutLayout(filterInput, refreshButton)

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

	libraryViewLayout := container.NewBorder(topWidgets, nil, nil, nil, libraryViewGrid)

	return container.NewTabItemWithIcon("Library", theme.StorageIcon(), libraryViewLayout)
}
