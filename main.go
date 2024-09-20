package main

import (
	"AletheiaDesktop/internal/ui/views"
	"AletheiaDesktop/pkg/util/cache"
	"AletheiaDesktop/pkg/util/config"
	"AletheiaDesktop/pkg/util/database"
	"AletheiaDesktop/pkg/util/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func checkFirstRun() {
	defaultConfigPath, _ := config.ConstructConfigLocation()
	exists, err := shared.Exists(defaultConfigPath)
	if !exists || err != nil {
		config.InitializeConfig()
		database.InitializeDatabase()
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Aletheia")
	checkFirstRun()
	cache.CreateCacheDir()

	tabs := container.NewAppTabs()

	searchView := views.CreateSearchView(myApp, myWindow)
	settingsView := views.CreateSettingsView(myApp)
	libraryView := views.CreateLibraryView(myApp, myWindow, tabs)
	bookmarksView := views.CreateBookmarksView(myApp, myWindow, tabs)
	downloadsView := views.CreateDownloadsView()

	tabs = container.NewAppTabs(
		searchView,
		libraryView,
		bookmarksView,
		downloadsView,
		settingsView,
	)

	// this is necessary to refresh library view
	// when the user downloads a book
	tabs.OnSelected = func(tab *container.TabItem) {
		if tab.Icon == theme.StorageIcon() {
			libraryView = views.CreateLibraryView(myApp, myWindow, tabs)
			tabs.Items[1] = libraryView
			tabs.Refresh()
		} else if tab.Icon == theme.ContentAddIcon() {
			bookmarksView = views.CreateBookmarksView(myApp, myWindow, tabs)
			tabs.Items[2] = bookmarksView
			tabs.Refresh()
		} else if tab.Icon == theme.DownloadIcon() {
			downloadsView = views.CreateDownloadsView()
			tabs.Items[3] = downloadsView
			tabs.Refresh()
			views.StartDownloadsAutoRefresh(tabs)
		}
	}

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(1200, 800))
	myWindow.ShowAndRun()
}
