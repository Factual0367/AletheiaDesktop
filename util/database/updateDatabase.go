package database

import (
	"AletheiaDesktop/search"
	"log"
)

func UpdateDatabase(Book search.Book, add bool) {
	existingDatabaseContent, databaseReadErr := ReadDatabaseFile()
	if databaseReadErr != nil {
		log.Fatalln(databaseReadErr.Error())
	}
	if add {
		existingDatabaseContent = addBook(&Book, existingDatabaseContent)
	} else {
		log.Println("Removing book from database")
		existingDatabaseContent = removeBook(&Book, existingDatabaseContent)
	}
	writeDatabaseErr := WriteDatabaseToFile(existingDatabaseContent)
	if writeDatabaseErr != nil {
		log.Fatalln(writeDatabaseErr.Error())
	}
}
