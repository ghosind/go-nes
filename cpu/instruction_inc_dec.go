package cpu

func (cpu *CPU) inc(value uint8) uint8 {
	result := value + 1
	cpu.ps.setZeroNeg(result)
	return result
}

func (cpu *CPU) inc_zp(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	val = cpu.inc(val)
	cpu.mem.WriteZeroPage(addr, val)
}

func (cpu *CPU) inc_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
	val := cpu.mem.ReadZeroPage(addr)
	val = cpu.inc(val)
	cpu.mem.WriteZeroPage(addr, val)
}

func (cpu *CPU) inc_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	val = cpu.inc(val)
	cpu.mem.WriteAbs(high, low, val)
}

func (cpu *CPU) inc_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.x)
	val = cpu.inc(val)
	cpu.mem.WriteAbsShift(high, low, cpu.x, val)
}

func (cpu *CPU) inx(operands ...uint8) {
	cpu.x = cpu.inc(cpu.x)
}

func (cpu *CPU) iny(operands ...uint8) {
	cpu.y = cpu.inc(cpu.y)
}

func (cpu *CPU) dec(value uint8) uint8 {
	result := value - 1
	cpu.ps.setZeroNeg(result)
	return result
}

func (cpu *CPU) dec_zp(operands ...uint8) {
	addr := operands[0]
	val := cpu.mem.ReadZeroPage(addr)
	val = cpu.dec(val)
	cpu.mem.WriteZeroPage(addr, val)
}

func (cpu *CPU) dec_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
	val := cpu.mem.ReadZeroPage(addr)
	val = cpu.dec(val)
	cpu.mem.WriteZeroPage(addr, val)
}

func (cpu *CPU) dec_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbs(high, low)
	val = cpu.dec(val)
	cpu.mem.WriteAbs(high, low, val)
}

func (cpu *CPU) dec_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	val := cpu.mem.ReadAbsShift(high, low, cpu.x)
	val = cpu.dec(val)
	cpu.mem.WriteAbsShift(high, low, cpu.x, val)
}

func (cpu *CPU) dex(operands ...uint8) {
	cpu.x = cpu.dec(cpu.x)
}

func (cpu *CPU) dey(operands ...uint8) {
	cpu.y = cpu.dec(cpu.y)
}
