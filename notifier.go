package gnotifier

import "time"

func Notify(title, message string, timeout time.Duration) error {
	return notify(title, message, timeout)
}
