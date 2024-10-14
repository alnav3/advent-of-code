package main

import (
	"bufio"
	"duel"
	"fmt"
	"os"
	"parser"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	duels := []duel.Duel{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else {
			duel := duel.Duel{}
			duel.Left = parser.Parse(line)
			if scanner.Scan() {
				duel.Right = parser.Parse(scanner.Text())
			}
			duels = append(duels, duel)
		}
	}
	counter := 0
	for i, duel := range duels {
		fmt.Println(duel.Left.ToString())
		fmt.Println(duel.Right.ToString())
		match := duel.Left.Match(duel.Right)
		if match == -1 {
			counter += i + 1
		}
		fmt.Println(match)
	}
	fmt.Println(counter)
}
