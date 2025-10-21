package cpu

import (
	"github.com/ghosind/go-nes/logger"
	"github.com/ghosind/go-nes/memory"
)

type CPU struct {
	PC     uint16            // Program Counter
	SP     uint8             // Stack Pointer
	A      uint8             // Accumulator
	X      uint8             // X Register
	Y      uint8             // Y Register
	PS     *ProcessorStatus  // Processor Status
	Mem    *memory.MemoryMap // Memory Bus
	Cycles uint64

	EnableTrace bool
	Logger      logger.Logger
}

func New(mmap *memory.MemoryMap) *CPU {
	cpu := new(CPU)
	cpu.PS = new(ProcessorStatus)
	cpu.Mem = mmap
	return cpu
}

func (cpu *CPU) Reset() {
	cpu.PC = uint16(cpu.Mem.Read(0xFFFD))<<8 | uint16(cpu.Mem.Read(0xFFFC))
	cpu.SP = 0xFD
	cpu.A = 0
	cpu.X = 0
	cpu.Y = 0
	*cpu.PS = psFlagUnused | psFlagInterrupt // Set unused flag to 1 and interrupt disable to 1
	cpu.Cycles = 0
}

func (cpu *CPU) Step() uint64 {
	// record current pc
	pc := cpu.PC

	// Fetch opcode
	opcode := cpu.fetch()

	// Decode
	instruction, exists := instructionMap[opcode]
	if !exists || instruction.execute == nil {
		// Handle unknown opcode (for simplicity, we'll just ignore it here)
		return 0
	}

	// Fetch operands
	operands, crossPageCycles := cpu.fetchOperands(instruction.addressing)

	if cpu.EnableTrace {
		cpu.trace(pc, opcode, instruction, operands...)
	}

	// hardcode cycles to simulate the real cpu cycles
	cycles := instruction.cycles
	if instruction.crossPageCycle && crossPageCycles > 0 {
		cycles += crossPageCycles
	}

	// Execute instruction and get additional execution cycles
	cycles += instruction.execute(cpu, operands...)

	// Update CPU cycles
	cpu.Cycles += cycles

	return cycles
}

func (cpu *CPU) fetch() uint8 {
	value := cpu.Mem.Read(cpu.PC)
	cpu.PC++
	return value
}
