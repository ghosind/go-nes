package apu

type APU struct {
}

func New() *APU {
	apu := new(APU)
	return apu
}

func (a *APU) Step() {
	// APU step
}
