package views

import (
	"AletheiaDesktop/util/config"
	"AletheiaDesktop/util/shared"
	"fmt"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateSettingsView() *container.TabItem {
	downloadDir := config.GetCurrentDownloadFolder()
	currentLibraryLocationMsg := "Current Library Location: "
	currentDownloadDirMsg := fmt.Sprintf("%s \n%s", currentLibraryLocationMsg, downloadDir)
	downloadDirLabel := widget.NewLabel(currentDownloadDirMsg)

	changeDownloadLocation := widget.NewButtonWithIcon("Change Library Location", theme.FolderIcon(), func() {
		newDownloadDir := shared.GetFolder()
		config.UpdateDownloadPath(newDownloadDir)
		downloadDirLabel.SetText(fmt.Sprintf("%s \n%s", currentLibraryLocationMsg, newDownloadDir))
	})

	padding := widget.NewLabel("")

	settingsInnerContent := container.NewGridWithRows(4, padding, downloadDirLabel, changeDownloadLocation)
	settingsContent := container.NewVBox(
		padding,
		container.NewGridWithColumns(3, padding, settingsInnerContent, padding),
	)
	settingsContentBordered := container.NewBorder(nil, nil, nil, nil, settingsContent)

	return container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), settingsContentBordered)
}
