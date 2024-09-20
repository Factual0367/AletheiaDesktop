package components

import (
	"AletheiaDesktop/internal/models"
	"AletheiaDesktop/pkg/util/email"
	"AletheiaDesktop/pkg/util/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateEmailButton(myApp fyne.App, book models.Book) *widget.Button {
	var emailButton *widget.Button
	emailButton = widget.NewButtonWithIcon("Email", theme.MailSendIcon(), func() {
		go func() {
			emailSent := email.SendBookEmail(book)
			if emailSent {
				shared.SendNotification(myApp, "Success", "Your book is emailed successfully.")
				emailButton.SetIcon(theme.ConfirmIcon())
			} else {
				shared.SendNotification(myApp, "Failed", "Your book could not be emailed. Check your credentials.")
				emailButton.SetIcon(theme.ErrorIcon())
			}
		}()
	})
	return emailButton
}
