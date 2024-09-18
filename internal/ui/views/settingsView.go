package views

import (
	config2 "AletheiaDesktop/pkg/util/config"
	email2 "AletheiaDesktop/pkg/util/email"
	shared2 "AletheiaDesktop/pkg/util/shared"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createDownloadLocationContainer() *fyne.Container {
	currentLibraryLocationMsg := "Current Library Location: "
	downloadDir := config2.GetCurrentDownloadFolder()

	currentDownloadDirLabel := widget.NewLabelWithStyle(currentLibraryLocationMsg, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	downloadDirLabel := widget.NewLabel(downloadDir)

	changeDownloadLocationButton := widget.NewButtonWithIcon("Change Library Location", theme.FolderIcon(), func() {
		newDownloadDir := shared2.GetFolder()
		if newDownloadDir != "" {
			config2.UpdateDownloadPath(newDownloadDir)
			downloadDirLabel.SetText(fmt.Sprintf("%s \n%s", currentLibraryLocationMsg, newDownloadDir))
		}
	})

	return container.NewVBox(currentDownloadDirLabel, downloadDirLabel, changeDownloadLocationButton)
}

func createEmailContainer() *fyne.Container {
	emailLabel := widget.NewLabelWithStyle("Email", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	emailEntry := widget.NewEntry()
	emailEntry.PlaceHolder = email2.GetUserEmail()

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.PlaceHolder = "Your app password"

	saveEmailButton := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		emailSaved := email2.SaveEmail(emailEntry.Text)
		passwordSaved := email2.SavePassword(passwordEntry.Text)
		if emailSaved && passwordSaved {
			shared2.SendNotification("Success", "Your email and password have been saved.")
		} else {
			shared2.SendNotification("Error", "Failed to save your email or password.")
		}
	})

	return container.NewVBox(emailLabel, emailEntry, passwordEntry, saveEmailButton)
}

func CreateSettingsView() *container.TabItem {

	downloadLocationContainer := createDownloadLocationContainer()
	emailContainer := createEmailContainer()

	padding := widget.NewLabel("")

	settingsInnerContent := container.NewGridWithRows(3, padding, downloadLocationContainer)
	settingsContent := container.NewVBox(
		container.NewGridWithColumns(3, padding, settingsInnerContent, padding),
		container.NewGridWithColumns(3, padding, emailContainer, padding),
	)
	settingsContentBordered := container.NewBorder(nil, nil, nil, nil, settingsContent)

	return container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), settingsContentBordered)
}
