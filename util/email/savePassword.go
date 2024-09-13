package email

import (
	"AletheiaDesktop/util/config"
	"fmt"
	"log"
)

func SavePassword(password string) bool {
	existingUserConfig, configReadErr := config.ReadConfigFile()
	if configReadErr != nil {
		log.Println(fmt.Sprintf("Error reading user config: %s", configReadErr.Error()))
	}

	existingUserConfig["userPassword"] = password
	configWriteErr := config.WriteConfigFile(existingUserConfig)
	if configWriteErr != nil {
		log.Println(fmt.Sprintf("Error writing user config: %s", configReadErr.Error()))
		return false
	}
	return true
}
