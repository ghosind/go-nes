package nes

import (
	"github.com/ghosind/go-nes/apu"
	"github.com/ghosind/go-nes/cpu"
	"github.com/ghosind/go-nes/logger"
	"github.com/ghosind/go-nes/memory"
	"github.com/ghosind/go-nes/ppu"
	"github.com/ghosind/go-nes/rom"
)

type NES struct {
	apu  *apu.APU
	cpu  *cpu.CPU
	mmap *memory.MemoryMap
	ppu  *ppu.PPU
	rom  *rom.ROM

	enableTrace bool
	logger      logger.Logger
}

func New(data []byte) (*NES, error) {
	nes := new(NES)

	var err error
	nes.rom, err = rom.New(data)
	if err != nil {
		return nil, err
	}
	nes.apu = apu.New()
	nes.ppu = ppu.New(nes.rom)
	nes.mmap = memory.NewMemoryMap(nes.rom, nes.ppu)
	nes.cpu = cpu.New(nes.mmap)
	nes.cpu.Reset()

	return nes, nil
}

func (n *NES) Step() uint64 {
	cycles := n.cpu.Step()

	for i := uint64(0); i < cycles*3; i++ {
		if i%3 == 0 {
			n.apu.Step()
		}

		n.ppu.Step()
	}

	return cycles
}

func (n *NES) EnableTrace(logger logger.Logger) {
	n.enableTrace = true
	n.logger = logger

	n.cpu.EnableTrace = true
	n.cpu.Logger = logger
}
