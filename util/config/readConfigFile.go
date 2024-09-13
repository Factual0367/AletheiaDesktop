package config

import (
	"AletheiaDesktop/util/shared"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ReadConfigFile() (map[string]string, error) {

	configPath, configPathConstructionErr := ConstructConfigLocation()

	if configPathConstructionErr != nil {

		log.Println(configPathConstructionErr.Error())
	}

	userConfigContent := map[string]string{}
	file, readFileErr := os.ReadFile(configPath)
	if readFileErr != nil {
		shared.SendNotification("Error", "Cannot read your config file. Please delete aletheia.cfg and start over.")
		log.Println(fmt.Errorf("Read config file error: %v", readFileErr))
		return nil, readFileErr
	}

	unmarshalErr := json.Unmarshal(file, &userConfigContent)
	if unmarshalErr != nil {
		shared.SendNotification("Error", "Cannot read your config file. Please delete aletheia.cfg and start over.")
		log.Println(fmt.Errorf("Unmarshal config file error: %v", unmarshalErr))
		return nil, unmarshalErr
	}
	return userConfigContent, nil
}
