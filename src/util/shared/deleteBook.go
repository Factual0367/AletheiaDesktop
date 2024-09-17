package shared

import (
	"AletheiaDesktop/src/models"
	"log"
	"os"
)

func DeleteBook(book models.Book) {
	// also needs to delete from the user json file
	// or set downloaded to False
	err := os.Remove(book.Filepath)
	if err != nil {
		log.Println("Unable to delete file", book.Filepath)
	}
}
