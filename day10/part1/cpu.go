package main

import (
	"bufio"
	"day10/part1/instructions"
	"fmt"
	"os"
)

func main() {
	cpu := &instructions.Cpu{0, 1, nil}
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	valuesInCycle := map[int]int{20: -1, 60: -1, 100: -1, 140: -1, 180: -1, 220: -1}
	for scanner.Scan() {
		line := scanner.Text()
		instructions.Execute(line, cpu, &valuesInCycle)
	}
	fmt.Println(calculate(valuesInCycle))
}

func calculate(valuesInCycle map[int]int) int {
	sum := 0
	for i, value := range valuesInCycle {
		sum += value * i
	}
	return sum
}
