package ppu

import (
	"github.com/ghosind/go-nes/rom"
)

type PPU struct {
	PPUCTRL   uint8
	PPUMASK   uint8
	PPUSTATUS uint8
	OAMADDR   uint8
	OAMDATA   uint8
	PPUSCROLL uint8
	PPUADDR   uint8
	PPUDATA   uint8

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

func (p *PPU) CPURead(addr uint16) uint8 {
	addr = (addr - 0x2000) & 0x07
	switch addr {
	case 0x02:
		return p.PPUSTATUS
	case 0x04:
		return p.OAMDATA
	case 0x07:
		return p.PPUDATA
	default:
		return 0
	}
}

func (p *PPU) CPUWrite(addr uint16, value uint8) {
	addr = ((addr & 0x3FFF) - 0x2000) % 8
	switch addr {
	case 0x00:
		p.PPUCTRL = value
	case 0x01:
		p.PPUMASK = value
	case 0x03:
		p.OAMADDR = value
	case 0x04:
		p.OAMDATA = value
	case 0x05:
		p.PPUSCROLL = value
	case 0x06:
		p.PPUADDR = value
	case 0x07:
		p.PPUDATA = value
	}
}
