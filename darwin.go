// +build darwin

package gnotifier

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/deckarep/gosx-notifier"
	"github.com/everdev/mack"
)

func notify(app, title, message string, timeout time.Duration, image string) error {
	n := gosxnotifier.NewNotification(message)
	n.Title = title

	n.Group = fmt.Sprintf("com.%s.gnotifier.notify", strings.ToLower(app))
	if image != "" {
		n.AppIcon = image
	}

	err := n.Push()
	if _, ok := err.(*exec.ExitError); err == nil || !ok {
		return err
	}

	return mack.Notify(message, title)
}
