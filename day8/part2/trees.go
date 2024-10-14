package main

import (
	"bufio"
	"fmt"
	"os"
	"parser"
	"visibility"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	forest := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		array := parser.ParseArray(line)
		forest = append(forest, array)
	}
	perfectTree := visibility.Score(forest)
	fmt.Println(perfectTree)
}
