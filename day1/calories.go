package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	calories, elder := []int{}, 0
	for scanner.Scan() {
		cal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			elder++
			calories = []int{}
		} else {
			calories[elder] += cal
		}
	}
	sort.Ints(calories)
	fmt.Println(calories[len(calories)-1])
}
