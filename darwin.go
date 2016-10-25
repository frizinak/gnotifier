// +build darwin

package gnotifier

import (
	"math"
	"strconv"
	"time"

	"github.com/everdev/mack"
)

func notify(title, message string, timeout time.Duration) error {
	_, err := mack.Alert(
		title,
		message,
		"informational",
		strconv.Itoa(int(math.Max(1.0, timeout.Seconds()))),
	)

	return err
}
