package database

import (
	"AletheiaDesktop/internal/models"
)

func LoadSavedBooks() (map[string]*models.Book, error) {
	userData, err := ReadDatabaseFile()
	if err != nil {
		return nil, err
	}

	if savedBooks, ok := userData["savedBooks"].(map[string]*models.Book); ok {
		return savedBooks, nil
	}
	return nil, nil
}
