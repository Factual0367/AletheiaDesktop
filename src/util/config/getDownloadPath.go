package config

import (
	"log"
)

func GetCurrentDownloadFolder() string {

	userConfigContent, configReadErr := ReadConfigFile()

	if configReadErr != nil {
		log.Println("Created a new configuration file because config file is corrupted:", configReadErr)
	}

	currentDownloadFolder := userConfigContent["downloadLocation"]

	return currentDownloadFolder
}
