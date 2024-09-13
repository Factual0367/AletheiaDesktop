// conversion_test.go
package conversion

import (
	"AletheiaDesktop/search"
	"github.com/onurhanak/libgenapi"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestConvertToFormat(t *testing.T) {
	_, err := exec.LookPath("ebook-convert")
	if err != nil {
		t.Skip("ebook-convert not found in PATH; skipping test.")
	}

	tempDir, err := ioutil.TempDir("", "conversion_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	inputFile := filepath.Join(tempDir, "testbook.txt")
	inputContent := "This is a test book content."
	err = ioutil.WriteFile(inputFile, []byte(inputContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write input file: %v", err)
	}

	book := search.Book{
		Filepath: inputFile,
		Book: libgenapi.Book{
			ID: "1234",
		},
	}

	targetFormat := "epub"
	success := ConvertToFormat(targetFormat, book)

	if !success {
		t.Errorf("ConvertToFormat returned false; expected true")
	}

	outputFile := inputFile[0:len(inputFile)-len(filepath.Ext(inputFile))] + "." + targetFormat
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Errorf("Expected output file %s to exist, but it does not.", outputFile)
	} else {
		defer os.Remove(outputFile)
	}

}
