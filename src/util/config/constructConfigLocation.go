package config

import (
	"fmt"
	"log"
	"os"
	"path"
)

const configFileName string = "aletheia.cfg"

func ConstructConfigLocation() (string, error) {
	homedir, err := os.UserConfigDir()

	if err != nil {
		log.Println(fmt.Sprintf("Error getting user config dir: %s", err))
		return "", err
	}

	fileLocation := path.Join(homedir, configFileName)

	return fileLocation, err
}
