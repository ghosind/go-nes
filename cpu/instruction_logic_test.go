package cpu

import (
	"testing"
)

func TestCPU_AND_IMM(t *testing.T) {
	vector := &instructionTestVector{
		name: "AND Immediate",
		memory: map[uint16]uint8{
			0x8000: 0x29, // AND immediate
			0x8001: 0x0F, // Operand
		},
		a:          0x3C,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	vector.test(t)
}

func TestCPU_AND_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "AND Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x25, // AND zero page
			0x8001: 0x10, // Operand
			0x0010: 0x0F, // Value at zero page address
		},
		a:          0x3C,
		cycles:     3,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	vector.test(t)
}

func TestCPU_AND_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "AND Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0x35, // AND zero page, X
			0x8001: 0x10, // Operand
			0x0015: 0x0F, // Value at zero page address + X (use addition for clarity)
		},
		a:          0x3C,
		x:          0x05,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	vector.test(t)
}

func TestCPU_AND_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "AND Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x2D, // AND absolute
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x10, // High byte of operand
			0x1000: 0x0F, // Value at absolute address
		},
		a:          0x3C,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	vector.test(t)
}

func TestCPU_AND_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "AND Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x3D, // AND absolute, X
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x10, // High byte of operand
			0x1005: 0x0F, // Value at absolute address + X
		},
		a:          0x3C,
		x:          0x05,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	vector.test(t)
}

func TestCPU_AND_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "AND Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0x39, // AND absolute, Y
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x10, // High byte of operand
			0x1005: 0x0F, // Value at absolute address + Y
		},
		a:          0x3C,
		y:          0x05,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	vector.test(t)
}

func TestCPU_AND_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "AND (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0x21, // AND (indirect, X)
			0x8001: 0x10, // Operand
			0x0015: 0x00, // Low byte of effective address (use addition for clarity)
			0x0016: 0x10, // High byte of effective address
			0x1000: 0x0F, // Value at effective address
		},
		a:          0x3C,
		x:          0x05,
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	vector.test(t)
}

func TestCPU_AND_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "AND (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0x31, // AND (indirect), Y
			0x8001: 0x10, // Operand
			0x0010: 0x00, // Low byte of effective address
			0x0011: 0x10, // High byte of effective address
			0x1005: 0x0F, // Value at effective address + Y
		},
		a:          0x3C,
		y:          0x05,
		cycles:     5,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	vector.test(t)
}

func TestCPU_EOR_IMM(t *testing.T) {
	vector := &instructionTestVector{
		name: "EOR Immediate",
		memory: map[uint16]uint8{
			0x8000: 0x49, // EOR immediate
			0x8001: 0x0F, // Operand
		},
		a:          0x3C,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	vector.test(t)
}

func TestCPU_EOR_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "EOR Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x45, // EOR zero page
			0x8001: 0x10, // Operand
			0x0010: 0x0F, // Value at zero page address
		},
		a:          0x3C,
		cycles:     3,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	vector.test(t)
}

func TestCPU_EOR_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "EOR Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0x55, // EOR zero page, X
			0x8001: 0x10, // Operand
			0x0015: 0x0F, // Value at zero page address + X
		},
		a:          0x3C,
		x:          0x05,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	vector.test(t)
}

func TestCPU_EOR_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "EOR Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x4D, // EOR absolute
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x10, // High byte of operand
			0x1000: 0x0F, // Value at absolute address
		},
		a:          0x3C,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	vector.test(t)
}

func TestCPU_EOR_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "EOR Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x5D, // EOR absolute, X
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x10, // High byte of operand
			0x1005: 0x0F, // Value at absolute address + X
		},
		a:          0x3C,
		x:          0x05,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	vector.test(t)
}

func TestCPU_EOR_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "EOR Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0x59, // EOR absolute, Y
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x10, // High byte of operand
			0x1005: 0x0F, // Value at absolute address + Y
		},
		a:          0x3C,
		y:          0x05,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	vector.test(t)
}

func TestCPU_EOR_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "EOR (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0x41, // EOR (indirect, X)
			0x8001: 0x10, // Operand
			0x0015: 0x00, // Low byte of effective address
			0x0016: 0x10, // High byte of effective address
			0x1000: 0x0F, // Value at effective address
		},
		a:          0x3C,
		x:          0x05,
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	vector.test(t)
}

func TestCPU_EOR_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "EOR (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0x51, // EOR (indirect), Y
			0x8001: 0x10, // Operand
			0x0010: 0x00, // Low byte of effective address
			0x0011: 0x10, // High byte of effective address
			0x1005: 0x0F, // Value at effective address + Y
		},
		a:          0x3C,
		y:          0x05,
		cycles:     5,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	vector.test(t)
}

func TestCPU_ORA_IMM(t *testing.T) {
	vector := &instructionTestVector{
		name: "ORA Immediate",
		memory: map[uint16]uint8{
			0x8000: 0x09, // ORA immediate
			0x8001: 0x0F, // Operand
		},
		a:          0x30,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x3F)),
	}

	vector.test(t)
}

func TestCPU_ORA_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "ORA Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x05, // ORA zero page
			0x8001: 0x10, // Operand
			0x0010: 0x0F, // Value at zero page address
		},
		a:          0x30,
		cycles:     3,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x3F)),
	}

	vector.test(t)
}

func TestCPU_ORA_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ORA Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0x15, // ORA zero page, X
			0x8001: 0x10, // Operand
			0x0015: 0x0F, // Value at zero page address + X
		},
		a:          0x30,
		x:          0x05,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x3F)),
	}

	vector.test(t)
}

func TestCPU_ORA_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "ORA Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x0D, // ORA absolute
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x10, // High byte of operand
			0x1000: 0x0F, // Value at absolute address
		},
		a:          0x30,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x3F)),
	}

	vector.test(t)
}

func TestCPU_ORA_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ORA Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x1D, // ORA absolute,X
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x10, // High byte of operand
			0x1005: 0x0F, // Value at absolute address + X
		},
		a:          0x30,
		x:          0x05,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x3F)),
	}

	vector.test(t)
}

func TestCPU_ORA_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "ORA Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0x19, // ORA absolute, Y
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x10, // High byte of operand
			0x1005: 0x0F, // Value at absolute address + Y
		},
		a:          0x30,
		y:          0x05,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x3F)),
	}

	vector.test(t)
}

func TestCPU_ORA_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ORA (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0x01, // ORA (indirect, X)
			0x8001: 0x10, // Operand
			0x0015: 0x00, // Low byte of effective address
			0x0016: 0x10, // High byte of effective address
			0x1000: 0x0F, // Value at effective address
		},
		a:          0x30,
		x:          0x05,
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x3F)),
	}

	vector.test(t)
}

func TestCPU_ORA_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "ORA (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0x11, // ORA (indirect), Y
			0x8001: 0x10, // Operand
			0x0010: 0x00, // Low byte of effective address
			0x0011: 0x10, // High byte of effective address
			0x1005: 0x0F, // Value at effective address + Y
		},
		a:          0x30,
		y:          0x05,
		cycles:     5,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x3F)),
	}

	vector.test(t)
}

func TestCPU_BIT_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "BIT Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x24, // BIT zero page
			0x8001: 0x10, // Operand
			0x0010: 0xC0, // Value at zero page address (sets N and V)
		},
		a:          0x40,
		cycles:     3,
		psMask:     psFlagZero | psFlagNegative | psFlagOverflow,
		expectedPS: psFlagNegative | psFlagOverflow,
	}

	vector.test(t)
}

func TestCPU_BIT_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "BIT Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x2C, // BIT absolute
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x10, // High byte of operand
			0x1000: 0xC0, // Value at absolute address (sets N and V)
		},
		a:          0x40,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagOverflow,
		expectedPS: psFlagNegative | psFlagOverflow,
	}

	vector.test(t)
}
