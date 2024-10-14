package blocks

import "strings"

func Recoverblocks(line string) []string {
	blocks := strings.Split(line, " ")
	result := []string{}
	blankSpaces := 0
	for _, block := range blocks {
		if len(block) != 0 {
			char := string(rune(block[1]))
			result = append(result, char)
			blankSpaces = 0
		} else {
			blankSpaces++
			if blankSpaces == 4 {
				result = append(result, "")
				blankSpaces = 0

			}
		}
	}
	return result
}

func GetHashmap(blocks [][]string) map[int][]string {
	result := map[int][]string{}
	for _, block := range blocks {
		for j, char := range block {
			if result[j] == nil {
				result[j] = []string{}
			}
			if char != "" {
				result[j] = append(result[j], char)
			}
		}
	}
	return reverse(result)
}

func reverse(nums map[int][]string) map[int][]string {
	for _, block := range nums {
		for i := 0; i < len(block)/2; i++ {
			block[i], block[len(block)-1-i] = block[len(block)-1-i], block[i]
		}
	}
	return nums
}
