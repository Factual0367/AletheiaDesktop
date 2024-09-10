package config

import "log"

func GetCurrentDownloadFolder() string {

	userConfigContent, configReadErr := ReadConfigFile()

	if configReadErr != nil {
		log.Fatalln("Could not read config file:", configReadErr)
	}

	currentDownloadFolder := userConfigContent["downloadFolder"]

	return currentDownloadFolder
}
