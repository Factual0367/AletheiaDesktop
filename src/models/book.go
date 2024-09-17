package models

import (
	"AletheiaDesktop/src/util/cache"
	"AletheiaDesktop/src/util/config"
	"fmt"
	"github.com/onurhanak/libgenapi"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Book struct {
	libgenapi.Book   // extend libgenapi.Book
	Filename         string
	Filepath         string
	CoverPath        string
	Downloaded       bool
	DownloadFolder   string // for tests
	DownloadProgress float64
}

func (book *Book) ConstructFilename() string {
	filename := fmt.Sprintf("%s - %s.%s", book.Author, book.Title, book.Extension)

	regex := regexp.MustCompile(`[\/\\:\*\?"<>\|]`)
	filename = regex.ReplaceAllString(filename, "_")

	filename = strings.TrimSpace(filename)
	book.Filename = filename
	return filename
}

func (book *Book) ConstructFilepath() string {
	downloadPath := config.GetCurrentDownloadFolder()
	filepath := filepath.Join(downloadPath, book.ConstructFilename())
	book.Filepath = filepath
	return filepath
}

func (book *Book) ConstructCoverPath() string {
	aletheiaCachePath := cache.GetAletheiaCache()
	book.CoverPath = filepath.Join(aletheiaCachePath, book.ID)
	return book.CoverPath
}

func (book *Book) SaveToFile(response *http.Response) bool {
	// create the file
	out, err := os.Create(book.Filepath)
	if err != nil {
		log.Println(fmt.Sprintf("Could not create a file for the book: %s", err))
		return false
	}
	defer out.Close()

	// copy the contents to the
	// newly created file
	_, err = io.Copy(out, response.Body)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (book *Book) Download() bool {
	const chunkSize = 32 * 1024 // 32 KB

	book.ConstructFilepath()
	response, err := http.Get(book.DownloadLink)
	if err != nil {
		// try alternative download link
		log.Println("Trying alternative download link")
		altDownloadLinkErr := book.AddSecondDownloadLink()
		response, err = http.Get(book.AlternativeDownloadLink)
		if err != nil || altDownloadLinkErr != nil {
			log.Println("First download link and second download link did not work.")
			return false
		}
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println(fmt.Errorf("failed to download file: %s", response.Status))
		return false
	}

	outFile, err := os.Create(book.Filepath)
	if err != nil {
		log.Println("Failed to create file:", err)
		return false
	}
	defer outFile.Close()

	totalSize, _ := strconv.Atoi(response.Header.Get("Content-Length"))
	book.DownloadProgress = 0

	buf := make([]byte, chunkSize)
	var downloaded int

	for {
		n, err := response.Body.Read(buf)
		if n > 0 {
			_, writeErr := outFile.Write(buf[:n])
			if writeErr != nil {
				log.Println("Failed to write to file:", writeErr)
				return false
			}

			downloaded += n
			book.DownloadProgress = float64(downloaded) / float64(totalSize)
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error during download:", err)
			return false
		}
	}

	book.Downloaded = true
	return book.Downloaded
}
