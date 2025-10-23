package cpu

import (
	"testing"
)

func TestCPU_INC_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "INC Zero Page",
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

	vector.test(t)
}

func TestCPU_INC_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "INC Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0xF6, // INC Zero Page, X
			0x8001: 0x10, // address $10
			0x0015: 0x05, // value $05 (address = $10 + X($05) = $15)
		},
		x:          0x05,
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0015: 0x06,
		},
	}

	vector.test(t)
}

func TestCPU_INC_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "INC Absolute",
		memory: map[uint16]uint8{
			0x8000: 0xEE, // INC Absolute
			0x8001: 0x00, // address $1000
			0x8002: 0x10, //
			0x1000: 0x05, // value $05
		},
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x06,
		},
	}

	vector.test(t)
}

func TestCPU_INC_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "INC Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0xFE, // INC Absolute, X
			0x8001: 0x00, // address $1000
			0x8002: 0x10, //
			0x1005: 0x05, // value $05 (address = $1000 + X($05) = $1005)
		},
		x:          0x05,
		cycles:     7,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x06,
		},
	}

	vector.test(t)
}

func TestCPU_INX(t *testing.T) {
	vector := &instructionTestVector{
		name: "INX",
		memory: map[uint16]uint8{
			0x8000: 0xE8, // INX
		},
		x:          0x05,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedX:  pointer(uint8(0x06)),
	}

	vector.test(t)
}

func TestCPU_INY(t *testing.T) {
	vector := &instructionTestVector{
		name: "INY",
		memory: map[uint16]uint8{
			0x8000: 0xC8, // INY
		},
		y:          0x05,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedY:  pointer(uint8(0x06)),
	}

	vector.test(t)
}

func TestCPU_DEC_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "DEC Zero Page",
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

	vector.test(t)
}

func TestCPU_DEC_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "DEC Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0xD6, // DEC Zero Page, X
			0x8001: 0x10, // address $10
			0x0015: 0x05, // value $05 (address = $10 + X($05) = $15)
		},
		x:          0x05,
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x0015: 0x04,
		},
	}

	vector.test(t)
}

func TestCPU_DEC_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "DEC Absolute",
		memory: map[uint16]uint8{
			0x8000: 0xCE, // DEC Absolute
			0x8001: 0x00, // address $2000
			0x8002: 0x10, //
			0x1000: 0x05, // value $05
		},
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1000: 0x04,
		},
	}

	vector.test(t)
}

func TestCPU_DEC_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "DEC Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0xDE, // DEC Absolute, X
			0x8001: 0x00, // address $1000
			0x8002: 0x10, //
			0x1005: 0x05, // value $05 (address = $1000 + X($05) = $1005)
		},
		x:          0x05,
		cycles:     7,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedMem: map[uint16]uint8{
			0x1005: 0x04,
		},
	}

	vector.test(t)
}

func TestCPU_DEX(t *testing.T) {
	vector := &instructionTestVector{
		name: "DEX",
		memory: map[uint16]uint8{
			0x8000: 0xCA, // DEX
		},
		x:          0x05,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedX:  pointer(uint8(0x04)),
	}

	vector.test(t)
}

func TestCPU_DEY(t *testing.T) {
	vector := &instructionTestVector{
		name: "DEY",
		memory: map[uint16]uint8{
			0x8000: 0x88, // DEY
		},
		y:          0x05,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative,
		expectedPS: 0,
		expectedY:  pointer(uint8(0x04)),
	}

	vector.test(t)
}
