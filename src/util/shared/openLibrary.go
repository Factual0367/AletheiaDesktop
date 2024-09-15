package shared

import (
	"AletheiaDesktop/src/util/config"
	"log"
)

func OpenLibraryFolder() {
	currentLibraryFolder := config.GetCurrentDownloadFolder()
	err := OpenWithDefaultApp(currentLibraryFolder)
	if err != nil {
		log.Println("Failed to open library folder.")
	}
}
