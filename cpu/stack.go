package cpu

const StackBaseAddress = 0x0100

func (cpu *CPU) pushStack(value uint8) {
	addr := StackBaseAddress + uint16(cpu.sp)
	cpu.mem.Write(addr, value)
	cpu.sp--
}

func (cpu *CPU) popStack() uint8 {
	cpu.sp++
	addr := StackBaseAddress + uint16(cpu.sp)
	return cpu.mem.Read(addr)
}
