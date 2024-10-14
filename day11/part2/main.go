package main

import (
	"bufio"
	"fmt"
	"monkey"
	"os"
	"parser"
	"rounds"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	monkeys := make([]*monkey.Monkey, 0)
	for scanner.Scan() {
		monkeys = parser.Parse(scanner.Text(), monkeys)
	}
	for i := 0; i < 10000; i++ {

		monkeys = rounds.Round(monkeys, i+1)
	}
	minimum, maximum := getMaximumFrom(monkeys)

	for _, Monkey := range monkeys {
		fmt.Println(Monkey)
	}
	fmt.Println(minimum * maximum)

}

func getMaximumFrom(monkeys []*monkey.Monkey) (int, int) {
	minimum, maximum := 0, 0
	for _, Monkey := range monkeys {
		if Monkey.ItemsInspected > maximum {
			minimum = maximum
			maximum = Monkey.ItemsInspected
		} else if Monkey.ItemsInspected > minimum {
			minimum = Monkey.ItemsInspected
		}

	}
	return minimum, maximum

}
