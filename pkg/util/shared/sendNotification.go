package shared

import (
	"fyne.io/fyne/v2"
)

func SendNotification(myApp fyne.App, notificationHeader, notificationContent string) {
	notification := fyne.NewNotification(notificationHeader, notificationContent)
	myApp.SendNotification(notification)
}
