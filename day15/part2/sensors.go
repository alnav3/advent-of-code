package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"parser"
	"sort"
	"structure"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	sensorBeaconMap := structure.SBMap{}

	line, _, err := reader.ReadLine()
	for err == nil {
		sensor, beacon := parser.Parse(string(line), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d")
		sensorBeaconMap.Sensor = append(sensorBeaconMap.Sensor, structure.Gamemap{Width: sensor[0], Height: sensor[1], Distance: parser.Distance(sensor, beacon)})
		sensorBeaconMap.Beacon = append(sensorBeaconMap.Beacon, structure.Gamemap{Width: beacon[0], Height: beacon[1]})
		line, _, err = reader.ReadLine()
	}

	if err != nil && err != io.EOF {
		fmt.Println("Error reading file:", err)
		return
	}

	minWidth, maxWidth := 0, 4000000
	ranges := make([]structure.Rangex, 0, len(sensorBeaconMap.Sensor))

	for y := minWidth; y <= maxWidth; y++ {
		for _, sensor := range sensorBeaconMap.Sensor {
			rangeX, err := sensor.Minmaxin(y)
			if err == nil {
				ranges = append(ranges, rangeX)
			}
		}
		ranges = join(ranges)
		if len(ranges) > 1 {
			x := ranges[0].Max + 1
			result := x*maxWidth + y
			panic(fmt.Errorf("finish! result = %d", result))
		}
		ranges = ranges[:0] // Clear the slice without reallocation
	}
}

func join(minmax []structure.Rangex) []structure.Rangex {
    // Sort the minmax slice based on the Min field
    sort.Slice(minmax, func(i, j int) bool {
        return minmax[i].Min < minmax[j].Min
    })

    // Initialize the result slice with the first element of minmax
    result := []structure.Rangex{minmax[0]}

    for _, rangex := range minmax[1:] {
        last := &result[len(result)-1]

        // If the current rangex overlaps with the last rangex in the result, merge them
        if rangex.Min <= last.Max+1 {
            if last.Max < rangex.Max {
                last.Max = rangex.Max
            }
        } else {
            // If the current rangex does not overlap, add it to the result
            result = append(result, rangex)
        }
    }

    return result
}
