package views

import (
	"AletheiaDesktop/src/models"
	"AletheiaDesktop/src/search"
	"AletheiaDesktop/src/util/cache"
	"AletheiaDesktop/src/util/shared"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/onurhanak/libgenapi"
	"log"
	"path"
	"strconv"
)

var (
	searchType      = "Default"
	numberOfResults = 25
)

func constructBookContainers(query *libgenapi.Query, detailsContainer *fyne.Container) *fyne.Container {
	bookGrid := container.NewVBox()

	for _, book := range query.Results {
		convertedBook := models.Book{
			Book:       book,
			Filename:   "",
			Filepath:   "",
			Downloaded: false,
		}
		convertedBook.ConstructFilename()
		convertedBook.ConstructFilepath()
		convertedBook.ConstructCoverPath()
		bookItem := CreateBookListContainer(convertedBook, detailsContainer)
		bookGrid.Add(bookItem)
	}

	return bookGrid
}

func createDefaultDetailsView() *fyne.Container {
	defaultBook := models.Book{
		Book: libgenapi.Book{
			ID:        "Default",
			Title:     "Select a book to view details.",
			CoverLink: "https://cdn.pixabay.com/photo/2013/07/13/13/34/book-161117_960_720.png",
		},
		Filename:   "",
		Filepath:   "",
		Downloaded: false,
		CoverPath:  path.Join(cache.GetAletheiaCache(), "Default"),
	}

	defaultDetailsView := CreateBookDetailsView(defaultBook, true)
	return container.NewVBox(defaultDetailsView)
}

func createSearchBar(onSearch func()) (*widget.Entry, *widget.Button, *widget.Select, *widget.Select) {
	searchInput := widget.NewEntry()
	searchInput.SetPlaceHolder("Enter search query")

	searchButton := widget.NewButtonWithIcon("", theme.SearchIcon(), onSearch)
	searchInput.OnSubmitted = func(text string) { onSearch() }

	searchTypeWidget := widget.NewSelect([]string{"Default", "Author", "Title"}, func(value string) {
		searchType = value
		log.Println("Search type set to", value)
	})
	searchTypeWidget.PlaceHolder = "Default"

	numberOfResultsSelector := widget.NewSelect([]string{"25", "50", "100"}, func(value string) {
		numberOfResults, _ = strconv.Atoi(value)
		log.Println("Number of results set to", value)
	})
	numberOfResultsSelector.PlaceHolder = "25"

	return searchInput, searchButton, searchTypeWidget, numberOfResultsSelector
}

func layoutTopContent(searchInput *widget.Entry, searchButton *widget.Button, searchTypeWidget *widget.Select, numberOfResultsSelector *widget.Select) *fyne.Container {
	topContent := container.NewGridWithColumns(2,
		container.NewStack(searchInput),
		container.NewHBox(searchButton, searchTypeWidget, numberOfResultsSelector, layout.NewSpacer()),
	)
	return topContent
}

func executeSearch(searchInput *widget.Entry, searchType string, resultsContainer, detailsContainer *fyne.Container) {
	resultsContainer.Objects = nil // Clear previous results

	go func() {
		query, err := search.SearchLibgen(searchInput.Text, searchType, numberOfResults)
		if err != nil {
			shared.SendNotification("Failed", "Library Genesis is not responding.")
			return
		}

		if query != nil {
			resultsContainer.Add(constructBookContainers(query, detailsContainer))
			resultsContainer.Refresh()
		}
	}()
}

func CreateSearchView() *container.TabItem {
	resultsContainer := container.NewVBox()
	detailsContainer := createDefaultDetailsView()
	var searchInput = widget.NewEntry()

	searchInput, searchButton, searchTypeWidget, numberOfResultsSelector := createSearchBar(func() {
		executeSearch(searchInput, searchType, resultsContainer, detailsContainer)
	})

	topContent := layoutTopContent(searchInput, searchButton, searchTypeWidget, numberOfResultsSelector)

	searchContent := container.NewBorder(
		topContent, nil, nil, nil,
		container.NewVScroll(resultsContainer), // Results are scrollable
	)

	splitView := container.NewHSplit(detailsContainer, searchContent)
	splitView.SetOffset(0.20) // Adjust the split ratio

	return container.NewTabItemWithIcon("Search", theme.SearchIcon(), splitView)
}
