package main

import (
	"bufio"
	"fmt"
	"os"
)

type QueueItem struct {
    valve Valve
    time  int
    path  []string
}


type travel struct {
    valvesOpened []string
    elephant player
    person player
}

type player struct {
	minutesSpent int
	totalpoints  int
	flowRate     int
    valve Valve
}

type Edge struct {
    Source      string
    Destination string
    Time        int
    FlowRate    int
}

type step struct {
	personRoute Edge
	elephantRoute Edge
}

func (t travel) score() int {
    personTotalPoints := t.person.totalpoints + ((26 - t.person.minutesSpent) * t.person.flowRate)
    elephantTotalPoints := t.elephant.totalpoints + ((26 - t.elephant.minutesSpent) * t.elephant.flowRate)
    return personTotalPoints + elephantTotalPoints
}

func (e Edge) print() {
    fmt.Println("Source: ", e.Source, ", destination: ", e.Destination, ", time: ", e.Time,", flowRate: ", e.FlowRate)
}

func (e Edge) totalPointsIn(time int) int {
    usableTime := time - e.Time
    return e.FlowRate * usableTime
}

func main() {
	valves := extractValves()
	actualvalve, _ := Find("AA", valves)
	// Create a channel to receive the result
	resultChan := make(chan int)

    routes, valves := createRoutes(valves)

    for _, route := range routes {
        for _, edge := range route {
            edge.print()
        }
    }

	// Start the dfs function in a new goroutine
	go func() {
		resultChan <- dfs(routes, valves, actualvalve)
	}()

	// Receive the result from the channel
	travels := <-resultChan

	println("dfs finished")

	println("total points: ", travels)
}

func dfs(routes map[string][]Edge, valves []Valve, start Valve) int {
    queue := []travel{{[]string{start.Name}, player{0, 0, 0, start}, player{0, 0, 0, start}}}
    maxPoints := 0
    score := make(map[string]int)

    for len(queue) > 0 {
        currentTravel := queue[0]
        queue = queue[1:]

        scoreName := currentTravel.person.valve.Name + currentTravel.elephant.valve.Name + fmt.Sprint(currentTravel.person.minutesSpent) + fmt.Sprint(currentTravel.elephant.minutesSpent)
        currentScore := currentTravel.score()

        if currentTravel.person.minutesSpent >= 26 && currentTravel.elephant.minutesSpent >= 26{
            totalpoints := currentTravel.person.totalpoints + currentTravel.elephant.totalpoints
            if totalpoints > maxPoints {
                maxPoints = totalpoints
            }
            continue
        } else if len(currentTravel.valvesOpened) == len(valves) {
            totalPoints := currentTravel.person.totalpoints + ((26 - currentTravel.person.minutesSpent) * currentTravel.person.flowRate)
            totalPoints += currentTravel.elephant.totalpoints + ((26 - currentTravel.elephant.minutesSpent) * currentTravel.elephant.flowRate)
            if totalPoints > maxPoints {
                maxPoints = totalPoints
            }
            continue
        } else if len(currentTravel.valvesOpened) == len(valves) - 1 {
            totalpoints := findShortestPathAndFinish(routes, valves, currentTravel)
            if totalpoints > maxPoints {
                maxPoints = totalpoints
            }
            continue
        } else if (score[scoreName] > currentScore) {
            continue
        }
        score[scoreName] = currentScore

        possibleRoutes := getPathsFrom(routes, currentTravel)
        // Create a channel to collect the results
        results := make(chan travel, len(possibleRoutes))

        for _, route := range possibleRoutes {
            route := route
            // Start a new goroutine for each route
            go func(route step) {
                newTravel := currentTravel.copy()

                // person modifiers
                newTravel.person.minutesSpent += route.personRoute.Time
                newTravel.person.totalpoints += route.personRoute.Time * newTravel.person.flowRate
                newTravel.person.flowRate += route.personRoute.FlowRate
                newTravel.person.valve, _ = Find(route.personRoute.Destination, valves)

                // elephant modifiers
                newTravel.elephant.minutesSpent += route.elephantRoute.Time
                newTravel.elephant.totalpoints += route.elephantRoute.Time * newTravel.elephant.flowRate
                newTravel.elephant.flowRate += route.elephantRoute.FlowRate
                newTravel.elephant.valve, _ = Find(route.elephantRoute.Destination, valves)

                // valves opened
                newTravel.valvesOpened = append(newTravel.valvesOpened, route.personRoute.Destination)
                newTravel.valvesOpened = append(newTravel.valvesOpened, route.elephantRoute.Destination)

                // Send the result to the channel
                results <- newTravel
            }(route)
        }

        // Collect the results
        for i := 0; i < len(possibleRoutes); i++ {
            queue = append(queue, <-results)
        }

        // Close the channel
        close(results)
    }
    return maxPoints
}

func findShortestPathAndFinish(routes map[string][]Edge, valves []Valve, currentTravel travel) int {
    valveName := findUniqueValve(valves, currentTravel)
    if valveName == "" {
        return -1
    }

    maxpoints := -1
    for _, route := range routes[currentTravel.person.valve.Name] {
        if route.Destination == valveName {
            // person modifiers
            currentTravel.person.minutesSpent += route.Time
            currentTravel.person.totalpoints += route.Time * currentTravel.person.flowRate
            currentTravel.person.flowRate += route.FlowRate

            // get final point per each
            currentTravel.person.totalpoints += (26 - currentTravel.person.minutesSpent) * currentTravel.person.flowRate
            currentTravel.elephant.totalpoints += (26 - currentTravel.elephant.minutesSpent) * currentTravel.elephant.flowRate
            maxpoints = currentTravel.person.totalpoints + currentTravel.elephant.totalpoints
        }
    }

    for _, route := range routes[currentTravel.elephant.valve.Name] {
        if route.Destination == valveName {
            // elephant modifiers
            currentTravel.elephant.minutesSpent += route.Time
            currentTravel.elephant.totalpoints += route.Time * currentTravel.elephant.flowRate
            currentTravel.elephant.flowRate += route.FlowRate

            // get final point per each
            currentTravel.person.totalpoints += (26 - currentTravel.person.minutesSpent) * currentTravel.person.flowRate
            currentTravel.elephant.totalpoints += (26 - currentTravel.elephant.minutesSpent) * currentTravel.elephant.flowRate
            if (currentTravel.person.totalpoints + currentTravel.elephant.totalpoints) > maxpoints {
                maxpoints = currentTravel.person.totalpoints + currentTravel.elephant.totalpoints
            }
        }
    }

    return maxpoints
}

func findUniqueValve(valves []Valve, currentTravel travel) string {
    m := make(map[string]int)

    for _, val := range valves {
        m[val.Name]++
    }

    for _, val := range currentTravel.valvesOpened {
        m[val]--
    }

    for key, val := range m {
        if val != 0 {
            return key
        }
    }
    return ""
}

func getPathsFrom(routes map[string][]Edge, currentTravel travel )[]step {
    steps := make([]step, 0)

    for _, route1 := range routes[currentTravel.person.valve.Name] {
        for _, route2 := range routes[currentTravel.elephant.valve.Name] {
            if route1.Destination != route2.Destination {
                if currentTravel.person.minutesSpent + route1.Time > 26 && currentTravel.elephant.minutesSpent + route2.Time > 26 {
                    continue
                } else if (find(currentTravel.valvesOpened, route1.Destination) || find(currentTravel.valvesOpened, route2.Destination)) {
                    continue
                } else {
                    steps = append(steps, step{route1, route2})
                }
            }
        }
    }

    return steps
}

func createRoutes(valves []Valve) (map[string][]Edge, []Valve) {
    edges := make(map[string][]Edge)
    temp := make([]Valve, 0)

    for _, valve := range valves {
        if valve.FlowRate != 0 || valve.Name == "AA" {
            temp = append(temp, valve)
        }
    }

    for _, source := range temp {
        for _, destination := range temp {
            source := source
            destination := destination
            if source.Name != destination.Name && destination.Name != "AA" {
                visitedValves := make([]string, 0)
                visitedValves = append(visitedValves, source.Name)
                time := getTimeSpent(source, destination, valves)
                if time != -1 {
                    edge := Edge{source.Name, destination.Name, time, destination.FlowRate}
                    edges[source.Name] = append(edges[source.Name], edge)
                }
            }
        }
    }
    return edges, temp
}

func getTimeSpent(source, destination Valve, valves []Valve) int {
    queue := []QueueItem{{source, 0, []string{}}}

    for len(queue) > 0 {
        item := queue[0]
        queue = queue[1:]

        if item.valve.Name == destination.Name {
            return item.time + 1
        }

        for _, nextValveName := range item.valve.LeadsTo {
            if nextValveName == source.Name || find(item.path, nextValveName) {
                continue
            }
            nextValve, _ := Find(nextValveName, valves)
            queue = append(queue, QueueItem{nextValve, item.time + 1, append(item.path, nextValveName)})
        }
    }

    return -1
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
        person: player {
            minutesSpent: t.person.minutesSpent,
            totalpoints:  t.person.totalpoints,
            flowRate:     t.person.flowRate,
            valve:        t.person.valve,
        },
        elephant: player {
            minutesSpent: t.elephant.minutesSpent,
            totalpoints:  t.elephant.totalpoints,
            flowRate:     t.elephant.flowRate,
            valve:        t.elephant.valve,
        },
		valvesOpened: append([]string{}, t.valvesOpened...),
	}
}

func nextSteps(valve Valve) []string {
    steps := append(make([]string, 0), valve.Name)

	return append(steps, valve.LeadsTo...)
}

func find(valves []string, valve string) bool {
	for _, v := range valves {
		if v == valve {
			return true
		}
	}
	return false
}

