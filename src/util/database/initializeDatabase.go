package database

import (
	"AletheiaDesktop/src/search"
	"log"
)

func InitializeDatabase() {
	initialEmptyDatabase := map[string]interface{}{
		"savedBooks":    map[string]*search.Book{},
		"favoriteBooks": map[string]*search.Book{},
	}

	fileWriteErr := WriteDatabaseToFile(initialEmptyDatabase)

	if fileWriteErr != nil {
		log.Fatalln("Unable to write to database file")
		panic(fileWriteErr)
	}
}
