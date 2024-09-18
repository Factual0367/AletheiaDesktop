package database

import (
	"AletheiaDesktop/internal/models"
)

func addBook(book *models.Book, existingDatabaseContent map[string]interface{}) map[string]interface{} {
	if existingDatabaseContent["savedBooks"] == nil {
		existingDatabaseContent["savedBooks"] = map[string]*models.Book{}
	}

	savedBooks := existingDatabaseContent["savedBooks"].(map[string]*models.Book)

	savedBooks[book.ID] = book

	return existingDatabaseContent
}
