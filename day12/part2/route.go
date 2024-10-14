package main

import (
	"bufio"
	"fmt"
	"maze"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	maze := maze.Maze{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		maze = append(maze, line)
	}
	shortestPath := maze.FindShortestPath()
	fmt.Println(shortestPath)
}
