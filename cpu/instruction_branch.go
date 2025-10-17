package cpu

func (cpu *CPU) bcc(operands ...uint8) {
	if cpu.PS.getCarry() {
		return
	}

	rel := operands[0]
	cpu.PC += uint16(rel)
}

func (cpu *CPU) bcs(operands ...uint8) {
	if !cpu.PS.getCarry() {
		return
	}

	rel := operands[0]
	cpu.PC += uint16(rel)
}

func (cpu *CPU) beq(operands ...uint8) {
	if !cpu.PS.getZero() {
		return
	}

	rel := operands[0]
	cpu.PC += uint16(rel)
}

func (cpu *CPU) bmi(operands ...uint8) {
	if !cpu.PS.getNegative() {
		return
	}

	rel := operands[0]
	cpu.PC += uint16(rel)
}

func (cpu *CPU) bne(operands ...uint8) {
	if cpu.PS.getZero() {
		return
	}

	rel := operands[0]
	cpu.PC += uint16(rel)
}

func (cpu *CPU) bpl(operands ...uint8) {
	if cpu.PS.getNegative() {
		return
	}

	rel := operands[0]
	cpu.PC += uint16(rel)
}

func (cpu *CPU) bvc(operands ...uint8) {
	if cpu.PS.getOverflow() {
		return
	}

	rel := operands[0]
	cpu.PC += uint16(rel)
}

func (cpu *CPU) bvs(operands ...uint8) {
	if !cpu.PS.getOverflow() {
		return
	}

	rel := operands[0]
	cpu.PC += uint16(rel)
}
