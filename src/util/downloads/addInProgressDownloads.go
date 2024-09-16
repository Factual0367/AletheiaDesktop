package downloads

import (
	"AletheiaDesktop/src/search"
)

var InProgressDownloads = make(map[string]search.Book)

func CheckInProgressDownloads(book search.Book) bool {

	_, ok := InProgressDownloads[book.ID]

	if ok {
		return true
	}
	InProgressDownloads[book.ID] = book

	return false
}
