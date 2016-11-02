// +build windows

package gnotifier

import (
	"time"

	"github.com/go-toast/toast"
)

func notify(app, title, message string, timeout time.Duration, image string) error {
	n := toast.Notification{
		AppID:    app,
		Title:    title,
		Message:  message,
		Duration: toast.Short,
	}

	// https://blogs.msdn.microsoft.com/tiles_and_toasts/2015/07/02/adaptive-and-interactive-toast-notifications-for-windows-10/
	// short = 7s
	// long = 25s
	if timeout > time.Duration(time.Second*16) {
		n.Duration = toast.Long
	}

	if image != "" {
		n.Icon = image
	}

	return n.Push()
}
