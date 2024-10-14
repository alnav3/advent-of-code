package rounds

import (
	"monkey"
)

func Round(monkeys []*monkey.Monkey) []*monkey.Monkey {
	for _, Monkey := range monkeys {
		for _, item := range Monkey.Items {
			item = Monkey.Operation(item)
			item = item / 3
			if Monkey.Test(item) {
				monkeys[Monkey.TrueMonkey].Items = append(monkeys[Monkey.TrueMonkey].Items, item)
			} else {
				monkeys[Monkey.FalseMonkey].Items = append(monkeys[Monkey.FalseMonkey].Items, item)
			}
			Monkey.ItemsInspected++

		}
		Monkey.Items = []int{}
	}
	return monkeys
}
