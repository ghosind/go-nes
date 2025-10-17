package cpu

func (cpu *CPU) asl(value uint8) uint8 {
	result := value << 1
	cpu.PS.setCarry(value&0x80 != 0)
	cpu.PS.setZeroNeg(result)
	return result
}

func (cpu *CPU) asl_acc(operands ...uint8) {
	cpu.A = cpu.asl(cpu.A)
}

func (cpu *CPU) asl_zp(operands ...uint8) {
	addr := operands[0]
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.asl(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) asl_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.X
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.asl(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) asl_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	value := cpu.mem.ReadAbs(high, low)
	result := cpu.asl(value)
	cpu.mem.WriteAbs(high, low, result)
}

func (cpu *CPU) asl_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.X)
	result := cpu.asl(value)
	cpu.mem.WriteAbsShift(high, low, cpu.X, result)
}

func (cpu *CPU) lsr(value uint8) uint8 {
	result := value >> 1
	cpu.PS.setCarry(value&0x01 != 0)
	cpu.PS.setZeroNeg(result)
	return result
}

func (cpu *CPU) lsr_acc(operands ...uint8) {
	cpu.A = cpu.lsr(cpu.A)
}

func (cpu *CPU) lsr_zp(operands ...uint8) {
	addr := operands[0]
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.lsr(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) lsr_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.X
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.lsr(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) lsr_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	value := cpu.mem.ReadAbs(high, low)
	result := cpu.lsr(value)
	cpu.mem.WriteAbs(high, low, result)
}

func (cpu *CPU) lsr_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.X)
	result := cpu.lsr(value)
	cpu.mem.WriteAbsShift(high, low, cpu.X, result)
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

func (cpu *CPU) rol_acc(operands ...uint8) {
	cpu.A = cpu.rol(cpu.A)
}

func (cpu *CPU) rol_zp(operands ...uint8) {
	addr := operands[0]
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.rol(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) rol_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.X
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.rol(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) rol_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	value := cpu.mem.ReadAbs(high, low)
	result := cpu.rol(value)
	cpu.mem.WriteAbs(high, low, result)
}

func (cpu *CPU) rol_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.X)
	result := cpu.rol(value)
	cpu.mem.WriteAbsShift(high, low, cpu.X, result)
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

func (cpu *CPU) ror_acc(operands ...uint8) {
	cpu.A = cpu.ror(cpu.A)
}

func (cpu *CPU) ror_zp(operands ...uint8) {
	addr := operands[0]
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.ror(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) ror_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.X
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.ror(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) ror_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	value := cpu.mem.ReadAbs(high, low)
	result := cpu.ror(value)
	cpu.mem.WriteAbs(high, low, result)
}

func (cpu *CPU) ror_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	value := cpu.mem.ReadAbsShift(high, low, cpu.X)
	result := cpu.ror(value)
	cpu.mem.WriteAbsShift(high, low, cpu.X, result)
}
