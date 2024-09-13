package search_test

import (
	"AletheiaDesktop/search"
	"bytes"
	"github.com/onurhanak/libgenapi"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestConstructFilename(t *testing.T) {
	book := search.Book{
		Book: libgenapi.Book{
			Author:    "John Doe",
			Title:     "Test Book",
			Extension: "pdf",
		},
	}

	expectedFilename := "John Doe - Test Book.pdf"
	filename := book.ConstructFilename()

	if filename != expectedFilename {
		t.Errorf("Expected filename to be %s, got %s", expectedFilename, filename)
	}
}

func TestSaveToFile(t *testing.T) {
	mockResponse := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte("file content"))),
	}

	tmpFile, err := ioutil.TempFile("", "testbook")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	book := search.Book{
		Filepath: tmpFile.Name(),
	}

	success := book.SaveToFile(mockResponse)
	if !success {
		t.Error("Expected SaveToFile to return true, but got false")
	}

	content, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := "file content"
	if string(content) != expectedContent {
		t.Errorf("Expected file content to be '%s', but got '%s'", expectedContent, string(content))
	}
}

func TestDownload(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("mock file content"))
	}))
	defer mockServer.Close()

	tmpDir := os.TempDir()
	defer os.RemoveAll(tmpDir)

	book := search.Book{
		Book: libgenapi.Book{
			Author:       "John Doe",
			Title:        "Test Book",
			Extension:    "pdf",
			DownloadLink: mockServer.URL,
		},
		Filepath: filepath.Join(tmpDir, "John Doe - Test Book.pdf"),
	}

	success := book.Download()
	if !success {
		t.Error("Expected download to succeed, but it failed")
	}

	content, err := os.ReadFile(filepath.Join(tmpDir, "John Doe - Test Book.pdf"))
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := "mock file content"
	if string(content) != expectedContent {
		t.Errorf("Expected file content to be '%s', but got '%s'", expectedContent, string(content))
	}
}
