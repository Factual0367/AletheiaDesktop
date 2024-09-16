package downloads

import "AletheiaDesktop/src/search"

var inProgressDownloads = make(map[string]search.Book)

func CheckInProgressDownloads(book search.Book) bool {
	_, ok := inProgressDownloads[book.ID]

	if ok {
		return true
	}
	inProgressDownloads[book.ID] = book
	return false
}
