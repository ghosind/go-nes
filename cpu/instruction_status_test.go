package cpu

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestCPU_CLC(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x18, // CLC opcode
		},
		ps:         psFlagCarry, // Initial PS with Carry flag set
		cycles:     2,
		psMask:     psFlagCarry,
		expectedPS: 0, // Expect Carry flag to be cleared
	}

	testCPUInstruction(a, vector)
}

func TestCPU_CLD(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xD8, // CLD opcode
		},
		ps:         psFlagDecimal, // Initial PS with Decimal flag set
		cycles:     2,
		psMask:     psFlagDecimal,
		expectedPS: 0, // Expect Decimal flag to be cleared
	}

	testCPUInstruction(a, vector)
}

func TestCPU_CLI(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x58, // CLI opcode
		},
		ps:         psFlagInterrupt, // Initial PS with Interrupt flag set
		cycles:     2,
		psMask:     psFlagInterrupt,
		expectedPS: 0, // Expect Interrupt flag to be cleared
	}

	testCPUInstruction(a, vector)
}

func TestCPU_CLV(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xB8, // CLV opcode
		},
		ps:         psFlagOverflow, // Initial PS with Overflow flag set
		cycles:     2,
		psMask:     psFlagOverflow,
		expectedPS: 0, // Expect Overflow flag to be cleared
	}

	testCPUInstruction(a, vector)
}

func TestCPU_SEC(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x38, // SEC opcode
		},
		cycles:     2,
		psMask:     psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	testCPUInstruction(a, vector)
}

func TestCPU_SED(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xF8, // SED opcode
		},
		cycles:     2,
		psMask:     psFlagDecimal,
		expectedPS: psFlagDecimal, // Expect Decimal flag to be set
	}

	testCPUInstruction(a, vector)
}

func TestCPU_SEI(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x78, // SEI opcode
		},
		cycles:     2,
		psMask:     psFlagInterrupt,
		expectedPS: psFlagInterrupt, // Expect Interrupt flag to be set
	}

	testCPUInstruction(a, vector)
}
