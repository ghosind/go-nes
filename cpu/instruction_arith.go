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
	return cpu.adc(val)
}

func (cpu *CPU) adc_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadZeroPage(addr)
	return cpu.adc(val)
}

func (cpu *CPU) adc_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	val := cpu.Mem.ReadZeroPage(addr)
	return cpu.adc(val)
}

func (cpu *CPU) adc_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbs(high, low)
	return cpu.adc(val)
}

func (cpu *CPU) adc_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	return cpu.adc(val)
}

func (cpu *CPU) adc_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbsShift(high, low, cpu.Y)
	return cpu.adc(val)
}

func (cpu *CPU) adc_ind_x(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadIndexedIndirect(addr, cpu.X)
	return cpu.adc(val)
}

func (cpu *CPU) adc_ind_y(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadIndirectIndexed(addr, cpu.Y)
	return cpu.adc(val)
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
	return cpu.sbc(val)
}

func (cpu *CPU) sbc_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadZeroPage(addr)
	return cpu.sbc(val)
}

func (cpu *CPU) sbc_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	val := cpu.Mem.ReadZeroPage(addr)
	return cpu.sbc(val)
}

func (cpu *CPU) sbc_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbs(high, low)
	return cpu.sbc(val)
}

func (cpu *CPU) sbc_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	return cpu.sbc(val)
}

func (cpu *CPU) sbc_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbsShift(high, low, cpu.Y)
	return cpu.sbc(val)
}

func (cpu *CPU) sbc_ind_x(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadIndexedIndirect(addr, cpu.X)
	return cpu.sbc(val)
}

func (cpu *CPU) sbc_ind_y(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadIndirectIndexed(addr, cpu.Y)
	return cpu.sbc(val)
}

func (cpu *CPU) cmp(a, b uint8) uint64 {
	tmp := uint16(a) + uint16(^b) + 1

	cpu.PS.setCarry(tmp > 0xFF)
	cpu.PS.setZeroNeg(uint8(tmp & 0xFF))

	return 0
}

func (cpu *CPU) cmp_imm(operands ...uint8) uint64 {
	val := operands[0]
	return cpu.cmp(cpu.A, val)
}

func (cpu *CPU) cmp_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadZeroPage(addr)
	return cpu.cmp(cpu.A, val)
}

func (cpu *CPU) cmp_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	val := cpu.Mem.ReadZeroPage(addr)
	return cpu.cmp(cpu.A, val)
}

func (cpu *CPU) cmp_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbs(high, low)
	return cpu.cmp(cpu.A, val)
}

func (cpu *CPU) cmp_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	return cpu.cmp(cpu.A, val)
}

func (cpu *CPU) cmp_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbsShift(high, low, cpu.Y)
	return cpu.cmp(cpu.A, val)
}

func (cpu *CPU) cmp_ind_x(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadIndexedIndirect(addr, cpu.X)
	return cpu.cmp(cpu.A, val)
}

func (cpu *CPU) cmp_ind_y(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadIndirectIndexed(addr, cpu.Y)
	return cpu.cmp(cpu.A, val)
}

func (cpu *CPU) cpx_imm(operands ...uint8) uint64 {
	val := operands[0]
	return cpu.cmp(cpu.X, val)
}

func (cpu *CPU) cpx_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadZeroPage(addr)
	return cpu.cmp(cpu.X, val)
}

func (cpu *CPU) cpx_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbs(high, low)
	return cpu.cmp(cpu.X, val)
}

func (cpu *CPU) cpy_imm(operands ...uint8) uint64 {
	val := operands[0]
	return cpu.cmp(cpu.Y, val)
}

func (cpu *CPU) cpy_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadZeroPage(addr)
	return cpu.cmp(cpu.Y, val)
}

func (cpu *CPU) cpy_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbs(high, low)
	return cpu.cmp(cpu.Y, val)
}
