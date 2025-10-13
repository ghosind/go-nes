package cpu

type CPU struct {
	pc  uint16           // Program Counter
	sp  uint8            // Stack Pointer
	a   uint8            // Accumulator
	x   uint8            // X Register
	y   uint8            // Y Register
	ps  *ProcessorStatus // Processor Status
	mem *Memory          // Memory
}

func NewCPU() *CPU {
	cpu := new(CPU)
	cpu.ps = new(ProcessorStatus)
	cpu.mem = new(Memory)
	return cpu
}

func (cpu *CPU) Reset() {
	cpu.pc = uint16(cpu.mem.Read(0xFFFD))<<8 | uint16(cpu.mem.Read(0xFFFC))
	cpu.sp = 0xFD
	cpu.a = 0
	cpu.x = 0
	cpu.y = 0
	*cpu.ps = 0x24 // Set unused flag to 1 and interrupt disable to 1
}

func (cpu *CPU) Step() {
	// Fetch opcode
	opcode := cpu.fetch()

	// Decode & Execute
	if instruction, exists := InstructionMap[opcode]; exists {
		instruction(cpu)
	} else {
		// Handle unknown opcode (for simplicity, we'll just ignore it here)
	}
}

func (cpu *CPU) fetch() uint8 {
	value := cpu.mem.Read(cpu.pc)
	cpu.pc++
	return value
}
