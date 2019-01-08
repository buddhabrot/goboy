package main

import "log"

// MMU The mmu
type MMU struct {
	bios   [0x100]byte
	rom    []byte
	vram   [0x2000]byte
	xram   []byte
	ram    [0x2000]byte
	sprite [0x0A0]byte
	mmio   [0x080]byte
	zram   [0x080]byte
	inBios bool

	romoffset uint16
	ramoffset uint16
}

func (mmu *MMU) rb(addr uint16) uint8 {
	switch 0xF000 & addr {
	// ROM
	case 0x0000:
		if mmu.inBios {
			if addr < 0x100 {
				return mmu.bios[addr]
			} else {

			}
		} else {
			return mmu.rom[addr]
		}
	// ROM 1
	case 0x1000, 0x2000, 0x3000, 0x4000, 0x5000, 0x6000, 0x7000:
		return mmu.rom[mmu.romoffset+(addr&0x3FFF)]
	// VRAM
	case 0x8000, 0x9000:
		return mmu.vram[addr&0x1FFF]
	// EXTRAM
	case 0xA000, 0xB000:
		return mmu.xram[mmu.ramoffset+(addr&0x1FFF)]
	// RAM
	case 0xC000, 0xD000, 0xE000:
		return mmu.ram[addr&0x1FFF]
	case 0xF000:
		switch 0x0F00 & addr {
		// Echo
		case 0x000, 0x100, 0x200, 0x300,
			0x400, 0x500, 0x600, 0x700,
			0x800, 0x900, 0xA00, 0xB00,
			0xC00, 0xD00:
			return mmu.ram[addr&0x1FFF]

		}
	}

	return 0x0
}

func (mmu *MMU) rbrom(addr uint16) uint8 {

}

// Init initializes the MMU
func (mmu *MMU) Init() {
	mmu.inBios = true
}

func (mmu *MMU) writeBios(offset uint16, length uint16, data []byte) {
	n := copy(mmu.bios[offset:length], data)
	if uint16(n) != length {
		log.Fatal("Unsupported memory write: source and dest do not match sizes")
	}
}
