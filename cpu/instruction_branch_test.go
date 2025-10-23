package cpu

import "testing"

func TestCPU_BCC(t *testing.T) {
	vectors := []*instructionTestVector{
		{
			name: "BCC branch not taken",
			memory: map[uint16]uint8{
				0x8000: 0x90, // BCC
				0x8001: 0x02, // Branch to $8004
			},
			ps:         psFlagCarry,
			cycles:     2,
			expectedPC: pointer(uint16(0x8002)),
		},
		{
			name: "BCC branch taken",
			memory: map[uint16]uint8{
				0x8000: 0x90, // BCC
				0x8001: 0x02, // Branch to $8004
			},
			ps:         0x00,
			cycles:     3,
			expectedPC: pointer(uint16(0x8004)),
		},
		{
			name: "BCC branch taken with page cross",
			memory: map[uint16]uint8{
				0xFFFC: 0xF0, // Reset Vector Low Byte
				0xFFFD: 0x80, // Reset Vector High Byte
				0x80F0: 0x90, // BCC
				0x80F1: 0x10, // Branch to $8102
			},
			ps:         0x00,
			cycles:     4,
			expectedPC: pointer(uint16(0x8102)),
		},
		{
			name: "BCC branch taken backwards",
			memory: map[uint16]uint8{
				0x8000: 0x90, // BCC
				0x8001: 0xFE, // Branch to $8000
			},
			ps:         0x00,
			cycles:     3,
			expectedPC: pointer(uint16(0x8000)),
		},
	}

	for _, vector := range vectors {
		vector.test(t)
	}
}

func TestCPU_BCS(t *testing.T) {
	vectors := []*instructionTestVector{
		{
			name: "BCS branch not taken",
			memory: map[uint16]uint8{
				0x8000: 0xB0, // BCS
				0x8001: 0x02, // Branch to $8004
			},
			ps:         0x00,
			cycles:     2,
			expectedPC: pointer(uint16(0x8002)),
		},
		{
			name: "BCS branch taken",
			memory: map[uint16]uint8{
				0x8000: 0xB0, // BCS
				0x8001: 0x02, // Branch to $8004
			},
			ps:         psFlagCarry,
			cycles:     3,
			expectedPC: pointer(uint16(0x8004)),
		},
		{
			name: "BCS branch taken with page cross",
			memory: map[uint16]uint8{
				0xFFFC: 0xF0, // Reset Vector Low Byte
				0xFFFD: 0x80, // Reset Vector High Byte
				0x80F0: 0xB0, // BCS
				0x80F1: 0x10, // Branch to $8102
			},
			ps:         psFlagCarry,
			cycles:     4,
			expectedPC: pointer(uint16(0x8102)),
		},
	}

	for _, vector := range vectors {
		vector.test(t)
	}
}

func TestCPU_BEQ(t *testing.T) {
	vectors := []*instructionTestVector{
		{
			name: "BEQ branch not taken",
			memory: map[uint16]uint8{
				0x8000: 0xF0, // BEQ
				0x8001: 0x02, // Branch to $8004
			},
			ps:         0x00,
			cycles:     2,
			expectedPC: pointer(uint16(0x8002)),
		},
		{
			name: "BEQ branch taken",
			memory: map[uint16]uint8{
				0x8000: 0xF0, // BEQ
				0x8001: 0x02, // Branch to $8004
			},
			ps:         psFlagZero,
			cycles:     3,
			expectedPC: pointer(uint16(0x8004)),
		},
		{
			name: "BEQ branch taken with page cross",
			memory: map[uint16]uint8{
				0xFFFC: 0xF0, // Reset Vector Low Byte
				0xFFFD: 0x80, // Reset Vector High Byte
				0x80F0: 0xF0, // BEQ
				0x80F1: 0x10, // Branch to $8102
			},
			ps:         psFlagZero,
			cycles:     4,
			expectedPC: pointer(uint16(0x8102)),
		},
	}

	for _, vector := range vectors {
		vector.test(t)
	}
}

func TestCPU_BNE(t *testing.T) {
	vectors := []*instructionTestVector{
		{
			name: "BNE branch not taken",
			memory: map[uint16]uint8{
				0x8000: 0xD0, // BNE
				0x8001: 0x02, // Branch to $8004
			},
			ps:         psFlagZero,
			cycles:     2,
			expectedPC: pointer(uint16(0x8002)),
		},
		{
			name: "BNE branch taken",
			memory: map[uint16]uint8{
				0x8000: 0xD0, // BNE
				0x8001: 0x02, // Branch to $8004
			},
			ps:         0x00,
			cycles:     3,
			expectedPC: pointer(uint16(0x8004)),
		},
		{
			name: "BNE branch taken with page cross",
			memory: map[uint16]uint8{
				0xFFFC: 0xF0, // Reset Vector Low Byte
				0xFFFD: 0x80, // Reset Vector High Byte
				0x80F0: 0xD0, // BNE
				0x80F1: 0x10, // Branch to $8102
			},
			ps:         0x00,
			cycles:     4,
			expectedPC: pointer(uint16(0x8102)),
		},
	}

	for _, vector := range vectors {
		vector.test(t)
	}
}

func TestCPU_BMI(t *testing.T) {
	vectors := []*instructionTestVector{
		{
			name: "BMI branch not taken",
			memory: map[uint16]uint8{
				0x8000: 0x30, // BMI
				0x8001: 0x02, // Branch to $8004
			},
			ps:         0x00,
			cycles:     2,
			expectedPC: pointer(uint16(0x8002)),
		},
		{
			name: "BMI branch taken",
			memory: map[uint16]uint8{
				0x8000: 0x30, // BMI
				0x8001: 0x02, // Branch to $8004
			},
			ps:         psFlagNegative,
			cycles:     3,
			expectedPC: pointer(uint16(0x8004)),
		},
		{
			name: "BMI branch taken with page cross",
			memory: map[uint16]uint8{
				0xFFFC: 0xF0, // Reset Vector Low Byte
				0xFFFD: 0x80, // Reset Vector High Byte
				0x80F0: 0x30, // BMI
				0x80F1: 0x10, // Branch to $8102
			},
			ps:         psFlagNegative,
			cycles:     4,
			expectedPC: pointer(uint16(0x8102)),
		},
	}

	for _, vector := range vectors {
		vector.test(t)
	}
}

func TestCPU_BPL(t *testing.T) {
	vectors := []*instructionTestVector{
		{
			name: "BPL branch not taken",
			memory: map[uint16]uint8{
				0x8000: 0x10, // BPL
				0x8001: 0x02, // Branch to $8004
			},
			ps:         psFlagNegative,
			cycles:     2,
			expectedPC: pointer(uint16(0x8002)),
		},
		{
			name: "BPL branch taken",
			memory: map[uint16]uint8{
				0x8000: 0x10, // BPL
				0x8001: 0x02, // Branch to $8004
			},
			ps:         0x00,
			cycles:     3,
			expectedPC: pointer(uint16(0x8004)),
		},
		{
			name: "BPL branch taken with page cross",
			memory: map[uint16]uint8{
				0xFFFC: 0xF0, // Reset Vector Low Byte
				0xFFFD: 0x80, // Reset Vector High Byte
				0x80F0: 0x10, // BPL
				0x80F1: 0x10, // Branch to $8102
			},
			ps:         0x00,
			cycles:     4,
			expectedPC: pointer(uint16(0x8102)),
		},
	}

	for _, vector := range vectors {
		vector.test(t)
	}
}

func TestCPU_BVC(t *testing.T) {
	vectors := []*instructionTestVector{
		{
			name: "BVC branch not taken",
			memory: map[uint16]uint8{
				0x8000: 0x50, // BVC
				0x8001: 0x02, // Branch to $8004
			},
			ps:         psFlagOverflow,
			cycles:     2,
			expectedPC: pointer(uint16(0x8002)),
		},
		{
			name: "BVC branch taken",
			memory: map[uint16]uint8{
				0x8000: 0x50, // BVC
				0x8001: 0x02, // Branch to $8004
			},
			ps:         0x00,
			cycles:     3,
			expectedPC: pointer(uint16(0x8004)),
		},
		{
			name: "BVC branch taken with page cross",
			memory: map[uint16]uint8{
				0xFFFC: 0xF0, // Reset Vector Low Byte
				0xFFFD: 0x80, // Reset Vector High Byte
				0x80F0: 0x50, // BVC
				0x80F1: 0x10, // Branch to $8102
			},
			ps:         0x00,
			cycles:     4,
			expectedPC: pointer(uint16(0x8102)),
		},
	}

	for _, vector := range vectors {
		vector.test(t)
	}
}

func TestCPU_BVS(t *testing.T) {
	vectors := []*instructionTestVector{
		{
			name: "BVS branch not taken",
			memory: map[uint16]uint8{
				0x8000: 0x70, // BVS
				0x8001: 0x02, // Branch to $8004
			},
			ps:         0x00,
			cycles:     2,
			expectedPC: pointer(uint16(0x8002)),
		},
		{
			name: "BVS branch taken",
			memory: map[uint16]uint8{
				0x8000: 0x70, // BVS
				0x8001: 0x02, // Branch to $8004
			},
			ps:         psFlagOverflow,
			cycles:     3,
			expectedPC: pointer(uint16(0x8004)),
		},
		{
			name: "BVS branch taken with page cross",
			memory: map[uint16]uint8{
				0xFFFC: 0xF0, // Reset Vector Low Byte
				0xFFFD: 0x80, // Reset Vector High Byte
				0x80F0: 0x70, // BVS
				0x80F1: 0x10, // Branch to $8102
			},
			ps:         psFlagOverflow,
			cycles:     4,
			expectedPC: pointer(uint16(0x8102)),
		},
	}

	for _, vector := range vectors {
		vector.test(t)
	}
}
