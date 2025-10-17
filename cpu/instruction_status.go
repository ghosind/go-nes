package cpu

func (cpu *CPU) clc(operands ...uint8) {
	cpu.PS.setCarry(false)
}

func (cpu *CPU) cld(operands ...uint8) {
	cpu.PS.setDecimal(false)
}

func (cpu *CPU) cli(operands ...uint8) {
	cpu.PS.setInterrupt(false)
}

func (cpu *CPU) clv(operands ...uint8) {
	cpu.PS.setOverflow(false)
}

func (cpu *CPU) sec(operands ...uint8) {
	cpu.PS.setCarry(true)
}

func (cpu *CPU) sed(operands ...uint8) {
	cpu.PS.setDecimal(true)
}

func (cpu *CPU) sei(operands ...uint8) {
	cpu.PS.setInterrupt(true)
}
