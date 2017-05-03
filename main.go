package main

import (
	"math/rand"
	"runtime"
	"time"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 1200
	screenHeight = 800
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	runtime.LockOSThread()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Rectangle", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	var dt float32
	for window.IsOpen() {
		start := time.Now()

		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}

		window.Clear(sf.ColorWhite)
		window.Display()

		elapsed := time.Since(start)
		dt = float32(elapsed) / float32(time.Second)
	}
}
