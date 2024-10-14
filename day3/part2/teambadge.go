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

	lineas := []string{}
	for scanner.Scan() {
		lineas = append(lineas, scanner.Text())
		if len(lineas) == 3 {
			for _, byte := range lineas[0] {
				if strings.ContainsRune(lineas[1], byte) && strings.ContainsRune(lineas[2], byte) {
					if int(byte) >= int('a') && int(byte) <= int('z') {
						totalValue += int(byte) - int('a') + 1
						break
					} else if int(byte) >= int('A') && int(byte) <= int('Z') {
						totalValue += int(byte) - int('A') + 27
						break
					}
				}
			}
			lineas = []string{}
		}
	}
	fmt.Println(totalValue)
}
