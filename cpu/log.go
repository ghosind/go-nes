package cpu

import "fmt"

func (cpu *CPU) trace(pc uint16, opcode uint8, ins cpuInstruction, operands ...uint8) {
	bstr := ""
	bstr += fmt.Sprintf("%02X ", opcode)
	for i := 0; i < len(operands); i++ {
		bstr += fmt.Sprintf("%02X ", operands[i])
	}
	for len(operands) < 3 {
		bstr += "   " // 补齐
		operands = append(operands, 0)
	}

	line := fmt.Sprintf("%04X  %-9s %-28s A:%02X X:%02X Y:%02X P:%02X SP:%02X CYC:%d\n",
		pc, bstr, ins.opcode, cpu.A, cpu.X, cpu.Y, *cpu.PS, cpu.SP, cpu.Cycles)

	cpu.Logger.Println(line)
}
