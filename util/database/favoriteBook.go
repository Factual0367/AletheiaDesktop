package database

import (
	"AletheiaDesktop/search"
)

func favoriteBook(book *search.Book, existingDatabaseContent map[string]interface{}) map[string]interface{} {
	// should check if book is already favorited
	if existingDatabaseContent["favoriteBooks"] == nil {
		existingDatabaseContent["favoriteBooks"] = map[string]*search.Book{}
	}
	favoriteBooks := existingDatabaseContent["favoriteBooks"].(map[string]*search.Book)

	favoriteBooks[book.ID] = book

	return existingDatabaseContent
}
