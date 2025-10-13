package cpu

func (cpu *CPU) lda_imm(operands ...uint8) {
	cpu.a = operands[0]
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_zp(operands ...uint8) {
	addr := operands[0]
	cpu.a = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
	cpu.a = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.a = cpu.mem.ReadAbs(high, low)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.a = cpu.mem.ReadAbsShift(high, low, cpu.x)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_abs_y(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.a = cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_ind_x(operands ...uint8) {
	addr := operands[0] + cpu.x
	low := cpu.mem.ReadZeroPage(addr)
	high := cpu.mem.ReadZeroPage(addr + 1)
	cpu.a = cpu.mem.ReadAbs(high, low)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) lda_ind_y(operands ...uint8) {
	addr := operands[0]
	low := cpu.mem.ReadZeroPage(addr)
	high := cpu.mem.ReadZeroPage(addr + 1)
	cpu.a = cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) ldx_imm(operands ...uint8) {
	cpu.x = operands[0]
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) ldx_zp(operands ...uint8) {
	addr := operands[0]
	cpu.x = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) ldx_zp_y(operands ...uint8) {
	addr := operands[0] + cpu.y
	cpu.x = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) ldx_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.x = cpu.mem.ReadAbs(high, low)
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) ldx_abs_y(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.x = cpu.mem.ReadAbsShift(high, low, cpu.y)
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) ldy_imm(operands ...uint8) {
	cpu.y = operands[0]
	cpu.ps.setZeroNeg(cpu.y)
}

func (cpu *CPU) ldy_zp(operands ...uint8) {
	addr := operands[0]
	cpu.y = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.y)
}

func (cpu *CPU) ldy_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
	cpu.y = cpu.mem.ReadZeroPage(addr)
	cpu.ps.setZeroNeg(cpu.y)
}

func (cpu *CPU) ldy_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.y = cpu.mem.ReadAbs(high, low)
	cpu.ps.setZeroNeg(cpu.y)
}

func (cpu *CPU) ldy_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.y = cpu.mem.ReadAbsShift(high, low, cpu.x)
	cpu.ps.setZeroNeg(cpu.y)
}

func (cpu *CPU) sta_zp(operands ...uint8) {
	addr := operands[0]
	cpu.mem.WriteZeroPage(addr, cpu.a)
}

func (cpu *CPU) sta_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
	cpu.mem.WriteZeroPage(addr, cpu.a)
}

func (cpu *CPU) sta_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.mem.WriteAbs(high, low, cpu.a)
}

func (cpu *CPU) sta_abs_x(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.mem.WriteAbsShift(high, low, cpu.x, cpu.a)
}

func (cpu *CPU) sta_abs_y(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.mem.WriteAbsShift(high, low, cpu.y, cpu.a)
}

func (cpu *CPU) sta_ind_x(operands ...uint8) {
	addr := operands[0] + cpu.x
	low := cpu.mem.ReadZeroPage(addr)
	high := cpu.mem.ReadZeroPage(addr + 1)
	cpu.mem.WriteAbs(high, low, cpu.a)
}

func (cpu *CPU) sta_ind_y(operands ...uint8) {
	addr := operands[0]
	low := cpu.mem.ReadZeroPage(addr)
	high := cpu.mem.ReadZeroPage(addr + 1)
	cpu.mem.WriteAbsShift(high, low, cpu.y, cpu.a)
}

func (cpu *CPU) stx_zp(operands ...uint8) {
	addr := operands[0]
	cpu.mem.WriteZeroPage(addr, cpu.x)
}

func (cpu *CPU) stx_zp_y(operands ...uint8) {
	addr := operands[0] + cpu.y
	cpu.mem.WriteZeroPage(addr, cpu.x)
}

func (cpu *CPU) stx_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.mem.WriteAbs(high, low, cpu.x)
}

func (cpu *CPU) sty_zp(operands ...uint8) {
	addr := operands[0]
	cpu.mem.WriteZeroPage(addr, cpu.y)
}

func (cpu *CPU) sty_zp_x(operands ...uint8) {
	addr := operands[0] + cpu.x
	cpu.mem.WriteZeroPage(addr, cpu.y)
}

func (cpu *CPU) sty_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.mem.WriteAbs(high, low, cpu.y)
}
