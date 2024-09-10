package database

import (
	"AletheiaDesktop/search"
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
		file, readFileErr := os.ReadFile(databasePath)
		if readFileErr != nil {
			return nil, fmt.Errorf("Error reading database file: %v", readFileErr)
		}

		if len(file) == 0 {
			return nil, fmt.Errorf("Database file is empty")
		}

		userData := map[string]interface{}{}
		unmarshalErr := json.Unmarshal(file, &userData)
		if unmarshalErr != nil {
			return nil, fmt.Errorf("Unmarshal config file error: %v", unmarshalErr)
		}

		if savedBooksRaw, ok := userData["savedBooks"]; ok {
			savedBooks := make(map[string]*search.Book)
			savedBooksBytes, marshalErr := json.Marshal(savedBooksRaw)
			if marshalErr != nil {
				return nil, fmt.Errorf("Error marshaling savedBooks: %v", marshalErr)
			}

			unmarshalErr = json.Unmarshal(savedBooksBytes, &savedBooks)
			if unmarshalErr != nil {
				return nil, fmt.Errorf("Unmarshal savedBooks error: %v", unmarshalErr)
			}

			userData["savedBooks"] = savedBooks
		}

		return userData, nil
	}

	return map[string]interface{}{}, nil
}
