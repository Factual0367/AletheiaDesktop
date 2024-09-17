package database

import (
	"AletheiaDesktop/src/models"
	"log"
)

func unfavoriteBook(book *models.Book, existingDatabaseContent map[string]interface{}) map[string]interface{} {
	if existingDatabaseContent["favoriteBooks"] == nil {
		log.Println("Database is already empty, cannot remove book.")
	}

	favoriteBooks := existingDatabaseContent["favoriteBooks"].(map[string]*models.Book)
	delete(favoriteBooks, book.ID)

	return existingDatabaseContent
}
