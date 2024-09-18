package conversion

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/pkg/util/database"
	"fmt"
	"os/exec"
	"path"
	"strings"
)

func ConvertToFormat(targetFormat string, book models.Book) bool {
	existingFilepath := book.Filepath
	extension := path.Ext(existingFilepath)
	outfile := existingFilepath[0:len(existingFilepath)-len(extension)] + "." + strings.ToLower(targetFormat)
	cmd := exec.Command("ebook-convert", book.Filepath, outfile)
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Run(); err != nil {
		return false
	}

	// new book to add to the library
	newBook := book
	newBook.Filepath = outfile
	newBook.Extension = targetFormat
	newBook.ID = fmt.Sprintf("%s-%s", newBook.ID, targetFormat)

	database.UpdateDatabase(newBook, true, "downloaded")

	return true
}
