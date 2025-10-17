package mapper

import "github.com/ghosind/go-nes/rom/ines"

type Mapper interface {
	CPURead(addr uint16) uint8
	CPUWrite(addr uint16, value uint8)
	PPURead(addr uint16) uint8
	PPUWrite(addr uint16, value uint8)
}

var validMappers = map[int]func(header *ines.INESHeader, data []byte) Mapper{
	0: NewMapper0,
}

func NewMapper(header *ines.INESHeader, data []byte) Mapper {
	if constructor, ok := validMappers[header.MapperID]; ok {
		return constructor(header, data)
	}
	return nil
}
