package cpu

type Memory [65536]uint8

func (m *Memory) Read(addr uint16) uint8 {
	return m[addr]
}

func (m *Memory) ReadZeroPage(addr uint8) uint8 {
	return m[uint16(addr)]
}

func (m *Memory) ReadAbs(high, low uint8) uint8 {
	addr := (uint16(high) << 8) | uint16(low)
	return m[addr]
}

func (m *Memory) ReadAbsShift(high, low, shift uint8) uint8 {
	addr := (uint16(high) << 8) | uint16(low) + uint16(shift)
	return m[addr]
}

func (m *Memory) Write(addr uint16, value uint8) {
	m[addr] = value
}
