package shared

import (
	"github.com/gen2brain/beeep"
	"log"
)

func SendNotification(notificationHeader, notificationContent string) {
	ok := beeep.Notify(notificationHeader, "Downloaded successfully", "")
	if ok != nil {
		log.Println("Could not send notification.")
	}
}
