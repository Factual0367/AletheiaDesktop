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

func constructBookContainers(query *libgenapi.Query, detailsContainer *fyne.Container) *fyne.Container {
	bookGrid := container.NewVBox()

	for _, book := range query.Results {
		convertedBook := search.Book{
			Book:       book, // assign libgenapi.Book
			Filename:   "",   // initialize with dummy values
			Filepath:   "",
			Downloaded: false,
		}
		convertedBook.ConstructFilename()
		convertedBook.ConstructFilepath()
		bookItem := book2.CreateBookListContainer(convertedBook, detailsContainer)
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
	defaultDetailsView := book2.CreateBookDetailsView(defaultBook)
	defaultDetailsViewContainer := container.NewVBox(defaultDetailsView)
	return defaultDetailsViewContainer
}

func createSearchBar(onSearch func()) (searchInput *widget.Entry, searchButton *widget.Button, searchTypeWidget *widget.Select) {
	searchInput = widget.NewEntry()
	searchInput.SetPlaceHolder("Enter search query")

	searchButton = widget.NewButtonWithIcon("", theme.SearchIcon(), func() { onSearch() })
	searchInput.OnSubmitted = func(text string) { onSearch() }

	searchTypeWidget = widget.NewSelect([]string{"Default", "Author", "Title"}, func(value string) {
		log.Println("Select set to", value)
	})
	searchTypeWidget.PlaceHolder = "Default"

	return searchInput, searchButton, searchTypeWidget
}

func layoutTopContent(searchInput *widget.Entry, searchButton *widget.Button, searchTypeWidget *widget.Select) *fyne.Container {
	topContent := container.NewWithoutLayout(searchInput, searchButton, searchTypeWidget)

	searchInput.Move(fyne.NewPos(5, 7))
	searchInput.Resize(fyne.NewSize(500, 40))

	searchButton.Move(fyne.NewPos(635, 7))
	searchButton.Resize(fyne.NewSize(50, 40))

	searchTypeWidget.Move(fyne.NewPos(510, 7))
	searchTypeWidget.Resize(fyne.NewSize(120, 40))

	return topContent
}

func executeSearch(searchInput *widget.Entry, searchType string, resultsContainer *fyne.Container, defaultDetailsContainer *fyne.Container) {
	resultsContainer.Objects = nil            // clear previous results
	resultsContainer.Add(widget.NewLabel("")) // padding

	go func() {
		query := search.SearchLibgen(searchInput.Text, searchType)
		if query != nil {
			resultsContainer.Add(constructBookContainers(query, defaultDetailsContainer))
		}
		resultsContainer.Refresh() // refresh to display new results
	}()
}

func CreateSearchView() *container.TabItem {
	resultsContainer := container.NewVBox()
	resultsContentScrollable := container.NewVScroll(resultsContainer)
	detailsContainer := createDefaultDetailsView()

	var searchType = "Default"

	var searchInput = widget.NewEntry()
	searchInput, searchButton, searchTypeWidget := createSearchBar(func() {
		executeSearch(searchInput, searchType, resultsContainer, detailsContainer)
	})

	topContent := layoutTopContent(searchInput, searchButton, searchTypeWidget)

	searchContent := container.NewBorder(
		topContent, nil, nil, nil, // bottom, left, right are nil
		resultsContentScrollable, // center content
	)

	splitView := container.NewHSplit(
		detailsContainer,
		searchContent,
	)
	splitView.SetOffset(0.25)

	return container.NewTabItemWithIcon("Search", theme.SearchIcon(), splitView)
}
