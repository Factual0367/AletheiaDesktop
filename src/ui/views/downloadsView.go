package views

import (
	"AletheiaDesktop/src/util/downloads"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"time"
)

var downloadRefreshTicker *time.Ticker

func RefreshDownloadsTab(appWindow fyne.Window, tabs *container.AppTabs) {
	if tabs.Selected().Text == "Downloads" {
		newDownloadsView := CreateDownloadsView(appWindow, tabs)
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

func CreateDownloadsView(appWindow fyne.Window, tabs *container.AppTabs) *container.TabItem {
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

func StartDownloadsAutoRefresh(appWindow fyne.Window, tabs *container.AppTabs) {
	if downloadRefreshTicker != nil {
		downloadRefreshTicker.Stop()
	}

	downloadRefreshTicker = time.NewTicker(2 * time.Second)
	go func() {
		for range downloadRefreshTicker.C {
			RefreshDownloadsTab(appWindow, tabs)
		}
	}()
}
