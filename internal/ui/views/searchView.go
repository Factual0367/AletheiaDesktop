package views

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/internal/search"
	"AletheiaDesktop/pkg/util/shared"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/onurhanak/libgenapi"
)

var (
	searchType      = "Default"
	numberOfResults = 25
)

func constructBookContainers(myApp fyne.App, query *libgenapi.Query, appWindow fyne.Window) *fyne.Container {
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
		bookItem := CreateBookListContainer(myApp, convertedBook, appWindow)
		bookGrid.Add(bookItem)
	}

	return bookGrid
}

func createSearchBar(onSearch func()) (*widget.Entry, *widget.Button, *widget.Select, *widget.Select) {
	searchInput := widget.NewEntry()
	searchInput.SetPlaceHolder("Enter search query")
	searchButton := widget.NewButtonWithIcon("", theme.SearchIcon(), onSearch)
	searchInput.OnSubmitted = func(text string) { onSearch() }

	searchTypeWidget := widget.NewSelect([]string{"Default", "Author", "Title"}, func(value string) {
		searchType = value
	})
	searchTypeWidget.PlaceHolder = "Default"

	numberOfResultsSelector := widget.NewSelect([]string{"25", "50", "100"}, func(value string) {
		numberOfResults, _ = strconv.Atoi(value)
	})
	numberOfResultsSelector.PlaceHolder = "25"

	return searchInput, searchButton, searchTypeWidget, numberOfResultsSelector
}

func layoutTopContent(searchInput *widget.Entry, searchButton *widget.Button, searchTypeWidget *widget.Select, numberOfResultsSelector *widget.Select) *fyne.Container {
	topContent := container.NewGridWithRows(1,
		container.NewStack(searchInput),
		container.NewHBox(searchButton, searchTypeWidget, numberOfResultsSelector),
	)
	return topContent
}

func executeSearch(myApp fyne.App, searchInput *widget.Entry, searchType string, resultsContainer *fyne.Container, appWindow fyne.Window) {
	resultsContainer.Objects = nil // Clear previous results

	go func() {
		query, err := search.SearchLibgen(searchInput.Text, searchType, numberOfResults)
		if err != nil {
			shared.SendNotification(myApp, "Failed", "Library Genesis is not responding.")
			return
		}

		if query != nil {
			resultsContainer.Add(constructBookContainers(myApp, query, appWindow))
			resultsContainer.Refresh()
		}
	}()
}

func CreateSearchView(myApp fyne.App, appWindow fyne.Window) *container.TabItem {
	resultsContainer := container.NewVBox()
	searchInput := widget.NewEntry()

	searchInput, searchButton, searchTypeWidget, numberOfResultsSelector := createSearchBar(func() {
		executeSearch(myApp, searchInput, searchType, resultsContainer, appWindow)
	})

	topContent := layoutTopContent(searchInput, searchButton, searchTypeWidget, numberOfResultsSelector)

	searchContent := container.NewBorder(
		topContent, nil, nil, nil,
		container.NewVScroll(resultsContainer), // Results are scrollable
	)

	searchContentView := container.NewVScroll(searchContent)

	return container.NewTabItemWithIcon("Search", theme.SearchIcon(), searchContentView)
}
