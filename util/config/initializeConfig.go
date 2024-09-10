package config

import (
	"log"
	"os"
)

func InitializeConfig() {
	configLocation, err := ConstructConfigLocation()

	if err != nil {
		log.Fatalln("Unable to create config location")
		panic(err)
	}

	initialDownloadDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln("Could not get home directory")
		panic(err)
	}

	initialUserConfig := map[string]string{
		"downloadLocation": initialDownloadDir,
	}

	fileWriteErr := WriteConfigFile(configLocation, initialUserConfig)

	if fileWriteErr != nil {
		log.Fatalln("Unable to write config file")
		panic(fileWriteErr)
	}
}
