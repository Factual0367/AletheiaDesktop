package search

import (
	"AletheiaDesktop/src/models"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/onurhanak/libgenapi"
)

func TestConstructFilename(t *testing.T) {
	book := &models.Book{
		Book: libgenapi.Book{
			Author:    "John/Smith",
			Title:     "Go Programming: A Comprehensive Guide?",
			Extension: "pdf",
		},
	}

	expectedFilename := "John_Smith - Go Programming_ A Comprehensive Guide_.pdf"
	filename := book.ConstructFilename()

	if filename != expectedFilename {
		t.Errorf("Expected filename '%s', got '%s'", expectedFilename, filename)
	}
}

func TestSaveToFile(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "savefile")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	book := &models.Book{
		Filepath: filepath.Join(tempDir, "testfile.txt"),
	}

	bodyContent := "Test content for SaveToFile method."
	response := &http.Response{
		Body: io.NopCloser(strings.NewReader(bodyContent)),
	}

	success := book.SaveToFile(response)
	if !success {
		t.Error("Expected SaveToFile to return true, got false")
	}

	data, err := ioutil.ReadFile(book.Filepath)
	if err != nil {
		t.Fatalf("Failed to read file '%s': %v", book.Filepath, err)
	}

	if string(data) != bodyContent {
		t.Errorf("Expected file content '%s', got '%s'", bodyContent, string(data))
	}
}

func TestDownload(t *testing.T) {
	serverContent := "Test content for Download method."
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(serverContent))
	}))
	defer ts.Close()

	tempDir, err := ioutil.TempDir("", "download")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	book := &models.Book{
		Book: libgenapi.Book{
			Author:       "Bob",
			Title:        "Mastering Go",
			Extension:    "pdf",
			DownloadLink: ts.URL,
		},
		DownloadFolder: tempDir,
	}

	success := book.Download()
	if !success {
		t.Error("Expected Download to return true, got false")
	}

	data, err := ioutil.ReadFile(book.Filepath)
	if err != nil {
		t.Fatalf("Failed to read file '%s': %v", book.Filepath, err)
	}

	if string(data) != serverContent {
		t.Errorf("Expected file content '%s', got '%s'", serverContent, string(data))
	}
}
