package views

import (
	"AletheiaDesktop/search"
	book2 "AletheiaDesktop/ui/book"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/onurhanak/libgenapi"
	"log"
)

func constructBookContainers(query *libgenapi.Query) *fyne.Container {

	var bookGrid *fyne.Container

	bookGrid = container.NewVBox()
	for _, book := range query.Results {
		convertedBook := &search.Book{
			Book:       book, // assign libgenapi.Book
			Filename:   "",   // initialize with dummy values
			Filepath:   "",
			Downloaded: false,
		}
		convertedBook.ConstructFilename()
		convertedBook.ConstructFilepath()
		bookItem := book2.CreateBookListContainer(convertedBook)
		bookGrid.Add(bookItem)
	}

	return container.NewVBox(bookGrid)
}

func CreateSearchView() *container.TabItem {
	searchInput := widget.NewEntry()
	searchInput.SetPlaceHolder("Enter search query")

	resultsContainer := container.NewVBox()

	var searchType = "Default"

	resultsContentScrollable := container.NewVScroll(resultsContainer)

	executeSearch := func() {
		resultsContainer.Objects = nil            // clear previous results
		resultsContainer.Add(widget.NewLabel("")) // padding
		go func() {
			query := search.SearchLibgen(searchInput.Text, searchType)
			if query != nil {
				resultsContainer.Add(constructBookContainers(query))
			}
			resultsContainer.Refresh() // refresh to display new results
		}()
	}

	searchButton := widget.NewButtonWithIcon("", theme.SearchIcon(), func() { executeSearch() })
	searchInput.OnSubmitted = func(text string) { executeSearch() }

	searchTypeWidget := widget.NewSelect([]string{"Default", "Author", "Title"}, func(value string) {
		searchType = value
		log.Println("Select set to", value)
	})
	searchTypeWidget.PlaceHolder = "Default"

	topContent := container.NewWithoutLayout(searchInput, searchButton, searchTypeWidget)
	searchTypeWidget.Move(fyne.NewPos(510, 7))
	searchTypeWidget.Resize(fyne.NewSize(120, 40))
	searchInput.Move(fyne.NewPos(5, 7))
	searchButton.Move(fyne.NewPos(635, 7))
	searchInput.Resize(fyne.NewSize(500, 40))
	searchButton.Resize(fyne.NewSize(50, 40))

	searchContent := container.NewBorder(
		topContent, nil, nil, nil, // bottom, left, right are nil
		resultsContentScrollable, // center content
	)

	return container.NewTabItemWithIcon("Search", theme.SearchIcon(), searchContent)
}
