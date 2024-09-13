package email

import (
	"AletheiaDesktop/util/database"
	"fmt"
	"log"
)

func GetUserEmail() string {
	existingDatabaseContent, databaseReadErr := database.ReadDatabaseFile()
	if databaseReadErr != nil {
		log.Println(databaseReadErr.Error())
	}
	userEmail := existingDatabaseContent["userEmail"]
	if userEmail == "" {
		return "Your email address"
	}
	userEmailText := fmt.Sprintf("%s", userEmail)
	return userEmailText
}
