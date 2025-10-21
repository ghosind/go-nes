package cpu

func (cpu *CPU) and(value uint8) uint64 {
	cpu.A &= value
	cpu.PS.setZeroNeg(cpu.A)
	return 0
}

func (cpu *CPU) and_imm(operand ...uint8) uint64 {
	value := operand[0]
	return cpu.and(value)
}

func (cpu *CPU) and_zp(operand ...uint8) uint64 {
	addr := operand[0]
	value := cpu.Mem.ReadZeroPage(addr)
	return cpu.and(value)
}

func (cpu *CPU) and_zp_x(operand ...uint8) uint64 {
	addr := (operand[0] + cpu.X)
	value := cpu.Mem.ReadZeroPage(addr)
	return cpu.and(value)
}

func (cpu *CPU) and_abs(operand ...uint8) uint64 {
	low := operand[0]
	high := operand[1]
	value := cpu.Mem.ReadAbs(high, low)
	return cpu.and(value)
}

func (cpu *CPU) and_abs_x(operand ...uint8) uint64 {
	low := operand[0]
	high := operand[1]
	value := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	return cpu.and(value)
}

func (cpu *CPU) and_abs_y(operand ...uint8) uint64 {
	low := operand[0]
	high := operand[1]
	value := cpu.Mem.ReadAbsShift(high, low, cpu.Y)
	return cpu.and(value)
}

func (cpu *CPU) and_ind_x(operand ...uint8) uint64 {
	ptr := operand[0]
	value := cpu.Mem.ReadIndexedIndirect(ptr, cpu.X)
	return cpu.and(value)
}

func (cpu *CPU) and_ind_y(operand ...uint8) uint64 {
	ptr := operand[0]
	value := cpu.Mem.ReadIndirectIndexed(ptr, cpu.Y)
	return cpu.and(value)
}

func (cpu *CPU) eor(value uint8) uint64 {
	cpu.A ^= value
	cpu.PS.setZeroNeg(cpu.A)
	return 0
}

func (cpu *CPU) eor_imm(operand ...uint8) uint64 {
	value := operand[0]
	return cpu.eor(value)
}

func (cpu *CPU) eor_zp(operand ...uint8) uint64 {
	addr := operand[0]
	value := cpu.Mem.ReadZeroPage(addr)
	return cpu.eor(value)
}

func (cpu *CPU) eor_zp_x(operand ...uint8) uint64 {
	addr := (operand[0] + cpu.X)
	value := cpu.Mem.ReadZeroPage(addr)
	return cpu.eor(value)
}

func (cpu *CPU) eor_abs(operand ...uint8) uint64 {
	low := operand[0]
	high := operand[1]
	value := cpu.Mem.ReadAbs(high, low)
	return cpu.eor(value)
}

func (cpu *CPU) eor_abs_x(operand ...uint8) uint64 {
	low := operand[0]
	high := operand[1]
	value := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	return cpu.eor(value)
}

func (cpu *CPU) eor_abs_y(operand ...uint8) uint64 {
	low := operand[0]
	high := operand[1]
	value := cpu.Mem.ReadAbsShift(high, low, cpu.Y)
	return cpu.eor(value)
}

func (cpu *CPU) eor_ind_x(operand ...uint8) uint64 {
	ptr := operand[0]
	value := cpu.Mem.ReadIndexedIndirect(ptr, cpu.X)
	return cpu.eor(value)
}

func (cpu *CPU) eor_ind_y(operand ...uint8) uint64 {
	ptr := operand[0]
	value := cpu.Mem.ReadIndirectIndexed(ptr, cpu.Y)
	return cpu.eor(value)
}

func (cpu *CPU) ora(value uint8) uint64 {
	cpu.A |= value
	cpu.PS.setZeroNeg(cpu.A)
	return 0
}

func (cpu *CPU) ora_imm(operand ...uint8) uint64 {
	value := operand[0]
	return cpu.ora(value)
}

func (cpu *CPU) ora_zp(operand ...uint8) uint64 {
	addr := operand[0]
	value := cpu.Mem.ReadZeroPage(addr)
	return cpu.ora(value)
}

func (cpu *CPU) ora_zp_x(operand ...uint8) uint64 {
	addr := (operand[0] + cpu.X)
	value := cpu.Mem.ReadZeroPage(addr)
	return cpu.ora(value)
}

func (cpu *CPU) ora_abs(operand ...uint8) uint64 {
	low := operand[0]
	high := operand[1]
	value := cpu.Mem.ReadAbs(high, low)
	return cpu.ora(value)
}

func (cpu *CPU) ora_abs_x(operand ...uint8) uint64 {
	low := operand[0]
	high := operand[1]
	value := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	return cpu.ora(value)
}

func (cpu *CPU) ora_abs_y(operand ...uint8) uint64 {
	low := operand[0]
	high := operand[1]
	value := cpu.Mem.ReadAbsShift(high, low, cpu.Y)
	return cpu.ora(value)
}

func (cpu *CPU) ora_ind_x(operand ...uint8) uint64 {
	ptr := operand[0]
	value := cpu.Mem.ReadIndexedIndirect(ptr, cpu.X)
	return cpu.ora(value)
}

func (cpu *CPU) ora_ind_y(operand ...uint8) uint64 {
	ptr := operand[0]
	value := cpu.Mem.ReadIndirectIndexed(ptr, cpu.Y)
	return cpu.ora(value)
}

func (cpu *CPU) bit(value uint8) uint64 {
	cpu.PS.setZero((cpu.A & value) == 0)
	cpu.PS.setNegByValue(value)
	cpu.PS.setOverflowByValue(value)
	return 0
}

func (cpu *CPU) bit_zp(operand ...uint8) uint64 {
	addr := operand[0]
	value := cpu.Mem.ReadZeroPage(addr)
	return cpu.bit(value)
}

func (cpu *CPU) bit_abs(operand ...uint8) uint64 {
	low := operand[0]
	high := operand[1]
	value := cpu.Mem.ReadAbs(high, low)
	return cpu.bit(value)
}
