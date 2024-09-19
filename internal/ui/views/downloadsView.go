package views

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/pkg/util/downloads"
	"log"
	"sort"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

var downloadRefreshTicker *time.Ticker

func RefreshDownloadsTab(tabs *container.AppTabs) {
	if tabs.Selected().Text == "Downloads" {
		newDownloadsView := CreateDownloadsView()
		tabs.Items[3] = newDownloadsView
		tabs.SelectIndex(3)
		tabs.Refresh()
	}
}

func StopDownloadsAutoRefresh() {
	if downloadRefreshTicker != nil {
		downloadRefreshTicker.Stop()
		downloadRefreshTicker = nil
	}
}

func StartDownloadsAutoRefresh(tabs *container.AppTabs) {
	if downloadRefreshTicker != nil {
		downloadRefreshTicker.Stop()
	}
	downloadRefreshTicker = time.NewTicker(2 * time.Second)
	go func() {
		for range downloadRefreshTicker.C {
			RefreshDownloadsTab(tabs)
		}
	}()
}

var (
	previousDownloadSize  int = 0
	previousDownloadBooks models.BookSlice
)

func CreateDownloadsView() *container.TabItem {
	downloadsViewContainer := container.NewVBox()
	shouldStopRefreshing := true

	currentDownloadSize := len(downloads.InProgressDownloads)

	// rebuild and sort bookSlice if the number of in-progress downloads has changed
	if currentDownloadSize != previousDownloadSize {
		log.Println("Number of activa downloads changed, rebuilding book slice.")
		previousDownloadSize = currentDownloadSize

		previousDownloadBooks = make(models.BookSlice, 0, currentDownloadSize)
		for _, book := range downloads.InProgressDownloads {
			previousDownloadBooks = append(previousDownloadBooks, book)
		}
		// srting to maintain order in the GUI
		sort.Sort(previousDownloadBooks)
	}

	for _, book := range previousDownloadBooks {
		if book.DownloadProgress < 1 {
			shouldStopRefreshing = false
		}
		bookDownloadsContainer := CreateBookDownloadsContainer(book)
		downloadsViewContainer.Add(bookDownloadsContainer)
	}

	if shouldStopRefreshing {
		StopDownloadsAutoRefresh()
	}

	return container.NewTabItemWithIcon("Downloads", theme.DownloadIcon(), downloadsViewContainer)
}
