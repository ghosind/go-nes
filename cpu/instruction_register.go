package cpu

func (cpu *CPU) tax(operands ...uint8) uint64 {
	cpu.X = cpu.A
	cpu.PS.setZeroNeg(cpu.X)

	return 0
}

func (cpu *CPU) tay(operands ...uint8) uint64 {
	cpu.Y = cpu.A
	cpu.PS.setZeroNeg(cpu.Y)

	return 0
}

func (cpu *CPU) txa(operands ...uint8) uint64 {
	cpu.A = cpu.X
	cpu.PS.setZeroNeg(cpu.A)

	return 0
}

func (cpu *CPU) tya(operands ...uint8) uint64 {
	cpu.A = cpu.Y
	cpu.PS.setZeroNeg(cpu.A)

	return 0
}
