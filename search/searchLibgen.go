package search

import (
	"fmt"
	"github.com/onurhanak/libgenapi"
	"log"
	"strings"
)

type QueryConstructor func(queryType, searchQuery string, numOfResults int) *libgenapi.Query

func SearchLibgen(searchQuery string, queryType string, numberOfResults int, newQuery QueryConstructor) (*libgenapi.Query, error) {
	query := newQuery(strings.ToLower(queryType), searchQuery, numberOfResults)
	err := query.Search()
	if err != nil {
		log.Println(fmt.Sprintf("Error : %s. Libgen API did not return any results.", err))
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
