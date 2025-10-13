package cpu

type cpuInstruction struct {
	execute    func(*CPU, ...uint8)
	addressing AddressingMode
	cycles     int
}

var (
	instructionMap map[uint8]cpuInstruction = map[uint8]cpuInstruction{
		// Load/Store Operations
		0xA9: {execute: (*CPU).lda_imm, addressing: addressingModeImmediate, cycles: 2},         // LDA Immediate
		0xA5: {execute: (*CPU).lda_zp, addressing: addressingModeZeroPage, cycles: 3},           // LDA Zero Page
		0xB5: {execute: (*CPU).lda_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // LDA Zero Page, X
		0xAD: {execute: (*CPU).lda_abs, addressing: addressingModeAbsolute, cycles: 4},          // LDA Absolute
		0xBD: {execute: (*CPU).lda_abs_x, addressing: addressingModeAbsoluteX, cycles: 4},       // LDA Absolute, X
		0xB9: {execute: (*CPU).lda_abs_y, addressing: addressingModeAbsoluteY, cycles: 4},       // LDA Absolute, Y
		0xA1: {execute: (*CPU).lda_ind_x, addressing: addressingModeIndexedIndirect, cycles: 6}, // LDA (Indirect, X)
		0xB1: {execute: (*CPU).lda_ind_y, addressing: addressingModeIndirectIndexed, cycles: 5}, // LDA (Indirect), Y
		0xA2: {execute: (*CPU).ldx_imm, addressing: addressingModeImmediate, cycles: 2},         // LDX Immediate
		0xA6: {execute: (*CPU).ldx_zp, addressing: addressingModeZeroPage, cycles: 3},           // LDX Zero Page
		0xB6: {execute: (*CPU).ldx_zp_y, addressing: addressingModeZeroPageY, cycles: 4},        // LDX Zero Page, Y
		0xAE: {execute: (*CPU).ldx_abs, addressing: addressingModeAbsolute, cycles: 4},          // LDX Absolute
		0xBE: {execute: (*CPU).ldx_abs_y, addressing: addressingModeAbsoluteY, cycles: 4},       // LDX Absolute, Y
		0xA0: {execute: (*CPU).ldy_imm, addressing: addressingModeImmediate, cycles: 2},         // LDY Immediate
		0xA4: {execute: (*CPU).ldy_zp, addressing: addressingModeZeroPage, cycles: 3},           // LDY Zero Page
		0xB4: {execute: (*CPU).ldy_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // LDY Zero Page, X
		0xAC: {execute: (*CPU).ldy_abs, addressing: addressingModeAbsolute, cycles: 4},          // LDY Absolute
		0xBC: {execute: (*CPU).ldy_abs_x, addressing: addressingModeAbsoluteX, cycles: 4},       // LDY Absolute, X
		0x85: {execute: (*CPU).sta_zp, addressing: addressingModeZeroPage, cycles: 3},           // STA Zero Page
		0x95: {execute: (*CPU).sta_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // STA Zero Page, X
		0x8D: {execute: (*CPU).sta_abs, addressing: addressingModeAbsolute, cycles: 4},          // STA Absolute
		0x9D: {execute: (*CPU).sta_abs_x, addressing: addressingModeAbsoluteX, cycles: 5},       // STA Absolute, X
		0x99: {execute: (*CPU).sta_abs_y, addressing: addressingModeAbsoluteY, cycles: 5},       // STA Absolute, Y
		0x81: {execute: (*CPU).sta_ind_x, addressing: addressingModeIndexedIndirect, cycles: 6}, // STA (Indirect, X)
		0x91: {execute: (*CPU).sta_ind_y, addressing: addressingModeIndirectIndexed, cycles: 6}, // STA (Indirect), Y
		0x86: {execute: (*CPU).stx_zp, addressing: addressingModeZeroPage, cycles: 3},           // STX Zero Page
		0x96: {execute: (*CPU).stx_zp_y, addressing: addressingModeZeroPageY, cycles: 4},        // STX Zero Page, Y
		0x8E: {execute: (*CPU).stx_abs, addressing: addressingModeAbsolute, cycles: 4},          // STX Absolute
		0x84: {execute: (*CPU).sty_zp, addressing: addressingModeZeroPage, cycles: 3},           // STY Zero Page
		0x94: {execute: (*CPU).sty_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // STY Zero Page, X
		0x8C: {execute: (*CPU).sty_abs, addressing: addressingModeAbsolute, cycles: 4},          // STY Absolute

		// Register Transfers
		0xAA: {}, // TAX - Transfer Accumulator to X
		0xA8: {}, // TAY - Transfer Accumulator to Y
		0x8A: {}, // TXA - Transfer X to Accumulator
		0x98: {}, // TYA - Transfer Y to Accumulator

		// Stack Operations
		0xBA: {}, // TSX - Transfer Stack Pointer to X
		0x9A: {}, // TXS - Transfer X to Stack Pointer
		0x48: {}, // PHA - Push Accumulator onto Stack
		0x68: {}, // PLA - Pull Accumulator from Stack
		0x08: {}, // PHP - Push Processor Status onto Stack
		0x28: {}, // PLP - Pull Processor Status from Stack

		// Logical Operations
		0x29: {}, // AND Immediate
		0x25: {}, // AND Zero Page
		0x35: {}, // AND Zero Page, X
		0x2D: {}, // AND Absolute
		0x3D: {}, // AND Absolute, X
		0x39: {}, // AND Absolute, Y
		0x21: {}, // AND (Indirect, X)
		0x31: {}, // AND (Indirect), Y
		0x49: {}, // EOR Immediate
		0x45: {}, // EOR Zero Page
		0x55: {}, // EOR Zero Page, X
		0x4D: {}, // EOR Absolute
		0x5D: {}, // EOR Absolute, X
		0x59: {}, // EOR Absolute, Y
		0x41: {}, // EOR (Indirect, X)
		0x51: {}, // EOR (Indirect), Y
		0x09: {}, // ORA Immediate
		0x05: {}, // ORA Zero Page
		0x15: {}, // ORA Zero Page, X
		0x0D: {}, // ORA Absolute
		0x1D: {}, // ORA Absolute, X
		0x19: {}, // ORA Absolute, Y
		0x01: {}, // ORA (Indirect, X)
		0x11: {}, // ORA (Indirect), Y
		0x24: {}, // BIT Zero Page
		0x2C: {}, // BIT Absolute

		// Arithmetic
		0x69: {}, // ADC Immediate
		0x65: {}, // ADC Zero Page
		0x75: {}, // ADC Zero Page, X
		0x6D: {}, // ADC Absolute
		0x7D: {}, // ADC Absolute, X
		0x79: {}, // ADC Absolute, Y
		0x61: {}, // ADC (Indirect, X)
		0x71: {}, // ADC (Indirect), Y
		0xE9: {}, // SBC Immediate
		0xE5: {}, // SBC Zero Page
		0xF5: {}, // SBC Zero Page, X
		0xED: {}, // SBC Absolute
		0xFD: {}, // SBC Absolute, X
		0xF9: {}, // SBC Absolute, Y
		0xE1: {}, // SBC (Indirect, X)
		0xF1: {}, // SBC (Indirect), Y
		0xC9: {}, // CMP Immediate
		0xC5: {}, // CMP Zero Page
		0xD5: {}, // CMP Zero Page, X
		0xCD: {}, // CMP Absolute
		0xDD: {}, // CMP Absolute, X
		0xD9: {}, // CMP Absolute, Y
		0xC1: {}, // CMP (Indirect, X)
		0xD1: {}, // CMP (Indirect), Y
		0xE0: {}, // CPX Immediate
		0xE4: {}, // CPX Zero Page
		0xEC: {}, // CPX Absolute
		0xC0: {}, // CPY Immediate
		0xC4: {}, // CPY Zero Page
		0xCC: {}, // CPY Absolute

		// Increments & Decrements
		0xE6: {}, // INC Zero Page
		0xF6: {}, // INC Zero Page, X
		0xEE: {}, // INC Absolute
		0xFE: {}, // INC Absolute, X
		0xE8: {}, // INX - Increment X
		0xC8: {}, // INY - Increment Y
		0xC6: {}, // DEC Zero Page
		0xD6: {}, // DEC Zero Page, X
		0xCE: {}, // DEC Absolute
		0xDE: {}, // DEC Absolute, X
		0xCA: {}, // DEX - Decrement X
		0x88: {}, // DEY - Decrement Y

		// Shifts
		0x0A: {}, // ASL Accumulator
		0x06: {}, // ASL Zero Page
		0x16: {}, // ASL Zero Page, X
		0x0E: {}, // ASL Absolute
		0x1E: {}, // ASL Absolute, X
		0x4A: {}, // LSR Accumulator
		0x46: {}, // LSR Zero Page
		0x56: {}, // LSR Zero Page, X
		0x4E: {}, // LSR Absolute
		0x5E: {}, // LSR Absolute, X
		0x2A: {}, // ROL Accumulator
		0x26: {}, // ROL Zero Page
		0x36: {}, // ROL Zero Page, X
		0x2E: {}, // ROL Absolute
		0x3E: {}, // ROL Absolute, X
		0x6A: {}, // ROR Accumulator
		0x66: {}, // ROR Zero Page
		0x76: {}, // ROR Zero Page, X
		0x6E: {}, // ROR Absolute
		0x7E: {}, // ROR Absolute, X

		// Jumps & Calls
		0x4C: {}, // JMP Absolute
		0x6C: {}, // JMP Indirect
		0x20: {}, // JSR - Jump to Subroutine
		0x60: {}, // RTS - Return from Subroutine

		// Branches
		0x90: {}, // BCC - Branch if Carry Clear
		0xB0: {}, // BCS - Branch if Carry Set
		0xF0: {}, // BEQ - Branch if Equal
		0x30: {}, // BMI - Branch if Minus
		0xD0: {}, // BNE - Branch if Not Equal
		0x10: {}, // BPL - Branch if Positive
		0x50: {}, // BVC - Branch if Overflow Clear
		0x70: {}, // BVS - Branch if Overflow Set

		// Status Flag
		0x18: {execute: (*CPU).clc, addressing: addressingModeImplied, cycles: 2}, // CLC - Clear Carry Flag
		0xD8: {execute: (*CPU).cld, addressing: addressingModeImplied, cycles: 2}, // CLD - Clear Decimal Mode
		0x58: {execute: (*CPU).cli, addressing: addressingModeImplied, cycles: 2}, // CLI - Clear Interrupt Disable
		0xB8: {execute: (*CPU).clv, addressing: addressingModeImplied, cycles: 2}, // CLV - Clear Overflow Flag
		0x38: {execute: (*CPU).sec, addressing: addressingModeImplied, cycles: 2}, // SEC - Set Carry Flag
		0xF8: {execute: (*CPU).sed, addressing: addressingModeImplied, cycles: 2}, // SED - Set Decimal Mode
		0x78: {execute: (*CPU).sei, addressing: addressingModeImplied, cycles: 2}, // SEI - Set Interrupt Disable

		// System Functions
		0x00: {},                                                                  // BRK - Force Interrupt
		0xEA: {execute: (*CPU).nop, addressing: addressingModeImplied, cycles: 2}, // NOP - No Operation
		0x40: {},                                                                  // RTI - Return from Interrupt
	}
)
