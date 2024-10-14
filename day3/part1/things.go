package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	totalValue := 0
	for scanner.Scan() {
		linea := scanner.Text()
		middle := len(linea) / 2
		first := linea[:middle]
		second := linea[middle:]
		for i := 0; i < len(first); i++ {
			if strings.Contains(second, string(first[i])) {
				if first[i] >= 'a' && first[i] <= 'z' {
					totalValue += int(first[i]) - 'a' + 1
				} else if first[i] >= 'A' && first[i] <= 'Z' {
					totalValue += int(first[i]) - 'A' + 27
				}
				break
			}
		}
	}
	fmt.Println(totalValue)
}
