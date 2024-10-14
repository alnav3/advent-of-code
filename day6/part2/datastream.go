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
	for scanner.Scan() {
		fmt.Println(findDataStream(scanner.Text()))
	}

}

func findDataStream(input string) (string, int) {
	result := ""
	len := 0
	startOfPacker := 0
	for i, ch := range input {
		len++
		if strings.Contains(result, string(ch)) {
			len, result = deleteRepeatedString(result, ch)
		} else if len == 14 {
			result += string(ch)
			startOfPacker = 1 + i
			break
		} else {
			result += string(ch)
		}
	}
	return result, startOfPacker

}

func deleteRepeatedString(input string, ch rune) (int, string) {
	result := ""
	index := strings.Index(input, string(ch))
	result += input[index+1:]
	result += string(ch)

	return len(result), result
}
