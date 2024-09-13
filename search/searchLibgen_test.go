// search_test.go
package search_test

import (
	"AletheiaDesktop/search"
	"testing"
)

func TestSearchLibgen_Success(t *testing.T) {
	searchQuery := "Go Programming"
	queryType := "title"
	numberOfResults := 3

	query, err := search.SearchLibgen(searchQuery, queryType, numberOfResults)
	if err != nil {
		t.Fatalf("SearchLibgen returned an error: %v", err)
	}

	if query == nil {
		t.Fatal("Expected query to be non-nil")
	}

	if len(query.Results) == 0 {
		t.Error("Expected results, got none")
	}

	for _, book := range query.Results {
		if book.CoverLink == "" {
			t.Errorf("Book '%s' has an empty CoverLink", book.Title)
		}
	}
}
