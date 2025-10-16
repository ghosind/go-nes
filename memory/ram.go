package memory

type RAM [2048]uint8

func (r *RAM) Read(addr uint16) uint8 {
	return r[addr%0x0800]
}

func (r *RAM) Write(addr uint16, value uint8) {
	r[addr%0x0800] = value
}
