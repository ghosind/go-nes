package cpu

var (
	InstructionMap map[uint8]func(*CPU) = map[uint8]func(*CPU){
		// Load/Store Operations
		0xA9: (*CPU).lda_imm,   // LDA Immediate
		0xA5: (*CPU).lda_zp,    // LDA Zero Page
		0xB5: (*CPU).lda_zp_x,  // LDA Zero Page, X
		0xAD: (*CPU).lda_abs,   // LDA Absolute
		0xBD: (*CPU).lda_abs_x, // LDA Absolute, X
		0xB9: (*CPU).lda_abs_y, // LDA Absolute, Y
		0xA1: (*CPU).lda_ind_x, // LDA (Indirect, X)
		0xB1: (*CPU).lda_ind_y, // LDA (Indirect), Y
		0xA2: (*CPU).ldx_imm,   // LDX Immediate
		0xA6: (*CPU).ldx_zp,    // LDX Zero Page
		0xB6: (*CPU).ldx_zp_y,  // LDX Zero Page, Y
		0xAE: (*CPU).ldx_abs,   // LDX Absolute
		0xBE: (*CPU).ldx_abs_y, // LDX Absolute, Y
		0xA0: (*CPU).ldy_imm,   // LDY Immediate
		0xA4: (*CPU).ldy_zp,    // LDY Zero Page
		0xB4: (*CPU).ldy_zp_x,  // LDY Zero Page, X
		0xAC: (*CPU).ldy_abs,   // LDY Absolute
		0xBC: (*CPU).ldy_abs_x, // LDY Absolute, X
		0x85: nil,              // STA Zero Page
		0x95: nil,              // STA Zero Page, X
		0x8D: nil,              // STA Absolute
		0x9D: nil,              // STA Absolute, X
		0x99: nil,              // STA Absolute, Y
		0x81: nil,              // STA (Indirect, X)
		0x91: nil,              // STA (Indirect), Y
		0x86: nil,              // STX Zero Page
		0x96: nil,              // STX Zero Page, Y
		0x8E: nil,              // STX Absolute
		0x84: nil,              // STY Zero Page
		0x94: nil,              // STY Zero Page, X
		0x8C: nil,              // STY Absolute

		// Register Transfers
		0xAA: nil, // TAX - Transfer Accumulator to X
		0xA8: nil, // TAY - Transfer Accumulator to Y
		0x8A: nil, // TXA - Transfer X to Accumulator
		0x98: nil, // TYA - Transfer Y to Accumulator

		// Stack Operations
		0xBA: nil, // TSX - Transfer Stack Pointer to X
		0x9A: nil, // TXS - Transfer X to Stack Pointer
		0x48: nil, // PHA - Push Accumulator onto Stack
		0x68: nil, // PLA - Pull Accumulator from Stack
		0x08: nil, // PHP - Push Processor Status onto Stack
		0x28: nil, // PLP - Pull Processor Status from Stack

		// Logical Operations
		0x29: nil, // AND Immediate
		0x25: nil, // AND Zero Page
		0x35: nil, // AND Zero Page, X
		0x2D: nil, // AND Absolute
		0x3D: nil, // AND Absolute, X
		0x39: nil, // AND Absolute, Y
		0x21: nil, // AND (Indirect, X)
		0x31: nil, // AND (Indirect), Y
		0x49: nil, // EOR Immediate
		0x45: nil, // EOR Zero Page
		0x55: nil, // EOR Zero Page, X
		0x4D: nil, // EOR Absolute
		0x5D: nil, // EOR Absolute, X
		0x59: nil, // EOR Absolute, Y
		0x41: nil, // EOR (Indirect, X)
		0x51: nil, // EOR (Indirect), Y
		0x09: nil, // ORA Immediate
		0x05: nil, // ORA Zero Page
		0x15: nil, // ORA Zero Page, X
		0x0D: nil, // ORA Absolute
		0x1D: nil, // ORA Absolute, X
		0x19: nil, // ORA Absolute, Y
		0x01: nil, // ORA (Indirect, X)
		0x11: nil, // ORA (Indirect), Y
		0x24: nil, // BIT Zero Page
		0x2C: nil, // BIT Absolute

		// Arithmetic
		0x69: nil, // ADC Immediate
		0x65: nil, // ADC Zero Page
		0x75: nil, // ADC Zero Page, X
		0x6D: nil, // ADC Absolute
		0x7D: nil, // ADC Absolute, X
		0x79: nil, // ADC Absolute, Y
		0x61: nil, // ADC (Indirect, X)
		0x71: nil, // ADC (Indirect), Y
		0xE9: nil, // SBC Immediate
		0xE5: nil, // SBC Zero Page
		0xF5: nil, // SBC Zero Page, X
		0xED: nil, // SBC Absolute
		0xFD: nil, // SBC Absolute, X
		0xF9: nil, // SBC Absolute, Y
		0xE1: nil, // SBC (Indirect, X)
		0xF1: nil, // SBC (Indirect), Y
		0xC9: nil, // CMP Immediate
		0xC5: nil, // CMP Zero Page
		0xD5: nil, // CMP Zero Page, X
		0xCD: nil, // CMP Absolute
		0xDD: nil, // CMP Absolute, X
		0xD9: nil, // CMP Absolute, Y
		0xC1: nil, // CMP (Indirect, X)
		0xD1: nil, // CMP (Indirect), Y
		0xE0: nil, // CPX Immediate
		0xE4: nil, // CPX Zero Page
		0xEC: nil, // CPX Absolute
		0xC0: nil, // CPY Immediate
		0xC4: nil, // CPY Zero Page
		0xCC: nil, // CPY Absolute

		// Increments & Decrements
		0xE6: nil, // INC Zero Page
		0xF6: nil, // INC Zero Page, X
		0xEE: nil, // INC Absolute
		0xFE: nil, // INC Absolute, X
		0xE8: nil, // INX - Increment X
		0xC8: nil, // INY - Increment Y
		0xC6: nil, // DEC Zero Page
		0xD6: nil, // DEC Zero Page, X
		0xCE: nil, // DEC Absolute
		0xDE: nil, // DEC Absolute, X
		0xCA: nil, // DEX - Decrement X
		0x88: nil, // DEY - Decrement Y

		// Shifts
		0x0A: nil, // ASL Accumulator
		0x06: nil, // ASL Zero Page
		0x16: nil, // ASL Zero Page, X
		0x0E: nil, // ASL Absolute
		0x1E: nil, // ASL Absolute, X
		0x4A: nil, // LSR Accumulator
		0x46: nil, // LSR Zero Page
		0x56: nil, // LSR Zero Page, X
		0x4E: nil, // LSR Absolute
		0x5E: nil, // LSR Absolute, X
		0x2A: nil, // ROL Accumulator
		0x26: nil, // ROL Zero Page
		0x36: nil, // ROL Zero Page, X
		0x2E: nil, // ROL Absolute
		0x3E: nil, // ROL Absolute, X
		0x6A: nil, // ROR Accumulator
		0x66: nil, // ROR Zero Page
		0x76: nil, // ROR Zero Page, X
		0x6E: nil, // ROR Absolute
		0x7E: nil, // ROR Absolute, X

		// Jumps & Calls
		0x4C: nil, // JMP Absolute
		0x6C: nil, // JMP Indirect
		0x20: nil, // JSR - Jump to Subroutine
		0x60: nil, // RTS - Return from Subroutine

		// Branches
		0x90: nil, // BCC - Branch if Carry Clear
		0xB0: nil, // BCS - Branch if Carry Set
		0xF0: nil, // BEQ - Branch if Equal
		0x30: nil, // BMI - Branch if Minus
		0xD0: nil, // BNE - Branch if Not Equal
		0x10: nil, // BPL - Branch if Positive
		0x50: nil, // BVC - Branch if Overflow Clear
		0x70: nil, // BVS - Branch if Overflow Set

		// Status Flag
		0x18: (*CPU).clc, // CLC - Clear Carry Flag
		0xD8: (*CPU).cld, // CLD - Clear Decimal Mode
		0x58: (*CPU).cli, // CLI - Clear Interrupt Disable
		0xB8: (*CPU).clv, // CLV - Clear Overflow Flag
		0x38: (*CPU).sec, // SEC - Set Carry Flag
		0xF8: (*CPU).sed, // SED - Set Decimal Mode
		0x78: (*CPU).sei, // SEI - Set Interrupt Disable

		// System Functions
		0x00: nil,        // BRK - Force Interrupt
		0xEA: (*CPU).nop, // NOP - No Operation
		0x40: nil,        // RTI - Return from Interrupt
	}
)

func (cpu *CPU) clc() {
	cpu.ps.setCarry(false)
}

func (cpu *CPU) cld() {
	cpu.ps.setDecimal(false)
}

func (cpu *CPU) cli() {
	cpu.ps.setInterrupt(false)
}

func (cpu *CPU) clv() {
	cpu.ps.setOverflow(false)
}

func (cpu *CPU) lda_imm() {
	cpu.a = cpu.fetch()
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_zp() {
	addr := cpu.fetch()
	cpu.a = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_zp_x() {
	addr := cpu.fetch() + cpu.x
	cpu.a = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_abs() {
	low := cpu.fetch()
	high := cpu.fetch()
	cpu.a = cpu.mem.ReadAbs(high, low)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_abs_x() {
	low := cpu.fetch()
	high := cpu.fetch()
	cpu.a = cpu.mem.ReadAbsShift(high, low, cpu.x)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_abs_y() {
	low := cpu.fetch()
	high := cpu.fetch()
	cpu.a = cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_ind_x() {
	addr := cpu.fetch() + cpu.x
	low := cpu.mem.ReadZeroPage(addr)
	high := cpu.mem.ReadZeroPage(addr + 1)
	cpu.a = cpu.mem.ReadAbs(high, low)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_ind_y() {
	addr := cpu.fetch()
	low := cpu.mem.ReadZeroPage(addr)
	high := cpu.mem.ReadZeroPage(addr + 1)
	cpu.a = cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) ldx_imm() {
	cpu.x = cpu.fetch()
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) ldx_zp() {
	addr := cpu.fetch()
	cpu.x = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) ldx_zp_y() {
	addr := cpu.fetch() + cpu.y
	cpu.x = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) ldx_abs() {
	low := cpu.fetch()
	high := cpu.fetch()
	cpu.x = cpu.mem.ReadAbs(high, low)
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) ldx_abs_y() {
	low := cpu.fetch()
	high := cpu.fetch()
	cpu.x = cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) ldy_imm() {
	cpu.y = cpu.fetch()
	cpu.ps.setZeroNeg(cpu.y)
}

func (cpu *CPU) ldy_zp() {
	addr := cpu.fetch()
	cpu.y = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.y)
}

func (cpu *CPU) ldy_zp_x() {
	addr := cpu.fetch() + cpu.x
	cpu.y = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.y)
}

func (cpu *CPU) ldy_abs() {
	low := cpu.fetch()
	high := cpu.fetch()
	cpu.y = cpu.mem.ReadAbs(high, low)
	cpu.ps.setZeroNeg(cpu.y)
}

func (cpu *CPU) ldy_abs_x() {
	low := cpu.fetch()
	high := cpu.fetch()
	cpu.y = cpu.mem.ReadAbsShift(high, low, cpu.x)
	cpu.ps.setZeroNeg(cpu.y)
}

func (cpu *CPU) nop() {
	// Do nothing
}

func (cpu *CPU) sec() {
	cpu.ps.setCarry(true)
}

func (cpu *CPU) sed() {
	cpu.ps.setDecimal(true)
}

func (cpu *CPU) sei() {
	cpu.ps.setInterrupt(true)
}
