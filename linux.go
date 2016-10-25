// +build linux

package gnotifier

import (
	"os/exec"
	"time"
)

var path string
var pathErr error

func init() {
	path, pathErr = exec.LookPath("notify-send")
}

func notify(title, message string, timeout time.Duration) error {
	if pathErr != nil {
		return pathErr
	}

	return exec.Command(path, title, message).Run()
}
