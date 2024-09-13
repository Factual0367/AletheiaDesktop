package email

import (
	"AletheiaDesktop/util/database"
	"AletheiaDesktop/util/shared"
	"fmt"
	"log"
)

func GetUserPassword() string {
	existingDatabaseContent, databaseReadErr := database.ReadDatabaseFile()
	if databaseReadErr != nil {
		log.Println(databaseReadErr.Error())
	}
	userPassword := existingDatabaseContent["userPassword"]
	if userPassword == "" {
		shared.SendNotification("Error", "You need to set up the app password to use this feature.")
	}
	userPasswordText := fmt.Sprintf("%s", userPassword)
	return userPasswordText
}
