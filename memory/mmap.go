package memory

type MemoryMap struct {
	mem [65536]uint8
	ram RAM
}

func NewMemoryMap() *MemoryMap {
	mmap := new(MemoryMap)
	return mmap
}

func (m *MemoryMap) Read(addr uint16) uint8 {
	if addr < 0x2000 {
		return m.ram.Read(addr)
	}
	return m.mem[addr]
}

func (m *MemoryMap) ReadZeroPage(addr uint8) uint8 {
	return m.Read(uint16(addr))
}

func (m *MemoryMap) ReadAbs(high, low uint8) uint8 {
	addr := (uint16(high) << 8) | uint16(low)
	return m.Read(addr)
}

func (m *MemoryMap) ReadAbsShift(high, low, shift uint8) uint8 {
	addr := (uint16(high) << 8) | uint16(low) + uint16(shift)
	return m.Read(addr)
}

func (m *MemoryMap) ReadIndexedIndirect(addr, shift uint8) uint8 {
	ptr := (addr + shift)
	low := m.ReadZeroPage(ptr)
	high := m.ReadZeroPage(ptr + 1)
	return m.ReadAbs(high, low)
}

func (m *MemoryMap) ReadIndirectIndexed(addr, shift uint8) uint8 {
	low := m.ReadZeroPage(addr)
	high := m.ReadZeroPage(addr + 1)
	return m.ReadAbsShift(high, low, shift)
}

func (m *MemoryMap) Write(addr uint16, value uint8) {
	if addr < 0x2000 {
		m.ram.Write(addr, value)
	} else {
		m.mem[addr] = value
	}
}

func (m *MemoryMap) WriteZeroPage(addr uint8, value uint8) {
	m.Write(uint16(addr), value)
}

func (m *MemoryMap) WriteAbs(high, low uint8, value uint8) {
	addr := (uint16(high) << 8) | uint16(low)
	m.Write(addr, value)
}

func (m *MemoryMap) WriteAbsShift(high, low, shift uint8, value uint8) {
	addr := (uint16(high) << 8) | uint16(low) + uint16(shift)
	m.Write(addr, value)
}

func (m *MemoryMap) WriteIndexedIndirect(addr, shift uint8, value uint8) {
	ptr := (addr + shift)
	low := m.ReadZeroPage(ptr)
	high := m.ReadZeroPage(ptr + 1)
	m.WriteAbs(high, low, value)
}

func (m *MemoryMap) WriteIndirectIndexed(addr, shift uint8, value uint8) {
	low := m.ReadZeroPage(addr)
	high := m.ReadZeroPage(addr + 1)
	m.WriteAbsShift(high, low, shift, value)
}
