package parser

import (
	"monkey"
	"strconv"
	"strings"
)

func Parse(line string, Monkey []*monkey.Monkey) []*monkey.Monkey {
	parts := strings.Split(line, " ")

	for len(parts) > 0 && parts[0] == "" {
		parts = parts[1:]
	}
	if len(parts) == 0 {
		return Monkey
	}

	actualMonkey := &monkey.Monkey{}
	if len(Monkey) != 0 {
		actualMonkey = Monkey[len(Monkey)-1]
	}

	switch parts[0] {
	case "Monkey":
		position := len(Monkey)
		Monkey = append(Monkey, &monkey.Monkey{Items: make([]int, 0), Operation: nil, Test: nil, Position: position})

	case "Starting":
		Starting(line, actualMonkey, Monkey)

	case "Operation:":
		Operation(line, actualMonkey)
	case "Test:":
		number, _ := strconv.Atoi(parts[3])
		Monkey[len(Monkey)-1].TestNumber = number
		Monkey[len(Monkey)-1].Test = func(value int) bool {
			return value%number == 0
		}
	case "If":
		ifCase(parts, actualMonkey)

	}
	return Monkey
}

func Starting(line string, actualMonkey *monkey.Monkey, Monkey []*monkey.Monkey) {

	objects := strings.Split(line, ":")[1]
	parts, worry := strings.Split(objects, ","), []string{}
	for _, part := range parts {
		numbers := strings.Split(part, " ")
		worry = append(worry, filter(numbers)...)
	}
	for _, worriness := range worry {
		value, _ := strconv.Atoi(worriness)
		actualMonkey.Items = append(Monkey[len(Monkey)-1].Items, value)
	}
}

func Operation(line string, actualMonkey *monkey.Monkey) {
	operation := strings.Split(line, ":")[1]
	operations := strings.Split(operation, " ")
	operationType := operations[4]
	operationValue, err := strconv.Atoi(operations[5])
	if err != nil {
		switch operationType {
		case "+":
			actualMonkey.Operation = func(value int) int {
				return value + value
			}
		default:
			actualMonkey.Operation = func(value int) int {
				return value * value
			}
		}
		return
	}
	switch operationType {
	case "+":
		actualMonkey.Operation = func(value int) int {
			return value + operationValue
		}
	default:
		actualMonkey.Operation = func(value int) int {
			return value * operationValue
		}
	}
}

func filter(arr []string) []string {
	result := []string{}
	for _, part := range arr {
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}

func ifCase(parts []string, actualMonkey *monkey.Monkey) {
	position, _ := strconv.Atoi(parts[5])
	switch parts[1] {
	case "true:":
		actualMonkey.TrueMonkey = position
	default:
		actualMonkey.FalseMonkey = position

	}

}
