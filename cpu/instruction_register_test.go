package cpu

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestCPU_TAX(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xAA, // TAX opcode
		},
		a:          0x42,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedX:  pointer(uint8(0x42)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_TAY(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xA8, // TAY opcode
		},
		a:          0x42,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedY:  pointer(uint8(0x42)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_TXA(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x8A, // TXA opcode
		},
		x:          0x42,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x42)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_TYA(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x98, // TYA opcode
		},
		y:          0x42,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x42)),
	}

	testCPUInstruction(a, vector)
}
