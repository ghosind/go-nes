package cpu

func (cpu *CPU) lda_imm(operands ...uint8) uint64 {
	cpu.A = operands[0]
	cpu.PS.setZeroNeg(cpu.A)

	return 0
}

func (cpu *CPU) lda_zp(operands ...uint8) uint64 {
	addr := operands[0]
	cpu.A = cpu.Mem.ReadZeroPage(addr)
	cpu.PS.setZeroNeg(cpu.A)

	return 0
}

func (cpu *CPU) lda_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	cpu.A = cpu.Mem.ReadZeroPage(addr)
	cpu.PS.setZeroNeg(cpu.A)

	return 0
}

func (cpu *CPU) lda_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.A = cpu.Mem.ReadAbs(high, low)
	cpu.PS.setZeroNeg(cpu.A)

	return 0
}

func (cpu *CPU) lda_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.A = cpu.Mem.ReadAbsShift(high, low, cpu.X)
	cpu.PS.setZeroNeg(cpu.A)

	return 0
}

func (cpu *CPU) lda_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.A = cpu.Mem.ReadAbsShift(high, low, cpu.Y)
	cpu.PS.setZeroNeg(cpu.A)

	return 0
}

func (cpu *CPU) lda_ind_x(operands ...uint8) uint64 {
	addr := operands[0]
	cpu.A = cpu.Mem.ReadIndexedIndirect(addr, cpu.X)
	cpu.PS.setZeroNeg(cpu.A)

	return 0
}

func (cpu *CPU) lda_ind_y(operands ...uint8) uint64 {
	addr := operands[0]
	cpu.A = cpu.Mem.ReadIndirectIndexed(addr, cpu.Y)
	cpu.PS.setZeroNeg(cpu.A)

	return 0
}

func (cpu *CPU) ldx_imm(operands ...uint8) uint64 {
	cpu.X = operands[0]
	cpu.PS.setZeroNeg(cpu.X)

	return 0
}

func (cpu *CPU) ldx_zp(operands ...uint8) uint64 {
	addr := operands[0]
	cpu.X = cpu.Mem.ReadZeroPage(addr)
	cpu.PS.setZeroNeg(cpu.X)

	return 0
}

func (cpu *CPU) ldx_zp_y(operands ...uint8) uint64 {
	addr := operands[0] + cpu.Y
	cpu.X = cpu.Mem.ReadZeroPage(addr)
	cpu.PS.setZeroNeg(cpu.X)

	return 0
}

func (cpu *CPU) ldx_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.X = cpu.Mem.ReadAbs(high, low)
	cpu.PS.setZeroNeg(cpu.X)

	return 0
}

func (cpu *CPU) ldx_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.X = cpu.Mem.ReadAbsShift(high, low, cpu.Y)
	cpu.PS.setZeroNeg(cpu.X)

	return 0
}

func (cpu *CPU) ldy_imm(operands ...uint8) uint64 {
	cpu.Y = operands[0]
	cpu.PS.setZeroNeg(cpu.Y)

	return 0
}

func (cpu *CPU) ldy_zp(operands ...uint8) uint64 {
	addr := operands[0]
	cpu.Y = cpu.Mem.ReadZeroPage(addr)
	cpu.PS.setZeroNeg(cpu.Y)

	return 0
}

func (cpu *CPU) ldy_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	cpu.Y = cpu.Mem.ReadZeroPage(addr)
	cpu.PS.setZeroNeg(cpu.Y)

	return 0
}

func (cpu *CPU) ldy_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.Y = cpu.Mem.ReadAbs(high, low)
	cpu.PS.setZeroNeg(cpu.Y)

	return 0
}

func (cpu *CPU) ldy_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.Y = cpu.Mem.ReadAbsShift(high, low, cpu.X)
	cpu.PS.setZeroNeg(cpu.Y)

	return 0
}

func (cpu *CPU) sta_zp(operands ...uint8) uint64 {
	addr := operands[0]
	cpu.Mem.WriteZeroPage(addr, cpu.A)

	return 0
}

func (cpu *CPU) sta_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	cpu.Mem.WriteZeroPage(addr, cpu.A)

	return 0
}

func (cpu *CPU) sta_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.Mem.WriteAbs(high, low, cpu.A)

	return 0
}

func (cpu *CPU) sta_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.Mem.WriteAbsShift(high, low, cpu.X, cpu.A)

	return 0
}

func (cpu *CPU) sta_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.Mem.WriteAbsShift(high, low, cpu.Y, cpu.A)

	return 0
}

func (cpu *CPU) sta_ind_x(operands ...uint8) uint64 {
	addr := operands[0]
	cpu.Mem.WriteIndexedIndirect(addr, cpu.X, cpu.A)

	return 0
}

func (cpu *CPU) sta_ind_y(operands ...uint8) uint64 {
	addr := operands[0]
	cpu.Mem.WriteIndirectIndexed(addr, cpu.Y, cpu.A)

	return 0
}

func (cpu *CPU) stx_zp(operands ...uint8) uint64 {
	addr := operands[0]
	cpu.Mem.WriteZeroPage(addr, cpu.X)

	return 0
}

func (cpu *CPU) stx_zp_y(operands ...uint8) uint64 {
	addr := operands[0] + cpu.Y
	cpu.Mem.WriteZeroPage(addr, cpu.X)

	return 0
}

func (cpu *CPU) stx_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.Mem.WriteAbs(high, low, cpu.X)

	return 0
}

func (cpu *CPU) sty_zp(operands ...uint8) uint64 {
	addr := operands[0]
	cpu.Mem.WriteZeroPage(addr, cpu.Y)

	return 0
}

func (cpu *CPU) sty_zp_x(operands ...uint8) uint64 {
	addr := operands[0] + cpu.X
	cpu.Mem.WriteZeroPage(addr, cpu.Y)

	return 0
}

func (cpu *CPU) sty_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	cpu.Mem.WriteAbs(high, low, cpu.Y)

	return 0
}
