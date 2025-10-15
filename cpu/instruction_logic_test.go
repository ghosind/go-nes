package cpu

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestCPU_AND_IMM(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x29, // AND immediate opcode
			0x8001: 0x0F, // Operand
		},
		cycles:     2,
		a:          0x3C,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_AND_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x25, // AND zero page opcode
			0x8001: 0x10, // Operand
			0x0010: 0x0F, // Value at zero page address
		},
		cycles:     3,
		a:          0x3C,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_AND_ZP_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x35, // AND zero page,X opcode
			0x8001: 0x10, // Operand
			0x0015: 0x0F, // Value at zero page address + X (use addition for clarity)
		},
		cycles:     4,
		a:          0x3C,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_AND_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x2D, // AND absolute opcode
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x20, // High byte of operand
			0x2000: 0x0F, // Value at absolute address
		},
		cycles:     4,
		a:          0x3C,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_AND_ABS_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x3D, // AND absolute,X opcode
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x20, // High byte of operand
			0x2005: 0x0F, // Value at absolute address + X
		},
		cycles:     4,
		a:          0x3C,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_AND_ABS_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x39, // AND absolute,Y opcode
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x20, // High byte of operand
			0x2005: 0x0F, // Value at absolute address + Y
		},
		cycles:     4,
		a:          0x3C,
		y:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_AND_IND_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x21, // AND (indirect,X) opcode
			0x8001: 0x10, // Operand
			0x0015: 0x00, // Low byte of effective address (use addition for clarity)
			0x0016: 0x20, // High byte of effective address
			0x2000: 0x0F, // Value at effective address
		},
		cycles:     6,
		a:          0x3C,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_AND_IND_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x31, // AND (indirect),Y opcode
			0x8001: 0x10, // Operand
			0x0010: 0x00, // Low byte of effective address
			0x0011: 0x20, // High byte of effective address
			0x2005: 0x0F, // Value at effective address + Y
		},
		cycles:     5,
		a:          0x3C,
		y:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x0C)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_EOR_IMM(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x49, // EOR immediate opcode
			0x8001: 0x0F, // Operand
		},
		cycles:     2,
		a:          0x3C,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_EOR_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x45, // EOR zero page opcode
			0x8001: 0x10, // Operand
			0x0010: 0x0F, // Value at zero page address
		},
		cycles:     3,
		a:          0x3C,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_EOR_ZP_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x55, // EOR zero page,X opcode
			0x8001: 0x10, // Operand
			0x0015: 0x0F, // Value at zero page address + X
		},
		cycles:     4,
		a:          0x3C,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_EOR_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x4D, // EOR absolute opcode
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x20, // High byte of operand
			0x2000: 0x0F, // Value at absolute address
		},
		cycles:     4,
		a:          0x3C,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_EOR_ABS_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x5D, // EOR absolute,X opcode
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x20, // High byte of operand
			0x2005: 0x0F, // Value at absolute address + X
		},
		cycles:     4,
		a:          0x3C,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_EOR_ABS_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x59, // EOR absolute,Y opcode
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x20, // High byte of operand
			0x2005: 0x0F, // Value at absolute address + Y
		},
		cycles:     4,
		a:          0x3C,
		y:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_EOR_IND_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x41, // EOR (indirect,X) opcode
			0x8001: 0x10, // Operand
			0x0015: 0x00, // Low byte of effective address
			0x0016: 0x20, // High byte of effective address
			0x2000: 0x0F, // Value at effective address
		},
		cycles:     6,
		a:          0x3C,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_EOR_IND_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x51, // EOR (indirect),Y opcode
			0x8001: 0x10, // Operand
			0x0010: 0x00, // Low byte of effective address
			0x0011: 0x20, // High byte of effective address
			0x2005: 0x0F, // Value at effective address + Y
		},
		cycles:     5,
		a:          0x3C,
		y:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x33)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_ORA_IMM(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x09, // ORA immediate opcode
			0x8001: 0x0F, // Operand
		},
		cycles:     2,
		a:          0x30,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedA:  pointer(uint8(0x3F)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_ORA_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x05, // ORA zero page opcode
			0x8001: 0x10, // Operand
			0x0010: 0x0F, // Value at zero page address
		},
		a:          0x30,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		cycles:     3,
		expectedA:  pointer(uint8(0x3F)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_ORA_ZP_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x15, // ORA zero page,X opcode
			0x8001: 0x10, // Operand
			0x0015: 0x0F, // Value at zero page address + X
		},
		a:          0x30,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		cycles:     4,
		expectedA:  pointer(uint8(0x3F)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_ORA_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x0D, // ORA absolute opcode
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x20, // High byte of operand
			0x2000: 0x0F, // Value at absolute address
		},
		a:          0x30,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		cycles:     4,
		expectedA:  pointer(uint8(0x3F)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_ORA_ABS_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x1D, // ORA absolute,X opcode
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x20, // High byte of operand
			0x2005: 0x0F, // Value at absolute address + X
		},
		a:          0x30,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		cycles:     4,
		expectedA:  pointer(uint8(0x3F)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_ORA_ABS_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x19, // ORA absolute,Y opcode
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x20, // High byte of operand
			0x2005: 0x0F, // Value at absolute address + Y
		},
		a:          0x30,
		y:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		cycles:     4,
		expectedA:  pointer(uint8(0x3F)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_ORA_IND_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x01, // ORA (indirect,X) opcode
			0x8001: 0x10, // Operand
			0x0015: 0x00, // Low byte of effective address
			0x0016: 0x20, // High byte of effective address
			0x2000: 0x0F, // Value at effective address
		},
		a:          0x30,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		cycles:     6,
		expectedA:  pointer(uint8(0x3F)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_ORA_IND_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x11, // ORA (indirect),Y opcode
			0x8001: 0x10, // Operand
			0x0010: 0x00, // Low byte of effective address
			0x0011: 0x20, // High byte of effective address
			0x2005: 0x0F, // Value at effective address + Y
		},
		a:          0x30,
		y:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		cycles:     5,
		expectedA:  pointer(uint8(0x3F)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_BIT_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x24, // BIT zero page opcode
			0x8001: 0x10, // Operand
			0x0010: 0xC0, // Value at zero page address (sets N and V)
		},
		cycles:     3,
		a:          0x40,
		psMask:     psFlagZero | psFlagNegative | psFlagOverflow,
		expectedPS: psFlagNegative | psFlagOverflow,
	}

	testCPUInstruction(a, vector)
}

func TestCPU_BIT_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x2C, // BIT absolute opcode
			0x8001: 0x00, // Low byte of operand
			0x8002: 0x20, // High byte of operand
			0x2000: 0xC0, // Value at absolute address (sets N and V)
		},
		cycles:     4,
		a:          0x40,
		psMask:     psFlagZero | psFlagNegative | psFlagOverflow,
		expectedPS: psFlagNegative | psFlagOverflow,
	}

	testCPUInstruction(a, vector)
}
