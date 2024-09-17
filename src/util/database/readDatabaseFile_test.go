// database_test.go
package database

import (
	"AletheiaDesktop/src/models"
	"encoding/json"
	"github.com/onurhanak/libgenapi"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func createTempDatabaseFile(t *testing.T, content []byte) string {
	databasePath, err := ConstructDatabaseLocation()
	if err != nil {
		t.Fatalf("ConstructDatabaseLocation returned an error: %v", err)
	}

	dir := filepath.Dir(databasePath)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory %s: %v", dir, err)
	}

	err = ioutil.WriteFile(databasePath, content, 0644)
	if err != nil {
		t.Fatalf("Failed to write database file: %v", err)
	}

	return databasePath
}

func removeDatabaseFile(t *testing.T, path string) {
	err := os.Remove(path)
	if err != nil && !os.IsNotExist(err) {
		t.Errorf("Failed to remove database file: %v", err)
	}
}

func TestReadDatabaseFile_Success(t *testing.T) {
	userData := map[string]interface{}{
		"savedBooks": map[string]*models.Book{
			"book1": {
				Book: libgenapi.Book{
					ID:    "book1",
					Title: "Test Book 1",
				},
			},
		},
		"favoriteBooks": map[string]*models.Book{
			"book2": {
				Book: libgenapi.Book{
					ID:    "book2",
					Title: "Test Book 2",
				},
			},
		},
	}
	content, err := json.Marshal(userData)
	if err != nil {
		t.Fatalf("Failed to marshal userData: %v", err)
	}

	databasePath := createTempDatabaseFile(t, content)
	defer removeDatabaseFile(t, databasePath)

	result, err := ReadDatabaseFile()
	if err != nil {
		t.Fatalf("ReadDatabaseFile returned an error: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 keys in result, got %d", len(result))
	}

	savedBooks, ok := result["savedBooks"].(map[string]*models.Book)
	if !ok {
		t.Errorf("Expected savedBooks to be map[string]*search.Book, got %T", result["savedBooks"])
	} else {
		if len(savedBooks) != 1 {
			t.Errorf("Expected 1 saved book, got %d", len(savedBooks))
		}
		if book, exists := savedBooks["book1"]; !exists || book.Title != "Test Book 1" {
			t.Errorf("Saved book data does not match expected values")
		}
	}

	favoriteBooks, ok := result["favoriteBooks"].(map[string]*models.Book)
	if !ok {
		t.Errorf("Expected favoriteBooks to be map[string]*search.Book, got %T", result["favoriteBooks"])
	} else {
		if len(favoriteBooks) != 1 {
			t.Errorf("Expected 1 favorite book, got %d", len(favoriteBooks))
		}
		if book, exists := favoriteBooks["book2"]; !exists || book.Title != "Test Book 2" {
			t.Errorf("Favorite book data does not match expected values")
		}
	}
}

func TestReadDatabaseFile_EmptyFile(t *testing.T) {
	databasePath := createTempDatabaseFile(t, []byte{})
	defer removeDatabaseFile(t, databasePath)

	result, err := ReadDatabaseFile()
	if err == nil {
		t.Fatal("Expected an error due to empty database file, got nil")
	}

	if result != nil {
		t.Errorf("Expected result to be nil when an error occurs, got %v", result)
	}
}

func TestReadDatabaseFile_InvalidJSON(t *testing.T) {
	databasePath := createTempDatabaseFile(t, []byte("{invalid json}"))
	defer removeDatabaseFile(t, databasePath)

	result, err := ReadDatabaseFile()
	if err == nil {
		t.Fatal("Expected an error due to invalid JSON, got nil")
	}

	if result != nil {
		t.Errorf("Expected result to be nil when an error occurs, got %v", result)
	}
}
