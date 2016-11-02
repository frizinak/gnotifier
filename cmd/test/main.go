package main

import (
	"time"

	"github.com/frizinak/gnotifier"
)

func main() {
	err := gnotifier.Notify("test", "hey", "you", time.Second*5, "./icon.png")
	panic(err)
}
