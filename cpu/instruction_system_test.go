package cpu

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestCPU_NOP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xEA, // NOP opcode
		},
		cycles: 2,
	}

	testCPUInstruction(a, vector)
}
