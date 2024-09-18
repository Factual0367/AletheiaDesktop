package database

import (
	"AletheiaDesktop/internal/models"
	"log"
)

func removeBook(book *models.Book, existingDatabaseContent map[string]interface{}) map[string]interface{} {
	if existingDatabaseContent["savedBooks"] == nil {
		log.Println("Database is already empty, cannot remove book.")
	}

	savedBooks := existingDatabaseContent["savedBooks"].(map[string]*models.Book)
	delete(savedBooks, book.ID)

	return existingDatabaseContent
}
