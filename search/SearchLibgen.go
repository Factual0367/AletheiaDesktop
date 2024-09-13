package search

import (
	"AletheiaDesktop/util/shared"
	"fmt"
	"github.com/onurhanak/libgenapi"
	"log"
	"strings"
)

func SearchLibgen(searchQuery string, queryType string, numberOfResults int) (*libgenapi.Query, error) {
	query := libgenapi.NewQuery(strings.ToLower(queryType), searchQuery, numberOfResults)
	err := query.Search()
	if err != nil {
		log.Println(fmt.Sprintf("Error : %s. Libgen API did not return any results.", err))
		shared.SendNotification("Failed", "Library Genesis is not responding.")
		query.Results = []libgenapi.Book{}
		return query, err
	}

	// add coverlinks if does not exist
	for i := range query.Results {
		book := &query.Results[i]
		if book.CoverLink == "" {
			book.CoverLink = "https://cdn.pixabay.com/photo/2013/07/13/13/34/book-161117_960_720.png"
		}
	}

	return query, nil
}
