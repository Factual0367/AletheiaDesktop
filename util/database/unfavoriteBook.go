package database

import (
	"AletheiaDesktop/search"
	"log"
)

func unfavoriteBook(book *search.Book, existingDatabaseContent map[string]interface{}) map[string]interface{} {
	if existingDatabaseContent["favoriteBooks"] == nil {
		log.Println("Database is already empty, cannot remove book.")
	}

	favoriteBooks := existingDatabaseContent["favoriteBooks"].(map[string]*search.Book)
	delete(favoriteBooks, book.ID)

	return existingDatabaseContent
}
