package views

import (
	"AletheiaDesktop/search"
	book2 "AletheiaDesktop/ui/book"
	"AletheiaDesktop/util/database"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"strings"
)

func loadSavedBooks() (map[string]*search.Book, error) {
	userData, err := database.ReadDatabaseFile()
	if err != nil {
		return nil, err
	}

	if savedBooks, ok := userData["savedBooks"].(map[string]*search.Book); ok {
		return savedBooks, nil
	}
	return nil, nil
}

func updateLibraryGrid(grid *fyne.Container, books map[string]*search.Book, filter string) {
	grid.Objects = nil

	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(filter)) {
			bookLibraryContainer := book2.CreateBookLibraryContainer(*book)
			grid.Add(bookLibraryContainer)
		}
	}
	grid.Refresh()
}

func CreateLibraryView() *container.TabItem {
	filterInput := widget.NewEntry()
	filterInput.PlaceHolder = "Filter"
	libraryViewGrid := container.NewVBox()

	savedBooks, err := loadSavedBooks()
	if err != nil {
		log.Printf("Could not read savedBooks %s", err)
	}

	if savedBooks != nil {
		updateLibraryGrid(libraryViewGrid, savedBooks, "")
	}

	filterInput.OnChanged = func(filter string) {
		if savedBooks != nil {
			updateLibraryGrid(libraryViewGrid, savedBooks, filter)
		}
	}

	libraryViewLayout := container.NewBorder(filterInput, nil, nil, nil, libraryViewGrid)

	return container.NewTabItemWithIcon("Library", theme.StorageIcon(), libraryViewLayout)
}
