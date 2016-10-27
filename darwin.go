// +build darwin

package gnotifier

import (
	"os/exec"
	"time"

	"github.com/deckarep/gosx-notifier"
	"github.com/everdev/mack"
)

func notify(title, message string, timeout time.Duration) error {
	n := gosxnotifier.NewNotification(message)
	n.Title = title
	// TODO
	// n.Group = ""
	// n.AppIcon
	err := n.Push()
	if _, ok := err.(*exec.ExitError); err == nil || !ok {
		return err
	}

	return mack.Notify(message, title)
}
