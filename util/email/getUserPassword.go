package email

import (
	"AletheiaDesktop/util/config"
	"fmt"
	"log"
)

func GetUserPassword() string {
	existingUserConfig, configErr := config.ReadConfigFile()

	if configErr != nil {
		log.Println(fmt.Sprintf("Error reading config file: %s", configErr))
	}
	userPassword := existingUserConfig["userPassword"]

	userPasswordText := fmt.Sprintf("%s", userPassword)
	return userPasswordText
}
