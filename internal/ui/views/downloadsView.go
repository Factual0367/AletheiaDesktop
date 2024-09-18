package views

import (
	"AletheiaDesktop/pkg/util/downloads"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"time"
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

	for _, book := range downloads.InProgressDownloads {
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
