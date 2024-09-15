package cache

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getCoverImage(coverLink string) http.Response {
	resp, err := http.Get(coverLink)
	if err != nil {
		log.Println(fmt.Sprintf("Could not get cover for the book: %s", err))
	}
	return *resp
}

func SaveCoverImage(bookCoverLink string, bookCoverPath string) {
	resp := getCoverImage(bookCoverLink)
	coverImageFile, err := os.Create(bookCoverPath)
	if err != nil {
		panic(err)
	}
	defer coverImageFile.Close()

	_, err = io.Copy(coverImageFile, resp.Body)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to save cover image: %s", err))
	}

}
