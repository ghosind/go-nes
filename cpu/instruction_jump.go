package cpu

func (cpu *CPU) jmp_abs(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	cpu.PC = uint16(high)<<8 | uint16(low)
}

func (cpu *CPU) jmp_ind(operands ...uint8) {
	low := operands[0]
	high := operands[1]
	addr := uint16(high)<<8 | uint16(low)

	// Emulate page boundary hardware bug
	var targetLow, targetHigh uint8
	if low == 0xFF {
		targetLow = cpu.mem.Read(addr)
		targetHigh = cpu.mem.Read(addr & 0xFF00) // Wrap around to the beginning of the page
	} else {
		targetLow = cpu.mem.Read(addr)
		targetHigh = cpu.mem.Read(addr + 1)
	}

	cpu.PC = uint16(targetHigh)<<8 | uint16(targetLow)
}

func (cpu *CPU) jsr(operands ...uint8) {
	low := operands[0]
	high := operands[1]

	// Push (PC - 1) onto the stack
	returnAddr := cpu.PC - 1
	cpu.pushStack(uint8((returnAddr >> 8) & 0xFF)) // High byte
	cpu.pushStack(uint8(returnAddr & 0xFF))        // Low byte

	// Set PC to target address
	cpu.PC = uint16(high)<<8 | uint16(low)
}

func (cpu *CPU) rts(operands ...uint8) {
	// Pull return address from stack
	low := cpu.popStack()
	high := cpu.popStack()

	cpu.PC = (uint16(high) << 8) | uint16(low)
	cpu.PC++ // Increment PC to point to the next instruction
}
