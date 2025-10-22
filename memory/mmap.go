package memory

import (
	"github.com/ghosind/go-nes/ppu"
	"github.com/ghosind/go-nes/rom"
)

// MemoryMap represents the NES memory map. It includes RAM, ROM, and I/O registers.
//
// The NES memory map 64KB address space is divided as follows:
//
// 0x0000 - 0x07FF: 2KB internal RAM.
// 0x0800 - 0x1FFF: Mirrors of 0x0000 - 0x07FF (repeats every 2KB).
// 0x2000 - 0x2007: PPU registers.
// 0x2008 - 0x3FFF: Mirrors of 0x2000 - 0x2007 (repeats every 8 bytes).
// 0x4000 - 0x4017: APU and I/O registers.
// 0x4018 - 0x401F: APU and I/O functionality that is normally disabled.
// 0x4020 - 0xFFFF: Cartridge space (PRG ROM, PRG RAM, and mapper registers).
type MemoryMap struct {
	// mem is the placeholder for the entire 64KB memory space.
	// In a complete implementation, this would be divided into RAM, ROM, and I/O registers.
	// For simplicity, we will use a single array here.
	mem [65536]uint8
	// ram represents the 2KB internal RAM of the NES.
	ram RAM
	// ppu represents the Picture Processing Unit.
	ppu *ppu.PPU
	// rom represents the cartridge ROM, which includes PRG ROM and PRG RAM.
	rom *rom.ROM
}

func NewMemoryMap(rom *rom.ROM, ppu *ppu.PPU) *MemoryMap {
	mmap := new(MemoryMap)
	mmap.rom = rom
	mmap.ppu = ppu
	return mmap
}

func (m *MemoryMap) Read(addr uint16) uint8 {
	if addr < 0x2000 {
		return m.ram.Read(addr)
	} else if addr < 0x4000 {
		return m.ppu.CPURead(addr)
	} else if addr >= 0x6000 {
		return m.rom.CPURead(addr)
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
	} else if addr >= 0x2000 && addr < 0x4000 {
		m.ppu.CPUWrite(addr, value)
	} else if addr >= 0x6000 {
		m.rom.CPUWrite(addr, value)
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
