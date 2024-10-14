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
	ropeArray := []rope.Rope{}
	for i := 0; i < 10; i++ {
		ropeArray = append(ropeArray, rope.Rope{0, 0})
	}
	m := movements.Map{}
	for scanner.Scan() {
		m, ropeArray = movements.Move(ropeArray, scanner.Text())
	}
	fmt.Println(m.Count())
}
