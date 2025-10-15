package cpu

const StackBaseAddress = 0x0100

func (cpu *CPU) pushStack(value uint8) {
	cpu.mem[StackBaseAddress+uint16(cpu.sp)] = value
	cpu.sp--
}

func (cpu *CPU) popStack() uint8 {
	cpu.sp++
	return cpu.mem[StackBaseAddress+uint16(cpu.sp)]
}
