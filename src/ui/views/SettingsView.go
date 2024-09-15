package views

import (
	config2 "AletheiaDesktop/src/util/config"
	email2 "AletheiaDesktop/src/util/email"
	shared2 "AletheiaDesktop/src/util/shared"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createDownloadLocationContainer() *fyne.Container {
	downloadDir := config2.GetCurrentDownloadFolder()
	currentLibraryLocationMsg := "Current Library Location: "
	currentDownloadDirLabel := widget.NewLabel(currentLibraryLocationMsg)
	currentDownloadDirLabel.TextStyle = fyne.TextStyle{Bold: true}
	downloadDirLabel := widget.NewLabel(downloadDir)

	changeDownloadLocationButton := widget.NewButtonWithIcon("Change Library Location", theme.FolderIcon(), func() {
		newDownloadDir := shared2.GetFolder()
		config2.UpdateDownloadPath(newDownloadDir)
		downloadDirLabel.SetText(fmt.Sprintf("%s \n%s", currentLibraryLocationMsg, newDownloadDir))
	})

	downloadLocationContainer := container.NewVBox(currentDownloadDirLabel, downloadDirLabel, changeDownloadLocationButton)
	return downloadLocationContainer
}

func createEmailContainer() *fyne.Container {
	emailLabel := widget.NewLabel("Email")
	emailLabel.TextStyle = fyne.TextStyle{Bold: true}

	emailEntry := widget.NewEntry()
	userEmail := email2.GetUserEmail()
	emailEntry.PlaceHolder = userEmail
	userPassword := widget.NewPasswordEntry()

	saveEmailButton := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		emailSaved := email2.SaveEmail(emailEntry.Text)
		passwordSaved := email2.SavePassword(userPassword.Text)
		if emailSaved && passwordSaved {
			shared2.SendNotification("Success", "Your email has been saved.")
		} else {
			shared2.SendNotification("Error", "Failed to save your email.")
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
