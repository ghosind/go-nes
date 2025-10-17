package mapper

import "github.com/ghosind/go-nes/rom/ines"

type mapper struct {
	header *ines.INESHeader
	prgRom []byte
	chrRom []byte
	prgRam []byte
}

func (m *mapper) init(header *ines.INESHeader, data []byte) error {
	m.header = header

	prgOffset := 16
	if m.header.Trainer {
		prgOffset += 512
	}

	prgSize := m.header.PrgRomBanks * 16 * 1024 // 16KB per bank
	chrSize := m.header.ChrRomBanks * 8 * 1024  // 8KB per bank
	ramSize := m.header.PrgRamUnits * 8 * 1024  // 8KB per unit

	if len(data) < prgOffset+prgSize+chrSize {
		return ines.ErrInvalidROM
	}

	m.prgRom = make([]byte, prgSize)
	copy(m.prgRom, data[prgOffset:prgOffset+prgSize])

	m.chrRom = make([]byte, chrSize)
	copy(m.chrRom, data[prgOffset+prgSize:prgOffset+prgSize+chrSize])

	m.prgRam = make([]byte, ramSize)

	return nil
}
