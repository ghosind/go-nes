package cpu

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestCPU_LDA_IMM(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xA9, // LDA Immediate opcode
			0x8001: 0x42, // Operand: Load the value 0x42 into A
		},
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     2,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.a, uint8(0x42))
}

func TestCPU_LDA_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xA5, // LDA Zero Page opcode
			0x8001: 0x10, // Operand: Zero page address 0x10
			0x0010: 0x37, // Value at Zero Page address 0x10
		},
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     3,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.a, uint8(0x37))
}

func TestCPU_LDA_ZP_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xB5, // LDA Zero Page, X opcode
			0x8001: 0x10, // Operand: Zero Page address 0x10
			0x0015: 0x58, // Value at Zero Page address 0x10 + X (0x15)
		},
		x:          0x05, // Set X register to 5
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.a, uint8(0x58))
}

func TestCPU_LDA_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xAD, // LDA Absolute opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address (0x2000)
			0x2000: 0x7A, // Value at address 0x2000
		},
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.a, uint8(0x7A))
}

func TestCPU_LDA_ABS_X(t *testing.T) {
	a := assert.New(t)

	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xBD, // LDA Absolute, X opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address (0x2000)
			0x2005: 0x9C, // Value at address 0x2000 + X (0x2005)
		},
		x:          0x05, // Set X register to 5
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.a, uint8(0x9C))
}

func TestCPU_LDA_ABS_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xB9, // LDA Absolute, Y opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address (0x2000)
			0x2003: 0xFF, // Value at address 0x2000 + Y (0x2003)
		},
		y:          0x03, // Set Y register to 3
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: psFlagNegative, // Expect Negative flag to be set
		cycles:     4,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.a, uint8(0xFF))
}

func TestCPU_LDA_IND_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xA1, // LDA (Indirect, X) opcode
			0x8001: 0x10, // Operand: Zero Page address 0x10
			0x0015: 0x00, // Low byte of effective address (0x3000)
			0x0016: 0x30, // High byte of effective address
			0x3000: 0x66, // Value at effective address 0x3000
		},
		x:          0x05, // Set X register to 5
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     6,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.a, uint8(0x66))
}

func TestCPU_LDA_IND_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xB1, // LDA (Indirect), Y opcode
			0x8001: 0x10, // Operand: Zero Page address 0x10
			0x0010: 0x00, // Low byte of base address (0x3000)
			0x0011: 0x30, // High byte of base address
			0x3002: 0xAB, // Value at effective address 0x3000 + Y (0x3002)
		},
		y:          0x02, // Set Y register to 2
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: psFlagNegative, // Expect Negative flag to be set
		cycles:     5,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.a, uint8(0xAB))
}

func TestCPU_LDX_IMM(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xA2, // LDX Immediate opcode
			0x8001: 0x55, // Operand: Load the value 0x55 into X
		},
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     2,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.x, uint8(0x55))
}

func TestCPU_LDX_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xA6, // LDX Zero Page opcode
			0x8001: 0x20, // Operand: Zero Page address 0x20
			0x0020: 0x33, // Value at Zero Page address 0x20
		},
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     3,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.x, uint8(0x33))
}

func TestCPU_LDX_ZP_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xB6, // LDX Zero Page, Y opcode
			0x8001: 0x20, // Operand: Zero Page address 0x20
			0x0025: 0x77, // Value at Zero Page address 0x20 + Y (0x25)
		},
		y:          0x05, // Set Y register to 5
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.x, uint8(0x77))
}

func TestCPU_LDX_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xAE, // LDX Absolute opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x30, // High byte of address (0x3000)
			0x3000: 0x88, // Value at address 0x3000
		},
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: psFlagNegative, // Expect Negative flag to be set
		cycles:     4,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.x, uint8(0x88))
}

func TestCPU_LDX_ABS_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xBE, // LDX Absolute, Y opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x30, // High byte of address (0x3000)
			0x3004: 0x22, // Value at address 0x3000 + Y (0x3004)
		},
		y:          0x04, // Set Y register to 4
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.x, uint8(0x22))
}

func TestCPU_LDY_IMM(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xA0, // LDY Immediate opcode
			0x8001: 0x99, // Operand: Load the value 0x99 into Y
		},
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: psFlagNegative, // Expect Negative flag to be set
		cycles:     2,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.y, uint8(0x99))
}

func TestCPU_LDY_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xA4, // LDY Zero Page opcode
			0x8001: 0x30, // Operand: Zero Page address 0x30
			0x0030: 0x44, // Value at Zero Page address 0x30
		},
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     3,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.y, uint8(0x44))
}

func TestCPU_LDY_ZP_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xB4, // LDY Zero Page, X opcode
			0x8001: 0x30, // Operand: Zero Page address 0x30
			0x0035: 0x11, // Value at Zero Page address 0x30 + X (0x35)
		},
		x:          0x05, // Set X register to 5
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.y, uint8(0x11))
}

func TestCPU_LDY_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xAC, // LDY Absolute opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x40, // High byte of address (0x4000)
			0x4000: 0xFE, // Value at address 0x4000
		},
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: psFlagNegative, // Expect Negative flag to be set
		cycles:     4,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.y, uint8(0xFE))
}

func TestCPU_LDY_ABS_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xBC, // LDY Absolute, X opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x40, // High byte of address (0x4000)
			0x4003: 0x66, // Value at address 0x4000 + X (0x4003)
		},
		x:          0x03, // Set X register to 3
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
	}

	cpu := testCPUInstruction(a, vector)
	a.EqualNow(cpu.y, uint8(0x66))
}

func TestCPU_STA_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x85, // STA Zero Page opcode
			0x8001: 0x10, // Operand: Zero Page address 0x10
		},
		a:          0x42, // Pre-load A with 0x42
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     3,
		expectedMem: map[uint16]uint8{
			0x0010: 0x42,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STA_ZP_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x95, // STA Zero Page, X opcode
			0x8001: 0x10, // Operand: Zero Page address 0x10
		},
		a:          0x37, // Pre-load A with 0x37
		x:          0x05, // Set X register to 5
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
		expectedMem: map[uint16]uint8{
			0x0015: 0x37,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STA_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x8D, // STA Absolute opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address (0x2000)
		},
		a:          0x99, // Pre-load A with 0x99
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: psFlagNegative, // Expect Negative flag to be set
		cycles:     4,
		expectedMem: map[uint16]uint8{
			0x2000: 0x99,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STA_ABS_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x9D, // STA Absolute, X opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address (0x2000)
		},
		a:          0x55, // Pre-load A with 0x55
		x:          0x05, // Set X register to 5
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     5,
		expectedMem: map[uint16]uint8{
			0x2005: 0x55,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STA_ABS_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x99, // STA Absolute, Y opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address (0x2000)
		},
		a:          0x77, // Pre-load A with 0x77
		y:          0x03, // Set Y register to 3
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     5,
		expectedMem: map[uint16]uint8{
			0x2003: 0x77,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STA_IND_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x81, // STA (Indirect, X) opcode
			0x8001: 0x10, // Operand: Zero Page address 0x10
			0x0015: 0x00, // Low byte of effective address (0x3000)
			0x0016: 0x30, // High byte of effective address
		},
		a:          0x42, // Pre-load A with 0x42
		x:          0x05, // Set X register to 5
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     6,
		expectedMem: map[uint16]uint8{
			0x3000: 0x42,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STA_IND_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x91, // STA (Indirect), Y opcode
			0x8001: 0x10, // Operand: Zero Page address 0x10
			0x0010: 0x00, // Low byte of base address (0x3000)
			0x0011: 0x30, // High byte of base address
		},
		a:          0x37, // Pre-load A with 0x37
		y:          0x02, // Set Y register to 2
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     6,
		expectedMem: map[uint16]uint8{
			0x3002: 0x37,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STX_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x86, // STX Zero Page opcode
			0x8001: 0x10, // Operand: Zero Page address 0x10
		},
		x:          0x55, // Pre-load X with 0x55
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     3,
		expectedMem: map[uint16]uint8{
			0x0010: 0x55,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STX_ZP_Y(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x96, // STX Zero Page, Y opcode
			0x8001: 0x10, // Operand: Zero Page address 0x10
		},
		x:          0x33, // Pre-load X with 0x33
		y:          0x05, // Set Y register to 5
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
		expectedMem: map[uint16]uint8{
			0x0015: 0x33,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STX_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x8E, // STX Absolute opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address (0x2000)
		},
		x:          0x77, // Pre-load X with 0x77
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
		expectedMem: map[uint16]uint8{
			0x2000: 0x77,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STY_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x84, // STY Zero Page opcode
			0x8001: 0x10, // Operand: Zero Page address 0x10
		},
		y:          0x99, // Pre-load Y with 0x99
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: psFlagNegative, // Expect Negative flag to be set
		cycles:     3,
		expectedMem: map[uint16]uint8{
			0x0010: 0x99,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STY_ZP_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x94, // STY Zero Page, X opcode
			0x8001: 0x10, // Operand: Zero Page address 0x10
		},
		y:          0x44, // Pre-load Y with 0x44
		x:          0x05, // Set X register to 5
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
		expectedMem: map[uint16]uint8{
			0x0015: 0x44,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_STY_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x8C, // STY Absolute opcode
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address (0x2000)
		},
		y:          0x22, // Pre-load Y with 0x22
		psMask:     psFlagZero & psFlagNegative,
		expectedPS: 0x00, // Expect neither Zero nor Negative flags to be set
		cycles:     4,
		expectedMem: map[uint16]uint8{
			0x2000: 0x22,
		},
	}

	testCPUInstruction(a, vector)
}
