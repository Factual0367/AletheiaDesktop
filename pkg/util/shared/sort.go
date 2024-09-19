package shared

import (
	"AletheiaDesktop/internal/models"
	"sort"
)

func SortBooksByTitle(books map[string]*models.Book) []*models.Book {
	sortedBooks := make([]*models.Book, 0, len(books))
	for _, book := range books {
		sortedBooks = append(sortedBooks, book)
	}

	sort.Slice(sortedBooks, func(i, j int) bool {
		return sortedBooks[i].Title < sortedBooks[j].Title
	})

	return sortedBooks
}
