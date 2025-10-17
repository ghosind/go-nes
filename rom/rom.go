package rom

import (
	"github.com/ghosind/go-nes/rom/ines"
	"github.com/ghosind/go-nes/rom/mapper"
)

type ROM struct {
	Mapper mapper.Mapper
}

func New(data []byte) (*ROM, error) {
	rom := new(ROM)

	if err := rom.parseROM(data); err != nil {
		return nil, err
	}

	return rom, nil
}

func (r *ROM) parseROM(data []byte) error {
	header, err := r.parseINesHeader(data)
	if err != nil {
		return err
	}

	mapper := mapper.NewMapper(header, data)
	if mapper == nil {
		return ines.ErrUnsupportedMapper
	}
	r.Mapper = mapper

	return nil
}

func (r *ROM) parseINesHeader(data []byte) (*ines.INESHeader, error) {
	if len(data) < 16 {
		return nil, ines.ErrInvalidROM
	}

	// Check for "NES" file signature
	if string(data[0:4]) != "NES\x1A" {
		return nil, ines.ErrInvalidROM
	}

	header := new(ines.INESHeader)
	header.Parse(data[0:16])

	return header, nil
}

func (r *ROM) CPURead(addr uint16) uint8 {
	return r.Mapper.CPURead(addr)
}

func (r *ROM) CPUWrite(addr uint16, value uint8) {
	r.Mapper.CPUWrite(addr, value)
}

func (r *ROM) PPURead(addr uint16) uint8 {
	return r.Mapper.PPURead(addr)
}

func (r *ROM) PPUWrite(addr uint16, value uint8) {
	r.Mapper.PPUWrite(addr, value)
}
