package database

import (
	"fmt"
	"log"
	"os"
	"path"
)

const databaseFilename string = "aletheia.json"

func ConstructDatabaseLocation() (string, error) {
	homedir, err := os.UserConfigDir()

	if err != nil {
		log.Println(fmt.Sprintf("Error getting user home directory: %s", err))
		return "", err
	}

	fileLocation := path.Join(homedir, databaseFilename)

	return fileLocation, err
}
