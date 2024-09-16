package views

import (
	"AletheiaDesktop/src/util/config"
	"AletheiaDesktop/src/util/email"
	"AletheiaDesktop/src/util/shared"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createDownloadLocationContainer() *fyne.Container {
	downloadDir := config.GetCurrentDownloadFolder()
	currentLibraryLocationMsg := "Current Library Location: "
	currentDownloadDirLabel := widget.NewLabel(currentLibraryLocationMsg)
	currentDownloadDirLabel.TextStyle = fyne.TextStyle{Bold: true}
	downloadDirLabel := widget.NewLabel(downloadDir)

	changeDownloadLocationButton := widget.NewButtonWithIcon("Change Library Location", theme.FolderIcon(), func() {
		newDownloadDir := shared.GetFolder()
		config.UpdateDownloadPath(newDownloadDir)
		downloadDirLabel.SetText(fmt.Sprintf("%s \n%s", currentLibraryLocationMsg, newDownloadDir))
	})

	downloadLocationContainer := container.NewVBox(currentDownloadDirLabel, downloadDirLabel, changeDownloadLocationButton)
	return downloadLocationContainer
}

func createEmailContainer() *fyne.Container {
	emailLabel := widget.NewLabel("Email")
	emailLabel.TextStyle = fyne.TextStyle{Bold: true}

	emailEntry := widget.NewEntry()
	userEmail := email.GetUserEmail()
	emailEntry.PlaceHolder = userEmail
	userPassword := widget.NewPasswordEntry()

	saveEmailButton := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		emailSaved := email.SaveEmail(emailEntry.Text)
		passwordSaved := email.SavePassword(userPassword.Text)
		if emailSaved && passwordSaved {
			shared.SendNotification("Success", "Your email has been saved.")
		} else {
			shared.SendNotification("Error", "Failed to save your email.")
		}
	})

	entryContainer := container.NewVBox(
		emailEntry,
		userPassword,
		saveEmailButton,
	)

	emailContainer := container.NewVBox(emailLabel, entryContainer)
	return emailContainer
}

func CreateSettingsView() *container.TabItem {

	downloadLocationContainer := createDownloadLocationContainer()
	emailContainer := createEmailContainer()

	padding := widget.NewLabel("")

	settingsInnerContent := container.NewGridWithRows(4, padding, downloadLocationContainer)
	settingsContent := container.NewVBox(
		container.NewGridWithColumns(3, padding, settingsInnerContent, padding),
		container.NewGridWithColumns(3, padding, emailContainer, padding),
	)
	settingsContentBordered := container.NewBorder(nil, nil, nil, nil, settingsContent)

	return container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), settingsContentBordered)
}
