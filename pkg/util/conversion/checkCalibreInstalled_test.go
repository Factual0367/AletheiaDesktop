package conversion

import (
	"os/exec"
	"testing"
)

func TestCheckCalibreInstalled(t *testing.T) {
	_, err := exec.LookPath("ebook-convert")
	expectedInstalled := err == nil

	actualInstalled := CheckCalibreInstalled()

	if expectedInstalled != actualInstalled {
		t.Errorf("Expected CheckCalibreInstalled to return %v, got %v", expectedInstalled, actualInstalled)
	}
}
