package email

import (
	"AletheiaDesktop/util/database"
	"log"
)

func SaveEmail(email string) bool {
	existingDatabaseContent, databaseReadErr := database.ReadDatabaseFile()
	if databaseReadErr != nil {
		log.Fatalln(databaseReadErr.Error())
	}

	existingDatabaseContent["userEmail"] = email
	writeDatabaseErr := database.WriteDatabaseToFile(existingDatabaseContent)
	if writeDatabaseErr != nil {
		log.Fatalln(writeDatabaseErr.Error())
		return false
	}
	return true
}
