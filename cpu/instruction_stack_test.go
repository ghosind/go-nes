package cpu

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestCPU_TSX(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xBA, // TSX opcode
		},
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: psFlagNegative, // Expect Negative flag to be set
		expectedX:  pointer(uint8(0xFD)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_TXS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x9A, // TXS opcode
		},
		x:          0x42,
		cycles:     2,
		expectedSP: pointer(uint8(0x42)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_PHA(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x48, // PHA opcode
		},
		a:      0x42,
		cycles: 3,
		expectedMem: map[uint16]uint8{
			0x01FD: 0x42, // Expect the value of A to be pushed onto the stack
		},
		expectedSP: pointer(uint8(0xFC)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_PHP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x08, // PHP opcode
		},
		ps:     psFlagNegative | psFlagOverflow, // Set Negative and Overflow flags
		cycles: 3,
		expectedMem: map[uint16]uint8{
			// Expect the value of PS with bits 4 and 5 set to be pushed onto the stack
			0x01FD: uint8(psFlagNegative | psFlagOverflow | 0x30),
		},
		expectedSP: pointer(uint8(0xFC)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_PLA(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000:       0x68, // PLA opcode
			0x100 | 0x43: 0x37, // Value to be pulled from stack
		},
		sp:         pointer(uint8(0x42)),
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x37)),
		expectedSP: pointer(uint8(0x43)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_PLP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000:       0x28, // PLP opcode
			0x100 | 0x43: 0b00111111,
		},
		sp:         pointer(uint8(0x42)),
		cycles:     4,
		psMask:     0xFF,
		expectedPS: 0b00111111, // Expect all flags to be set as per the value pulled from stack
		expectedSP: pointer(uint8(0x43)),
	}

	testCPUInstruction(a, vector)
}
