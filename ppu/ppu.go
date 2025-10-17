package ppu

type PPU struct {
}

func New() *PPU {
	ppu := new(PPU)
	return ppu
}

func (p *PPU) Step() {
	// PPU step
}
