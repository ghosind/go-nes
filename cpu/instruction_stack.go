package cpu

func (cpu *CPU) tsx(operands ...uint8) {
	cpu.x = cpu.sp
	cpu.ps.setZeroNeg(cpu.x)
}

func (cpu *CPU) txs(operands ...uint8) {
	cpu.sp = cpu.x
	// Note: TXS does not affect any flags
}

func (cpu *CPU) pha(operands ...uint8) {
	cpu.pushStack(cpu.a)
}

func (cpu *CPU) php(operands ...uint8) {
	// When pushing the status register onto the stack, bits 4 and 5 are set to 1
	status := uint8(*cpu.ps) | 0x30
	cpu.pushStack(status)
}

func (cpu *CPU) pla(operands ...uint8) {
	cpu.a = cpu.popStack()
	cpu.ps.setZeroNeg(cpu.a)
}

func (cpu *CPU) plp(operands ...uint8) {
	status := cpu.popStack()
	// Preserve bits 4 and 5 as unused (always set to 1)
	status |= 0x30
	*cpu.ps = ProcessorStatus(status)
}
