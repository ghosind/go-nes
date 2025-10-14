package cpu

import "github.com/ghosind/go-assert"

type instructionTestVector struct {
	// Initial CPU state
	memory map[uint16]uint8
	a      uint8
	x      uint8
	y      uint8
	sp     uint8
	ps     ProcessorStatus

	// expected CPU state after instruction execution
	cycles      int
	psMask      ProcessorStatus
	expectedPS  ProcessorStatus
	expectedMem map[uint16]uint8
}

func testCPUInstruction(a *assert.Assertion, vector instructionTestVector) *CPU {
	cpu := NewCPU()

	// Load initial memory state
	for addr, value := range vector.memory {
		cpu.mem[addr] = value
	}

	// Set the reset vector to point to 0x8000
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80

	cpu.Reset()

	// Set initial CPU registers
	cpu.a = vector.a
	cpu.x = vector.x
	cpu.y = vector.y
	cpu.sp = vector.sp
	*cpu.ps = vector.ps

	cycles := cpu.Step()

	a.EqualNow(cycles, vector.cycles)
	if vector.psMask != 0 {
		a.EqualNow(*cpu.ps&vector.psMask, vector.expectedPS&vector.psMask)
	}
	for addr, value := range vector.expectedMem {
		a.EqualNow(cpu.mem[addr], value)
	}

	return cpu
}
