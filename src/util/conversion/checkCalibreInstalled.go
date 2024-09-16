package conversion

import (
	"fmt"
	"log"
	"os/exec"
)

func CheckCalibreInstalled() bool {
	cmd := exec.Command("ebook-convert", "--version")
	if err := cmd.Run(); err != nil {
		log.Println(fmt.Sprintf("Calibre is not installed."))
		return false
	}
	return true
}
