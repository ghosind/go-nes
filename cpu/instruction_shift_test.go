package cpu

import "testing"

func TestCPU_ASL_ACC(t *testing.T) {
	vector := &instructionTestVector{
		name: "ASL Accumulator",
		memory: map[uint16]uint8{
			0x8000: 0x0A, // ASL Accumulator
		},
		a:          0x01,
		cycles:     2,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_ASL_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "ASL Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x06, // ASL Zero Page
			0x8001: 0x10, // Address $10
			0x0010: 0x01, // Value at $10
		},
		cycles:     5,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0010: 0x02,
		},
	}

	vector.test(t)
}

func TestCPU_ASL_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ASL Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0x16, // ASL Zero Page, X
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
	}

	vector.test(t)
}

func TestCPU_ASL_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "ASL Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x0E, // ASL Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($2000)
			0x1000: 0x01, // Value at $2000
		},
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x02,
		},
	}

	vector.test(t)
}

func TestCPU_ASL_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ASL Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x1E, // ASL Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($2000)
			0x1005: 0x01, // Value at $2000 + X (X=5)
		},
		x:          0x05,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x02,
		},
	}

	vector.test(t)
}

func TestCPU_LSR_ACC(t *testing.T) {
	vector := &instructionTestVector{
		name: "LSR Accumulator",
		memory: map[uint16]uint8{
			0x8000: 0x4A, // LSR Accumulator
		},
		a:          0x02,
		cycles:     2,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x01)),
	}

	vector.test(t)
}

func TestCPU_LSR_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "LSR Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x46, // LSR Zero Page
			0x8001: 0x10, // Address $10
			0x0010: 0x02, // Value at $10
		},
		cycles:     5,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0010: 0x01,
		},
	}

	vector.test(t)
}

func TestCPU_LSR_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "LSR Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0x56, // LSR Zero Page, X
			0x8001: 0x10, // Address $10
			0x0015: 0x02, // Value at $10 + X (X=5)
		},
		x:          0x05,
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0015: 0x01,
		},
	}

	vector.test(t)
}

func TestCPU_LSR_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "LSR Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x4E, // LSR Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($2000)
			0x1000: 0x02, // Value at $2000
		},
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x01,
		},
	}

	vector.test(t)
}

func TestCPU_LSR_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "LSR Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x5E, // LSR Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($2000)
			0x1005: 0x02, // Value at $2000 + X (X=5)
		},
		x:          0x05,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x01,
		},
	}

	vector.test(t)
}

func TestCPU_ROL_ACC(t *testing.T) {
	vector := &instructionTestVector{
		name: "ROL Accumulator",
		memory: map[uint16]uint8{
			0x8000: 0x2A, // ROL Accumulator
		},
		a:          0x01,
		cycles:     2,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x02)),
	}

	vector.test(t)
}

func TestCPU_ROL_ACC_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "ROL Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x26, // ROL Zero Page
			0x8001: 0x10, // Address $10
			0x0010: 0x01, // Value at $10
		},
		cycles:     5,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0010: 0x02,
		},
	}

	vector.test(t)
}

func TestCPU_ROL_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ROL Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0x36, // ROL Zero Page, X
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
	}

	vector.test(t)
}

func TestCPU_ROL_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "ROL Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x2E, // ROL Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($2000)
			0x1000: 0x01, // Value at $2000
		},
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x02,
		},
	}

	vector.test(t)
}

func TestCPU_ROL_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ROL Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x3E, // ROL Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($2000)
			0x1005: 0x01, // Value at $2000 + X (X=5)
		},
		x:          0x05,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x02,
		},
	}

	vector.test(t)
}

func TestCPU_ROR_ACC(t *testing.T) {
	vector := &instructionTestVector{
		name: "ROR Accumulator",
		memory: map[uint16]uint8{
			0x8000: 0x6A, // ROR Accumulator
		},
		a:          0x02,
		cycles:     2,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x01)),
	}

	vector.test(t)
}

func TestCPU_ROR_ACC_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "ROR Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x66, // ROR Zero Page
			0x8001: 0x10, // Address $10
			0x0010: 0x02, // Value at $10
		},
		cycles:     5,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0010: 0x01,
		},
	}

	vector.test(t)
}

func TestCPU_ROR_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ROR Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0x76, // ROR Zero Page, X
			0x8001: 0x10, // Address $10
			0x0015: 0x02, // Value at $10 + X (X=5)
		},
		x:          0x05,
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0015: 0x01,
		},
	}

	vector.test(t)
}

func TestCPU_ROR_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "ROR Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x6E, // ROR Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($2000)
			0x1000: 0x02, // Value at $2000
		},
		cycles:     6,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x01,
		},
	}

	vector.test(t)
}

func TestCPU_ROR_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ROR Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x7E, // ROR Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x10, // High byte of address ($2000)
			0x1005: 0x02, // Value at $2000 + X (X=5)
		},
		x:          0x05,
		cycles:     7,
		psMask:     psFlagCarry | psFlagNegative | psFlagZero,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x01,
		},
	}

	vector.test(t)
}
