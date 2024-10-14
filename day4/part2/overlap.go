package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntSlice []int

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	overlapCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		areas := strings.Split(line, ",")
		workers := refactor(areas)
		if overlap(workers[0], workers[1]) {
			overlapCount++
		}
	}
	fmt.Println(overlapCount)

}

func overlap(worker1, worker2 IntSlice) bool {
	if worker1.contains(worker2) || worker2.contains(worker1) {
		return true
	}
	return false
}

func refactor(workers []string) []IntSlice {
	result := []IntSlice{}
	for _, worker := range workers {
		areas := strings.Split(worker, "-")
		fromArea, _ := strconv.Atoi(areas[0])
		toArea, _ := strconv.Atoi(areas[1])
		result = append(result, generate(fromArea, toArea))
	}

	return result
}

func generate(from, to int) IntSlice {
	sections := []int{}
	for i := from; i <= to; i++ {
		sections = append(sections, i)
	}
	return sections
}

func (array IntSlice) contains(elements IntSlice) bool {
	for _, element := range elements {
		if array.contain(element) {
			return true
		}
	}
	return false
}

func (array IntSlice) contain(element int) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}
