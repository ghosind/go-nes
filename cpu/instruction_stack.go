package cpu

func (cpu *CPU) tsx(operands ...uint8) {
	cpu.X = cpu.SP
	cpu.PS.setZeroNeg(cpu.X)
}

func (cpu *CPU) txs(operands ...uint8) {
	cpu.SP = cpu.X
	// Note: TXS does not affect any flags
}

func (cpu *CPU) pha(operands ...uint8) {
	cpu.pushStack(cpu.A)
}

func (cpu *CPU) php(operands ...uint8) {
	// When pushing the status register onto the stack, bits 4 and 5 are set to 1
	status := uint8(*cpu.PS) | 0x30
	cpu.pushStack(status)
}

func (cpu *CPU) pla(operands ...uint8) {
	cpu.A = cpu.popStack()
	cpu.PS.setZeroNeg(cpu.A)
}

func (cpu *CPU) plp(operands ...uint8) {
	status := cpu.popStack()
	// Preserve bits 4 and 5 as unused (always set to 1)
	status |= 0x30
	*cpu.PS = ProcessorStatus(status)
}
