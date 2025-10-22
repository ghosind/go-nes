package ppu

import (
	"github.com/ghosind/go-nes/rom"
)

type VRAM struct {
	rom *rom.ROM
}

func NewVRAM(rom *rom.ROM) *VRAM {
	return &VRAM{
		rom: rom,
	}
}

func (vram *VRAM) Read(addr uint16) uint8 {
	// Mirror address to 0x0000-0x3FFF range
	addr = addr & 0x3FFF

	switch {
	// TODO: Implement NameTable and Palette RAM reads
	case addr < 0x2000:
		return vram.rom.PPURead(addr)
	}
	return 0
}

func (vram *VRAM) Write(addr uint16, value uint8) {
	// Mirror address to 0x0000-0x3FFF range
	addr = addr & 0x3FFF

	switch {
	// TODO: Implement NameTable and Palette RAM writes
	case addr < 0x2000:
		vram.rom.PPUWrite(addr, value)
	}
}
