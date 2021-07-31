package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

const WindowWidth = 1280
const WindowHeight = 720

func main() {

	// Initialize SDL components
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal("error initializing ", err)
	}

	// Create window
	window, err := sdl.CreateWindow("Pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(WindowWidth), int32(WindowHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal("error creating window ", err)
	}

	//Create renderer
	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		log.Fatal("error creating renderer ", err)
	}

	// Create the ball
	ball := NewBall(WindowWidth/2.0,
		WindowHeight/2.0)

	// Draw the ball
	renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
	ball.Draw(renderer)
	renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

	// Present the backbuffer
	renderer.Present()

	// Color the window
	if err := renderer.SetDrawColor(0x0, 0x0, 0x0, 0xFF); err != nil {
		return
	}
	//renderer.Clear()

	// Set the draw color to be white
	renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

	// Draw the net
	var y int32
	for y = 0; y < WindowHeight; y++ {
		if y%5 == 0 {
			if err := renderer.DrawPoint(WindowWidth/2, y); err != nil {
				log.Fatal("failed to draw point!")
			}
		}
	}
	renderer.Present()

	//Clean up
	defer window.Destroy()
	defer renderer.Destroy()

	// Game Loop
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
	}
}
