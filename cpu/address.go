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

func (cpu *CPU) fetchOperands(mode AddressingMode) ([]uint8, int) {
	additionalCycles := 0

	switch mode {
	case addressingModeImmediate, addressingModeZeroPage, addressingModeZeroPageX,
		addressingModeZeroPageY, addressingModeIndexedIndirect, addressingModeIndirectIndexed,
		addressingModeRelative:
		op := cpu.fetch()
		return []uint8{op}, additionalCycles
	case addressingModeAbsolute, addressingModeIndirect:
		low := cpu.fetch()
		high := cpu.fetch()
		return []uint8{low, high}, additionalCycles
	case addressingModeAbsoluteX:
		low := cpu.fetch()
		high := cpu.fetch()
		if (uint16(low) + uint16(cpu.x)) > 0xFF {
			additionalCycles = 1
		}
		return []uint8{low, high}, additionalCycles
	case addressingModeAbsoluteY:
		low := cpu.fetch()
		high := cpu.fetch()
		if (uint16(low) + uint16(cpu.y)) > 0xFF {
			additionalCycles = 1
		}
		return []uint8{low, high}, additionalCycles
	default:
		return nil, 0
	}
}
