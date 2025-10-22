package mapper

import (
	"github.com/ghosind/go-nes/rom/ines"
)

type mapper0 struct {
	mapper
}

func NewMapper0(header *ines.INESHeader, data []byte) Mapper {
	m := new(mapper0)
	m.init(header, data)

	if m.header.ChrRomBanks == 0 {
		m.chrRam = make([]byte, 8*1024) // Default to 8KB CHR RAM
	}
	return m
}

func (m *mapper0) CPURead(addr uint16) uint8 {
	if addr >= 0x8000 {
		if m.header.PrgRomBanks == 1 {
			// Mirror 16KB bank if only one bank is present
			return m.prgRom[addr&0x3FFF]
		}
		return m.prgRom[addr&0x7FFF]
	} else if addr >= 0x6000 {
		return m.prgRam[addr-0x6000]
	}

	return 0
}

func (m *mapper0) CPUWrite(addr uint16, value uint8) {
	if addr >= 0x6000 && addr < 0x8000 {
		m.prgRam[addr-0x6000] = value
	}
	// Writes to ROM area are ignored
}

func (m *mapper0) PPURead(addr uint16) uint8 {
	if addr < 0x2000 {
		if m.header.ChrRomBanks == 0 {
			return m.chrRam[addr]
		}
		return m.chrRom[addr]
	}
	return 0
}

func (m *mapper0) PPUWrite(addr uint16, value uint8) {
	if addr < 0x2000 && m.header.ChrRomBanks == 0 {
		m.chrRam[addr] = value
	}
}
