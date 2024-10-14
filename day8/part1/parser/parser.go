package parser

import "strconv"

func ParseArray(line string) []int {
	array := []int{}
	for i := 0; i < len(line); i++ {
		height, _ := strconv.Atoi(string(line[i]))
		array = append(array, height)
	}
	return array
}
