package cpu

import (
	"testing"
)

func TestCPU_NOP(t *testing.T) {
	vector := &instructionTestVector{
		name: "NOP",
		memory: map[uint16]uint8{
			0x8000: 0xEA, // NOP
		},
		cycles: 2,
	}

	vector.test(t)
}
