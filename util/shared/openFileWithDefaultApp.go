package shared

import (
	"github.com/skratchdot/open-golang/open"
)

func OpenBookWithDefaultApp(filepath string) error {
	err := open.Run(filepath)
	return err
}
