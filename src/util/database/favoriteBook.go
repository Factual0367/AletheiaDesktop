package database

import (
	"AletheiaDesktop/src/models"
)

func favoriteBook(book *models.Book, existingDatabaseContent map[string]interface{}) map[string]interface{} {
	// should check if book is already favorited
	if existingDatabaseContent["favoriteBooks"] == nil {
		existingDatabaseContent["favoriteBooks"] = map[string]*models.Book{}
	}
	favoriteBooks := existingDatabaseContent["favoriteBooks"].(map[string]*models.Book)

	favoriteBooks[book.ID] = book

	return existingDatabaseContent
}
