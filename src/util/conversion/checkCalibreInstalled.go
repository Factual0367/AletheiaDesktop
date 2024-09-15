package conversion

import (
	"fmt"
	"os/exec"
)

func CheckCalibreInstalled() bool {
	cmd := exec.Command("ebook-convert", "--version")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
