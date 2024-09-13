package database

import (
	"AletheiaDesktop/search"
	"log"
)

func InitializeDatabase() {
	initialEmptyDatabase := map[string]interface{}{
		"savedBooks":    map[string]*search.Book{},
		"favoriteBooks": map[string]*search.Book{},
		"userEmail":     "",
		"userPassword":  "",
	}

	fileWriteErr := WriteDatabaseToFile(initialEmptyDatabase)

	if fileWriteErr != nil {
		log.Fatalln("Unable to write to database file")
		panic(fileWriteErr)
	}
}
