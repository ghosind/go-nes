package cpu

func (cpu *CPU) adc(value uint8) {
	carry := uint8(0)
	if cpu.ps.getCarry() {
		carry = 1
	}

	tmp := uint16(cpu.a) + uint16(value) + uint16(carry)
	result := uint8(tmp & 0xFF)

	cpu.ps.setCarry(tmp > 0xFF)
	cpu.ps.setOverflow((^(cpu.a ^ value) & (cpu.a ^ result) & 0x80) != 0)
	cpu.ps.setZeroNeg(result)

	cpu.a = result
}

func (cpu *CPU) adc_imm(operands ...uint8) {
	val := operands[0]
	cpu.adc(val)
}

func (cpu *CPU) adc_zp(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	cpu.adc(val)
}

func (cpu *CPU) adc_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
	val := cpu.mem.ReadZeroPage(addr)
	cpu.adc(val)
}

func (cpu *CPU) adc_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	cpu.adc(val)
}

func (cpu *CPU) adc_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.x)
	cpu.adc(val)
}

func (cpu *CPU) adc_abs_y(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.adc(val)
}

func (cpu *CPU) adc_ind_x(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadIndexedIndirect(addr, cpu.x)
	cpu.adc(val)
}

func (cpu *CPU) adc_ind_y(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadIndirectIndexed(addr, cpu.y)
	cpu.adc(val)
}

func (cpu *CPU) sbc(value uint8) {
	carry := uint8(0)
	if cpu.ps.getCarry() {
		carry = 1
	}

	tmp := uint16(cpu.a) + uint16(^value) + uint16(carry)
	result := uint8(tmp & 0xFF)

	cpu.ps.setCarry(tmp > 0xFF)
	cpu.ps.setOverflow((^(cpu.a ^ value) & (cpu.a ^ result) & 0x80) != 0)
	cpu.ps.setZeroNeg(result)

	cpu.a = result
}

func (cpu *CPU) sbc_imm(operands ...uint8) {
	val := operands[0]
	cpu.sbc(val)
}

func (cpu *CPU) sbc_zp(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	cpu.sbc(val)
}

func (cpu *CPU) sbc_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
	val := cpu.mem.ReadZeroPage(addr)
	cpu.sbc(val)
}

func (cpu *CPU) sbc_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	cpu.sbc(val)
}

func (cpu *CPU) sbc_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.x)
	cpu.sbc(val)
}

func (cpu *CPU) sbc_abs_y(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.sbc(val)
}

func (cpu *CPU) sbc_ind_x(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadIndexedIndirect(addr, cpu.x)
	cpu.sbc(val)
}

func (cpu *CPU) sbc_ind_y(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadIndirectIndexed(addr, cpu.y)
	cpu.sbc(val)
}

func (cpu *CPU) cmp(a, b uint8) {
	tmp := uint16(a) + uint16(^b) + 1

	cpu.ps.setCarry(tmp > 0xFF)
	cpu.ps.setZeroNeg(uint8(tmp & 0xFF))
}

func (cpu *CPU) cmp_imm(operands ...uint8) {
	val := operands[0]
	cpu.cmp(cpu.a, val)
}

func (cpu *CPU) cmp_zp(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	cpu.cmp(cpu.a, val)
}

func (cpu *CPU) cmp_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
	val := cpu.mem.ReadZeroPage(addr)
	cpu.cmp(cpu.a, val)
}

func (cpu *CPU) cmp_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	cpu.cmp(cpu.a, val)
}

func (cpu *CPU) cmp_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.x)
	cpu.cmp(cpu.a, val)
}

func (cpu *CPU) cmp_abs_y(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.cmp(cpu.a, val)
}

func (cpu *CPU) cmp_ind_x(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadIndexedIndirect(addr, cpu.x)
	cpu.cmp(cpu.a, val)
}

func (cpu *CPU) cmp_ind_y(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadIndirectIndexed(addr, cpu.y)
	cpu.cmp(cpu.a, val)
}

func (cpu *CPU) cpx_imm(operands ...uint8) {
	val := operands[0]
	cpu.cmp(cpu.x, val)
}

func (cpu *CPU) cpx_zp(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	cpu.cmp(cpu.x, val)
}

func (cpu *CPU) cpx_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	cpu.cmp(cpu.x, val)
}

func (cpu *CPU) cpy_imm(operands ...uint8) {
	val := operands[0]
	cpu.cmp(cpu.y, val)
}

func (cpu *CPU) cpy_zp(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	cpu.cmp(cpu.y, val)
}

func (cpu *CPU) cpy_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	cpu.cmp(cpu.y, val)
}
