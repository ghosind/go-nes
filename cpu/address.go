package cpu

type AddressingMode int

const (
	addressingModeImmediate AddressingMode = iota
	addressingModeZeroPage
	addressingModeZeroPageX
	addressingModeZeroPageY
	addressingModeAbsolute
	addressingModeAbsoluteX
	addressingModeAbsoluteY
	addressingModeIndirect
	addressingModeIndexedIndirect
	addressingModeIndirectIndexed
	addressingModeAccumulator
	addressingModeImplied
	addressingModeRelative
)

func (cpu *CPU) fetchOperands(mode AddressingMode) ([]uint8, uint64) {
	additionalCycles := uint64(0)

	switch mode {
	case addressingModeImmediate, addressingModeZeroPage, addressingModeZeroPageX,
		addressingModeZeroPageY, addressingModeIndexedIndirect, addressingModeRelative:
		op := cpu.fetch()
		return []uint8{op}, additionalCycles
	case addressingModeAbsolute, addressingModeIndirect:
		low := cpu.fetch()
		high := cpu.fetch()
		return []uint8{low, high}, additionalCycles
	case addressingModeAbsoluteX:
		low := cpu.fetch()
		high := cpu.fetch()
		if (uint16(low) + uint16(cpu.X)) > 0xFF {
			additionalCycles = 1
		}
		return []uint8{low, high}, additionalCycles
	case addressingModeAbsoluteY:
		low := cpu.fetch()
		high := cpu.fetch()
		if (uint16(low) + uint16(cpu.Y)) > 0xFF {
			additionalCycles = 1
		}
		return []uint8{low, high}, additionalCycles
	case addressingModeIndirectIndexed:
		addr := cpu.fetch()
		low := cpu.mem.ReadZeroPage(addr)
		if (uint16(low) + uint16(cpu.Y)) > 0xFF {
			additionalCycles = 1
		}
		return []uint8{addr}, additionalCycles
	default:
		return nil, 0
	}
}
