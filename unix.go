// +build linux openbsd freebsd netbsd

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

func notify(app, title, message string, timeout time.Duration, image string) error {
	if pathErr != nil {
		return pathErr
	}

	args := make([]string, 2, 8)
	args[0] = "-a"
	args[1] = app

	if ms := int(timeout.Nanoseconds()) / 1e6; ms > 0 {
		args = append(args, "-t")
		args = append(args, strconv.Itoa(ms))
	}

	if image != "" {
		args = append(args, "-i")
		args = append(args, image)
	}

	args = append(args, title)
	args = append(args, message)

	return exec.Command(path, args...).Run()
}
