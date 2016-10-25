package main

import (
	"time"

	"github.com/frizinak/gnotifier"
)

func main() {
	err := gnotifier.Notify("hey", "you", time.Second*5)
	panic(err)
}
