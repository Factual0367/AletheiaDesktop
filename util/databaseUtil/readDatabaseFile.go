package configUtil

import (
	"AletheiaDesktop/util/shared"
	"encoding/json"
	"fmt"
	"os"
)

func ReadDatabaseFile() (map[string]interface{}, error) {

	databasePath, databasePathConstructionErr := ConstructDatabaseLocation()
	if databasePathConstructionErr != nil {
		return nil, databasePathConstructionErr
	}

	exists, err := shared.Exists(databasePath)
	if err != nil {
		return nil, err
	}

	if exists {
		userData := map[string]interface{}{}
		file, readFileErr := os.ReadFile(databasePath)
		if readFileErr != nil {
			return nil, readFileErr
		}

		unmarshalErr := json.Unmarshal(file, &userData)
		if unmarshalErr != nil {
			return nil, fmt.Errorf("Unmarshal config file error: %v", unmarshalErr)
		}
		return userData, nil
	}

	return map[string]interface{}{}, nil
}
