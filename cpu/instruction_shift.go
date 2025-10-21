package cpu

func (cpu *CPU) asl(value uint8) uint8 {
	result := value << 1
	cpu.PS.setCarry(value&0x80 != 0)
	cpu.PS.setZeroNeg(result)
	return result
}

func (cpu *CPU) asl_acc(operands ...uint8) uint64 {
	cpu.A = cpu.asl(cpu.A)
	return 0
}

func (cpu *CPU) asl_zp(operands ...uint8) uint64 {
	addr := operands[0]
	value := cpu.Mem.ReadZeroPage(addr)
	result := cpu.asl(value)
	cpu.Mem.WriteZeroPage(addr, result)
	return 0
}

func (cpu *CPU) asl_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	value := cpu.Mem.ReadZeroPage(addr)
	result := cpu.asl(value)
	cpu.Mem.WriteZeroPage(addr, result)
	return 0
}

func (cpu *CPU) asl_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	value := cpu.Mem.ReadAbs(high, low)
	result := cpu.asl(value)
	cpu.Mem.WriteAbs(high, low, result)
	return 0
}

func (cpu *CPU) asl_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	value := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	result := cpu.asl(value)
	cpu.Mem.WriteAbsShift(high, low, cpu.X, result)
	return 0
}

func (cpu *CPU) lsr(value uint8) uint8 {
	result := value >> 1
	cpu.PS.setCarry(value&0x01 != 0)
	cpu.PS.setZeroNeg(result)
	return result
}

func (cpu *CPU) lsr_acc(operands ...uint8) uint64 {
	cpu.A = cpu.lsr(cpu.A)
	return 0
}

func (cpu *CPU) lsr_zp(operands ...uint8) uint64 {
	addr := operands[0]
	value := cpu.Mem.ReadZeroPage(addr)
	result := cpu.lsr(value)
	cpu.Mem.WriteZeroPage(addr, result)
	return 0
}

func (cpu *CPU) lsr_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	value := cpu.Mem.ReadZeroPage(addr)
	result := cpu.lsr(value)
	cpu.Mem.WriteZeroPage(addr, result)
	return 0
}

func (cpu *CPU) lsr_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	value := cpu.Mem.ReadAbs(high, low)
	result := cpu.lsr(value)
	cpu.Mem.WriteAbs(high, low, result)
	return 0
}

func (cpu *CPU) lsr_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	value := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	result := cpu.lsr(value)
	cpu.Mem.WriteAbsShift(high, low, cpu.X, result)
	return 0
}

func (cpu *CPU) rol(value uint8) uint8 {
	carry := uint8(0)
	if cpu.PS.getCarry() {
		carry = 1
	}
	result := (value << 1) | carry
	cpu.PS.setCarry(value&0x80 != 0)
	cpu.PS.setZeroNeg(result)
	return result
}

func (cpu *CPU) rol_acc(operands ...uint8) uint64 {
	cpu.A = cpu.rol(cpu.A)
	return 0
}

func (cpu *CPU) rol_zp(operands ...uint8) uint64 {
	addr := operands[0]
	value := cpu.Mem.ReadZeroPage(addr)
	result := cpu.rol(value)
	cpu.Mem.WriteZeroPage(addr, result)
	return 0
}

func (cpu *CPU) rol_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	value := cpu.Mem.ReadZeroPage(addr)
	result := cpu.rol(value)
	cpu.Mem.WriteZeroPage(addr, result)
	return 0
}

func (cpu *CPU) rol_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	value := cpu.Mem.ReadAbs(high, low)
	result := cpu.rol(value)
	cpu.Mem.WriteAbs(high, low, result)
	return 0
}

func (cpu *CPU) rol_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	value := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	result := cpu.rol(value)
	cpu.Mem.WriteAbsShift(high, low, cpu.X, result)
	return 0
}

func (cpu *CPU) ror(value uint8) uint8 {
	carry := uint8(0)
	if cpu.PS.getCarry() {
		carry = 0x80
	}
	result := (value >> 1) | carry
	cpu.PS.setCarry(value&0x01 != 0)
	cpu.PS.setZeroNeg(result)
	return result
}

func (cpu *CPU) ror_acc(operands ...uint8) uint64 {
	cpu.A = cpu.ror(cpu.A)
	return 0
}

func (cpu *CPU) ror_zp(operands ...uint8) uint64 {
	addr := operands[0]
	value := cpu.Mem.ReadZeroPage(addr)
	result := cpu.ror(value)
	cpu.Mem.WriteZeroPage(addr, result)
	return 0
}

func (cpu *CPU) ror_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	value := cpu.Mem.ReadZeroPage(addr)
	result := cpu.ror(value)
	cpu.Mem.WriteZeroPage(addr, result)
	return 0
}

func (cpu *CPU) ror_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	value := cpu.Mem.ReadAbs(high, low)
	result := cpu.ror(value)
	cpu.Mem.WriteAbs(high, low, result)
	return 0
}

func (cpu *CPU) ror_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	value := cpu.Mem.ReadAbsShift(high, low, cpu.X)
	result := cpu.ror(value)
	cpu.Mem.WriteAbsShift(high, low, cpu.X, result)
	return 0
}
