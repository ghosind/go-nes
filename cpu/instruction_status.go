package cpu

func (cpu *CPU) clc(operands ...uint8) uint64 {
	cpu.PS.setCarry(false)
	return 0
}

func (cpu *CPU) cld(operands ...uint8) uint64 {
	cpu.PS.setDecimal(false)
	return 0
}

func (cpu *CPU) cli(operands ...uint8) uint64 {
	cpu.PS.setInterrupt(false)
	return 0
}

func (cpu *CPU) clv(operands ...uint8) uint64 {
	cpu.PS.setOverflow(false)
	return 0
}

func (cpu *CPU) sec(operands ...uint8) uint64 {
	cpu.PS.setCarry(true)
	return 0
}

func (cpu *CPU) sed(operands ...uint8) uint64 {
	cpu.PS.setDecimal(true)
	return 0
}

func (cpu *CPU) sei(operands ...uint8) uint64 {
	cpu.PS.setInterrupt(true)
	return 0
}
