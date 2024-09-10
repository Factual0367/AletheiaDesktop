package config

import (
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
	}

	fileWriteErr := WriteConfigFile(initialUserConfig)

	if fileWriteErr != nil {
		log.Fatalln("Unable to write config file")
		panic(fileWriteErr)
	}
}
