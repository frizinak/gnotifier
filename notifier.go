package gnotifier

import (
	"path/filepath"
	"time"
)

func Notify(app, title, message string, timeout time.Duration, image string) error {
	if image != "" && !filepath.IsAbs(image) {
		var err error
		if image, err = filepath.Abs(image); err != nil {
			return err
		}
	}

	return notify(app, title, message, timeout, image)
}
