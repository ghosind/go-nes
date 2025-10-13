package cpu

type Memory [65536]uint8

func (m *Memory) Read(addr uint16) uint8 {
	return m[addr]
}

func (m *Memory) ReadZeroPage(addr uint8) uint8 {
	return m.Read(uint16(addr))
}

func (m *Memory) ReadAbs(high, low uint8) uint8 {
	addr := (uint16(high) << 8) | uint16(low)
	return m.Read(addr)
}

func (m *Memory) ReadAbsShift(high, low, shift uint8) uint8 {
	addr := (uint16(high) << 8) | uint16(low) + uint16(shift)
	return m.Read(addr)
}

func (m *Memory) ReadIndexedIndirect(addr, shift uint8) uint8 {
	ptr := (addr + shift)
	low := m.ReadZeroPage(ptr)
	high := m.ReadZeroPage(ptr + 1)
	return m.ReadAbs(high, low)
}

func (m *Memory) ReadIndirectIndexed(addr, shift uint8) uint8 {
	low := m.ReadZeroPage(addr)
	high := m.ReadZeroPage(addr + 1)
	return m.ReadAbsShift(high, low, shift)
}

func (m *Memory) Write(addr uint16, value uint8) {
	m[addr] = value
}

func (m *Memory) WriteZeroPage(addr uint8, value uint8) {
	m.Write(uint16(addr), value)
}

func (m *Memory) WriteAbs(high, low uint8, value uint8) {
	addr := (uint16(high) << 8) | uint16(low)
	m.Write(addr, value)
}

func (m *Memory) WriteAbsShift(high, low, shift uint8, value uint8) {
	addr := (uint16(high) << 8) | uint16(low) + uint16(shift)
	m.Write(addr, value)
}

func (m *Memory) WriteIndexedIndirect(addr, shift uint8, value uint8) {
	ptr := (addr + shift)
	low := m.ReadZeroPage(ptr)
	high := m.ReadZeroPage(ptr + 1)
	m.WriteAbs(high, low, value)
}

func (m *Memory) WriteIndirectIndexed(addr, shift uint8, value uint8) {
	low := m.ReadZeroPage(addr)
	high := m.ReadZeroPage(addr + 1)
	m.WriteAbsShift(high, low, shift, value)
}
