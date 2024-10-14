package parser

import (
	"duel"
	"strconv"
	"strings"
)

func Parse(line string) *duel.PlayerSlice {
	line = line[1 : len(line)-1]
	p := &duel.PlayerSlice{}
	parts := strings.Split(line, ",")
	queue := parts
	for len(queue) > 0 {
		part := queue[0]
		queue = queue[1:]
		if part == "" {
			continue
		} else if part[0] == '[' {
			for len(queue) > 0 && len(part) > 0 && !isArrayFinised(part) {
				part += "," + queue[0]
				queue = queue[1:]
			}
			*p = append(*p, Parse(part))
		} else {
			number, _ := strconv.Atoi(part)
			player := duel.PlayerInt(number)
			*p = append(*p, &player)
		}
	}
	return p
}

func isArrayFinised(line string) bool {
	countOpen := 0
	countClose := 0
	for i := 0; i < len(line); i++ {
		if line[i] == '[' {
			countOpen++
		} else if line[i] == ']' {
			countClose++
		}
	}
	return countOpen == countClose
}
