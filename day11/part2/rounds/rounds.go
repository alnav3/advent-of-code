package rounds

import (
	"monkey"
)

func Round(monkeys []*monkey.Monkey, round int) []*monkey.Monkey {
	for _, Monkey := range monkeys {
		for _, item := range Monkey.Items {
			item = Monkey.Operation(item)
			if Monkey.Test(item) {
				monkeys[Monkey.TrueMonkey].Items = append(monkeys[Monkey.TrueMonkey].Items, item)
			} else {
				monkeys[Monkey.FalseMonkey].Items = append(monkeys[Monkey.FalseMonkey].Items, item)
			}
			Monkey.ItemsInspected++

		}
		Monkey.Items = []int{}
	}
	monkeys = normalizeItems(monkeys)
	return monkeys
}

func normalizeItems(monkeys []*monkey.Monkey) []*monkey.Monkey {
	newItems := []int{}
	normalizeNumber := 1

	for _, Monkey := range monkeys {
		normalizeNumber *= Monkey.TestNumber
	}

	for _, Monkey := range monkeys {
		for _, item := range Monkey.Items {
			newItems = append(newItems, item%normalizeNumber)
		}
		Monkey.Items = newItems
		newItems = []int{}
	}

	return monkeys
}
