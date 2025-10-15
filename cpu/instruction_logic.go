package cpu

func (cpu *CPU) and_imm(operand ...uint8) {
	cpu.a &= operand[0]
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) and_zp(operand ...uint8) {
	addr := operand[0]
	value := cpu.mem.ReadZeroPage(addr)
	cpu.a &= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) and_zp_x(operand ...uint8) {
	addr := (operand[0] + cpu.x)
	value := cpu.mem.ReadZeroPage(addr)
	cpu.a &= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) and_abs(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbs(high, low)
	cpu.a &= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) and_abs_x(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.x)
	cpu.a &= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) and_abs_y(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.a &= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) and_ind_x(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndexedIndirect(ptr, cpu.x)
	cpu.a &= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) and_ind_y(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndirectIndexed(ptr, cpu.y)
	cpu.a &= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) eor_imm(operand ...uint8) {
	cpu.a ^= operand[0]
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) eor_zp(operand ...uint8) {
	addr := operand[0]
	value := cpu.mem.ReadZeroPage(addr)
	cpu.a ^= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) eor_zp_x(operand ...uint8) {
	addr := (operand[0] + cpu.x)
	value := cpu.mem.ReadZeroPage(addr)
	cpu.a ^= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) eor_abs(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbs(high, low)
	cpu.a ^= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) eor_abs_x(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.x)
	cpu.a ^= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) eor_abs_y(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.a ^= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) eor_ind_x(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndexedIndirect(ptr, cpu.x)
	cpu.a ^= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) eor_ind_y(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndirectIndexed(ptr, cpu.y)
	cpu.a ^= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) ora_imm(operand ...uint8) {
	cpu.a |= operand[0]
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) ora_zp(operand ...uint8) {
	addr := operand[0]
	value := cpu.mem.ReadZeroPage(addr)
	cpu.a |= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) ora_zp_x(operand ...uint8) {
	addr := (operand[0] + cpu.x)
	value := cpu.mem.ReadZeroPage(addr)
	cpu.a |= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) ora_abs(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbs(high, low)
	cpu.a |= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) ora_abs_x(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.x)
	cpu.a |= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) ora_abs_y(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.a |= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) ora_ind_x(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndexedIndirect(ptr, cpu.x)
	cpu.a |= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) ora_ind_y(operand ...uint8) {
	ptr := operand[0]
	value := cpu.mem.ReadIndirectIndexed(ptr, cpu.y)
	cpu.a |= value
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) bit_zp(operand ...uint8) {
	addr := operand[0]
	value := cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZero((cpu.a & value) == 0)
	cpu.ps.setNegByValue(value)
	cpu.ps.setOverflowByValue(value)
}

func (cpu *CPU) bit_abs(operand ...uint8) {
	low := operand[0]
	high := operand[1]
	value := cpu.mem.ReadAbs(high, low)
	cpu.ps.setZero((cpu.a & value) == 0)
	cpu.ps.setNegByValue(value)
	cpu.ps.setOverflowByValue(value)
}
