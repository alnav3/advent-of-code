package parser

import (
	"fmt"
)

func Parse(line string, templ string) ([]int, []int) {
	var xSensor, ySensor, xBeacon, yBeacon int
	fmt.Sscanf(line,templ, &xSensor, &ySensor, &xBeacon, &yBeacon)
	Sensor := []int{xSensor, ySensor}
	Beacon := []int{xBeacon, yBeacon}
	return Sensor, Beacon
}

func Distance(a []int, b []int) int {
	return Abs(a[0]-b[0]) + Abs(a[1]-b[1])
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
