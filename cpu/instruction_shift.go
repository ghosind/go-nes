package cpu

func (cpu *CPU) asl(value uint8) uint8 {
	result := value << 1
	cpu.ps.setCarry(value&0x80 != 0)
	cpu.ps.setZeroNeg(result)
	return result
}

func (cpu *CPU) asl_acc(operands ...uint8) {
	cpu.a = cpu.asl(cpu.a)
}

func (cpu *CPU) asl_zp(operands ...uint8) {
	addr := operands[0]
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.asl(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) asl_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
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
	value := cpu.mem.ReadAbsShift(high, low, cpu.x)
	result := cpu.asl(value)
	cpu.mem.WriteAbsShift(high, low, cpu.x, result)
}

func (cpu *CPU) lsr(value uint8) uint8 {
	result := value >> 1
	cpu.ps.setCarry(value&0x01 != 0)
	cpu.ps.setZeroNeg(result)
	return result
}

func (cpu *CPU) lsr_acc(operands ...uint8) {
	cpu.a = cpu.lsr(cpu.a)
}

func (cpu *CPU) lsr_zp(operands ...uint8) {
	addr := operands[0]
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.lsr(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) lsr_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
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
	value := cpu.mem.ReadAbsShift(high, low, cpu.x)
	result := cpu.lsr(value)
	cpu.mem.WriteAbsShift(high, low, cpu.x, result)
}

func (cpu *CPU) rol(value uint8) uint8 {
	carry := uint8(0)
	if cpu.ps.getCarry() {
		carry = 1
	}
	result := (value << 1) | carry
	cpu.ps.setCarry(value&0x80 != 0)
	cpu.ps.setZeroNeg(result)
	return result
}

func (cpu *CPU) rol_acc(operands ...uint8) {
	cpu.a = cpu.rol(cpu.a)
}

func (cpu *CPU) rol_zp(operands ...uint8) {
	addr := operands[0]
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.rol(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) rol_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
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
	value := cpu.mem.ReadAbsShift(high, low, cpu.x)
	result := cpu.rol(value)
	cpu.mem.WriteAbsShift(high, low, cpu.x, result)
}

func (cpu *CPU) ror(value uint8) uint8 {
	carry := uint8(0)
	if cpu.ps.getCarry() {
		carry = 0x80
	}
	result := (value >> 1) | carry
	cpu.ps.setCarry(value&0x01 != 0)
	cpu.ps.setZeroNeg(result)
	return result
}

func (cpu *CPU) ror_acc(operands ...uint8) {
	cpu.a = cpu.ror(cpu.a)
}

func (cpu *CPU) ror_zp(operands ...uint8) {
	addr := operands[0]
	value := cpu.mem.ReadZeroPage(addr)
	result := cpu.ror(value)
	cpu.mem.WriteZeroPage(addr, result)
}

func (cpu *CPU) ror_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
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
	value := cpu.mem.ReadAbsShift(high, low, cpu.x)
	result := cpu.ror(value)
	cpu.mem.WriteAbsShift(high, low, cpu.x, result)
}
