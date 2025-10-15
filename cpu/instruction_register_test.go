package cpu

import (
	"testing"
)

func TestCPU_TAX(t *testing.T) {
	vector := &instructionTestVector{
		name: "TAX",
		memory: map[uint16]uint8{
			0x8000: 0xAA, // TAX
		},
		a:          0x42,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedX:  pointer(uint8(0x42)),
	}

	vector.test(t)
}

func TestCPU_TAY(t *testing.T) {
	vector := &instructionTestVector{
		name: "TAY",
		memory: map[uint16]uint8{
			0x8000: 0xA8, // TAY
		},
		a:          0x42,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedY:  pointer(uint8(0x42)),
	}

	vector.test(t)
}

func TestCPU_TXA(t *testing.T) {
	vector := &instructionTestVector{
		name: "TXA",
		memory: map[uint16]uint8{
			0x8000: 0x8A, // TXA
		},
		x:          0x42,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x42)),
	}

	vector.test(t)
}

func TestCPU_TYA(t *testing.T) {
	vector := &instructionTestVector{
		name: "TYA",
		memory: map[uint16]uint8{
			0x8000: 0x98, // TYA
		},
		y:          0x42,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x42)),
	}

	vector.test(t)
}
