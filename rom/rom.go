package rom

import (
	"errors"

	"github.com/ghosind/go-nes/mapper"
)

var (
	ErrInvalidROM = errors.New("invalid ROM")
)

type ROM struct {
	mapper mapper.Mapper
}

func New(data []byte) (*ROM, error) {
	rom := new(ROM)

	if len(data) < 16 {
		return nil, ErrInvalidROM
	}

	// Check for "NES" file signature
	if data[0] != 'N' || data[1] != 'E' || data[2] != 'S' || data[3] != 0x1A {
		return nil, ErrInvalidROM
	}

	return rom, nil
}

func (r *ROM) CPURead(addr uint16) uint8 {
	return r.mapper.CPURead(addr)
}

func (r *ROM) CPUWrite(addr uint16, value uint8) {
	r.mapper.CPUWrite(addr, value)
}

func (r *ROM) PPURead(addr uint16) uint8 {
	return r.mapper.PPURead(addr)
}

func (r *ROM) PPUWrite(addr uint16, value uint8) {
	r.mapper.PPUWrite(addr, value)
}
