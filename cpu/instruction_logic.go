package cpu

func (cpu *CPU) and(value uint8) {
	cpu.A &= value
	cpu.PS.setZeroNeg(cpu.A)
}

func (cpu *CPU) and_imm(operand ...uint8) {
	value := operand[0]
	cpu.and(value)
}

func (cpu *CPU) and_zp(operand ...uint8) {
	addr := operand[0]
	value := cpu.mem.ReadZeroPage(addr)
	cpu.and(value)
}

func (cpu *CPU) and_zp_x(operand ...uint8) {
	addr := (operand[0] + cpu.X)
	value := cpu.mem.ReadZeroPage(addr)
	cpu.and(value)
}

func (cpu *CPU) and_abs(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbs(high, low)
	cpu.and(value)
}

func (cpu *CPU) and_abs_x(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.X)
	cpu.and(value)
}

func (cpu *CPU) and_abs_y(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.Y)
	cpu.and(value)
}

func (cpu *CPU) and_ind_x(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndexedIndirect(ptr, cpu.X)
	cpu.and(value)
}

func (cpu *CPU) and_ind_y(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndirectIndexed(ptr, cpu.Y)
	cpu.and(value)
}

func (cpu *CPU) eor(value uint8) {
	cpu.A ^= value
	cpu.PS.setZeroNeg(cpu.A)
}

func (cpu *CPU) eor_imm(operand ...uint8) {
	value := operand[0]
	cpu.eor(value)
}

func (cpu *CPU) eor_zp(operand ...uint8) {
	addr := operand[0]
	value := cpu.mem.ReadZeroPage(addr)
	cpu.eor(value)
}

func (cpu *CPU) eor_zp_x(operand ...uint8) {
	addr := (operand[0] + cpu.X)
	value := cpu.mem.ReadZeroPage(addr)
	cpu.eor(value)
}

func (cpu *CPU) eor_abs(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbs(high, low)
	cpu.eor(value)
}

func (cpu *CPU) eor_abs_x(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.X)
	cpu.eor(value)
}

func (cpu *CPU) eor_abs_y(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.Y)
	cpu.eor(value)
}

func (cpu *CPU) eor_ind_x(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndexedIndirect(ptr, cpu.X)
	cpu.eor(value)
}

func (cpu *CPU) eor_ind_y(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndirectIndexed(ptr, cpu.Y)
	cpu.eor(value)
}

func (cpu *CPU) ora(value uint8) {
	cpu.A |= value
	cpu.PS.setZeroNeg(cpu.A)
}

func (cpu *CPU) ora_imm(operand ...uint8) {
	value := operand[0]
	cpu.ora(value)
}

func (cpu *CPU) ora_zp(operand ...uint8) {
	addr := operand[0]
	value := cpu.mem.ReadZeroPage(addr)
	cpu.ora(value)
}

func (cpu *CPU) ora_zp_x(operand ...uint8) {
	addr := (operand[0] + cpu.X)
	value := cpu.mem.ReadZeroPage(addr)
	cpu.ora(value)
}

func (cpu *CPU) ora_abs(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbs(high, low)
	cpu.ora(value)
}

func (cpu *CPU) ora_abs_x(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.X)
	cpu.ora(value)
}

func (cpu *CPU) ora_abs_y(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.Y)
	cpu.ora(value)
}

func (cpu *CPU) ora_ind_x(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndexedIndirect(ptr, cpu.X)
	cpu.ora(value)
}

func (cpu *CPU) ora_ind_y(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndirectIndexed(ptr, cpu.Y)
	cpu.ora(value)
}

func (cpu *CPU) bit(value uint8) {
	cpu.PS.setZero((cpu.A & value) == 0)
	cpu.PS.setNegByValue(value)
	cpu.PS.setOverflowByValue(value)
}

func (cpu *CPU) bit_zp(operand ...uint8) {
	addr := operand[0]
	value := cpu.mem.ReadZeroPage(addr)
	cpu.bit(value)
}

func (cpu *CPU) bit_abs(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbs(high, low)
	cpu.bit(value)
}
