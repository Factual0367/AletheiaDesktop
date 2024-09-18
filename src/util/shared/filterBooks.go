package shared

import (
	"AletheiaDesktop/src/models"
	"strings"
)

func FilterBooks(books map[string]*models.Book, filter string) []*models.Book {
	filter = strings.ToLower(filter)
	var filteredBooks []*models.Book
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), filter) {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return filteredBooks
}
