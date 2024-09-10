package search

import (
	"fmt"
	"github.com/onurhanak/libgenapi"
	"log"
	"strings"
)

func SearchLibgen(searchQuery string, queryType string) *libgenapi.Query {
	query := libgenapi.NewQuery(strings.ToLower(queryType), searchQuery, 25)
	err := query.Search()
	if err != nil {
		log.Println(fmt.Sprintf("Error : %s. Libgen API did not return any results.", err))
		query.Results = []libgenapi.Book{}
	}
	return query
}
