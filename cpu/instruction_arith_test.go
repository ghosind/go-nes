package cpu

import (
	"testing"
)

func TestCPU_ADC_IMM(t *testing.T) {
	vector := &instructionTestVector{
		name: "ADC Immediate",
		memory: map[uint16]uint8{
			0x8000: 0x69, // ADC Immediate
			0x8001: 0x10, // Value to add
		},
		a:          0x20,
		ps:         psFlagCarry,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x31)),
	}

	vector.test(t)
}

func TestCPU_ADC_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "ADC Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0x65, // ADC Zero Page
			0x8001: 0x10, // Zero Page Address
			0x0010: 0x10, // Value to add
		},
		a:          0x20,
		cycles:     3,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x30)),
	}

	vector.test(t)
}

func TestCPU_ADC_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ADC Zero Page,X",
		memory: map[uint16]uint8{
			0x8000: 0x75, // ADC Zero Page,X
			0x8001: 0x10, // Zero Page Address
			0x0015: 0x10, // Value to add (0x10 + X(5) = 0x15)
		},
		a:          0x20,
		x:          5,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x30)),
	}

	vector.test(t)
}

func TestCPU_ADC_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "ADC Absolute",
		memory: map[uint16]uint8{
			0x8000: 0x6D, // ADC Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address
			0x2000: 0x10, // Value to add
		},
		a:          0x20,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x30)),
	}

	vector.test(t)
}

func TestCPU_ADC_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ADC Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0x7D, // ADC Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address
			0x2005: 0x10, // Value to add (0x2000 + X(5) = 0x2005)
		},
		a:          0x20,
		x:          5,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x30)),
	}

	vector.test(t)
}

func TestCPU_ADC_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "ADC Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0x79, // ADC Absolute, Y
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address
			0x2005: 0x10, // Value to add (0x2000 + Y(5) = 0x2005)
		},
		a:          0x20,
		y:          5,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x30)),
	}

	vector.test(t)
}

func TestCPU_ADC_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "ADC (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0x61, // ADC (Indirect, X)
			0x8001: 0x10, // Zero Page Address
			0x0015: 0x00, // Low byte of effective address (0x0015 = 0x10 + X(5))
			0x0016: 0x20, // High byte of effective address
			0x2000: 0x10, // Value to add
		},
		a:          0x20,
		x:          5,
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x30)),
	}

	vector.test(t)
}

func TestCPU_ADC_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "ADC (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0x71, // ADC (Indirect), Y
			0x8001: 0x10, // Zero Page Address
			0x0010: 0x00, // Low byte of effective address
			0x0011: 0x20, // High byte of effective address
			0x2005: 0x10, // Value to add (0x2000 + Y(5) = 0x2005)
		},
		a:          0x20,
		y:          5,
		cycles:     5,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: 0, // Expect Zero and Negative flags to be cleared
		expectedA:  pointer(uint8(0x30)),
	}

	vector.test(t)
}

func TestCPU_SBC_IMM(t *testing.T) {
	vector := &instructionTestVector{
		name: "SBC Immediate",
		memory: map[uint16]uint8{
			0x8000: 0xE9, // SBC Immediate
			0x8001: 0x10, // Value to subtract
		},
		a:          0x20,
		ps:         psFlagCarry,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
		expectedA:  pointer(uint8(0x10)),
	}

	vector.test(t)
}

func TestCPU_SBC_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "SBC Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0xE5, // SBC Zero Page
			0x8001: 0x10, // Zero Page Address
			0x0010: 0x10, // Value to subtract
		},
		a:          0x20,
		ps:         psFlagCarry,
		cycles:     3,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
		expectedA:  pointer(uint8(0x10)),
	}

	vector.test(t)
}

func TestCPU_SBC_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "SBC Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0xF5, // SBC Zero Page, X
			0x8001: 0x10, // Zero Page Address
			0x0015: 0x10, // Value to subtract (0x10 + X(5) = 0x15)
		},
		a:          0x20,
		x:          5,
		ps:         psFlagCarry,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
		expectedA:  pointer(uint8(0x10)),
	}

	vector.test(t)
}

func TestCPU_SBC_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "SBC Absolute",
		memory: map[uint16]uint8{
			0x8000: 0xED, // SBC Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address
			0x2000: 0x10, // Value to subtract
		},
		a:          0x20,
		ps:         psFlagCarry,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
		expectedA:  pointer(uint8(0x10)),
	}

	vector.test(t)
}

func TestCPU_SBC_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "SBC Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0xFD, // SBC Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address
			0x2005: 0x10, // Value to subtract (0x2000 + X(5) = 0x2005)
		},
		a:          0x20,
		x:          5,
		ps:         psFlagCarry,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
		expectedA:  pointer(uint8(0x10)),
	}

	vector.test(t)
}

func TestCPU_SBC_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "SBC Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0xF9, // SBC Absolute,Y
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address
			0x2005: 0x10, // Value to subtract (0x2000 + Y(5) = 0x2005)
		},
		a:          0x20,
		y:          5,
		ps:         psFlagCarry,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
		expectedA:  pointer(uint8(0x10)),
	}

	vector.test(t)
}

func TestCPU_SBC_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "SBC (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0xE1, // SBC (Indirect, X)
			0x8001: 0x10, // Zero Page Address
			0x0015: 0x00, // Low byte of effective address (0x0015 = 0x10 + X(5))
			0x0016: 0x20, // High byte of effective address
			0x2000: 0x10, // Value to subtract
		},
		a:          0x20,
		x:          5,
		ps:         psFlagCarry,
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
		expectedA:  pointer(uint8(0x10)),
	}

	vector.test(t)
}

func TestCPU_SBC_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "SBC (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0xF1, // SBC (Indirect),Y
			0x8001: 0x10, // Zero Page Address
			0x0010: 0x00, // Low byte of effective address
			0x0011: 0x20, // High byte of effective address
			0x2005: 0x10, // Value to subtract (0x2000 + Y(5) = 0x2005)
		},
		a:          0x20,
		y:          5,
		ps:         psFlagCarry,
		cycles:     5,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry | psFlagOverflow,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
		expectedA:  pointer(uint8(0x10)),
	}

	vector.test(t)
}

func TestCPU_CMP_IMM(t *testing.T) {
	vector := &instructionTestVector{
		name: "CMP Immediate",
		memory: map[uint16]uint8{
			0x8000: 0xC9, // CMP Immediate
			0x8001: 0x10, // Value to compare
		},
		a:          0x20,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CMP_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "CMP Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0xC5, // CMP Zero Page
			0x8001: 0x10, // Zero Page Address
			0x0010: 0x10, // Value to compare
		},
		a:          0x20,
		cycles:     3,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CMP_ZP_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "CMP Zero Page, X",
		memory: map[uint16]uint8{
			0x8000: 0xD5, // CMP Zero Page, X
			0x8001: 0x10, // Zero Page Address
			0x0015: 0x10, // Value to compare (0x10 + X(5) = 0x15)
		},
		a:          0x20,
		x:          5,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CMP_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "CMP Absolute",
		memory: map[uint16]uint8{
			0x8000: 0xCD, // CMP Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address
			0x2000: 0x10, // Value to compare
		},
		a:          0x20,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CMP_ABS_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "CMP Absolute, X",
		memory: map[uint16]uint8{
			0x8000: 0xDD, // CMP Absolute, X
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address
			0x2005: 0x10, // Value to compare (0x2000 + X(5) = 0x2005)
		},
		a:          0x20,
		x:          5,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CMP_ABS_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "CMP Absolute, Y",
		memory: map[uint16]uint8{
			0x8000: 0xD9, // CMP Absolute, Y
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address
			0x2005: 0x10, // Value to compare (0x2000 + Y(5) = 0x2005)
		},
		a:          0x20,
		y:          5,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CMP_IND_X(t *testing.T) {
	vector := &instructionTestVector{
		name: "CMP (Indirect, X)",
		memory: map[uint16]uint8{
			0x8000: 0xC1, // CMP (Indirect, X)
			0x8001: 0x10, // Zero Page Address
			0x0015: 0x00, // Low byte of effective address (0x0015 = 0x10 + X(5))
			0x0016: 0x20, // High byte of effective address
			0x2000: 0x10, // Value to compare
		},
		a:          0x20,
		x:          5,
		cycles:     6,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CMP_IND_Y(t *testing.T) {
	vector := &instructionTestVector{
		name: "CMP (Indirect), Y",
		memory: map[uint16]uint8{
			0x8000: 0xD1, // CMP (Indirect), Y
			0x8001: 0x10, // Zero Page Address
			0x0010: 0x00, // Low byte of effective address
			0x0011: 0x20, // High byte of effective address
			0x2005: 0x10, // Value to compare (0x2000 + Y(5) = 0x2005)
		},
		a:          0x20,
		y:          5,
		cycles:     5,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CPX_IMM(t *testing.T) {
	vector := &instructionTestVector{
		name: "CPX Immediate",
		memory: map[uint16]uint8{
			0x8000: 0xE0, // CPX Immediate
			0x8001: 0x10, // Value to compare
		},
		x:          0x20,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CPX_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "CPX Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0xE4, // CPX Zero Page
			0x8001: 0x10, // Zero Page Address
			0x0010: 0x10, // Value to compare
		},
		x:          0x20,
		cycles:     3,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CPX_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "CPX Absolute",
		memory: map[uint16]uint8{
			0x8000: 0xEC, // CPX Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address
			0x2000: 0x10, // Value to compare
		},
		x:          0x20,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CPY_IMM(t *testing.T) {
	vector := &instructionTestVector{
		name: "CPY Immediate",
		memory: map[uint16]uint8{
			0x8000: 0xC0, // CPY Immediate
			0x8001: 0x10, // Value to compare
		},
		y:          0x20,
		cycles:     2,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CPY_ZP(t *testing.T) {
	vector := &instructionTestVector{
		name: "CPY Zero Page",
		memory: map[uint16]uint8{
			0x8000: 0xC4, // CPY Zero Page
			0x8001: 0x10, // Zero Page Address
			0x0010: 0x10, // Value to compare
		},
		y:          0x20,
		cycles:     3,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}

func TestCPU_CPY_ABS(t *testing.T) {
	vector := &instructionTestVector{
		name: "CPY Absolute",
		memory: map[uint16]uint8{
			0x8000: 0xCC, // CPY Absolute
			0x8001: 0x00, // Low byte of address
			0x8002: 0x20, // High byte of address
			0x2000: 0x10, // Value to compare
		},
		y:          0x20,
		cycles:     4,
		psMask:     psFlagZero | psFlagNegative | psFlagCarry,
		expectedPS: psFlagCarry, // Expect Carry flag to be set
	}

	vector.test(t)
}
