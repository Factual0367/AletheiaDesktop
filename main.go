package main

import (
	views2 "AletheiaDesktop/src/ui/views"
	"AletheiaDesktop/src/util/cache"
	config2 "AletheiaDesktop/src/util/config"
	"AletheiaDesktop/src/util/database"
	"AletheiaDesktop/src/util/shared"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func checkFirstRun() {
	defaultConfigPath, _ := config2.ConstructConfigLocation()
	exists, err := shared.Exists(defaultConfigPath)
	if !exists || err != nil {
		config2.InitializeConfig()
		database.InitializeDatabase()
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Aletheia")
	checkFirstRun()
	cache.CreateCacheDir()

	tabs := container.NewAppTabs()

	searchView := views2.CreateSearchView()
	settingsView := views2.CreateSettingsView()
	libraryView := views2.CreateLibraryView(myWindow, tabs)
	bookmarksView := views2.CreateBookmarksView(myWindow, tabs)

	tabs = container.NewAppTabs(
		searchView,
		libraryView,
		bookmarksView,
		settingsView,
	)

	// this is necessary to refresh library view
	// when the user downloads a book
	tabs.OnSelected = func(tab *container.TabItem) {
		if tab.Icon == theme.StorageIcon() {
			libraryView = views2.CreateLibraryView(myWindow, tabs)
			tabs.Items[1] = libraryView
			tabs.Refresh()
		} else if tab.Icon == theme.ContentAddIcon() {
			bookmarksView = views2.CreateBookmarksView(myWindow, tabs)
			tabs.Items[2] = bookmarksView
			tabs.Refresh()
		}
	}

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(1200, 800))
	myWindow.ShowAndRun()
}
