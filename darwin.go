// +build darwin

package gnotifier

import (
	"time"

	"github.com/everdev/mack"
)

func notify(title, message string, timeout time.Duration) error {
	return mack.Notify(message, title)
}
