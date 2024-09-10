package main

import (
	"AletheiaDesktop/ui/views"
	"AletheiaDesktop/util/config"
	"AletheiaDesktop/util/database"
	"AletheiaDesktop/util/shared"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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

	searchPage := views.CreateSearchView()
	//settingsPage := ui.CreateSettingsTab()
	//libraryPage := ui.CreateLibraryView()
	tabs := container.NewAppTabs(
		searchPage,
		//libraryPage,
		//settingsPage,
	)

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(900, 600)) // Set a decent window size
	myWindow.ShowAndRun()
}
