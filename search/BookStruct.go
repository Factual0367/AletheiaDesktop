package search

import (
	"AletheiaDesktop/util/configUtil"
	"fmt"
	"github.com/onurhanak/libgenapi"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Book struct {
	libgenapi.Book // extend libgenapi.Book
	Filename       string
	Filepath       string
	Downloaded     bool
}

func (book *Book) constructFilename() string {
	filename := fmt.Sprintf("%s - %s.%s", book.Author, book.Title, book.Extension)
	filename = strings.ReplaceAll(filename, "/", "_")
	filename = strings.TrimSpace(filename)
	book.Filename = filename
	return filename
}

func (book *Book) constructFilepath() string {
	downloadPath := configUtil.GetCurrentDownloadFolder()
	filepath := filepath.Join(downloadPath, book.constructFilename())
	book.Filepath = filepath
	return filepath
}

func (book *Book) saveToFile(response *http.Response) bool {
	// create the file
	out, err := os.Create(book.Filepath)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	defer out.Close()

	// copy the contents to the
	// newly created file
	_, err = io.Copy(out, response.Body)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	return true
}

func (book *Book) download() bool {
	response, err := http.Get(book.DownloadLink)
	if err != nil {
		log.Println("Could not download book")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println(fmt.Errorf("failed to download file: %s", response.Status))
	}

	book.Downloaded = book.saveToFile(response)

	return book.Downloaded
}
