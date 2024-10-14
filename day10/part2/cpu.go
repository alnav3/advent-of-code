package main

import (
	"bufio"
	"day10/part2/instructions"
	"fmt"
	"os"
)

func main() {
	cpu := &instructions.Cpu{Cycles: 0, Value: 1, Instruction: nil}
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	crtScreen := make([][]string, 6)
	for i := 0; i < 6; i++ {
		crtScreen[i] = make([]string, 40)
	}
	for scanner.Scan() {
		line := scanner.Text()
		instructions.Execute(line, cpu, &crtScreen)
	}
	for _, line := range crtScreen {
		fmt.Println(line)
	}
}
