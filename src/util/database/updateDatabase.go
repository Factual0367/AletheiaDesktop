package database

import (
	"AletheiaDesktop/src/models"
	"fmt"
	"log"
)

func UpdateDatabase(Book models.Book, add bool, saveType string) {
	existingDatabaseContent, databaseReadErr := ReadDatabaseFile()
	if databaseReadErr != nil {
		log.Fatalln(databaseReadErr.Error())
	}
	log.Println(fmt.Sprintf("Updating database"))
	if saveType == "downloaded" {
		if add {
			existingDatabaseContent = addBook(&Book, existingDatabaseContent)
		} else {
			log.Println("Removing book from database")
			existingDatabaseContent = removeBook(&Book, existingDatabaseContent)
		}
	} else if saveType == "favorited" {
		if add {
			existingDatabaseContent = favoriteBook(&Book, existingDatabaseContent)
		} else {
			log.Println("Removing book from database")
			existingDatabaseContent = unfavoriteBook(&Book, existingDatabaseContent)
		}
	}

	writeDatabaseErr := WriteDatabaseToFile(existingDatabaseContent)
	if writeDatabaseErr != nil {
		log.Fatalln(writeDatabaseErr.Error())
	}
}
