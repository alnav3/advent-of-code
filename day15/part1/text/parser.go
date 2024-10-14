package parser

import (
	"strconv"
	"strings"
)

func Parse(line string) ([]int, []int) {
	parts := strings.Split(line, ":")

	partsSensor := strings.Split(parts[0], "Sensor at x=")
	partsSensor = strings.Split(partsSensor[1], ", y=")
	partsBeacon := strings.Split(parts[1], "closest beacon is at x=")
	partsBeacon = strings.Split(partsBeacon[1], ", y=")
	xSensor, _ := strconv.Atoi(partsSensor[0])
	ySensor, _ := strconv.Atoi(partsSensor[1])
	xBeacon, _ := strconv.Atoi(partsBeacon[0])
	yBeacon, _ := strconv.Atoi(partsBeacon[1])

	Sensor := []int{xSensor, ySensor}
	Beacon := []int{xBeacon, yBeacon}
	return Sensor, Beacon
}

func Distance(a []int, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
