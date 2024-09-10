package database

import (
	"AletheiaDesktop/search"
	"log"
)

func UpdateDatabase(Book search.Book) {
	existingDatabaseContent, databaseReadErr := ReadDatabaseFile()
	if databaseReadErr != nil {
		log.Fatalln(databaseReadErr.Error())
	}
	existingDatabaseContent = addBook(&Book, existingDatabaseContent)
	writeDatabaseErr := WriteDatabaseToFile(existingDatabaseContent)
	if writeDatabaseErr != nil {
		log.Fatalln(writeDatabaseErr.Error())
	}
}
