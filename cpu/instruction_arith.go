package cpu

func (cpu *CPU) adc(value uint8) uint64 {
	carry := uint8(0)
	if cpu.PS.getCarry() {
		carry = 1
	}

	tmp := uint16(cpu.A) + uint16(value) + uint16(carry)
	result := uint8(tmp & 0xFF)

	cpu.PS.setCarry(tmp > 0xFF)
	cpu.PS.setOverflow((^(cpu.A ^ value) & (cpu.A ^ result) & 0x80) != 0)
	cpu.PS.setZeroNeg(result)

	cpu.A = result
	return 0
}

func (cpu *CPU) adc_imm(operands ...uint8) uint64 {
	val := operands[0]
	cpu.adc(val)
	return 0
}

func (cpu *CPU) adc_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	cpu.adc(val)
	return 0
}

func (cpu *CPU) adc_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	val := cpu.mem.ReadZeroPage(addr)
	cpu.adc(val)
	return 0
}

func (cpu *CPU) adc_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	cpu.adc(val)
	return 0
}

func (cpu *CPU) adc_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.X)
	cpu.adc(val)
	return 0
}

func (cpu *CPU) adc_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.Y)
	cpu.adc(val)
	return 0
}

func (cpu *CPU) adc_ind_x(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.mem.ReadIndexedIndirect(addr, cpu.X)
	cpu.adc(val)
	return 0
}

func (cpu *CPU) adc_ind_y(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.mem.ReadIndirectIndexed(addr, cpu.Y)
	cpu.adc(val)
	return 0
}

func (cpu *CPU) sbc(value uint8) uint64 {
	carry := uint8(0)
	if cpu.PS.getCarry() {
		carry = 1
	}

	tmp := uint16(cpu.A) + uint16(^value) + uint16(carry)
	result := uint8(tmp & 0xFF)

	cpu.PS.setCarry(tmp > 0xFF)
	cpu.PS.setOverflow(((cpu.A ^ result) & ((value ^ 0xFF) ^ result) & 0x80) != 0)
	cpu.PS.setZeroNeg(result)

	cpu.A = result
	return 0
}

func (cpu *CPU) sbc_imm(operands ...uint8) uint64 {
	val := operands[0]
	cpu.sbc(val)
	return 0
}

func (cpu *CPU) sbc_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	cpu.sbc(val)
	return 0
}

func (cpu *CPU) sbc_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	val := cpu.mem.ReadZeroPage(addr)
	cpu.sbc(val)
	return 0
}

func (cpu *CPU) sbc_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	cpu.sbc(val)
	return 0
}

func (cpu *CPU) sbc_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.X)
	cpu.sbc(val)
	return 0
}

func (cpu *CPU) sbc_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.Y)
	cpu.sbc(val)
	return 0
}

func (cpu *CPU) sbc_ind_x(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.mem.ReadIndexedIndirect(addr, cpu.X)
	cpu.sbc(val)
	return 0
}

func (cpu *CPU) sbc_ind_y(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.mem.ReadIndirectIndexed(addr, cpu.Y)
	cpu.sbc(val)
	return 0
}

func (cpu *CPU) cmp(a, b uint8) uint64 {
	tmp := uint16(a) + uint16(^b) + 1

	cpu.PS.setCarry(tmp > 0xFF)
	cpu.PS.setZeroNeg(uint8(tmp & 0xFF))
	return 0
}

func (cpu *CPU) cmp_imm(operands ...uint8) uint64 {
	val := operands[0]
	cpu.cmp(cpu.A, val)
	return 0
}

func (cpu *CPU) cmp_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	cpu.cmp(cpu.A, val)
	return 0
}

func (cpu *CPU) cmp_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	val := cpu.mem.ReadZeroPage(addr)
	cpu.cmp(cpu.A, val)
	return 0
}

func (cpu *CPU) cmp_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	cpu.cmp(cpu.A, val)
	return 0
}

func (cpu *CPU) cmp_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.X)
	cpu.cmp(cpu.A, val)
	return 0
}

func (cpu *CPU) cmp_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.Y)
	cpu.cmp(cpu.A, val)
	return 0
}

func (cpu *CPU) cmp_ind_x(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.mem.ReadIndexedIndirect(addr, cpu.X)
	cpu.cmp(cpu.A, val)
	return 0
}

func (cpu *CPU) cmp_ind_y(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.mem.ReadIndirectIndexed(addr, cpu.Y)
	cpu.cmp(cpu.A, val)
	return 0
}

func (cpu *CPU) cpx_imm(operands ...uint8) uint64 {
	val := operands[0]
	cpu.cmp(cpu.X, val)
	return 0
}

func (cpu *CPU) cpx_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	cpu.cmp(cpu.X, val)
	return 0
}

func (cpu *CPU) cpx_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	cpu.cmp(cpu.X, val)
	return 0
}

func (cpu *CPU) cpy_imm(operands ...uint8) uint64 {
	val := operands[0]
	cpu.cmp(cpu.Y, val)
	return 0
}

func (cpu *CPU) cpy_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	cpu.cmp(cpu.Y, val)
	return 0
}

func (cpu *CPU) cpy_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	cpu.cmp(cpu.Y, val)
	return 0
}
