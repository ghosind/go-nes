package cpu

func (cpu *CPU) lax(value uint8) uint64 {
	cpu.A = value
	cpu.X = value
	cpu.PS.setZeroNeg(value)

	return 0
}

func (cpu *CPU) lax_zp(operands ...uint8) uint64 {
	addr := operands[0]
	value := cpu.Mem.ReadZeroPage(addr)
	return cpu.lax(value)
}

func (cpu *CPU) lax_zp_y(operands ...uint8) uint64 {
	addr := operands[0] + cpu.Y
	value := cpu.Mem.ReadZeroPage(addr)
	return cpu.lax(value)
}

func (cpu *CPU) lax_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	value := cpu.Mem.ReadAbs(high, low)
	return cpu.lax(value)
}

func (cpu *CPU) lax_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	value := cpu.Mem.ReadAbsShift(high, low, cpu.Y)
	return cpu.lax(value)
}

func (cpu *CPU) lax_ind_x(operands ...uint8) uint64 {
	addr := operands[0]
	value := cpu.Mem.ReadIndexedIndirect(addr, cpu.X)
	return cpu.lax(value)
}

func (cpu *CPU) lax_ind_y(operands ...uint8) uint64 {
	addr := operands[0]
	value := cpu.Mem.ReadIndirectIndexed(addr, cpu.Y)
	return cpu.lax(value)
}

func (cpu *CPU) sax_zp(operands ...uint8) uint64 {
	addr := operands[0]
	value := cpu.A & cpu.X
	cpu.Mem.WriteZeroPage(addr, value)

	return 0
}

func (cpu *CPU) sax_zp_y(operands ...uint8) uint64 {
	addr := operands[0] + cpu.Y
	value := cpu.A & cpu.X
	cpu.Mem.WriteZeroPage(addr, value)

	return 0
}

func (cpu *CPU) sax_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	value := cpu.A & cpu.X
	cpu.Mem.WriteAbs(high, low, value)

	return 0
}

func (cpu *CPU) sax_ind_x(operands ...uint8) uint64 {
	addr := operands[0]
	value := cpu.A & cpu.X
	cpu.Mem.WriteIndexedIndirect(addr, cpu.X, value)

	return 0
}

func (cpu *CPU) dcp(addr uint16) uint64 {
	value := cpu.Mem.Read(addr)
	value = value - 1
	cpu.Mem.Write(addr, value)

	cpu.cmp(cpu.A, value)

	return 0
}

func (cpu *CPU) dcp_zp(operands ...uint8) uint64 {
	addr := uint16(operands[0])
	return cpu.dcp(addr)
}

func (cpu *CPU) dcp_zp_x(operands ...uint8) uint64 {
	addr := uint16(operands[0] + cpu.X)
	return cpu.dcp(addr)
}

func (cpu *CPU) dcp_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.dcp(addr)
}

func (cpu *CPU) dcp_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.X)
	return cpu.dcp(addr)
}

func (cpu *CPU) dcp_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.dcp(addr)
}

func (cpu *CPU) dcp_ind_x(operands ...uint8) uint64 {
	addr8 := operands[0]
	ptr := uint16(addr8 + cpu.X)
	low := cpu.Mem.ReadZeroPage(uint8(ptr))
	high := cpu.Mem.ReadZeroPage(uint8(ptr + 1))
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.dcp(addr)
}

func (cpu *CPU) dcp_ind_y(operands ...uint8) uint64 {
	addr8 := operands[0]
	low := cpu.Mem.ReadZeroPage(addr8)
	high := cpu.Mem.ReadZeroPage(addr8 + 1)
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.dcp(addr)
}

func (cpu *CPU) isb(addr uint16) uint64 {
	value := cpu.Mem.Read(addr)
	value = value + 1
	cpu.Mem.Write(addr, value)

	return cpu.sbc(value)
}

func (cpu *CPU) isb_zp(operands ...uint8) uint64 {
	addr := uint16(operands[0])
	return cpu.isb(addr)
}

func (cpu *CPU) isb_zp_x(operands ...uint8) uint64 {
	addr := uint16(operands[0] + cpu.X)
	return cpu.isb(addr)
}

func (cpu *CPU) isb_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.isb(addr)
}

func (cpu *CPU) isb_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.X)
	return cpu.isb(addr)
}

func (cpu *CPU) isb_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.isb(addr)
}

func (cpu *CPU) isb_ind_x(operands ...uint8) uint64 {
	addr8 := operands[0]
	ptr := uint16(addr8 + cpu.X)
	low := cpu.Mem.ReadZeroPage(uint8(ptr))
	high := cpu.Mem.ReadZeroPage(uint8(ptr + 1))
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.isb(addr)
}

func (cpu *CPU) isb_ind_y(operands ...uint8) uint64 {
	addr8 := operands[0]
	low := cpu.Mem.ReadZeroPage(addr8)
	high := cpu.Mem.ReadZeroPage(addr8 + 1)
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.isb(addr)
}

func (cpu *CPU) slo(addr uint16) uint64 {
	value := cpu.Mem.Read(addr)
	value = cpu.asl(value)
	cpu.Mem.Write(addr, value)

	return cpu.ora(value)
}

func (cpu *CPU) slo_zp(operands ...uint8) uint64 {
	addr := uint16(operands[0])
	return cpu.slo(addr)
}

func (cpu *CPU) slo_zp_x(operands ...uint8) uint64 {
	addr := uint16(operands[0] + cpu.X)
	return cpu.slo(addr)
}

func (cpu *CPU) slo_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.slo(addr)
}

func (cpu *CPU) slo_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.X)
	return cpu.slo(addr)
}

func (cpu *CPU) slo_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.slo(addr)
}

func (cpu *CPU) slo_ind_x(operands ...uint8) uint64 {
	addr8 := operands[0]
	ptr := uint16(addr8 + cpu.X)
	low := cpu.Mem.ReadZeroPage(uint8(ptr))
	high := cpu.Mem.ReadZeroPage(uint8(ptr + 1))
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.slo(addr)
}

func (cpu *CPU) slo_ind_y(operands ...uint8) uint64 {
	addr8 := operands[0]
	low := cpu.Mem.ReadZeroPage(addr8)
	high := cpu.Mem.ReadZeroPage(addr8 + 1)
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.slo(addr)
}

func (cpu *CPU) rla(addr uint16) uint64 {
	value := cpu.Mem.Read(addr)
	value = cpu.rol(value)
	cpu.Mem.Write(addr, value)

	return cpu.and(value)
}

func (cpu *CPU) rla_zp(operands ...uint8) uint64 {
	addr := uint16(operands[0])
	return cpu.rla(addr)
}

func (cpu *CPU) rla_zp_x(operands ...uint8) uint64 {
	addr := uint16(operands[0] + cpu.X)
	return cpu.rla(addr)
}

func (cpu *CPU) rla_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.rla(addr)
}

func (cpu *CPU) rla_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.X)
	return cpu.rla(addr)
}

func (cpu *CPU) rla_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.rla(addr)
}

func (cpu *CPU) rla_ind_x(operands ...uint8) uint64 {
	addr8 := operands[0]
	ptr := uint16(addr8 + cpu.X)
	low := cpu.Mem.ReadZeroPage(uint8(ptr))
	high := cpu.Mem.ReadZeroPage(uint8(ptr + 1))
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.rla(addr)
}

func (cpu *CPU) rla_ind_y(operands ...uint8) uint64 {
	addr8 := operands[0]
	low := cpu.Mem.ReadZeroPage(addr8)
	high := cpu.Mem.ReadZeroPage(addr8 + 1)
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.rla(addr)
}

func (cpu *CPU) sre(addr uint16) uint64 {
	value := cpu.Mem.Read(addr)
	value = cpu.lsr(value)
	cpu.Mem.Write(addr, value)

	return cpu.eor(value)
}

func (cpu *CPU) sre_zp(operands ...uint8) uint64 {
	addr := uint16(operands[0])
	return cpu.sre(addr)
}

func (cpu *CPU) sre_zp_x(operands ...uint8) uint64 {
	addr := uint16(operands[0] + cpu.X)
	return cpu.sre(addr)
}
func (cpu *CPU) sre_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.sre(addr)
}

func (cpu *CPU) sre_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.X)
	return cpu.sre(addr)
}

func (cpu *CPU) sre_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.sre(addr)
}

func (cpu *CPU) sre_ind_x(operands ...uint8) uint64 {
	addr8 := operands[0]
	ptr := uint16(addr8 + cpu.X)
	low := cpu.Mem.ReadZeroPage(uint8(ptr))
	high := cpu.Mem.ReadZeroPage(uint8(ptr + 1))
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.sre(addr)
}

func (cpu *CPU) sre_ind_y(operands ...uint8) uint64 {
	addr8 := operands[0]
	low := cpu.Mem.ReadZeroPage(addr8)
	high := cpu.Mem.ReadZeroPage(addr8 + 1)
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.sre(addr)
}

func (cpu *CPU) rra(addr uint16) uint64 {
	value := cpu.Mem.Read(addr)
	value = cpu.ror(value)
	cpu.Mem.Write(addr, value)

	return cpu.adc(value)
}

func (cpu *CPU) rra_zp(operands ...uint8) uint64 {
	addr := uint16(operands[0])
	return cpu.rra(addr)
}

func (cpu *CPU) rra_zp_x(operands ...uint8) uint64 {
	addr := uint16(operands[0] + cpu.X)
	return cpu.rra(addr)
}

func (cpu *CPU) rra_abs(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.rra(addr)
}

func (cpu *CPU) rra_abs_x(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.X)
	return cpu.rra(addr)
}

func (cpu *CPU) rra_abs_y(operands ...uint8) uint64 {
	low := operands[0]
	high := operands[1]
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.rra(addr)
}

func (cpu *CPU) rra_ind_x(operands ...uint8) uint64 {
	addr8 := operands[0]
	ptr := uint16(addr8 + cpu.X)
	low := cpu.Mem.ReadZeroPage(uint8(ptr))
	high := cpu.Mem.ReadZeroPage(uint8(ptr + 1))
	addr := (uint16(high) << 8) | uint16(low)
	return cpu.rra(addr)
}

func (cpu *CPU) rra_ind_y(operands ...uint8) uint64 {
	addr8 := operands[0]
	low := cpu.Mem.ReadZeroPage(addr8)
	high := cpu.Mem.ReadZeroPage(addr8 + 1)
	addr := (uint16(high) << 8) | uint16(low) + uint16(cpu.Y)
	return cpu.rra(addr)
}
