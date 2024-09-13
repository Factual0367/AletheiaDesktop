package config

import (
	"AletheiaDesktop/util/shared"
	"log"
	"os"
)

func InitializeConfig() {
	// add checks
	initialDownloadDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln("Could not get home directory")
		panic(err)
	}

	initialUserConfig := map[string]string{
		"downloadLocation": initialDownloadDir,
		"userEmail":        "",
		"userPassword":     "",
	}

	fileWriteErr := WriteConfigFile(initialUserConfig)

	if fileWriteErr != nil {
		log.Println("Unable to write config file")
		shared.SendNotification("Error", "Aletheia cannot create a configuration file.")
		panic(fileWriteErr)
	}
}
