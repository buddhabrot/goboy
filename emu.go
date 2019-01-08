package main

import (
	"log"
	"os"
	"time"
)

const (
	// HBLANKTiming time gpu/screen spends in HBLANK
	HBLANKTiming = 204
	// VBLANKTiming time gpu/screen spends in VBLANK for one line
	VBLANKTiming = 456
	// VBLANK10Timing time gpu/screen spends in VBLANK for ten lines
	VBLANK10Timing = 4560
	// FULLFRAMETiming time gpu/screen spends in VBLANK for ten lines
	FULLFRAMETiming = 70224
	// OAMTiming time gpu/screen spends in OAM
	OAMTiming = 80
	// VRAMTiming time gpu/screen spends in VRAM
	VRAMTiming = 172
)

// step: Emulates one step of the cpu
func step(z80 *Z80, mmu *MMU, gpu *GPU, video Video) bool {

	stepCPU(z80, mmu)
	stepDisplay(z80, mmu, gpu, video)

	return true
}

func stepCPU(z80 *Z80, mmu *MMU) {
	// CPU step
	z80.r.r &= 127
	z80.r.pc++
	var pc = z80.r.pc
	var nextOpCode = mmu.rb(pc)
	z80.Exec(nextOpCode, mmu)
	z80.r.pc &= 65535
	z80.c.m += z80.r.lc.m
	z80.c.t += z80.r.lc.t
	if mmu.inBios && z80.r.pc == 0x0100 {
		mmu.inBios = false
	}
}

func stepDisplay(z80 *Z80, mmu *MMU, gpu *GPU, video Video) {
	gpu.mclock += z80.r.lc.t

	switch mode := gpu.mode; mode {
	case HBLANK:
		if gpu.mclock >= HBLANKTiming {
			gpu.mclock = 0
			gpu.line++

			if gpu.line == 143 {
				gpu.mode = VBLANK
				video.Draw(gpu.screen)
			} else {
				gpu.mode = OAM
			}
		}
	case VBLANK:
		if gpu.mclock >= VBLANKTiming {
			gpu.mclock = 0
			gpu.line++

			if gpu.line > 153 {
				gpu.mode = OAM
				gpu.line = 0
			}
		}
	case OAM:
		if gpu.mclock >= OAMTiming {
			gpu.mclock = 0
			gpu.mode = VRAM
		}
	case VRAM:
		if gpu.mclock >= VRAMTiming {
			gpu.mclock = 0
			gpu.mode = HBLANK

			gpu.RenderLine()
		}
	}
}

func loadBios(mmu *MMU) {
	file, err := os.Open("roms/bios.gb")
	if err != nil {
		log.Fatal(err)
	}
	bios := make([]byte, 0x100, 0x100)
	n, err := file.Read(bios)
	if err != nil {
		log.Fatal(err)
	}
	if n != 0x100 {
		log.Fatal("Bad length read: bios file")
	}

	mmu.writeBios(0x0, 0x100, bios)
}

func main() {
	var mmu MMU
	mmu.Init()
	var z80 Z80
	z80.Init()
	var gpu GPU
	gpu.Init()

	var video WebviewVideo
	video.Setup()

	loadBios(&mmu)

	for {
		s := step(&z80, &mmu, &gpu, &video)
		if !s {
			break
		}
		time.Sleep(60)
	}
}
