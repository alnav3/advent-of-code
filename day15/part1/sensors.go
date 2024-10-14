package main

import (
	"bufio"
	"os"
	"parser"
	"structure"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	Gamemap := []structure.Gamemap{}
	for scanner.Scan() {
		line := scanner.Text()
		Sensor, Beacon := parser.Parse(line)
		Gamemap = append(Gamemap, structure.Gamemap{Width: Sensor[0], Height: Sensor[1], Type: "S", Distance: parser.Distance(Sensor, Beacon)})
		Gamemap = append(Gamemap, structure.Gamemap{Width: Beacon[0], Height: Beacon[1], Type: "B"})
	}
	minWidth := structure.GetMinWidth(Gamemap) * -1
	maxWidth := structure.GetMaxWidth(Gamemap) * 2

	freePositions := 0
	for i := minWidth; i <= maxWidth; i++ {
		for j := 0; j < len(Gamemap); j++ {
			if Gamemap[j].Type == "S" {
				if parser.Distance([]int{Gamemap[j].Width, Gamemap[j].Height}, []int{i, 2000000}) <= Gamemap[j].Distance && !structure.Find(Gamemap, i, 2000000) {
					freePositions++
					break
				}
			}
		}
	}
	println(freePositions)
}
