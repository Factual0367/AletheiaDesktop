package shared

import (
	"github.com/skratchdot/open-golang/open"
)

func OpenWithDefaultApp(filepath string) error {
	err := open.Run(filepath)
	return err
}
