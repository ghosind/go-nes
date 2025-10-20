package cpu

func (cpu *CPU) brk(operands ...uint8) uint64 {
	cpu.PC++ // BRK has an implied padding byte
	// Push PC to stack
	pcHigh := uint8((cpu.PC) >> 8)
	pcLow := uint8((cpu.PC) & 0xFF)
	cpu.pushStack(pcHigh)
	cpu.pushStack(pcLow)

	// Set Break and Unused flags and push status to stack
	cpu.PS.setBreak(true)
	cpu.PS.setUnused(true)
	cpu.pushStack(uint8(*cpu.PS))

	// Set Interrupt Disable flag
	cpu.PS.setInterrupt(true)

	// Set PC to the address at the IRQ/BRK vector ($FFFE/$FFFF)
	low := cpu.mem.Read(0xFFFE)
	high := cpu.mem.Read(0xFFFF)
	cpu.PC = uint16(high)<<8 | uint16(low)
	return 0
}

func (cpu *CPU) nop(operands ...uint8) uint64 {
	// Do nothing
	return 0
}

func (cpu *CPU) rti(operands ...uint8) uint64 {
	// Pull status from stack
	status := cpu.popStack()
	*cpu.PS = ProcessorStatus(status)

	// Pull PC from stack
	low := cpu.popStack()
	high := cpu.popStack()
	cpu.PC = uint16(high)<<8 | uint16(low)
	return 0
}
