package cpu

// ProcessorStatus represents the status register of the CPU.
// Each bit in the status register represents a different flag.
// The flags are as follows:
// | Negative | Overflow |   -   | Break | Decimal | Interrupt | Zero | Carry |
type ProcessorStatus uint8

const (
	psFlagCarry     ProcessorStatus = 1 << 0 // Carry Flag
	psFlagZero      ProcessorStatus = 1 << 1 // Zero Flag
	psFlagInterrupt ProcessorStatus = 1 << 2 // Interrupt Disable
	psFlagDecimal   ProcessorStatus = 1 << 3 // Decimal Mode (not used in NES)
	psFlagUnused    ProcessorStatus = 1 << 4 // Unused, always set to 1
	psFlagBreak     ProcessorStatus = 1 << 4 // Break Command
	psFlagOverflow  ProcessorStatus = 1 << 6 // Overflow Flag
	psFlagNegative  ProcessorStatus = 1 << 7 // Negative Flag
)

func (ps *ProcessorStatus) set(flag ProcessorStatus, value bool) {
	if value {
		*ps |= flag
	} else {
		*ps = *ps & (^flag)
	}
}

func (ps ProcessorStatus) get(flag ProcessorStatus) bool {
	return ps&flag != 0
}

func (ps *ProcessorStatus) setCarry(value bool) {
	ps.set(psFlagCarry, value)
}

func (ps ProcessorStatus) getCarry() bool {
	return ps.get(psFlagCarry)
}

func (ps *ProcessorStatus) setZero(value bool) {
	ps.set(psFlagZero, value)
}

func (ps ProcessorStatus) getZero() bool {
	return ps.get(psFlagZero)
}

func (ps *ProcessorStatus) setInterrupt(value bool) {
	ps.set(psFlagInterrupt, value)
}

func (ps ProcessorStatus) getInterrupt() bool {
	return ps.get(psFlagInterrupt)
}

func (ps *ProcessorStatus) setDecimal(value bool) {
	ps.set(psFlagDecimal, value)
}

func (ps ProcessorStatus) getDecimal() bool {
	return ps.get(psFlagDecimal)
}

func (ps *ProcessorStatus) setUnused(value bool) {
	ps.set(psFlagUnused, value)
}

func (ps *ProcessorStatus) setBreak(value bool) {
	ps.set(psFlagBreak, value)
}

func (ps ProcessorStatus) getBreak() bool {
	return ps.get(psFlagBreak)
}

func (ps *ProcessorStatus) setOverflow(value bool) {
	ps.set(psFlagOverflow, value)
}

func (ps ProcessorStatus) getOverflow() bool {
	return ps.get(psFlagOverflow)
}

func (ps *ProcessorStatus) setNegative(value bool) {
	ps.set(psFlagNegative, value)
}

func (ps ProcessorStatus) getNegative() bool {
	return ps.get(psFlagNegative)
}

func (ps *ProcessorStatus) setZeroNeg(value uint8) {
	ps.setZero(value == 0)
	ps.setNegative(value&0x80 != 0)
}

func (ps *ProcessorStatus) setNegByValue(value uint8) {
	ps.setNegative(value&0x80 != 0)
}

func (ps *ProcessorStatus) setOverflowByValue(value uint8) {
	ps.setOverflow(value&0x40 != 0)
}
