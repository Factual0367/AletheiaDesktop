package database

import "AletheiaDesktop/search"

func addBook(book *search.Book, existingDatabaseContent map[string]interface{}) map[string]interface{} {
	if existingDatabaseContent["savedBooks"] == nil {
		existingDatabaseContent["savedBooks"] = map[string]*search.Book{}
	}

	savedBooks := existingDatabaseContent["savedBooks"].(map[string]*search.Book)

	savedBooks[book.ID] = book

	return existingDatabaseContent
}
