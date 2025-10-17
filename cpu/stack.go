package cpu

const StackBaseAddress = 0x0100

func (cpu *CPU) pushStack(value uint8) {
	addr := StackBaseAddress + uint16(cpu.SP)
	cpu.mem.Write(addr, value)
	cpu.SP--
}

func (cpu *CPU) popStack() uint8 {
	cpu.SP++
	addr := StackBaseAddress + uint16(cpu.SP)
	return cpu.mem.Read(addr)
}
