package shared

import (
	"fmt"
	"github.com/sqweek/dialog"
	"log"
)

func GetFolder() string {
	directory, err := dialog.Directory().Title("Select Folder").Browse()
	if err != nil {
		log.Println(fmt.Sprintf("Could not get folder path from directory selection."))
	}
	return directory
}
