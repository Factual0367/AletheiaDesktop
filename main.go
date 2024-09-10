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

	/*defaultBook := libgenapi.Book{
		ID:           "",
		MD5:          "",
		Title:        "Select a book to view details",
		Author:       "",
		Publisher:    "",
		Year:         "",
		Language:     "",
		Pages:        "",
		Size:         "",
		Extension:    "",
		DownloadLink: "",
		CoverLink:    "https://cdn.pixabay.com/photo/2013/07/13/13/34/book-161117_960_720.png",
	} */

	searchPage := views.CreateSearchView()
	//settingsPage := ui.CreateSettingsTab()
	//libraryPage := ui.CreateLibraryView()
	tabs := container.NewAppTabs(
		searchPage,
		//libraryPage,
		//settingsPage,
	)

	tabs.SetTabLocation(container.TabLocationTop)

	mainContainer := container.NewHSplit(container.NewVBox(), tabs)
	mainContainer.SetOffset(0.25)

	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(900, 600)) // Set a decent window size
	myWindow.ShowAndRun()
}
