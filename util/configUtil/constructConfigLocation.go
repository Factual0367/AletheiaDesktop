package configUtil

import (
	"fmt"
	"os"
	"path"
)

const configFileName string = "aletheia.cfg"

func ConstructConfigLocation() (string, error) {
	homedir, err := os.UserConfigDir()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fileLocation := path.Join(homedir, configFileName)

	return fileLocation, err
}
