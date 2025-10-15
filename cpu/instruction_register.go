package cpu

func (cpu *CPU) tax(operands ...uint8) {
	cpu.x = cpu.a
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) tay(operands ...uint8) {
	cpu.y = cpu.a
	cpu.ps.setZeroNeg(cpu.y)
}

func (cpu *CPU) txa(operands ...uint8) {
	cpu.a = cpu.x
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) tya(operands ...uint8) {
	cpu.a = cpu.y
	cpu.ps.setZeroNeg(cpu.a)
}
