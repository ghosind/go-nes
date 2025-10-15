package cpu

import (
	"testing"
)

func TestCPU_JMP_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "JMP Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x4C, // JMP Absolute
			0x8001: 0x00, // Low byte of target address
			0x8002: 0x90, // High byte of target address
		},
		cycles:     3,
		expectedPC: pointer(uint16(0x9000)),
	}

	vector.test(t)
}

func TestCPU_JMP_IND(t *testing.T) {
	vector := &instructionTestVector{
		name: "JMP Indirect",
		memory: map[uint16]uint8{
			0x8000: 0x6C, // JMP Indirect
			0x8001: 0x00, // Low byte of pointer address
			0x8002: 0x90, // High byte of pointer address
			0x9000: 0x34, // Low byte of target address
			0x9001: 0x12, // High byte of target address
		},
		cycles:     5,
		expectedPC: pointer(uint16(0x1234)),
	}

	vector.test(t)
}

func TestCPU_JSR(t *testing.T) {
	vector := &instructionTestVector{
		name: "JSR",
		memory: map[uint16]uint8{
			0x8000: 0x20, // JSR
			0x8001: 0x00, // Low byte of target address
			0x8002: 0x90, // High byte of target address
		},
		cycles: 6,
		expectedMem: map[uint16]uint8{
			0x01FD: 0x80, // High byte of return address
			0x01FC: 0x02, // Low byte of return address
		},
		expectedPC: pointer(uint16(0x9000)),
		expectedSP: pointer(uint8(0xFB)), // Stack pointer should be decremented by 2
	}

	vector.test(t)
}

func TestCPU_RTS(t *testing.T) {
	vector := &instructionTestVector{
		name: "RTS",
		memory: map[uint16]uint8{
			0x8000: 0x60, // RTS
			0x01FC: 0x02, // Low byte of return address
			0x01FD: 0x80, // High byte of return address
		},
		sp:         pointer(uint8(0xFB)), // Stack pointer set to point to the return address
		cycles:     6,
		expectedPC: pointer(uint16(0x8003)), // PC should be set to return address + 1
		expectedSP: pointer(uint8(0xFD)),    // Stack pointer should be incremented by 2
	}

	vector.test(t)
}
