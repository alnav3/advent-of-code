package orders

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("vim-go")
}

func ExecuteOrder(move, from, to int, blocks map[int][]string) map[int][]string {
	blocksToMove := getBlocksToMove(move, blocks[from])

	blocks[to] = append(blocks[to], blocksToMove...)
	blocks[from] = blocks[from][:len(blocks[from])-move]
	return blocks
}

func MapOrder(line string) (int, int, int) {
	orders := strings.Split(line, " ")
	move, _ := strconv.Atoi(orders[1])
	from, _ := strconv.Atoi(orders[3])
	to, _ := strconv.Atoi(orders[5])
	return move, from - 1, to - 1
}

func getBlocksToMove(move int, blocks []string) []string {
	result := []string{}
	for i := len(blocks) - move; i < len(blocks); i++ {
		result = append(result, blocks[i])
	}
	return result
}
