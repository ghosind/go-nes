package cpu

func (cpu *CPU) tax(operands ...uint8) {
	cpu.X = cpu.A
	cpu.PS.setZeroNeg(cpu.X)
}

func (cpu *CPU) tay(operands ...uint8) {
	cpu.Y = cpu.A
	cpu.PS.setZeroNeg(cpu.Y)
}

func (cpu *CPU) txa(operands ...uint8) {
	cpu.A = cpu.X
	cpu.PS.setZeroNeg(cpu.A)
}

func (cpu *CPU) tya(operands ...uint8) {
	cpu.A = cpu.Y
	cpu.PS.setZeroNeg(cpu.A)
}
