package monkey

type Monkey struct {
	Items          []int
	Operation      func(int) int
	Test           func(int) bool
	TestNumber     int
	TrueMonkey     int
	FalseMonkey    int
	Position       int
	ItemsInspected int
}
