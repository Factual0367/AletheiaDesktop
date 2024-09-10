package configUtil

import (
	"encoding/json"
	"os"
)

func WriteDatabaseToFile(content map[string]interface{}) error {
	databasePath, err := ConstructDatabaseLocation()
	databaseData, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(databasePath, databaseData, 0644)
	return err
}
