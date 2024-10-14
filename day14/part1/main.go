package main

import (
	"bufio"
	"day14/part1/parser"
	"fmt"
	"os"
	"player"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	mapInput := []*parser.GameInput{}
	for scanner.Scan() {
		line := scanner.Text()
		mapInput = parser.MapInput(mapInput, line)
	}
	droplet := player.Droplet{Coordinates: [2]int{500, 0}}
	isGameOver, dropletsCount := false, -1
	for !isGameOver {
		mapInput, isGameOver = droplet.MoveDown(mapInput)
		droplet = player.Droplet{Coordinates: [2]int{500, 0}}
		dropletsCount++
	}
	fmt.Println(dropletsCount)
}
