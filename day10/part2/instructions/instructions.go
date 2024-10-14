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

func Execute(line string, cpu *Cpu, crtScreen *[][]string) {

	if cpu.Instruction != nil {
		printCRT(cpu, crtScreen)
		cpu.Value += *cpu.Instruction.Addvalue
		cpu.Cycles++
		cpu.Instruction = nil
	}

	printCRT(cpu, crtScreen)
	add := getinstruction(line)
	if add != nil {
		cpu.Instruction = add
		cpu.Cycles++
	} else {
		cpu.Cycles++
	}
}

func printCRT(cpu *Cpu, crtScreen *[][]string) {
	valueDraw := []int{}
	for i := cpu.Value - 1; i <= cpu.Value+1; i++ {
		valueDraw = append(valueDraw, i)
	}
	if exists(cpu.Cycles%40, valueDraw) {
		(*crtScreen)[(cpu.Cycles-1)/40][cpu.Cycles%40] = "#"
	} else {
		(*crtScreen)[(cpu.Cycles-1)/40][cpu.Cycles%40] = "."
	}
}

func exists(index int, value []int) bool {
	for _, v := range value {
		if v == index {
			return true
		}
	}
	return false
}
