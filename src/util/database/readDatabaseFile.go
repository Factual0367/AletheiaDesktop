package database

import (
	"AletheiaDesktop/src/models"
	"AletheiaDesktop/src/util/shared"
	"encoding/json"
	"fmt"
	"log"
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

		unmarshalBooks := func(key string) error {
			if booksRaw, ok := userData[key]; ok {
				books := make(map[string]*models.Book)
				booksBytes, marshalErr := json.Marshal(booksRaw)
				if marshalErr != nil {
					return fmt.Errorf("Error marshaling %s: %v", key, marshalErr)
				}

				unmarshalErr = json.Unmarshal(booksBytes, &books)
				if unmarshalErr != nil {
					return fmt.Errorf("Unmarshal %s error: %v", key, unmarshalErr)
				}

				userData[key] = books
			}
			return nil
		}

		if err = unmarshalBooks("savedBooks"); err != nil {
			log.Println(fmt.Sprintf("Error unmarshaling saved books: %v", err))
			return map[string]interface{}{}, err
		}
		if err = unmarshalBooks("favoriteBooks"); err != nil {
			log.Println(fmt.Sprintf("Error unmarshaling favorite books: %v", err))
			return map[string]interface{}{}, err
		}

		return userData, nil
	}

	return map[string]interface{}{}, nil
}
