// +build linux

package gnotifier

import (
	"os/exec"
	"strconv"
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

	args := make([]string, 2, 4)
	args[0] = title
	args[1] = message

	if ms := int(timeout.Nanoseconds()) / 1e6; ms >= 1 {
		args = append(args, "-t")
		args = append(args, strconv.Itoa(ms))
	}

	return exec.Command(path, args...).Run()
}
