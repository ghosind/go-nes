package cpu

import (
	"testing"
)

func TestCPU_CLC(t *testing.T) {
	vector := &instructionTestVector{
		name: "CLC",
		memory: map[uint16]uint8{
			0x8000: 0x18, // CLC
		},
		ps:         psFlagCarry, // Initial PS with Carry flag set
		cycles:     2,
		psMask:     psFlagCarry,
		expectedPS: 0, // Expect Carry flag to be cleared
	}

	vector.test(t)
}

func TestCPU_CLD(t *testing.T) {
	vector := &instructionTestVector{
		name: "CLD",
		memory: map[uint16]uint8{
			0x8000: 0xD8, // CLD
		},
		ps:         psFlagDecimal, // Initial PS with Decimal flag set
		cycles:     2,
		psMask:     psFlagDecimal,
		expectedPS: 0, // Expect Decimal flag to be cleared
	}

	vector.test(t)
}

func TestCPU_CLI(t *testing.T) {
	vector := &instructionTestVector{
		name: "CLI",
		memory: map[uint16]uint8{
			0x8000: 0x58, // CLI
		},
		ps:         psFlagInterrupt, // Initial PS with Interrupt flag set
		cycles:     2,
		psMask:     psFlagInterrupt,
		expectedPS: 0, // Expect Interrupt flag to be cleared
	}

	vector.test(t)
}

func TestCPU_CLV(t *testing.T) {
	vector := &instructionTestVector{
		name: "CLV",
		memory: map[uint16]uint8{
			0x8000: 0xB8, // CLV
		},
		ps:         psFlagOverflow, // Initial PS with Overflow flag set
		cycles:     2,
		psMask:     psFlagOverflow,
		expectedPS: 0, // Expect Overflow flag to be cleared
	}

	vector.test(t)
}

func TestCPU_SEC(t *testing.T) {
	vector := &instructionTestVector{
		name: "SEC",
		memory: map[uint16]uint8{
			0x8000: 0x38, // SEC
		},
		cycles:     2,
		psMask:     psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_SED(t *testing.T) {
	vector := &instructionTestVector{
		name: "SED",
		memory: map[uint16]uint8{
			0x8000: 0xF8, // SED
		},
		cycles:     2,
		psMask:     psFlagDecimal,
		expectedPS: psFlagDecimal, // Expect Decimal flag to be set
	}

	vector.test(t)
}

func TestCPU_SEI(t *testing.T) {
	vector := &instructionTestVector{
		name: "SEI",
		memory: map[uint16]uint8{
			0x8000: 0x78, // SEI
		},
		cycles:     2,
		psMask:     psFlagInterrupt,
		expectedPS: psFlagInterrupt, // Expect Interrupt flag to be set
	}

	vector.test(t)
}
