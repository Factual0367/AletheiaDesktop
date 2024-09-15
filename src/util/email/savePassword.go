package email

import (
	config2 "AletheiaDesktop/src/util/config"
	"fmt"
	"log"
)

func SavePassword(password string) bool {
	existingUserConfig, configReadErr := config2.ReadConfigFile()
	if configReadErr != nil {
		log.Println(fmt.Sprintf("Error reading user config: %s", configReadErr.Error()))
	}

	existingUserConfig["userPassword"] = password
	configWriteErr := config2.WriteConfigFile(existingUserConfig)
	if configWriteErr != nil {
		log.Println(fmt.Sprintf("Error writing user config: %s", configReadErr.Error()))
		return false
	}
	return true
}
