package cpu

func (cpu *CPU) bcc(operands ...uint8) uint64 {
	if cpu.PS.getCarry() {
		return 0
	}

	rel := operands[0]
	if rel&0x80 != 0 {
		rel = 0xFF - rel + 1
		cpu.PC -= uint16(rel)
	} else {
		cpu.PC += uint16(rel)
	}

	return 1
}

func (cpu *CPU) bcs(operands ...uint8) uint64 {
	if !cpu.PS.getCarry() {
		return 0
	}

	rel := operands[0]
	if rel&0x80 != 0 {
		rel = 0xFF - rel + 1
		cpu.PC -= uint16(rel)
	} else {
		cpu.PC += uint16(rel)
	}

	return 1
}

func (cpu *CPU) beq(operands ...uint8) uint64 {
	if !cpu.PS.getZero() {
		return 0
	}

	rel := operands[0]
	if rel&0x80 != 0 {
		rel = 0xFF - rel + 1
		cpu.PC -= uint16(rel)
	} else {
		cpu.PC += uint16(rel)
	}

	return 1
}

func (cpu *CPU) bmi(operands ...uint8) uint64 {
	if !cpu.PS.getNegative() {
		return 0
	}

	rel := operands[0]
	if rel&0x80 != 0 {
		rel = 0xFF - rel + 1
		cpu.PC -= uint16(rel)
	} else {
		cpu.PC += uint16(rel)
	}

	return 1
}

func (cpu *CPU) bne(operands ...uint8) uint64 {
	if cpu.PS.getZero() {
		return 0
	}

	rel := operands[0]
	if rel&0x80 != 0 {
		rel = 0xFF - rel + 1
		cpu.PC -= uint16(rel)
	} else {
		cpu.PC += uint16(rel)
	}

	return 1
}

func (cpu *CPU) bpl(operands ...uint8) uint64 {
	if cpu.PS.getNegative() {
		return 0
	}

	rel := operands[0]
	if rel&0x80 != 0 {
		rel = 0xFF - rel + 1
		cpu.PC -= uint16(rel)
	} else {
		cpu.PC += uint16(rel)
	}

	return 1
}

func (cpu *CPU) bvc(operands ...uint8) uint64 {
	if cpu.PS.getOverflow() {
		return 0
	}

	rel := operands[0]
	if rel&0x80 != 0 {
		rel = 0xFF - rel + 1
		cpu.PC -= uint16(rel)
	} else {
		cpu.PC += uint16(rel)
	}

	return 1
}

func (cpu *CPU) bvs(operands ...uint8) uint64 {
	if !cpu.PS.getOverflow() {
		return 0
	}

	rel := operands[0]
	if rel&0x80 != 0 {
		rel = 0xFF - rel + 1
		cpu.PC -= uint16(rel)
	} else {
		cpu.PC += uint16(rel)
	}

	return 1
}
