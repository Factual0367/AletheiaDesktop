package database

import "AletheiaDesktop/internal/models"

func LoadFavoriteBooks() (map[string]*models.Book, error) {
	userData, err := ReadDatabaseFile()
	if err != nil || len(userData) == 0 {
		return nil, err
	}

	if favoriteBooks, ok := userData["favoriteBooks"].(map[string]*models.Book); ok {
		return favoriteBooks, nil
	}
	return nil, nil
}
