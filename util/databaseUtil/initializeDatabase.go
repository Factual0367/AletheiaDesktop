package configUtil

import (
	"github.com/onurhanak/libgenapi"
	"log"
)

func InitializeDatabase() {
	initialEmptyDatabase := map[string]interface{}{
		"savedBooks": map[string]*libgenapi.Book{},
	}

	fileWriteErr := WriteDatabaseToFile(initialEmptyDatabase)

	if fileWriteErr != nil {
		log.Fatalln("Unable to write to database file")
		panic(fileWriteErr)
	}
}
