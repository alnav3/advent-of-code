package main

import (
	"bufio"
	"fmt"
	"os"
    "services"
)



func main() {

    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    points := 0
    for scanner.Scan() {
        points+= services.PlayGame(scanner.Text())
    }
    fmt.Println(points)
}



