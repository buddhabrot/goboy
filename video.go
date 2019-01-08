package main

// Video draws the screen
type Video interface {
	Setup()
	Draw(screen *Screen)
}
