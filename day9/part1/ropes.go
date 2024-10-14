package main

import (
	"bufio"
	"fmt"
	movements "map"
	"os"
	"rope"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	head := rope.Rope{0, 0}
	tail := rope.Rope{0, 0}
	m := movements.Map{}
	for scanner.Scan() {
		m, head, tail = movements.Move(head, tail, scanner.Text())
	}
	fmt.Println(m.Count())
}
