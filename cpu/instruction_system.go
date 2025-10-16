package cpu

func (cpu *CPU) brk(operands ...uint8) {
	cpu.pc++ // BRK has an implied padding byte
	// Push PC to stack
	pcHigh := uint8((cpu.pc) >> 8)
	pcLow := uint8((cpu.pc) & 0xFF)
	cpu.pushStack(pcHigh)
	cpu.pushStack(pcLow)

	// Set Break and Unused flags and push status to stack
	cpu.ps.setBreak(true)
	cpu.ps.setUnused(true)
	cpu.pushStack(uint8(*cpu.ps))

	// Set Interrupt Disable flag
	cpu.ps.setInterrupt(true)

	// Set PC to the address at the IRQ/BRK vector ($FFFE/$FFFF)
	low := cpu.mem.Read(0xFFFE)
	high := cpu.mem.Read(0xFFFF)
	cpu.pc = uint16(high)<<8 | uint16(low)
}

func (cpu *CPU) nop(operands ...uint8) {
	// Do nothing
}

func (cpu *CPU) rti(operands ...uint8) {
	// Pull status from stack
	status := cpu.popStack()
	*cpu.ps = ProcessorStatus(status)

	// Pull PC from stack
	low := cpu.popStack()
	high := cpu.popStack()
	cpu.pc = uint16(high)<<8 | uint16(low)
}
