package main

import (
	"AletheiaDesktop/ui/views"
	"AletheiaDesktop/util/config"
	"AletheiaDesktop/util/database"
	"AletheiaDesktop/util/shared"

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

	tabs := container.NewAppTabs()

	searchView := views.CreateSearchView()
	settingsView := views.CreateSettingsView()
	libraryView := views.CreateLibraryView(myWindow, tabs)
	bookmarksView := views.CreateBookmarksView(myWindow, tabs)

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
			libraryView = views.CreateLibraryView(myWindow, tabs)
			tabs.Items[1] = libraryView
			tabs.Refresh()
		} else if tab.Icon == theme.ContentAddIcon() {
			bookmarksView = views.CreateBookmarksView(myWindow, tabs)
			tabs.Items[2] = bookmarksView
			tabs.Refresh()
		}
	}

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(1200, 800))
	myWindow.ShowAndRun()
}
