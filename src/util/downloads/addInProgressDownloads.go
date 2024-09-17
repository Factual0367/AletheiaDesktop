package downloads

import (
	"AletheiaDesktop/src/models"
)

var InProgressDownloads = make(map[string]*models.Book)

func AddInProgressDownloads(book *models.Book) bool {

	_, ok := InProgressDownloads[book.ID]

	if ok {
		return true
	}
	InProgressDownloads[book.ID] = book

	return false
}
