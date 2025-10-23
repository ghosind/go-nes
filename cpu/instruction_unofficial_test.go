package cpu

import "testing"

func TestCPU_LAX_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "LAX Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0xA7, // LAX Zero Page
			0x8001: 0x10, // Address $10
			0x0010: 0x42, // Value at $10
		},
		cycles:     3,
		psMask:     psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x42)),
		expectedX:  pointer(uint8(0x42)),
	}

	vector.test(t)
}

func TestCPU_LAX_ZP_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "LAX Zero Page, Y",
		memory: map[uint16]uint8{
			0x8000: 0xB7, // LAX Zero Page, Y
			0x8001: 0x10, // Address $10
			0x0015: 0x37, // Value at $10 + Y (Y=5)
		},
		y:          0x05,
		cycles:     4,
		psMask:     psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x37)),
		expectedX:  pointer(uint8(0x37)),
	}

	vector.test(t)
}

func TestCPU_LAX_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "LAX Absolute",
		memory: map[uint16]uint8{
			0x8000: 0xAF, // LAX Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address
			0x1000: 0x99, // Value at $1000
		},
		cycles:     4,
		psMask:     psFlagNegative | psFlagZero,
		expectedPS: psFlagNegative,
		expectedA:  pointer(uint8(0x99)),
		expectedX:  pointer(uint8(0x99)),
	}

	vector.test(t)
}

func TestCPU_LAX_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "LAX Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0xBF, // LAX Absolute, Y
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address
			0x1005: 0x77, // Value at $1000 + Y (Y=5)
		},
		y:          0x05,
		cycles:     4,
		psMask:     psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x77)),
		expectedX:  pointer(uint8(0x77)),
	}

	vector.test(t)
}

func TestCPU_LAX_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "LAX (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0xA3, // LAX (Indirect, X)
			0x8001: 0x10, // Address $10
			0x0015: 0x00, // Low byte of target address (0x1000) at $10 + X (X=5)
			0x0016: 0x10, // High byte of target address at $11 + X
			0x1000: 0x66, // Value at $1000
		},
		x:          0x05,
		cycles:     6,
		psMask:     psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x66)),
		expectedX:  pointer(uint8(0x66)),
	}

	vector.test(t)
}

func TestCPU_LAX_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "LAX (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0xB3, // LAX (Indirect), Y
			0x8001: 0x10, // Address $10
			0x0010: 0x00, // Low byte of base address (0x1000) at $10
			0x0011: 0x10, // High byte of base address at $11
			0x1005: 0x88, // Value at $1000 + Y (Y=5)
		},
		y:          0x05,
		cycles:     5,
		psMask:     psFlagNegative | psFlagZero,
		expectedPS: psFlagNegative,
		expectedA:  pointer(uint8(0x88)),
		expectedX:  pointer(uint8(0x88)),
	}

	vector.test(t)
}

func TestCPU_SAX_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "SAX Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x87, // SAX Zero Page
			0x8001: 0x10, // Address $10
		},
		a:      0xAA,
		x:      0x0F,
		cycles: 3,
		expectedMem: map[uint16]uint8{
			0x0010: 0x0A, // A & X = 0xAA & 0x0F = 0x0A
		},
	}

	vector.test(t)
}

func TestCPU_SAX_ZP_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "SAX Zero Page, Y",
		memory: map[uint16]uint8{
			0x8000: 0x97, // SAX Zero Page, Y
			0x8001: 0x10, // Address $10
		},
		a:      0xFF,
		x:      0xF0,
		y:      0x05,
		cycles: 4,
		expectedMem: map[uint16]uint8{
			0x0015: 0xF0, // A & X = 0xFF & 0xF0 = 0xF0
		},
	}

	vector.test(t)
}

func TestCPU_SAX_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "SAX Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x8F, // SAX Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address
		},
		a:      0x55,
		x:      0x0F,
		cycles: 4,
		expectedMem: map[uint16]uint8{
			0x1000: 0x05, // A & X = 0x55 & 0x0F = 0x05
		},
	}

	vector.test(t)
}

func TestCPU_SAX_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "SAX (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0x83, // SAX (Indirect, X)
			0x8001: 0x10, // Address $10
			0x0015: 0x00, // Low byte of target address (0x1000) at $10 + X (X=5)
			0x0016: 0x10, // High byte of target address at $11 + X
		},
		a:      0xCC,
		x:      0x33,
		cycles: 6,
		expectedMem: map[uint16]uint8{
			0x1000: 0x00, // A & X = 0xCC & 0x33 = 0x00
		},
	}

	vector.test(t)
}

func TestCPU_DCP_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "DCP Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0xC7, // DCP Zero Page
			0x8001: 0x10, // Address $10
			0x0010: 0x05, // Value at $10
		},
		a:      0x04,
		cycles: 5,
		expectedMem: map[uint16]uint8{
			0x0010: 0x04, // Decremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_DCP_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "DCP Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0xD7, // DCP Zero Page, X
			0x8001: 0x10, // Address $10
			0x0015: 0x03, // Value at $10 + X (X=5)
		},
		a:      0x02,
		x:      0x05,
		cycles: 6,
		expectedMem: map[uint16]uint8{
			0x0015: 0x02, // Decremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_DCP_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "DCP Absolute",
		memory: map[uint16]uint8{
			0x8000: 0xCF, // DCP Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address
			0x1000: 0x02, // Value at $1000
		},
		a:      0x01,
		cycles: 6,
		expectedMem: map[uint16]uint8{
			0x1000: 0x01, // Decremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_DCP_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "DCP Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0xDF, // DCP Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address
			0x1005: 0x01, // Value at $1000 + X (X=5)
		},
		a:      0x00,
		x:      0x05,
		cycles: 7,
		expectedMem: map[uint16]uint8{
			0x1005: 0x00, // Decremented value
		},
		expectedPS: psFlagCarry | psFlagZero,
	}

	vector.test(t)
}

func TestCPU_DCP_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "DCP Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0xDB, // DCP Absolute, Y
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address
			0x1005: 0x03, // Value at $1000 + Y (Y=5)
		},
		a:      0x02,
		y:      0x05,
		cycles: 7,
		expectedMem: map[uint16]uint8{
			0x1005: 0x02, // Decremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_DCP_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "DCP (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0xC3, // DCP (Indirect, X)
			0x8001: 0x10, // Address $10
			0x0015: 0x00, // Low byte of target address (0x1000) at $10 + X (X=5)
			0x0016: 0x10, // High byte of target address at $11 + X
			0x1000: 0x04, // Value at $1000
		},
		a:      0x03,
		x:      0x05,
		cycles: 8,
		expectedMem: map[uint16]uint8{
			0x1000: 0x03, // Decremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_DCP_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "DCP (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0xD3, // DCP (Indirect), Y
			0x8001: 0x10, // Address $10
			0x0010: 0x00, // Low byte of base address (0x1000) at $10
			0x0011: 0x10, // High byte of base address at $11
			0x1005: 0x05, // Value at $1000 + Y (Y=5)
		},
		a:      0x04,
		y:      0x05,
		cycles: 8,
		expectedMem: map[uint16]uint8{
			0x1005: 0x04, // Decremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_ISB_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "ISB Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0xE7, // ISB Zero Page
			0x8001: 0x10, // Address $10
			0x0010: 0x05, // Value at $10
		},
		a:      0x06,
		cycles: 5,
		expectedMem: map[uint16]uint8{
			0x0010: 0x06, // Incremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_ISB_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ISB Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0xF7, // ISB Zero Page, X
			0x8001: 0x10, // Address $10
			0x0015: 0x03, // Value at $10 + X (X=5)
		},
		a:      0x04,
		x:      0x05,
		cycles: 6,
		expectedMem: map[uint16]uint8{
			0x0015: 0x04, // Incremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_ISB_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "ISB Absolute",
		memory: map[uint16]uint8{
			0x8000: 0xEF, // ISB Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address
			0x1000: 0x02, // Value at $1000
		},
		a:      0x03,
		cycles: 6,
		expectedMem: map[uint16]uint8{
			0x1000: 0x03, // Incremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_ISB_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ISB Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0xFF, // ISB Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address
			0x1005: 0x01, // Value at $1000 + X (X=5)
		},
		a:      0x02,
		x:      0x05,
		cycles: 7,
		expectedMem: map[uint16]uint8{
			0x1005: 0x02, // Incremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_ISB_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "ISB Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0xFB, // ISB Absolute, Y
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address
			0x1005: 0x03, // Value at $1000 + Y (Y=5)
		},
		a:      0x04,
		y:      0x05,
		cycles: 7,
		expectedMem: map[uint16]uint8{
			0x1005: 0x04, // Incremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_ISB_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ISB (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0xE3, // ISB (Indirect, X)
			0x8001: 0x10, // Address $10
			0x0015: 0x00, // Low byte of target address (0x1000) at $10 + X (X=5)
			0x0016: 0x10, // High byte of target address at $11 + X
			0x1000: 0x04, // Value at $1000
		},
		a:      0x05,
		x:      0x05,
		cycles: 8,
		expectedMem: map[uint16]uint8{
			0x1000: 0x05, // Incremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_ISB_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "ISB (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0xF3, // ISB (Indirect), Y
			0x8001: 0x10, // Address $10
			0x0010: 0x00, // Low byte of base address (0x1000) at $10
			0x0011: 0x10, // High byte of base address at $11
			0x1005: 0x05, // Value at $1000 + Y (Y=5)
		},
		a:      0x06,
		y:      0x05,
		cycles: 8,
		expectedMem: map[uint16]uint8{
			0x1005: 0x06, // Incremented value
		},
		expectedPS: psFlagCarry,
	}

	vector.test(t)
}

func TestCPU_SLO_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "SLO Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x07, // SLO Zero Page
			0x8001: 0x10, // Address $10
			0x0010: 0x01, // Value at $10
		},
		cycles:     5,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0010: 0x02,
		},
		expectedA: pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_SLO_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "SLO Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0x17, // SLO Zero Page, X
			0x8001: 0x10, // Address $10
			0x0015: 0x01, // Value at $10 + X (X=5)
		},
		x:          0x05,
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0015: 0x02,
		},
		expectedA: pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_SLO_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "SLO Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x0F, // SLO Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1000: 0x01, // Value at $1000
		},
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x02,
		},
		expectedA: pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_SLO_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "SLO Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x1F, // SLO Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1005: 0x01, // Value at $1000 + X (X=5)
		},
		x:          0x05,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x02,
		},
		expectedA: pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_SLO_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "SLO Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0x1B, // SLO Absolute, Y
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1005: 0x01, // Value at $1000 + Y (Y=5)
		},
		y:          0x05,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x02,
		},
		expectedA: pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_SLO_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "SLO (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0x03, // SLO (Indirect, X)
			0x8001: 0x10, // Address $10
			0x0015: 0x00, // Low byte of target address (0x1000) at $10 + X (X=5)
			0x0016: 0x10, // High byte of target address at $11 + X
			0x1000: 0x01, // Value at $1000
		},
		x:          0x05,
		cycles:     8,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x02,
		},
		expectedA: pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_SLO_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "SLO (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0x13, // SLO (Indirect), Y
			0x8001: 0x10, // Address $10
			0x0010: 0x00, // Low byte of base address (0x1000) at $10
			0x0011: 0x10, // High byte of base address at $11
			0x1005: 0x01, // Value at $1000 + Y (Y=5)
		},
		y:          0x05,
		cycles:     8,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x02,
		},
		expectedA: pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_RLA_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "RLA Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x27, // RLA Zero Page
			0x8001: 0x10, // Address $10
			0x0010: 0x80, // Value at $10
		},
		a:          0x40,
		cycles:     5,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: psFlagCarry | psFlagZero,
		expectedMem: map[uint16]uint8{
			0x0010: 0x00,
		},
		expectedA: pointer(uint8(0x00)),
	}

	vector.test(t)
}

func TestCPU_RLA_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "RLA Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0x37, // RLA Zero Page, X
			0x8001: 0x10, // Address $10
			0x0015: 0x01, // Value at $10 + X (X=5)
		},
		x:          0x05,
		a:          0x7F,
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0015: 0x02,
		},
		expectedA: pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_RLA_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "RLA Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x2F, // RLA Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1000: 0x40, // Value at $1000
		},
		a:          0x20,
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: psFlagZero,
		expectedMem: map[uint16]uint8{
			0x1000: 0x80,
		},
		expectedA: pointer(uint8(0x00)),
	}

	vector.test(t)
}

func TestCPU_RLA_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "RLA Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x3F, // RLA Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1005: 0x01, // Value at $1000 + X (X=5)
		},
		x:          0x05,
		a:          0x7F,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x02,
		},
		expectedA: pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_RLA_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "RLA Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0x3B, // RLA Absolute, Y
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1005: 0x01, // Value at $1000 + Y (Y=5)
		},
		y:          0x05,
		a:          0x7F,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x02,
		},
		expectedA: pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_RLA_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "RLA (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0x23, // RLA (Indirect, X)
			0x8001: 0x10, // Address $10
			0x0015: 0x00, // Low byte of target address (0x1000) at $10 + X (X=5)
			0x0016: 0x10, // High byte of target address at $11 + X
			0x1000: 0x01, // Value at $1000
		},
		x:          0x05,
		a:          0x7F,
		cycles:     8,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x02,
		},
		expectedA: pointer(uint8(02)),
	}

	vector.test(t)
}

func TestCPU_RLA_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "RLA (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0x33, // RLA (Indirect), Y
			0x8001: 0x10, // Address $10
			0x0010: 0x00, // Low byte of base address (0x1000) at $10
			0x0011: 0x10, // High byte of base address at $11
			0x1005: 0x01, // Value at $1000 + Y (Y=5)
		},
		y:          0x05,
		a:          0x7F,
		cycles:     8,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x02,
		},
		expectedA: pointer(uint8(02)),
	}

	vector.test(t)
}

func TestCPU_SRE_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "SRE Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x47, // SRE Zero Page
			0x8001: 0x10, // Address $10
			0x0010: 0x04, // Value at $10
		},
		cycles:     5,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0010: 0x02,
		},
		expectedA: pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_SRE_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "SRE Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0x57, // SRE Zero Page, X
			0x8001: 0x10, // Address $10
			0x0015: 0x08, // Value at $10 + X (X=5)
		},
		x:          0x05,
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0015: 0x04,
		},
		expectedA: pointer(uint8(0x04)),
	}

	vector.test(t)
}

func TestCPU_SRE_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "SRE Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x4F, // SRE Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1000: 0x10, // Value at $1000
		},
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x08,
		},
		expectedA: pointer(uint8(0x08)),
	}

	vector.test(t)
}

func TestCPU_SRE_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "SRE Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x5F, // SRE Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1005: 0x20, // Value at $1000 + X (X=5)
		},
		x:          0x05,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x10,
		},
		expectedA: pointer(uint8(0x10)),
	}

	vector.test(t)
}

func TestCPU_SRE_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "SRE Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0x5B, // SRE Absolute, Y
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1005: 0x20, // Value at $1000 + Y (Y=5)
		},
		y:          0x05,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x10,
		},
		expectedA: pointer(uint8(0x10)),
	}

	vector.test(t)
}

func TestCPU_SRE_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "SRE (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0x43, // SRE (Indirect, X)
			0x8001: 0x10, // Address $10
			0x0015: 0x00, // Low byte of target address (0x1000) at $10 + X (X=5)
			0x0016: 0x10, // High byte of target address at $11 + X
			0x1000: 0x10, // Value at $1000
		},
		x:          0x05,
		cycles:     8,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x08,
		},
		expectedA: pointer(uint8(0x08)),
	}

	vector.test(t)
}

func TestCPU_SRE_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "SRE (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0x53, // SRE (Indirect), Y
			0x8001: 0x10, // Address $10
			0x0010: 0x00, // Low byte of base address (0x1000) at $10
			0x0011: 0x10, // High byte of base address at $11
			0x1005: 0x10, // Value at $1000 + Y (Y=5)
		},
		y:          0x05,
		cycles:     8,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x08,
		},
		expectedA: pointer(uint8(0x08)),
	}

	vector.test(t)
}

func TestCPU_RRA_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "RRA Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x67, // RRA Zero Page
			0x8001: 0x10, // Address $10
			0x0010: 0x14, // Value at $10
		},
		a:          0x10,
		cycles:     5,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0010: 0x0A,
		},
		expectedA: pointer(uint8(0x1A)),
	}

	vector.test(t)
}

func TestCPU_RRA_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "RRA Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0x77, // RRA Zero Page, X
			0x8001: 0x10, // Address $10
			0x0015: 0x08, // Value at $10 + X (X=5)
		},
		x:          0x05,
		a:          0x04,
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0015: 0x04,
		},
		expectedA: pointer(uint8(0x08)),
	}

	vector.test(t)
}

func TestCPU_RRA_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "RRA Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x6F, // RRA Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1000: 0x20, // Value at $1000
		},
		a:          0x10,
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x10,
		},
		expectedA: pointer(uint8(0x20)),
	}

	vector.test(t)
}

func TestCPU_RRA_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "RRA Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x7F, // RRA Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1005: 0x08, // Value at $1000 + X (X=5)
		},
		x:          0x05,
		a:          0x04,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x04,
		},
		expectedA: pointer(uint8(0x08)),
	}

	vector.test(t)
}

func TestCPU_RRA_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "RRA Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0x7B, // RRA Absolute, Y
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($1000)
			0x1005: 0x08, // Value at $1000 + Y (Y=5)
		},
		y:          0x05,
		a:          0x04,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x04,
		},
		expectedA: pointer(uint8(0x08)),
	}

	vector.test(t)
}

func TestCPU_RRA_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "RRA (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0x63, // RRA (Indirect, X)
			0x8001: 0x10, // Address $10
			0x0015: 0x00, // Low byte of target address (0x1000) at $10 + X (X=5)
			0x0016: 0x10, // High byte of target address at $11 + X
			0x1000: 0x08, // Value at $1000
		},
		x:          0x05,
		a:          0x04,
		cycles:     8,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x04,
		},
		expectedA: pointer(uint8(0x08)),
	}

	vector.test(t)
}

func TestCPU_RRA_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "RRA (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0x73, // RRA (Indirect), Y
			0x8001: 0x10, // Address $10
			0x0010: 0x00, // Low byte of base address (0x1000) at $10
			0x0011: 0x10, // High byte of base address at $11
			0x1005: 0x08, // Value at $1000 + Y (Y=5)
		},
		y:          0x05,
		a:          0x04,
		cycles:     8,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x04,
		},
		expectedA: pointer(uint8(0x08)),
	}

	vector.test(t)
}
