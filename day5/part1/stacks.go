package main

import (
	"blocks"
	"bufio"
	"fmt"
	"orders"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	letters := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.Split(line, "[")) == 1 {
			break
		}
		letters = append(letters, blocks.Recoverblocks(line))
	}
	stacks := blocks.GetHashmap(letters)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.Split(line, "move")) != 1 {
			line := scanner.Text()
			move, from, to := orders.MapOrder(line)
			stacks = orders.ExecuteOrder(move, from, to, stacks)
		}
	}
	printlaststack(stacks)

	fmt.Println(stacks)
	file.Close()
}

func printlaststack(stacks map[int][]string) string {
	result := ""
	for _, stack := range stacks {
		result += stack[len(stack)-1]
	}
	return result

}
