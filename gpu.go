package main

// GPUMode different modes of the GPU
type GPUMode int

const (
	// HBLANK hblank
	HBLANK = 0
	// VBLANK vblank
	VBLANK = 1
	// OAM Accessing OAM
	OAM = 2
	// VRAM Accessing VRAM
	VRAM = 3
)

// GPU represents the gpu
type GPU struct {
	screen *Screen
	mmu    *MMU
	mode   GPUMode
	mclock uint16
	line   uint8
}

// Init initializes the GPU
func (gpu *GPU) Init() {
	gpu.screen.Init()
}

// RenderLine renders a scanline to the screen
func (gpu *GPU) RenderLine() {

}
