package nes

import (
	"os"
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-nes/cpu"
)

type testLogger struct {
	a *assert.Assertion
}

func (l testLogger) Println(v ...any) {
	l.a.Log(v...)
}

func TestNESByNESTest(t *testing.T) {
	a := assert.New(t)

	rom, err := os.ReadFile("testdata/nestest.nes")
	a.NilNow(err)
	nes, err := New(rom)
	a.NilNow(err)

	nes.cpu.Cycles = 7
	nes.cpu.PC = 0xC000
	*nes.cpu.PS = cpu.ProcessorStatus(0x24) // Set unused and interrupt disable flags

	logger := testLogger{a}

	nes.EnableTrace(logger)
	maxSteps := 100

	for i := 0; i < maxSteps; i++ {
		nes.Step()
	}
}
