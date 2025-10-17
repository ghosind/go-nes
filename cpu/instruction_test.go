package cpu

import (
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-nes/memory"
	"github.com/ghosind/go-nes/rom"
)

type instructionTestVector struct {
	// Descriptive name for the test case
	name string

	// Initial CPU state
	memory map[uint16]uint8
	a      uint8
	x      uint8
	y      uint8
	sp     *uint8
	ps     ProcessorStatus

	// expected CPU state after instruction execution
	cycles      uint64
	psMask      ProcessorStatus
	expectedPS  ProcessorStatus
	expectedMem map[uint16]uint8
	expectedA   *uint8
	expectedX   *uint8
	expectedY   *uint8
	expectedSP  *uint8
	expectedPC  *uint16
}

type testMapper struct {
	rom [65536]byte
}

func (m *testMapper) CPURead(addr uint16) uint8 { return m.rom[addr] }

func (m *testMapper) CPUWrite(addr uint16, value uint8) { m.rom[addr] = value }

func (m *testMapper) PPURead(addr uint16) uint8 { return 0 }

func (m *testMapper) PPUWrite(addr uint16, value uint8) {}

func pointer[T any](v T) *T {
	return &v
}

func (vector *instructionTestVector) test(t *testing.T) *CPU {
	// Create a new assertion context and mark it as a helper
	a := assert.New(t)
	a.Helper()

	// Create a new CPU instance
	cpu := New(memory.NewMemoryMap(&rom.ROM{
		Mapper: new(testMapper),
	}))

	// Load initial memory state
	for addr, value := range vector.memory {
		cpu.mem.Write(addr, value)
	}

	// Set the reset vector to point to 0x8000
	cpu.mem.Write(0xFFFC, 0x00)
	cpu.mem.Write(0xFFFD, 0x80)

	// Reset the CPU to initialize PC and other registers
	cpu.Reset()

	// Set initial CPU registers
	cpu.a = vector.a
	cpu.x = vector.x
	cpu.y = vector.y
	*cpu.ps = vector.ps

	if vector.sp != nil {
		cpu.sp = *vector.sp
	}

	// Execute the instruction in 0x8000
	cycles := cpu.Step()

	// Validate the clock cycles, processor status, and memory state
	if vector.cycles != 0 {
		a.EqualNow(cycles, vector.cycles,
			"%s: expected %d cycles, got %d", vector.name, vector.cycles, cycles)
	}
	if vector.psMask != 0 {
		actual := *cpu.ps & vector.psMask
		expected := vector.expectedPS & vector.psMask
		a.EqualNow(actual, expected,
			"%s: expected PS flags %08b, got %08b", vector.name, expected, actual)
	}
	for addr, value := range vector.expectedMem {
		actual := cpu.mem.Read(addr)
		a.EqualNow(actual, value,
			"%s: expected memory at 0x%04X to be 0x%02X, got 0x%02X",
			vector.name, addr, value, actual)
	}

	// Validate CPU registers if expected values are provided
	if vector.expectedA != nil {
		a.EqualNow(cpu.a, *vector.expectedA, "%s: expected A to be 0x%02X, got 0x%02X", vector.name, *vector.expectedA, cpu.a)
	}
	if vector.expectedX != nil {
		a.EqualNow(cpu.x, *vector.expectedX, "%s: expected X to be 0x%02X, got 0x%02X", vector.name, *vector.expectedX, cpu.x)
	}
	if vector.expectedY != nil {
		a.EqualNow(cpu.y, *vector.expectedY, "%s: expected Y to be 0x%02X, got 0x%02X", vector.name, *vector.expectedY, cpu.y)
	}
	if vector.expectedSP != nil {
		a.EqualNow(cpu.sp, *vector.expectedSP, "%s: expected SP to be 0x%02X, got 0x%02X", vector.name, *vector.expectedSP, cpu.sp)
	}
	if vector.expectedPC != nil {
		a.EqualNow(cpu.pc, *vector.expectedPC, "%s: expected PC to be 0x%04X, got 0x%04X", vector.name, *vector.expectedPC, cpu.pc)
	}

	// Return the CPU instance to allow further assertions if needed
	return cpu
}
