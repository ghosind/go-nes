package cpu

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestCPU_CLC(t *testing.T) {
	a := assert.New(t)

	cpu := NewCPU()
	cpu.mem[0xFFFC] = 0x00 // Reset vector low byte
	cpu.mem[0xFFFD] = 0x80 // Reset vector high byte
	cpu.mem[0x8000] = 0x18 // CLC opcode
	cpu.Reset()
	cpu.ps.setCarry(true) // Set carry flag to true

	cycles := cpu.Step()
	cpu.Step()

	a.EqualNow(cpu.ps.getCarry(), false)
	a.EqualNow(cycles, 2)
}

func TestCPU_CLD(t *testing.T) {
	a := assert.New(t)

	cpu := NewCPU()
	cpu.mem[0xFFFC] = 0x00 // Reset vector low byte
	cpu.mem[0xFFFD] = 0x80 // Reset vector high byte
	cpu.mem[0x8000] = 0xD8 // CLD opcode
	cpu.Reset()
	cpu.ps.setDecimal(true) // Set decimal flag to true

	cycles := cpu.Step()
	cpu.Step()

	a.EqualNow(cpu.ps.getDecimal(), false)
	a.EqualNow(cycles, 2)
}

func TestCPU_CLI(t *testing.T) {
	a := assert.New(t)

	cpu := NewCPU()
	cpu.mem[0xFFFC] = 0x00 // Reset vector low byte
	cpu.mem[0xFFFD] = 0x80 // Reset vector high byte
	cpu.mem[0x8000] = 0x58 // CLI opcode
	cpu.Reset()
	cpu.ps.setInterrupt(true) // Set interrupt flag to true

	cycles := cpu.Step()
	cpu.Step()

	a.EqualNow(cpu.ps.getInterrupt(), false)
	a.EqualNow(cycles, 2)
}

func TestCPU_CLV(t *testing.T) {
	a := assert.New(t)

	cpu := NewCPU()
	cpu.mem[0xFFFC] = 0x00 // Reset vector low byte
	cpu.mem[0xFFFD] = 0x80 // Reset vector high byte
	cpu.mem[0x8000] = 0xB8 // CLV opcode
	cpu.Reset()
	cpu.ps.setOverflow(true) // Set overflow flag to true

	cycles := cpu.Step()
	cpu.Step()

	a.EqualNow(cpu.ps.getOverflow(), false)
	a.EqualNow(cycles, 2)
}

func TestCPU_SEC(t *testing.T) {
	a := assert.New(t)

	cpu := NewCPU()
	cpu.mem[0xFFFC] = 0x00 // Reset vector low byte
	cpu.mem[0xFFFD] = 0x80 // Reset vector high byte
	cpu.mem[0x8000] = 0x38 // SEC opcode
	cpu.Reset()
	cpu.ps.setCarry(false) // Set carry flag to false

	cycles := cpu.Step()
	cpu.Step()

	a.EqualNow(cpu.ps.getCarry(), true)
	a.EqualNow(cycles, 2)
}

func TestCPU_SED(t *testing.T) {
	a := assert.New(t)

	cpu := NewCPU()
	cpu.mem[0xFFFC] = 0x00 // Reset vector low byte
	cpu.mem[0xFFFD] = 0x80 // Reset vector high byte
	cpu.mem[0x8000] = 0xF8 // SED opcode
	cpu.Reset()
	cpu.ps.setDecimal(false) // Set decimal flag to false

	cycles := cpu.Step()
	cpu.Step()

	a.EqualNow(cpu.ps.getDecimal(), true)
	a.EqualNow(cycles, 2)
}

func TestCPU_SEI(t *testing.T) {
	a := assert.New(t)

	cpu := NewCPU()
	cpu.mem[0xFFFC] = 0x00 // Reset vector low byte
	cpu.mem[0xFFFD] = 0x80 // Reset vector high byte
	cpu.mem[0x8000] = 0x78 // SEI opcode
	cpu.Reset()
	cpu.ps.setInterrupt(false) // Set interrupt flag to false

	cycles := cpu.Step()
	cpu.Step()

	a.EqualNow(cpu.ps.getInterrupt(), true)
	a.EqualNow(cycles, 2)
}
