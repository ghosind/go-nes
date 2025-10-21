package cpu

import (
	"fmt"
	"strings"
)

func (cpu *CPU) trace(pc uint16, opcode uint8, ins cpuInstruction, operands ...uint8) {
	bstr := ""
	bstr += fmt.Sprintf("%02X ", opcode)
	for i := 0; i < len(operands); i++ {
		bstr += fmt.Sprintf("%02X ", operands[i])
	}

	asm := cpu.disassemble(pc, ins, operands...)

	ppuTicks := cpu.Cycles * 3
	ppux := ppuTicks % 341
	ppuy := (ppuTicks / 341) % 262

	fmt := "%04X  %-9s %-31s A:%02X X:%02X Y:%02X P:%02X SP:%02X PPU:%3d,%3d CYC:%d"
	if strings.HasPrefix(asm, "*") {
		fmt = "%04X  %-9s%-32s A:%02X X:%02X Y:%02X P:%02X SP:%02X PPU:%3d,%3d CYC:%d"
	}

	cpu.Logger.Printf(fmt, pc, bstr, asm, cpu.A, cpu.X, cpu.Y, *cpu.PS, cpu.SP, ppuy, ppux, cpu.Cycles)
}

func (cpu *CPU) disassemble(pc uint16, ins cpuInstruction, operands ...uint8) string {
	buf := make([]byte, 0, 31)
	buf = append(buf, []byte(ins.opcode)...)

	switch ins.addressing {
	case addressingModeImmediate:
		buf = append(buf, fmt.Sprintf(" #$%02X", operands[0])...)
	case addressingModeZeroPage:
		buf = append(buf, fmt.Sprintf(" $%02X = %02X",
			operands[0], cpu.mem.ReadZeroPage(operands[0]))...)
	case addressingModeZeroPageX:
		buf = append(buf, fmt.Sprintf(" $%02X,X @ %02X = %02X",
			operands[0], operands[0]+cpu.X,
			cpu.mem.ReadZeroPage(operands[0]+cpu.X))...)
	case addressingModeZeroPageY:
		buf = append(buf, fmt.Sprintf(" $%02X,Y @ %02X = %02X",
			operands[0], operands[0]+cpu.Y,
			cpu.mem.ReadZeroPage(operands[0]+cpu.Y))...)
	case addressingModeAbsolute:
		if ins.opcode == "JMP" || ins.opcode == "JSR" {
			buf = append(buf, fmt.Sprintf(" $%02X%02X", operands[1], operands[0])...)
		} else {
			buf = append(buf, fmt.Sprintf(" $%02X%02X = %02X",
				operands[1], operands[0], cpu.mem.ReadAbs(operands[1], operands[0]))...)
		}
	case addressingModeAbsoluteX:
		addr := (uint16(operands[1]) << 8) | uint16(operands[0]) + uint16(cpu.X)
		value := cpu.mem.ReadAbsShift(operands[1], operands[0], cpu.X)

		buf = append(buf, fmt.Sprintf(" $%02X%02X,X @ %04X = %02X",
			operands[1], operands[0], addr, value)...)
	case addressingModeAbsoluteY:
		addr := (uint16(operands[1]) << 8) | uint16(operands[0]) + uint16(cpu.Y)
		value := cpu.mem.ReadAbsShift(operands[1], operands[0], cpu.Y)

		buf = append(buf, fmt.Sprintf(" $%02X%02X,Y @ %04X = %02X",
			operands[1], operands[0], addr, value)...)
	case addressingModeIndirect:
		buf = append(buf, fmt.Sprintf(" ($%02X%02X) = %02X%02X",
			operands[1], operands[0],
			cpu.mem.ReadAbs(operands[1], operands[0]+1),
			cpu.mem.ReadAbs(operands[1], operands[0]))...)
	case addressingModeIndexedIndirect:
		addr := operands[0] + cpu.X
		low := cpu.mem.ReadZeroPage(addr)
		high := cpu.mem.ReadZeroPage(addr + 1)

		buf = append(buf, fmt.Sprintf(" ($%02X,X) @ %02X = %02X%02X = %02X",
			operands[0], addr, high, low, cpu.mem.ReadAbs(high, low))...)
	case addressingModeIndirectIndexed:
		low := cpu.mem.ReadZeroPage(operands[0])
		high := cpu.mem.ReadZeroPage(operands[0] + 1)
		addr := uint16(high)<<8 | uint16(low) + uint16(cpu.Y)

		buf = append(buf, fmt.Sprintf(" ($%02X),Y = %02X%02X @ %04X = %02X",
			operands[0], high, low, addr, cpu.mem.ReadAbsShift(high, low, cpu.Y))...)
	case addressingModeAccumulator:
		buf = append(buf, " A"...)
	case addressingModeImplied:
		// nothing to do
	case addressingModeRelative:
		addr := pc + 2
		if operands[0] < 0x80 {
			addr += uint16(operands[0])
		} else {
			addr -= 0x100 - uint16(operands[0])
		}
		buf = append(buf, fmt.Sprintf(" $%02X", addr)...)
	default:
		buf = append(buf, " ???"...)
	}

	return string(buf)
}
