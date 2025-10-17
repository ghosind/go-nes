package cpu

type cpuInstruction struct {
	opcode     string
	execute    func(*CPU, ...uint8)
	addressing AddressingMode
	cycles     uint64
}

var (
	instructionMap map[uint8]cpuInstruction = map[uint8]cpuInstruction{
		// Load/Store Operations
		0xA9: {opcode: "LDA", execute: (*CPU).lda_imm, addressing: addressingModeImmediate, cycles: 2},         // LDA Immediate
		0xA5: {opcode: "LDA", execute: (*CPU).lda_zp, addressing: addressingModeZeroPage, cycles: 3},           // LDA Zero Page
		0xB5: {opcode: "LDA", execute: (*CPU).lda_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // LDA Zero Page, X
		0xAD: {opcode: "LDA", execute: (*CPU).lda_abs, addressing: addressingModeAbsolute, cycles: 4},          // LDA Absolute
		0xBD: {opcode: "LDA", execute: (*CPU).lda_abs_x, addressing: addressingModeAbsoluteX, cycles: 4},       // LDA Absolute, X
		0xB9: {opcode: "LDA", execute: (*CPU).lda_abs_y, addressing: addressingModeAbsoluteY, cycles: 4},       // LDA Absolute, Y
		0xA1: {opcode: "LDA", execute: (*CPU).lda_ind_x, addressing: addressingModeIndexedIndirect, cycles: 6}, // LDA (Indirect, X)
		0xB1: {opcode: "LDA", execute: (*CPU).lda_ind_y, addressing: addressingModeIndirectIndexed, cycles: 5}, // LDA (Indirect), Y
		0xA2: {opcode: "LDX", execute: (*CPU).ldx_imm, addressing: addressingModeImmediate, cycles: 2},         // LDX Immediate
		0xA6: {opcode: "LDX", execute: (*CPU).ldx_zp, addressing: addressingModeZeroPage, cycles: 3},           // LDX Zero Page
		0xB6: {opcode: "LDX", execute: (*CPU).ldx_zp_y, addressing: addressingModeZeroPageY, cycles: 4},        // LDX Zero Page, Y
		0xAE: {opcode: "LDX", execute: (*CPU).ldx_abs, addressing: addressingModeAbsolute, cycles: 4},          // LDX Absolute
		0xBE: {opcode: "LDX", execute: (*CPU).ldx_abs_y, addressing: addressingModeAbsoluteY, cycles: 4},       // LDX Absolute, Y
		0xA0: {opcode: "LDY", execute: (*CPU).ldy_imm, addressing: addressingModeImmediate, cycles: 2},         // LDY Immediate
		0xA4: {opcode: "LDY", execute: (*CPU).ldy_zp, addressing: addressingModeZeroPage, cycles: 3},           // LDY Zero Page
		0xB4: {opcode: "LDY", execute: (*CPU).ldy_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // LDY Zero Page, X
		0xAC: {opcode: "LDY", execute: (*CPU).ldy_abs, addressing: addressingModeAbsolute, cycles: 4},          // LDY Absolute
		0xBC: {opcode: "LDY", execute: (*CPU).ldy_abs_x, addressing: addressingModeAbsoluteX, cycles: 4},       // LDY Absolute, X
		0x85: {opcode: "STA", execute: (*CPU).sta_zp, addressing: addressingModeZeroPage, cycles: 3},           // STA Zero Page
		0x95: {opcode: "STA", execute: (*CPU).sta_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // STA Zero Page, X
		0x8D: {opcode: "STA", execute: (*CPU).sta_abs, addressing: addressingModeAbsolute, cycles: 4},          // STA Absolute
		0x9D: {opcode: "STA", execute: (*CPU).sta_abs_x, addressing: addressingModeAbsoluteX, cycles: 5},       // STA Absolute, X
		0x99: {opcode: "STA", execute: (*CPU).sta_abs_y, addressing: addressingModeAbsoluteY, cycles: 5},       // STA Absolute, Y
		0x81: {opcode: "STA", execute: (*CPU).sta_ind_x, addressing: addressingModeIndexedIndirect, cycles: 6}, // STA (Indirect, X)
		0x91: {opcode: "STA", execute: (*CPU).sta_ind_y, addressing: addressingModeIndirectIndexed, cycles: 6}, // STA (Indirect), Y
		0x86: {opcode: "STX", execute: (*CPU).stx_zp, addressing: addressingModeZeroPage, cycles: 3},           // STX Zero Page
		0x96: {opcode: "STX", execute: (*CPU).stx_zp_y, addressing: addressingModeZeroPageY, cycles: 4},        // STX Zero Page, Y
		0x8E: {opcode: "STX", execute: (*CPU).stx_abs, addressing: addressingModeAbsolute, cycles: 4},          // STX Absolute
		0x84: {opcode: "STY", execute: (*CPU).sty_zp, addressing: addressingModeZeroPage, cycles: 3},           // STY Zero Page
		0x94: {opcode: "STY", execute: (*CPU).sty_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // STY Zero Page, X
		0x8C: {opcode: "STY", execute: (*CPU).sty_abs, addressing: addressingModeAbsolute, cycles: 4},          // STY Absolute

		// Register Transfers
		0xAA: {opcode: "TAX", execute: (*CPU).tax, addressing: addressingModeImplied, cycles: 2}, // TAX - Transfer Accumulator to X
		0xA8: {opcode: "TAY", execute: (*CPU).tay, addressing: addressingModeImplied, cycles: 2}, // TAY - Transfer Accumulator to Y
		0x8A: {opcode: "TXA", execute: (*CPU).txa, addressing: addressingModeImplied, cycles: 2}, // TXA - Transfer X to Accumulator
		0x98: {opcode: "TYA", execute: (*CPU).tya, addressing: addressingModeImplied, cycles: 2}, // TYA - Transfer Y to Accumulator

		// Stack Operations
		0xBA: {opcode: "TSX", execute: (*CPU).tsx, addressing: addressingModeImplied, cycles: 2}, // TSX - Transfer Stack Pointer to X
		0x9A: {opcode: "TXS", execute: (*CPU).txs, addressing: addressingModeImplied, cycles: 2}, // TXS - Transfer X to Stack Pointer
		0x48: {opcode: "PHA", execute: (*CPU).pha, addressing: addressingModeImplied, cycles: 3}, // PHA - Push Accumulator onto Stack
		0x68: {opcode: "PLA", execute: (*CPU).pla, addressing: addressingModeImplied, cycles: 4}, // PLA - Pull Accumulator from Stack
		0x08: {opcode: "PHP", execute: (*CPU).php, addressing: addressingModeImplied, cycles: 3}, // PHP - Push Processor Status onto Stack
		0x28: {opcode: "PLP", execute: (*CPU).plp, addressing: addressingModeImplied, cycles: 4}, // PLP - Pull Processor Status from Stack

		// Logical Operations
		0x29: {opcode: "AND", execute: (*CPU).and_imm, addressing: addressingModeImmediate, cycles: 2},         // AND Immediate
		0x25: {opcode: "AND", execute: (*CPU).and_zp, addressing: addressingModeZeroPage, cycles: 3},           // AND Zero Page
		0x35: {opcode: "AND", execute: (*CPU).and_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // AND Zero Page, X
		0x2D: {opcode: "AND", execute: (*CPU).and_abs, addressing: addressingModeAbsolute, cycles: 4},          // AND Absolute
		0x3D: {opcode: "AND", execute: (*CPU).and_abs_x, addressing: addressingModeAbsoluteX, cycles: 4},       // AND Absolute, X
		0x39: {opcode: "AND", execute: (*CPU).and_abs_y, addressing: addressingModeAbsoluteY, cycles: 4},       // AND Absolute, Y
		0x21: {opcode: "AND", execute: (*CPU).and_ind_x, addressing: addressingModeIndexedIndirect, cycles: 6}, // AND (Indirect, X)
		0x31: {opcode: "AND", execute: (*CPU).and_ind_y, addressing: addressingModeIndirectIndexed, cycles: 5}, // AND (Indirect), Y
		0x49: {opcode: "EOR", execute: (*CPU).eor_imm, addressing: addressingModeImmediate, cycles: 2},         // EOR Immediate
		0x45: {opcode: "EOR", execute: (*CPU).eor_zp, addressing: addressingModeZeroPage, cycles: 3},           // EOR Zero Page
		0x55: {opcode: "EOR", execute: (*CPU).eor_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // EOR Zero Page, X
		0x4D: {opcode: "EOR", execute: (*CPU).eor_abs, addressing: addressingModeAbsolute, cycles: 4},          // EOR Absolute
		0x5D: {opcode: "EOR", execute: (*CPU).eor_abs_x, addressing: addressingModeAbsoluteX, cycles: 4},       // EOR Absolute, X
		0x59: {opcode: "EOR", execute: (*CPU).eor_abs_y, addressing: addressingModeAbsoluteY, cycles: 4},       // EOR Absolute, Y
		0x41: {opcode: "EOR", execute: (*CPU).eor_ind_x, addressing: addressingModeIndexedIndirect, cycles: 6}, // EOR (Indirect, X)
		0x51: {opcode: "EOR", execute: (*CPU).eor_ind_y, addressing: addressingModeIndirectIndexed, cycles: 5}, // EOR (Indirect), Y
		0x09: {opcode: "ORA", execute: (*CPU).ora_imm, addressing: addressingModeImmediate, cycles: 2},         // ORA Immediate
		0x05: {opcode: "ORA", execute: (*CPU).ora_zp, addressing: addressingModeZeroPage, cycles: 3},           // ORA Zero Page
		0x15: {opcode: "ORA", execute: (*CPU).ora_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // ORA Zero Page, X
		0x0D: {opcode: "ORA", execute: (*CPU).ora_abs, addressing: addressingModeAbsolute, cycles: 4},          // ORA Absolute
		0x1D: {opcode: "ORA", execute: (*CPU).ora_abs_x, addressing: addressingModeAbsoluteX, cycles: 4},       // ORA Absolute, X
		0x19: {opcode: "ORA", execute: (*CPU).ora_abs_y, addressing: addressingModeAbsoluteY, cycles: 4},       // ORA Absolute, Y
		0x01: {opcode: "ORA", execute: (*CPU).ora_ind_x, addressing: addressingModeIndexedIndirect, cycles: 6}, // ORA (Indirect, X)
		0x11: {opcode: "ORA", execute: (*CPU).ora_ind_y, addressing: addressingModeIndirectIndexed, cycles: 5}, // ORA (Indirect), Y
		0x24: {opcode: "ORA", execute: (*CPU).bit_zp, addressing: addressingModeZeroPage, cycles: 3},           // BIT Zero Page
		0x2C: {opcode: "ORA", execute: (*CPU).bit_abs, addressing: addressingModeAbsolute, cycles: 4},          // BIT Absolute

		// Arithmetic
		0x69: {opcode: "ADC", execute: (*CPU).adc_imm, addressing: addressingModeImmediate, cycles: 2},         // ADC Immediate
		0x65: {opcode: "ADC", execute: (*CPU).adc_zp, addressing: addressingModeZeroPage, cycles: 3},           // ADC Zero Page
		0x75: {opcode: "ADC", execute: (*CPU).adc_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // ADC Zero Page, X
		0x6D: {opcode: "ADC", execute: (*CPU).adc_abs, addressing: addressingModeAbsolute, cycles: 4},          // ADC Absolute
		0x7D: {opcode: "ADC", execute: (*CPU).adc_abs_x, addressing: addressingModeAbsoluteX, cycles: 4},       // ADC Absolute, X
		0x79: {opcode: "ADC", execute: (*CPU).adc_abs_y, addressing: addressingModeAbsoluteY, cycles: 4},       // ADC Absolute, Y
		0x61: {opcode: "ADC", execute: (*CPU).adc_ind_x, addressing: addressingModeIndexedIndirect, cycles: 6}, // ADC (Indirect, X)
		0x71: {opcode: "ADC", execute: (*CPU).adc_ind_y, addressing: addressingModeIndirectIndexed, cycles: 5}, // ADC (Indirect), Y
		0xE9: {opcode: "SBC", execute: (*CPU).sbc_imm, addressing: addressingModeImmediate, cycles: 2},         // SBC Immediate
		0xE5: {opcode: "SBC", execute: (*CPU).sbc_zp, addressing: addressingModeZeroPage, cycles: 3},           // SBC Zero Page
		0xF5: {opcode: "SBC", execute: (*CPU).sbc_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // SBC Zero Page, X
		0xED: {opcode: "SBC", execute: (*CPU).sbc_abs, addressing: addressingModeAbsolute, cycles: 4},          // SBC Absolute
		0xFD: {opcode: "SBC", execute: (*CPU).sbc_abs_x, addressing: addressingModeAbsoluteX, cycles: 4},       // SBC Absolute, X
		0xF9: {opcode: "SBC", execute: (*CPU).sbc_abs_y, addressing: addressingModeAbsoluteY, cycles: 4},       // SBC Absolute, Y
		0xE1: {opcode: "SBC", execute: (*CPU).sbc_ind_x, addressing: addressingModeIndexedIndirect, cycles: 6}, // SBC (Indirect, X)
		0xF1: {opcode: "SBC", execute: (*CPU).sbc_ind_y, addressing: addressingModeIndirectIndexed, cycles: 5}, // SBC (Indirect), Y
		0xC9: {opcode: "CMP", execute: (*CPU).cmp_imm, addressing: addressingModeImmediate, cycles: 2},         // CMP Immediate
		0xC5: {opcode: "CMP", execute: (*CPU).cmp_zp, addressing: addressingModeZeroPage, cycles: 3},           // CMP Zero Page
		0xD5: {opcode: "CMP", execute: (*CPU).cmp_zp_x, addressing: addressingModeZeroPageX, cycles: 4},        // CMP Zero Page, X
		0xCD: {opcode: "CMP", execute: (*CPU).cmp_abs, addressing: addressingModeAbsolute, cycles: 4},          // CMP Absolute
		0xDD: {opcode: "CMP", execute: (*CPU).cmp_abs_x, addressing: addressingModeAbsoluteX, cycles: 4},       // CMP Absolute, X
		0xD9: {opcode: "CMP", execute: (*CPU).cmp_abs_y, addressing: addressingModeAbsoluteY, cycles: 4},       // CMP Absolute, Y
		0xC1: {opcode: "CMP", execute: (*CPU).cmp_ind_x, addressing: addressingModeIndexedIndirect, cycles: 6}, // CMP (Indirect, X)
		0xD1: {opcode: "CMP", execute: (*CPU).cmp_ind_y, addressing: addressingModeIndirectIndexed, cycles: 5}, // CMP (Indirect), Y
		0xE0: {opcode: "CPX", execute: (*CPU).cpx_imm, addressing: addressingModeImmediate, cycles: 2},         // CPX Immediate
		0xE4: {opcode: "CPX", execute: (*CPU).cpx_zp, addressing: addressingModeZeroPage, cycles: 3},           // CPX Zero Page
		0xEC: {opcode: "CPX", execute: (*CPU).cpx_abs, addressing: addressingModeAbsolute, cycles: 4},          // CPX Absolute
		0xC0: {opcode: "CPY", execute: (*CPU).cpy_imm, addressing: addressingModeImmediate, cycles: 2},         // CPY Immediate
		0xC4: {opcode: "CPY", execute: (*CPU).cpy_zp, addressing: addressingModeZeroPage, cycles: 3},           // CPY Zero Page
		0xCC: {opcode: "CPY", execute: (*CPU).cpy_abs, addressing: addressingModeAbsolute, cycles: 4},          // CPY Absolute

		// Increments & Decrements
		0xE6: {opcode: "INC", execute: (*CPU).inc_zp, addressing: addressingModeZeroPage, cycles: 5},     // INC Zero Page
		0xF6: {opcode: "INC", execute: (*CPU).inc_zp_x, addressing: addressingModeZeroPageX, cycles: 6},  // INC Zero Page, X
		0xEE: {opcode: "INC", execute: (*CPU).inc_abs, addressing: addressingModeAbsolute, cycles: 6},    // INC Absolute
		0xFE: {opcode: "INC", execute: (*CPU).inc_abs_x, addressing: addressingModeAbsoluteX, cycles: 7}, // INC Absolute, X
		0xE8: {opcode: "INX", execute: (*CPU).inx, addressing: addressingModeImplied, cycles: 2},         // INX - Increment X
		0xC8: {opcode: "INY", execute: (*CPU).iny, addressing: addressingModeImplied, cycles: 2},         // INY - Increment Y
		0xC6: {opcode: "DEC", execute: (*CPU).dec_zp, addressing: addressingModeZeroPage, cycles: 5},     // DEC Zero Page
		0xD6: {opcode: "DEC", execute: (*CPU).dec_zp_x, addressing: addressingModeZeroPageX, cycles: 6},  // DEC Zero Page, X
		0xCE: {opcode: "DEC", execute: (*CPU).dec_abs, addressing: addressingModeAbsolute, cycles: 6},    // DEC Absolute
		0xDE: {opcode: "DEC", execute: (*CPU).dec_abs_x, addressing: addressingModeAbsoluteX, cycles: 7}, // DEC Absolute, X
		0xCA: {opcode: "DEX", execute: (*CPU).dex, addressing: addressingModeImplied, cycles: 2},         // DEX - Decrement X
		0x88: {opcode: "DEY", execute: (*CPU).dey, addressing: addressingModeImplied, cycles: 2},         // DEY - Decrement Y

		// Shifts
		0x0A: {opcode: "ASL", execute: (*CPU).asl_acc, addressing: addressingModeAccumulator, cycles: 2}, // ASL Accumulator
		0x06: {opcode: "ASL", execute: (*CPU).asl_zp, addressing: addressingModeZeroPage, cycles: 5},     // ASL Zero Page
		0x16: {opcode: "ASL", execute: (*CPU).asl_zp_x, addressing: addressingModeZeroPageX, cycles: 6},  // ASL Zero Page, X
		0x0E: {opcode: "ASL", execute: (*CPU).asl_abs, addressing: addressingModeAbsolute, cycles: 6},    // ASL Absolute
		0x1E: {opcode: "ASL", execute: (*CPU).asl_abs_x, addressing: addressingModeAbsoluteX, cycles: 7}, // ASL Absolute, X
		0x4A: {opcode: "LSR", execute: (*CPU).lsr_acc, addressing: addressingModeAccumulator, cycles: 2}, // LSR Accumulator
		0x46: {opcode: "LSR", execute: (*CPU).lsr_zp, addressing: addressingModeZeroPage, cycles: 5},     // LSR Zero Page
		0x56: {opcode: "LSR", execute: (*CPU).lsr_zp_x, addressing: addressingModeZeroPageX, cycles: 6},  // LSR Zero Page, X
		0x4E: {opcode: "LSR", execute: (*CPU).lsr_abs, addressing: addressingModeAbsolute, cycles: 6},    // LSR Absolute
		0x5E: {opcode: "LSR", execute: (*CPU).lsr_abs_x, addressing: addressingModeAbsoluteX, cycles: 7}, // LSR Absolute, X
		0x2A: {opcode: "ROL", execute: (*CPU).rol_acc, addressing: addressingModeAccumulator, cycles: 2}, // ROL Accumulator
		0x26: {opcode: "ROL", execute: (*CPU).rol_zp, addressing: addressingModeZeroPage, cycles: 5},     // ROL Zero Page
		0x36: {opcode: "ROL", execute: (*CPU).rol_zp_x, addressing: addressingModeZeroPageX, cycles: 6},  // ROL Zero Page, X
		0x2E: {opcode: "ROL", execute: (*CPU).rol_abs, addressing: addressingModeAbsolute, cycles: 6},    // ROL Absolute
		0x3E: {opcode: "ROL", execute: (*CPU).rol_abs_x, addressing: addressingModeAbsoluteX, cycles: 7}, // ROL Absolute, X
		0x6A: {opcode: "ROR", execute: (*CPU).ror_acc, addressing: addressingModeAccumulator, cycles: 2}, // ROR Accumulator
		0x66: {opcode: "ROR", execute: (*CPU).ror_zp, addressing: addressingModeZeroPage, cycles: 5},     // ROR Zero Page
		0x76: {opcode: "ROR", execute: (*CPU).ror_zp_x, addressing: addressingModeZeroPageX, cycles: 6},  // ROR Zero Page, X
		0x6E: {opcode: "ROR", execute: (*CPU).ror_abs, addressing: addressingModeAbsolute, cycles: 6},    // ROR Absolute
		0x7E: {opcode: "ROR", execute: (*CPU).ror_abs_x, addressing: addressingModeAbsoluteX, cycles: 7}, // ROR Absolute, X

		// Jumps & Calls
		0x4C: {opcode: "JMP", execute: (*CPU).jmp_abs, addressing: addressingModeAbsolute, cycles: 3}, // JMP Absolute
		0x6C: {opcode: "JMP", execute: (*CPU).jmp_ind, addressing: addressingModeIndirect, cycles: 5}, // JMP Indirect
		0x20: {opcode: "JSR", execute: (*CPU).jsr, addressing: addressingModeAbsolute, cycles: 6},     // JSR - Jump to Subroutine
		0x60: {opcode: "RTS", execute: (*CPU).rts, addressing: addressingModeImplied, cycles: 6},      // RTS - Return from Subroutine

		// Branches
		0x90: {opcode: "BCC", execute: (*CPU).bcc, addressing: addressingModeRelative, cycles: 2}, // BCC - Branch if Carry Clear
		0xB0: {opcode: "BCS", execute: (*CPU).bcs, addressing: addressingModeRelative, cycles: 2}, // BCS - Branch if Carry Set
		0xF0: {opcode: "BEQ", execute: (*CPU).beq, addressing: addressingModeRelative, cycles: 2}, // BEQ - Branch if Equal
		0x30: {opcode: "BMI", execute: (*CPU).bmi, addressing: addressingModeRelative, cycles: 2}, // BMI - Branch if Minus
		0xD0: {opcode: "BNE", execute: (*CPU).bne, addressing: addressingModeRelative, cycles: 2}, // BNE - Branch if Not Equal
		0x10: {opcode: "BPL", execute: (*CPU).bpl, addressing: addressingModeRelative, cycles: 2}, // BPL - Branch if Positive
		0x50: {opcode: "BVC", execute: (*CPU).bvc, addressing: addressingModeRelative, cycles: 2}, // BVC - Branch if Overflow Clear
		0x70: {opcode: "BVS", execute: (*CPU).bvs, addressing: addressingModeRelative, cycles: 2}, // BVS - Branch if Overflow Set

		// Status Flag
		0x18: {opcode: "CLC", execute: (*CPU).clc, addressing: addressingModeImplied, cycles: 2}, // CLC - Clear Carry Flag
		0xD8: {opcode: "CLD", execute: (*CPU).cld, addressing: addressingModeImplied, cycles: 2}, // CLD - Clear Decimal Mode
		0x58: {opcode: "CLI", execute: (*CPU).cli, addressing: addressingModeImplied, cycles: 2}, // CLI - Clear Interrupt Disable
		0xB8: {opcode: "CLV", execute: (*CPU).clv, addressing: addressingModeImplied, cycles: 2}, // CLV - Clear Overflow Flag
		0x38: {opcode: "SEC", execute: (*CPU).sec, addressing: addressingModeImplied, cycles: 2}, // SEC - Set Carry Flag
		0xF8: {opcode: "SED", execute: (*CPU).sed, addressing: addressingModeImplied, cycles: 2}, // SED - Set Decimal Mode
		0x78: {opcode: "SEI", execute: (*CPU).sei, addressing: addressingModeImplied, cycles: 2}, // SEI - Set Interrupt Disable

		// System Functions
		0x00: {opcode: "BRK", execute: (*CPU).brk, addressing: addressingModeImplied, cycles: 7}, // BRK - Force Interrupt
		0xEA: {opcode: "NOP", execute: (*CPU).nop, addressing: addressingModeImplied, cycles: 2}, // NOP - No Operation
		0x40: {opcode: "RTI", execute: (*CPU).rti, addressing: addressingModeImplied, cycles: 6}, // RTI - Return from Interrupt
	}
)
