package memory

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestMemoryMap_Read(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.Write(0x1000, val)
	a.Equal(val, mmap.Read(0x1000))
}

func TestMemoryMap_ReadZeroPage(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.Write(0x0050, val)
	a.Equal(val, mmap.ReadZeroPage(0x50))
}

func TestMemoryMap_ReadAbs(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.Write(0x1000, val)
	a.Equal(val, mmap.ReadAbs(0x10, 0x00))
}

func TestMemoryMap_ReadAbsShift(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.Write(0x1005, val)
	a.Equal(val, mmap.ReadAbsShift(0x10, 0x00, 0x05))
}

func TestMemoryMap_ReadIndexedIndirect(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.Write(0x1234, val)
	mmap.WriteZeroPage(0x22, 0x34)
	mmap.WriteZeroPage(0x23, 0x12)
	a.Equal(val, mmap.ReadIndexedIndirect(0x20, 0x02))
}

func TestMemoryMap_ReadIndirectIndexed(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.Write(0x1234, val)
	mmap.WriteZeroPage(0x20, 0x34)
	mmap.WriteZeroPage(0x21, 0x12)
	a.Equal(val, mmap.ReadIndirectIndexed(0x20, 0x00))
}

func TestMemoryMap_Write(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.Write(0x1000, val)
	a.Equal(val, mmap.Read(0x1000))
}

func TestMemoryMap_WriteZeroPage(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.WriteZeroPage(0x50, val)
	a.Equal(val, mmap.Read(0x0050))
}

func TestMemoryMap_WriteAbs(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.WriteAbs(0x10, 0x00, val)
	a.Equal(val, mmap.Read(0x1000))
}

func TestMemoryMap_WriteAbsShift(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.WriteAbsShift(0x10, 0x00, 0x05, val)
	a.Equal(val, mmap.Read(0x1005))
}

func TestMemoryMap_WriteIndexedIndirect(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.WriteZeroPage(0x22, 0x34)
	mmap.WriteZeroPage(0x23, 0x12)
	mmap.WriteIndexedIndirect(0x20, 0x02, val)
	a.Equal(val, mmap.Read(0x1234))
}

func TestMemoryMap_WriteIndirectIndexed(t *testing.T) {
	a := assert.New(t)

	mmap := new(MemoryMap)

	val := uint8(0x42)
	mmap.WriteZeroPage(0x20, 0x34)
	mmap.WriteZeroPage(0x21, 0x12)
	mmap.WriteIndirectIndexed(0x20, 0x00, val)
	a.Equal(val, mmap.Read(0x1234))
}
