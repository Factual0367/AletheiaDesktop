package shared

import (
	"AletheiaDesktop/search"
	"AletheiaDesktop/util/database"
	"log"
	"os"
)

func DeleteBook(book search.Book) {
	// also needs to delete from the user json file
	// or set downloaded to False
	err := os.Remove(book.Filepath)
	if err != nil {
		log.Println("Unable to delete file", book.Filepath)
	}
	database.UpdateDatabase(book, false) // false to remove the book
}
