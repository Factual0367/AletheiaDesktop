package configUtil

import (
	"fmt"
	"os"
	"path"
)

const databaseFilename string = "aletheia.json"

func ConstructDatabaseLocation() (string, error) {
	homedir, err := os.UserConfigDir()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fileLocation := path.Join(homedir, databaseFilename)

	return fileLocation, err
}
