package main

import (
	"bufio"
	"duel"
	"fmt"
	"os"
	"parser"
	"sort"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	players := duel.PlayerSlice{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else {
			players = append(players, parser.Parse(line))
		}
	}
	players = append(players, parser.Parse("[[2]]"))
	players = append(players, parser.Parse("[[6]]"))
	sort.Slice(players, func(i, j int) bool {
		return players[i].Match(players[j]) == -1
	})
	count := 1
	for i := range players {
		if players[i].ToString() == "[[2]]" || players[i].ToString() == "[[6]]" {
			count *= (i + 1)
		}
		fmt.Println(players[i].ToString())
	}
	fmt.Println(count)

}
