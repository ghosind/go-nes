package cpu

import (
	"testing"
)

func TestCPU_TSX(t *testing.T) {
	vector := &instructionTestVector{
		name: "TSX",
		memory: map[uint16]uint8{
			0x8000: 0xBA, // TSX
		},
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: psFlagNegative, // Expect Negative flag to be set
		expectedX:  pointer(uint8(0xFD)),
	}

	vector.test(t)
}

func TestCPU_TXS(t *testing.T) {
	vector := &instructionTestVector{
		name: "TXS",
		memory: map[uint16]uint8{
			0x8000: 0x9A, // TXS
		},
		x:          0x42,
		cycles:     2,
		expectedSP: pointer(uint8(0x42)),
	}

	vector.test(t)
}

func TestCPU_PHA(t *testing.T) {
	vector := &instructionTestVector{
		name: "PHA",
		memory: map[uint16]uint8{
			0x8000: 0x48, // PHA
		},
		a:      0x42,
		cycles: 3,
		expectedMem: map[uint16]uint8{
			0x01FD: 0x42, // Expect the value of A to be pushed onto the stack
		},
		expectedSP: pointer(uint8(0xFC)),
	}

	vector.test(t)
}

func TestCPU_PHP(t *testing.T) {
	vector := &instructionTestVector{
		name: "PHP",
		memory: map[uint16]uint8{
			0x8000: 0x08, // PHP
		},
		ps:     psFlagNegative | psFlagOverflow, // Set Negative and Overflow flags
		cycles: 3,
		expectedMem: map[uint16]uint8{
			// Expect the value of PS with bits 4 and 5 set to be pushed onto the stack
			0x01FD: uint8(psFlagNegative | psFlagOverflow | 0x30),
		},
		expectedSP: pointer(uint8(0xFC)),
	}

	vector.test(t)
}

func TestCPU_PLA(t *testing.T) {
	vector := &instructionTestVector{
		name: "PLA",
		memory: map[uint16]uint8{
			0x8000:       0x68, // PLA
			0x100 | 0x43: 0x37, // Value to be pulled from stack
		},
		sp:         pointer(uint8(0x42)),
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x37)),
		expectedSP: pointer(uint8(0x43)),
	}

	vector.test(t)
}

func TestCPU_PLP(t *testing.T) {
	vector := &instructionTestVector{
		name: "PLP",
		memory: map[uint16]uint8{
			0x8000:       0x28,       // PLP
			0x100 | 0x43: 0b00111111, // Value to be pulled from stack
		},
		sp:         pointer(uint8(0x42)),
		cycles:     4,
		psMask:     0xFF,       // Check all flags
		expectedPS: 0b00101111, // Expect all flags to be set as per the value pulled from stack
		expectedSP: pointer(uint8(0x43)),
	}

	vector.test(t)
}
