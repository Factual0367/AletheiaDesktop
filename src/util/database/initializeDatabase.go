package database

import (
	"AletheiaDesktop/src/models"
	"log"
)

func InitializeDatabase() map[string]interface{} {
	initialEmptyDatabase := map[string]interface{}{
		"savedBooks":    map[string]*models.Book{},
		"favoriteBooks": map[string]*models.Book{},
	}

	fileWriteErr := WriteDatabaseToFile(initialEmptyDatabase)

	if fileWriteErr != nil {
		log.Fatalln("Unable to write to database file")
		panic(fileWriteErr)
	}
	return initialEmptyDatabase
}
