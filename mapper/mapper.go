package mapper

type Mapper interface {
	CPURead(addr uint16) uint8
	CPUWrite(addr uint16, value uint8)
	PPURead(addr uint16) uint8
	PPUWrite(addr uint16, value uint8)
}
