package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type travel struct {
    minutesSpent int
    totalpoints int
    flowRate int
    valvesOpened []string
}


func main() {
    valves := extractValves()
    actualvalve, err := Find("AA", valves)
    if err != nil {
        panic("it should return valve with AA as name")
    }
    // Create a channel to receive the result
    resultChan := make(chan []travel)

    // Start the dfs function in a new goroutine
    go func() {
        resultChan <- dfs(valves, actualvalve, travel{minutesSpent: 0, totalpoints: 0, flowRate: 0})
    }()

    // Receive the result from the channel
    travels := <-resultChan

    println("dfs finished")
    sort.Slice(travels, func(i, j int) bool {
        return travels[i].totalpoints > travels[j].totalpoints
    })

    println("total points: ", travels[0].totalpoints)
}

func sortTravel(travels []travel) {
    sort.Slice(travels, func(i, j int) bool {
        if travels[i].flowRate != travels[j].flowRate {
            return travels[i].flowRate < travels[j].flowRate
        }
        if travels[i].minutesSpent != travels[j].minutesSpent {
            return travels[i].minutesSpent > travels[j].minutesSpent
        }
        return travels[i].totalpoints < travels[j].totalpoints
    })
}

func extractValves() []Valve {
    file, _ := os.Open("input.txt")
    defer file.Close()

    reader := bufio.NewReader(file)

    line, _, err := reader.ReadLine()
    valves := make([]Valve, 0)
    for err == nil {
        valves = append(valves, parseValve(string(line)))

        line, _, err = reader.ReadLine()
    }
    return valves
}
func (t travel) copy() travel {
    return travel{
        minutesSpent: t.minutesSpent,
        totalpoints: t.totalpoints,
        flowRate: t.flowRate,
        valvesOpened: append([]string{}, t.valvesOpened...),
    }
}

type TravelWithValve struct {
    travel travel
    valve  Valve
}

func dfs(valves []Valve, startValve Valve, startTravel travel) []travel {
    var travels []travel
    stack := []TravelWithValve{{startTravel, startValve}}
    bestScores := make(map[string]int)

    for len(stack) > 0 {
        actual := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        actualTravel := actual.travel
        actualValve := actual.valve

        // Pruning: if we've seen this state before with a better score, skip it
        state := actualValve.Name + strconv.Itoa(actualTravel.minutesSpent)
        if score, seen := bestScores[state]; seen && score >= actualTravel.totalpoints {
            continue
        }
        bestScores[state] = actualTravel.totalpoints

        if (actualTravel.minutesSpent >= 30) {
            travels = append(travels, actualTravel)
            continue
        }


        if actualValve.FlowRate > 0 && !find(actualTravel.valvesOpened, actualValve.Name) {
            fmt.Println(actualTravel.valvesOpened)
            newTravel := actualTravel.copy()
            newTravel.totalpoints += newTravel.flowRate
            newTravel.flowRate += actualValve.FlowRate
            newTravel.minutesSpent++
            newTravel.valvesOpened = append(newTravel.valvesOpened, actualValve.Name)

            stack = append(stack, TravelWithValve{newTravel, actualValve})
        }

        for _, leadsTo := range actualValve.LeadsTo {
            newValve, _ := Find(leadsTo, valves)
            newTravel := actualTravel.copy()
            newTravel.minutesSpent++
            newTravel.totalpoints += newTravel.flowRate

            stack = append(stack, TravelWithValve{newTravel, newValve})
        }
    }

    return travels
}

func find(valves []string, valve string) bool {
    for _, v := range valves {
        if v == valve {
            return true
        }
    }
    return false
}


