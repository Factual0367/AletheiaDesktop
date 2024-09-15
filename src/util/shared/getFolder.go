package shared

import (
	"fmt"
	"github.com/sqweek/dialog"
)

func GetFolder() string {
	directory, err := dialog.Directory().Title("Select Folder").Browse()
	if err != nil {
		fmt.Println(err)
	}
	return directory
}
