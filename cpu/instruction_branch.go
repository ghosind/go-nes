package cpu

func (cpu *CPU) branch_jump(rel uint8) uint64 {
	cycles := uint64(1)
	newPc := cpu.PC

	if rel&0x80 != 0 {
		rel = 0xFF - rel + 1
		newPc -= uint16(rel)
	} else {
		newPc += uint16(rel)
	}

	if (newPc & 0xFF00) != (cpu.PC & 0xFF00) {
		cycles++
	}

	cpu.PC = newPc

	return cycles
}

func (cpu *CPU) bcc(operands ...uint8) uint64 {
	if cpu.PS.getCarry() {
		return 0
	}

	return cpu.branch_jump(operands[0])
}

func (cpu *CPU) bcs(operands ...uint8) uint64 {
	if !cpu.PS.getCarry() {
		return 0
	}

	return cpu.branch_jump(operands[0])
}

func (cpu *CPU) beq(operands ...uint8) uint64 {
	if !cpu.PS.getZero() {
		return 0
	}

	return cpu.branch_jump(operands[0])
}

func (cpu *CPU) bmi(operands ...uint8) uint64 {
	if !cpu.PS.getNegative() {
		return 0
	}

	return cpu.branch_jump(operands[0])
}

func (cpu *CPU) bne(operands ...uint8) uint64 {
	if cpu.PS.getZero() {
		return 0
	}

	return cpu.branch_jump(operands[0])
}

func (cpu *CPU) bpl(operands ...uint8) uint64 {
	if cpu.PS.getNegative() {
		return 0
	}

	return cpu.branch_jump(operands[0])
}

func (cpu *CPU) bvc(operands ...uint8) uint64 {
	if cpu.PS.getOverflow() {
		return 0
	}

	return cpu.branch_jump(operands[0])
}

func (cpu *CPU) bvs(operands ...uint8) uint64 {
	if !cpu.PS.getOverflow() {
		return 0
	}

	return cpu.branch_jump(operands[0])
}
