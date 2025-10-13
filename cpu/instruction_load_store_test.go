package cpu

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestCPU_LDA_IMM(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDA Immediate instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xA9 // LDA Immediate opcode
	cpu.mem[0x8001] = 0x42 // Operand: Load the value 0x42 into A
	cpu.Reset()

	cpu.Step()

	a.EqualNow(cpu.a, uint8(0x42))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}

func TestCPU_LDA_ZP(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDA Zero Page instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xA5 // LDA Zero Page opcode
	cpu.mem[0x8001] = 0x10 // Operand: Zero Page address 0x10
	cpu.mem[0x0010] = 0x37 // Value at Zero Page address 0x10
	cpu.Reset()

	cpu.Step()

	a.EqualNow(cpu.a, uint8(0x37))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}

func TestCPU_LDA_ZP_X(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDA Zero Page, X instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xB5 // LDA Zero Page, X opcode
	cpu.mem[0x8001] = 0x10 // Operand: Zero Page address 0x10
	cpu.mem[0x0015] = 0x58 // Value at Zero Page address 0x10 + X (0x15)
	cpu.Reset()
	cpu.x = 0x05 // Set X register to 5

	cpu.Step()

	a.EqualNow(cpu.a, uint8(0x58))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}

func TestCPU_LDA_ABS(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDA Absolute instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xAD // LDA Absolute opcode
	cpu.mem[0x8001] = 0x00 // Low byte of address
	cpu.mem[0x8002] = 0x20 // High byte of address (0x2000)
	cpu.mem[0x2000] = 0x7A // Value at address 0x2000
	cpu.Reset()

	cpu.Step()

	a.EqualNow(cpu.a, uint8(0x7A))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}

func TestCPU_LDA_ABS_X(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDA Absolute, X instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xBD // LDA Absolute, X opcode
	cpu.mem[0x8001] = 0x00 // Low byte of address
	cpu.mem[0x8002] = 0x20 // High byte of address (0x2000)
	cpu.mem[0x2005] = 0x9C // Value at address 0x2000 + X (0x2005)
	cpu.Reset()
	cpu.x = 0x05 // Set X register to 5

	cpu.Step()

	a.EqualNow(cpu.a, uint8(0x9C))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), true)
}

func TestCPU_LDA_ABS_Y(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDA Absolute, Y instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xB9 // LDA Absolute, Y opcode
	cpu.mem[0x8001] = 0x00 // Low byte of address
	cpu.mem[0x8002] = 0x20 // High byte of address (0x2000)
	cpu.mem[0x2003] = 0xFF // Value at address 0x2000 + Y (0x2003)
	cpu.Reset()
	cpu.y = 0x03 // Set Y register to 3

	cpu.Step()

	a.EqualNow(cpu.a, uint8(0xFF))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), true)
}

func TestCPU_LDA_IND_X(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDA (Indirect, X) instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xA1 // LDA (Indirect, X) opcode
	cpu.mem[0x8001] = 0x10 // Operand: Zero Page address 0x10
	// Set up the indirect address
	cpu.mem[0x0015] = 0x00 // Low byte of effective address (0x3000)
	cpu.mem[0x0016] = 0x30 // High byte of effective address
	cpu.mem[0x3000] = 0x66 // Value at effective address 0x3000
	cpu.Reset()
	cpu.x = 0x05 // Set X register to 5

	cpu.Step()

	a.EqualNow(cpu.a, uint8(0x66))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}

func TestCPU_LDA_IND_Y(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDA (Indirect), Y instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xB1 // LDA (Indirect), Y opcode
	cpu.mem[0x8001] = 0x10 // Operand: Zero Page address 0x10
	// Set up the indirect address
	cpu.mem[0x0010] = 0x00 // Low byte of base address (0x3000)
	cpu.mem[0x0011] = 0x30 // High byte of base address
	cpu.mem[0x3002] = 0xAB // Value at effective address 0x3000 + Y (0x3002)
	cpu.Reset()
	cpu.y = 0x02 // Set Y register to 2

	cpu.Step()

	a.EqualNow(cpu.a, uint8(0xAB))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), true)
}

func TestCPU_LDX_IMM(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDX Immediate instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xA2 // LDX Immediate opcode
	cpu.mem[0x8001] = 0x55 // Operand: Load the value 0x55 into X
	cpu.Reset()

	cpu.Step()

	a.EqualNow(cpu.x, uint8(0x55))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}

func TestCPU_LDX_ZP(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDX Zero Page instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xA6 // LDX Zero Page opcode
	cpu.mem[0x8001] = 0x20 // Operand: Zero Page address 0x20
	cpu.mem[0x0020] = 0x33 // Value at Zero Page address 0x20
	cpu.Reset()

	cpu.Step()

	a.EqualNow(cpu.x, uint8(0x33))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}

func TestCPU_LDX_ZP_Y(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDX Zero Page, Y instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xB6 // LDX Zero Page, Y opcode
	cpu.mem[0x8001] = 0x20 // Operand: Zero Page address 0x20
	cpu.mem[0x0025] = 0x77 // Value at Zero Page address 0x20 + Y (0x25)
	cpu.Reset()
	cpu.y = 0x05 // Set Y register to 5

	cpu.Step()

	a.EqualNow(cpu.x, uint8(0x77))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}

func TestCPU_LDX_ABS(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDX Absolute instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xAE // LDX Absolute opcode
	cpu.mem[0x8001] = 0x00 // Low byte of address
	cpu.mem[0x8002] = 0x30 // High byte of address (0x3000)
	cpu.mem[0x3000] = 0x88 // Value at address 0x3000
	cpu.Reset()

	cpu.Step()

	a.EqualNow(cpu.x, uint8(0x88))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), true)
}

func TestCPU_LDX_ABS_Y(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDX Absolute, Y instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xBE // LDX Absolute, Y opcode
	cpu.mem[0x8001] = 0x00 // Low byte of address
	cpu.mem[0x8002] = 0x30 // High byte of address (0x3000)
	cpu.mem[0x3004] = 0x22 // Value at address 0x3000 + Y (0x3004)
	cpu.Reset()
	cpu.y = 0x04 // Set Y register to 4

	cpu.Step()

	a.EqualNow(cpu.x, uint8(0x22))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}

func TestCPU_LDY_IMM(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDY Immediate instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xA0 // LDY Immediate opcode
	cpu.mem[0x8001] = 0x99 // Operand: Load the value 0x99 into Y
	cpu.Reset()

	cpu.Step()

	a.EqualNow(cpu.y, uint8(0x99))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), true)
}

func TestCPU_LDY_ZP(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDY Zero Page instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xA4 // LDY Zero Page opcode
	cpu.mem[0x8001] = 0x30 // Operand: Zero Page address 0x30
	cpu.mem[0x0030] = 0x44 // Value at Zero Page address 0x30
	cpu.Reset()

	cpu.Step()

	a.EqualNow(cpu.y, uint8(0x44))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}

func TestCPU_LDY_ZP_X(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDY Zero Page, X instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xB4 // LDY Zero Page, X opcode
	cpu.mem[0x8001] = 0x30 // Operand: Zero Page address 0x30
	cpu.mem[0x0035] = 0x11 // Value at Zero Page address 0x30 + X (0x35)
	cpu.Reset()
	cpu.x = 0x05 // Set X register to 5

	cpu.Step()

	a.EqualNow(cpu.y, uint8(0x11))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}

func TestCPU_LDY_ABS(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDY Absolute instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xAC // LDY Absolute opcode
	cpu.mem[0x8001] = 0x00 // Low byte of address
	cpu.mem[0x8002] = 0x40 // High byte of address (0x4000)
	cpu.mem[0x4000] = 0xFE // Value at address 0x4000
	cpu.Reset()

	cpu.Step()

	a.EqualNow(cpu.y, uint8(0xFE))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), true)
}

func TestCPU_LDY_ABS_X(t *testing.T) {
	a := assert.New(t)
	cpu := NewCPU()

	// Load LDY Absolute, X instruction at the reset vector
	cpu.mem[0xFFFC] = 0x00
	cpu.mem[0xFFFD] = 0x80
	cpu.mem[0x8000] = 0xBC // LDY Absolute, X opcode
	cpu.mem[0x8001] = 0x00 // Low byte of address
	cpu.mem[0x8002] = 0x40 // High byte of address (0x4000)
	cpu.mem[0x4003] = 0x66 // Value at address 0x4000 + X (0x4003)
	cpu.Reset()
	cpu.x = 0x03 // Set X register to 3

	cpu.Step()

	a.EqualNow(cpu.y, uint8(0x66))
	a.EqualNow(cpu.ps.getZero(), false)
	a.EqualNow(cpu.ps.getNegative(), false)
}
