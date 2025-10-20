package nes

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ghosind/go-assert"
	"github.com/ghosind/go-nes/cpu"
)

type testLogger struct {
	buf   *bytes.Buffer
	a     *assert.Assertion
	debug bool
}

func (l testLogger) Printf(format string, v ...any) {
	fmt.Fprintf(l.buf, format, v...)
}

func TestNESByNESTest(t *testing.T) {
	a := assert.New(t)

	rom, err := os.ReadFile("testdata/nestest.nes")
	a.NilNow(err)
	nes, err := New(rom)
	a.NilNow(err)

	logs, err := os.ReadFile("testdata/nestest.log")
	a.NilNow(err)

	lines := strings.Split(string(logs), "\n")

	nes.cpu.Cycles = 7
	nes.cpu.PC = 0xC000
	*nes.cpu.PS = cpu.ProcessorStatus(0x24) // Set unused and interrupt disable flags

	logger := &testLogger{
		buf: bytes.NewBuffer(make([]byte, 0, 91)),
		a:   a,
	}

	nes.EnableTrace(logger)

	for i, expectedLog := range lines {
		nes.Step()
		actualLog := strings.TrimSpace(logger.buf.String())
		a.EqualNow(actualLog, expectedLog, "step %d, expected log: \"%s\", actual log: \"%s\"",
			i+1, expectedLog, actualLog)
		logger.buf.Reset()
	}
}
