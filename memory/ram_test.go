package memory

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestRAMMirror(t *testing.T) {
	a := assert.New(t)
	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.Write(0x0005, val)
	a.Equal(val, mmap.Read(0x0005))
	a.Equal(val, mmap.Read(0x0005+0x800))
	a.NotEqualNow(val, mmap.Read(0x0005+0x2000))
}
