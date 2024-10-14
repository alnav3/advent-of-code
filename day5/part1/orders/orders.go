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

	for i := 0; i < move; i++ {
		blocks[to] = append(blocks[to], blocks[from][len(blocks[from])-1])
		blocks[from] = blocks[from][:len(blocks[from])-1]
	}
	return blocks
}

func MapOrder(line string) (int, int, int) {
	orders := strings.Split(line, " ")
	move, _ := strconv.Atoi(orders[1])
	from, _ := strconv.Atoi(orders[3])
	to, _ := strconv.Atoi(orders[5])
	return move, from - 1, to - 1
}
