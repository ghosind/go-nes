package ppu

import (
	"github.com/ghosind/go-nes/rom"
)

type VRAM struct {
	rom        *rom.ROM
	nameTables [4][0x400]uint8
	palettes   [32]uint8
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
	case addr < 0x2000:
		// Pattern Table ROM
		return vram.rom.PPURead(addr)
	case addr >= 0x2000 && addr < 0x3F00:
		addr = (addr - 0x2000) % 0x1000
		// Name Table RAM
		index := addr / 0x400
		offset := addr % 0x400
		return vram.nameTables[index][offset]
	default:
		// Palette RAM
		return vram.palettes[(addr-0x3F00)%32]
	}
}

func (vram *VRAM) Write(addr uint16, value uint8) {
	// Mirror address to 0x0000-0x3FFF range
	addr = addr & 0x3FFF

	switch {
	case addr < 0x2000:
		// Pattern Table ROM
		vram.rom.PPUWrite(addr, value)
	case addr >= 0x2000 && addr < 0x3F00:
		// Name Table RAM
		addr = (addr - 0x2000) % 0x1000
		index := addr / 0x400
		offset := addr % 0x400
		vram.nameTables[index][offset] = value
	default:
		// Palette RAM
		vram.palettes[(addr-0x3F00)%32] = value
	}
}
