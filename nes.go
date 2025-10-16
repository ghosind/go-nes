package nes

import (
	"github.com/ghosind/go-nes/cpu"
	"github.com/ghosind/go-nes/memory"
)

type NES struct {
	cpu  *cpu.CPU
	mmap *memory.MemoryMap
}

func New(rom []byte) *NES {
	nes := new(NES)

	nes.mmap = memory.NewMemoryMap()
	nes.cpu = cpu.NewCPU(nes.mmap)

	return nes
}

func (n *NES) Step() uint64 {
	cycles := n.cpu.Step()

	for i := uint64(0); i < cycles*3; i++ {
		// PPU step
	}

	return cycles
}
