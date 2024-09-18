package email

import (
	"AletheiaDesktop/pkg/util/config"
	"fmt"
	"log"
)

func SaveEmail(email string) bool {
	existingUserConfig, configReadErr := config.ReadConfigFile()
	if configReadErr != nil {
		log.Println(fmt.Sprintf("Error reading user config: %s", configReadErr.Error()))
	}

	existingUserConfig["userEmail"] = email
	configWriteErr := config.WriteConfigFile(existingUserConfig)
	if configWriteErr != nil {
		log.Println(fmt.Sprintf("Error writing user config: %s", configReadErr.Error()))
		return false
	}
	return true
}
