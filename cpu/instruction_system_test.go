package cpu

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestCPU_NOP(t *testing.T) {
	a := assert.New(t)

	cpu := NewCPU()
	cpu.mem[0xFFFC] = 0x00 // Reset vector low byte
	cpu.mem[0xFFFD] = 0x80 // Reset vector high byte
	cpu.mem[0x8000] = 0xEA // NOP opcode
	cpu.Reset()

	cycles := cpu.Step()

	a.EqualNow(cycles, 2)
}
