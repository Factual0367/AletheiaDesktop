package database

import (
	"AletheiaDesktop/src/search"
	"log"
)

func removeBook(book *search.Book, existingDatabaseContent map[string]interface{}) map[string]interface{} {
	if existingDatabaseContent["savedBooks"] == nil {
		log.Println("Database is already empty, cannot remove book.")
	}

	savedBooks := existingDatabaseContent["savedBooks"].(map[string]*search.Book)
	delete(savedBooks, book.ID)

	return existingDatabaseContent
}
