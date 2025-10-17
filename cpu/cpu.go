package cpu

import "github.com/ghosind/go-nes/memory"

type CPU struct {
	pc  uint16            // Program Counter
	sp  uint8             // Stack Pointer
	a   uint8             // Accumulator
	x   uint8             // X Register
	y   uint8             // Y Register
	ps  *ProcessorStatus  // Processor Status
	mem *memory.MemoryMap // Memory Bus
}

func New(bus *memory.MemoryMap) *CPU {
	cpu := new(CPU)
	cpu.ps = new(ProcessorStatus)
	cpu.mem = bus
	return cpu
}

func (cpu *CPU) Reset() {
	cpu.pc = uint16(cpu.mem.Read(0xFFFD))<<8 | uint16(cpu.mem.Read(0xFFFC))
	cpu.sp = 0xFD
	cpu.a = 0
	cpu.x = 0
	cpu.y = 0
	*cpu.ps = psFlagUnused | psFlagInterrupt // Set unused flag to 1 and interrupt disable to 1
}

func (cpu *CPU) Step() uint64 {
	// Fetch opcode
	opcode := cpu.fetch()

	// Decode
	instruction, exists := instructionMap[opcode]
	if !exists || instruction.execute == nil {
		// Handle unknown opcode (for simplicity, we'll just ignore it here)
		return 0
	}

	// Fetch operands
	operands, addCycles := cpu.fetchOperands(instruction.addressing)
	// hardcode cycles to simulate the real cpu cycles
	cycles := instruction.cycles + addCycles

	// Execute instruction
	instruction.execute(cpu, operands...)

	return cycles
}

func (cpu *CPU) fetch() uint8 {
	value := cpu.mem.Read(cpu.pc)
	cpu.pc++
	return value
}
