package main

import (
	"fmt"
	"strings"
)

type Valve struct {
	Name     string
	FlowRate int
	LeadsTo  []string
}

func (v Valve) FlowRateIn(minutes, distance int) int {
    return v.FlowRate * (minutes - distance + 1)
}

func Find(name string, valves []Valve) (Valve, error) {
    for _, valve := range valves {
        if valve.Name == name {
            return valve, nil
        }
    }
    return Valve{}, fmt.Errorf("Valve %s not found", name)
}

func parseValve(s string) Valve {
	var name string
	var flowRate int

	fmt.Sscanf(s, "Valve %s has flow rate=%d", &name, &flowRate)

	leadsToStr := strings.Split(s, "valve")[1]
	leadsTo := strings.Split(leadsToStr, ", ")

    // control case that leadsTo starts with valves instead of valve
    if strings.HasPrefix(leadsTo[0], "s ") {
        leadsTo[0] = strings.TrimPrefix(leadsTo[0], "s ")
    } else {
        leadsTo[0] = strings.TrimPrefix(leadsTo[0], " ")
    }

	for i, valve := range leadsTo {
		leadsTo[i] = strings.TrimSpace(valve)
	}

    return Valve{
		Name:     name,
		FlowRate: flowRate,
		LeadsTo:  leadsTo,
	}
}

