package instructions

import (
	"strconv"
	"strings"
)

type instruction struct {
	Command  string
	Addvalue *int
}

type Cpu struct {
	Cycles      int
	Value       int
	Instruction *instruction
}

func getinstruction(line string) *instruction {
	parts := strings.Split(line, " ")
	if len(parts) > 1 {
		addvalue, _ := strconv.Atoi(parts[1])
		return &instruction{parts[0], &addvalue}
	}
	return nil
}

func checkValuesInCycle(valuesInCycle *map[int]int, cpu *Cpu) {
	if (*valuesInCycle)[cpu.Cycles] == -1 {
		(*valuesInCycle)[cpu.Cycles] = cpu.Value
	}
}
func Execute(line string, cpu *Cpu, valuesInCycle *map[int]int) {
	cpu.Cycles++
	checkValuesInCycle(valuesInCycle, cpu)

	if cpu.Instruction != nil {
		cpu.Value += *cpu.Instruction.Addvalue
		cpu.Cycles++
		cpu.Instruction = nil
		checkValuesInCycle(valuesInCycle, cpu)
	}

	add := getinstruction(line)
	if add != nil {
		cpu.Instruction = add
	}
}
