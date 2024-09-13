package views

import (
	"AletheiaDesktop/search"
	"AletheiaDesktop/util/shared"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/onurhanak/libgenapi"
	"log"
	"strconv"
)

var searchType string = "Default"
var numberOfResults int = 25

func constructBookContainers(query *libgenapi.Query, detailsContainer *fyne.Container) *fyne.Container {
	bookGrid := container.NewVBox()

	for _, book := range query.Results {
		convertedBook := search.Book{
			Book:       book, // extend libgenapi.Book
			Filename:   "",   // initialize with dummy values
			Filepath:   "",
			Downloaded: false,
		}
		convertedBook.ConstructFilename()
		convertedBook.ConstructFilepath()
		bookItem := CreateBookListContainer(convertedBook, detailsContainer)
		bookGrid.Add(bookItem)
	}

	return bookGrid
}

func createDefaultDetailsView() *fyne.Container {
	defaultBook := search.Book{
		Book: libgenapi.Book{
			Title:     "Select a book to view details.",
			CoverLink: "https://cdn.pixabay.com/photo/2013/07/13/13/34/book-161117_960_720.png",
		},
		Filename:   "",
		Filepath:   "",
		Downloaded: false,
	}
	defaultDetailsView := CreateBookDetailsView(defaultBook, true)
	defaultDetailsViewContainer := container.NewVBox(defaultDetailsView)
	return defaultDetailsViewContainer
}

func createSearchBar(onSearch func()) (searchInput *widget.Entry, searchButton *widget.Button, searchTypeWidget *widget.Select, numberOfResultsSelector *widget.Select) {
	searchInput = widget.NewEntry()
	searchInput.SetPlaceHolder("Enter search query")

	searchButton = widget.NewButtonWithIcon("", theme.SearchIcon(), func() { onSearch() })
	searchInput.OnSubmitted = func(text string) { onSearch() }

	searchTypeWidget = widget.NewSelect([]string{"Default", "Author", "Title"}, func(value string) {
		searchType = value
		log.Println("Select set to", value)
	})
	searchTypeWidget.PlaceHolder = "Default"

	numberOfResultsSelector = widget.NewSelect([]string{"25", "50", "100"}, func(value string) {
		numberOfResults, _ = strconv.Atoi(value) // handle this
		log.Println("Number of results set to", value)
	})
	numberOfResultsSelector.PlaceHolder = "25"

	return searchInput, searchButton, searchTypeWidget, numberOfResultsSelector
}

func layoutTopContent(searchInput *widget.Entry, searchButton *widget.Button, searchTypeWidget *widget.Select, numberOfResultsSelector *widget.Select) *fyne.Container {

	searchInputContainer := container.NewStack(searchInput)
	searchInputContainer.MinSize()
	topContent := container.NewGridWithColumns(2, searchInputContainer, container.NewHBox(searchButton, searchTypeWidget, numberOfResultsSelector, layout.NewSpacer()))
	return topContent
}

func executeSearch(searchInput *widget.Entry, searchType string, resultsContainer *fyne.Container, defaultDetailsContainer *fyne.Container) {
	resultsContainer.Objects = nil // clear previous results

	go func() {
		query, err := search.SearchLibgen(searchInput.Text, searchType, numberOfResults)
		if err != nil {
			shared.SendNotification("Failed", "Library Genesis is not responding.")
		}

		if query != nil {
			resultsContainer.Add(constructBookContainers(query, defaultDetailsContainer))
			resultsContainer.Refresh() // refresh to display new results
		}
	}()
}

func CreateSearchView() *container.TabItem {
	resultsContainer := container.NewVBox()
	resultsContentScrollable := container.NewVScroll(resultsContainer)
	detailsContainer := createDefaultDetailsView()

	var searchInput = widget.NewEntry()

	searchInput, searchButton, searchTypeWidget, numberOfResultsSelector := createSearchBar(func() {
		executeSearch(searchInput, searchType, resultsContainer, detailsContainer)
	})

	topContent := layoutTopContent(searchInput, searchButton, searchTypeWidget, numberOfResultsSelector)

	searchContent := container.NewBorder(
		topContent, nil, nil, nil, // bottom, left, right are nil
		resultsContentScrollable, // center content
	)

	splitView := container.NewHSplit(
		detailsContainer,
		searchContent,
	)
	splitView.SetOffset(0.20)

	return container.NewTabItemWithIcon("Search", theme.SearchIcon(), splitView)
}
