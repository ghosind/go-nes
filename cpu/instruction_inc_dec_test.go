package cpu

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestCPU_INC_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xE6, // INC Zero Page
			0x8001: 0x10, // address $10
			0x0010: 0x05, // value $05
		},
		cycles:     5,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0010: 0x06,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_INC_ZP_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xF6, // INC Zero Page,X
			0x8001: 0x10, // address $10
			0x0015: 0x05, // value $05 (address = $10 + X($05) = $15)
		},
		cycles:     6,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0015: 0x06,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_INC_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xEE, // INC Absolute
			0x8001: 0x00, // address $2000
			0x8002: 0x20, //
			0x2000: 0x05, // value $05
		},
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x2000: 0x06,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_INC_ABS_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xFE, // INC Absolute,X
			0x8001: 0x00, // address $2000
			0x8002: 0x20, //
			0x2005: 0x05, // value $05 (address = $2000 + X($05) = $2005)
		},
		cycles:     7,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x2005: 0x06,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_INX(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xE8, // INX
		},
		cycles:     2,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedX:  pointer(uint8(0x06)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_INY(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xC8, // INY
		},
		cycles:     2,
		y:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedY:  pointer(uint8(0x06)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_DEC_ZP(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xC6, // DEC Zero Page
			0x8001: 0x10, // address $10
			0x0010: 0x05, // value $05
		},
		cycles:     5,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0010: 0x04,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_DEC_ZP_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xD6, // DEC Zero Page,X
			0x8001: 0x10, // address $10
			0x0015: 0x05, // value $05 (address = $10 + X($05) = $15)
		},
		cycles:     6,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0015: 0x04,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_DEC_ABS(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xCE, // DEC Absolute
			0x8001: 0x00, // address $2000
			0x8002: 0x20, //
			0x2000: 0x05, // value $05
		},
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x2000: 0x04,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_DEC_ABS_X(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xDE, // DEC Absolute,X
			0x8001: 0x00, // address $2000
			0x8002: 0x20, //
			0x2005: 0x05, // value $05 (address = $2000 + X($05) = $2005)
		},
		cycles:     7,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x2005: 0x04,
		},
	}

	testCPUInstruction(a, vector)
}

func TestCPU_DEX(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0xCA, // DEX
		},
		cycles:     2,
		x:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedX:  pointer(uint8(0x04)),
	}

	testCPUInstruction(a, vector)
}

func TestCPU_DEY(t *testing.T) {
	a := assert.New(t)
	vector := instructionTestVector{
		memory: map[uint16]uint8{
			0x8000: 0x88, // DEY
		},
		cycles:     2,
		y:          0x05,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedY:  pointer(uint8(0x04)),
	}

	testCPUInstruction(a, vector)
}
