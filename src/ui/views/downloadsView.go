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

func CreateDownloadsView(appWindow fyne.Window, tabs *container.AppTabs) *container.TabItem {
	downloadsViewContainer := container.NewVBox()
	for _, book := range downloads.InProgressDownloads {
		bookDownloadsContainer := CreateBookDownloadsContainer(book)
		downloadsViewContainer.Add(bookDownloadsContainer)
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

func StopDownloadsAutoRefresh() {
	if downloadRefreshTicker != nil {
		downloadRefreshTicker.Stop()
		downloadRefreshTicker = nil
	}
}
