package ppu

import (
	"github.com/ghosind/go-nes/rom"
)

type PPU struct {
	VRAM *VRAM

	rom *rom.ROM
}

func New(rom *rom.ROM) *PPU {
	ppu := new(PPU)

	ppu.rom = rom
	ppu.VRAM = NewVRAM(rom)

	return ppu
}

func (p *PPU) Step() {
	// PPU step
}
