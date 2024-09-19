package views

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/pkg/util/downloads"
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

func CreateDownloadsView() *container.TabItem {
	downloadsViewContainer := container.NewVBox()
	shouldStopRefreshing := true

	bookSlice := make(models.BookSlice, 0, len(downloads.InProgressDownloads))
	for _, book := range downloads.InProgressDownloads {
		bookSlice = append(bookSlice, book)
	}
	// sorting is necessasry here for the time being
	// to keep order in the gui
	sort.Sort(bookSlice)

	for _, book := range bookSlice {
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
