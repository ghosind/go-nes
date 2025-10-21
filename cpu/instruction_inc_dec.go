package cpu

func (cpu *CPU) inc(value uint8) uint8 {
	result := value + 1
	cpu.PS.setZeroNeg(result)
	return result
}

func (cpu *CPU) inc_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadZeroPage(addr)
	val = cpu.inc(val)
	cpu.Mem.WriteZeroPage(addr, val)
	return 0
}

func (cpu *CPU) inc_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	val := cpu.Mem.ReadZeroPage(addr)
	val = cpu.inc(val)
	cpu.Mem.WriteZeroPage(addr, val)
	return 0
}

func (cpu *CPU) inc_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbs(high, low)
	val = cpu.inc(val)
	cpu.Mem.WriteAbs(high, low, val)
	return 0
}

func (cpu *CPU) inc_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	val = cpu.inc(val)
	cpu.Mem.WriteAbsShift(high, low, cpu.X, val)
	return 0
}

func (cpu *CPU) inx(operands ...uint8) uint64 {
	cpu.X = cpu.inc(cpu.X)
	return 0
}

func (cpu *CPU) iny(operands ...uint8) uint64 {
	cpu.Y = cpu.inc(cpu.Y)
	return 0
}

func (cpu *CPU) dec(value uint8) uint8 {
	result := value - 1
	cpu.PS.setZeroNeg(result)
	return result
}

func (cpu *CPU) dec_zp(operands ...uint8) uint64 {
	addr := operands[0]
	val := cpu.Mem.ReadZeroPage(addr)
	val = cpu.dec(val)
	cpu.Mem.WriteZeroPage(addr, val)
	return 0
}

func (cpu *CPU) dec_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	val := cpu.Mem.ReadZeroPage(addr)
	val = cpu.dec(val)
	cpu.Mem.WriteZeroPage(addr, val)
	return 0
}

func (cpu *CPU) dec_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbs(high, low)
	val = cpu.dec(val)
	cpu.Mem.WriteAbs(high, low, val)
	return 0
}

func (cpu *CPU) dec_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	val := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	val = cpu.dec(val)
	cpu.Mem.WriteAbsShift(high, low, cpu.X, val)
	return 0
}

func (cpu *CPU) dex(operands ...uint8) uint64 {
	cpu.X = cpu.dec(cpu.X)
	return 0
}

func (cpu *CPU) dey(operands ...uint8) uint64 {
	cpu.Y = cpu.dec(cpu.Y)
	return 0
}
