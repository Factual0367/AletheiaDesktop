package configUtil

import (
	"encoding/json"
	"os"
)

func WriteConfigFile(content map[string]string) error {
	configPath, err := ConstructConfigLocation()
	fileData, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, fileData, 0644)
	return err
}
