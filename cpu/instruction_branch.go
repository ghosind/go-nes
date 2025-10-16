package cpu

func (cpu *CPU) bcc(operands ...uint8) {
	if cpu.ps.getCarry() {
		return
	}

	rel := operands[0]
	cpu.pc += uint16(rel)
}

func (cpu *CPU) bcs(operands ...uint8) {
	if !cpu.ps.getCarry() {
		return
	}

	rel := operands[0]
	cpu.pc += uint16(rel)
}

func (cpu *CPU) beq(operands ...uint8) {
	if !cpu.ps.getZero() {
		return
	}

	rel := operands[0]
	cpu.pc += uint16(rel)
}

func (cpu *CPU) bmi(operands ...uint8) {
	if !cpu.ps.getNegative() {
		return
	}

	rel := operands[0]
	cpu.pc += uint16(rel)
}

func (cpu *CPU) bne(operands ...uint8) {
	if cpu.ps.getZero() {
		return
	}

	rel := operands[0]
	cpu.pc += uint16(rel)
}

func (cpu *CPU) bpl(operands ...uint8) {
	if cpu.ps.getNegative() {
		return
	}

	rel := operands[0]
	cpu.pc += uint16(rel)
}

func (cpu *CPU) bvc(operands ...uint8) {
	if cpu.ps.getOverflow() {
		return
	}

	rel := operands[0]
	cpu.pc += uint16(rel)
}

func (cpu *CPU) bvs(operands ...uint8) {
	if !cpu.ps.getOverflow() {
		return
	}

	rel := operands[0]
	cpu.pc += uint16(rel)
}
