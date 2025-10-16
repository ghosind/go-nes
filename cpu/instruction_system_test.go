package cpu

import (
	"testing"
)

func TestCPU_BRK(t *testing.T) {
	vector := &instructionTestVector{
		name: "BRK",
		memory: map[uint16]uint8{
			0x8000: 0x00, // BRK
			0xFFFE: 0x10, // IRQ/BRK vector low byte
			0xFFFF: 0x80, // IRQ/BRK vector high byte
		},
		cycles:     7,
		psMask:     psFlagInterrupt,
		expectedPS: psFlagInterrupt,
		expectedPC: pointer(uint16(0x8010)),
		expectedMem: map[uint16]uint8{
			0x01FD: 0x80,
			0x01FC: 0x02,
			0x01FB: uint8(psFlagBreak | psFlagUnused),
		},
		expectedSP: pointer(uint8(0xFA)),
	}

	vector.test(t)
}

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

func TestCPU_RTI(t *testing.T) {
	vector := &instructionTestVector{
		name: "RTI",
		memory: map[uint16]uint8{
			0x8000: 0x40,                              // RTI
			0x01FD: 0x80,                              // PCH
			0x01FC: 0x10,                              // PCL
			0x01FB: uint8(psFlagBreak | psFlagUnused), // Status
		},
		sp:         pointer(uint8(0xFA)),
		cycles:     6,
		expectedPC: pointer(uint16(0x8010)),
		expectedPS: psFlagBreak | psFlagUnused,
		expectedSP: pointer(uint8(0xFD)),
	}

	vector.test(t)
}
