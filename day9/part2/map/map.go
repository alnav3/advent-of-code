package movements

import (
	"rope"
	"strconv"
	"strings"
)

type coordinates struct {
	x, y int
}
type Map map[coordinates]bool

var m Map = make(Map)

func (m Map) Count() int {
	return len(m)
}

func Move(ropeArray []rope.Rope, line string) (Map, []rope.Rope) {
	parts := strings.Split(line, " ")
	direction := parts[0]
	times, _ := strconv.Atoi(parts[1])
	for i := 0; i < times; i++ {
		head, tail := &ropeArray[0], &ropeArray[len(ropeArray)-1]
		switch direction {
		case "U":
			head.J++
		case "D":
			head.J--
		case "L":
			head.I--
		case "R":
			head.I++
		}
		for i := 1; i < len(ropeArray); i++ {
			ropeArray[i].Move(ropeArray[i-1])
		}
		m[coordinates{tail.I, tail.J}] = true
	}
	return m, ropeArray
}
