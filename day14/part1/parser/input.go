package parser

import (
	"strconv"
	"strings"
)

type GameInput struct {
	Input       string
	Coordinates [2]int
}

func MapInput(GameMap []*GameInput, line string) []*GameInput {
	parts := strings.Split(line, " -> ")
	coordinatesMap := [2]int{}
	tempMap := []*GameInput{}
	for _, Coordinates := range parts {
		splittedCoordinates := strings.Split(Coordinates, ",")
		coordinatesMap[0], _ = strconv.Atoi(splittedCoordinates[0])
		coordinatesMap[1], _ = strconv.Atoi(splittedCoordinates[1])
		tempMap = append(tempMap, &GameInput{
			Coordinates: coordinatesMap,
		})
	}
	tempMap = CreateMap(tempMap)
	GameMap = append(GameMap, tempMap...)

	return GameMap
}

func CreateMap(GameMap []*GameInput) []*GameInput {
	resultMap := []*GameInput{}
	for i := 0; i < len(GameMap)-1; i++ {
		if GameMap[i].Coordinates[0] == GameMap[i+1].Coordinates[0] {
			maximum := max(GameMap[i].Coordinates[1], GameMap[i+1].Coordinates[1])
			minimum := min(GameMap[i].Coordinates[1], GameMap[i+1].Coordinates[1])
			for j := minimum; j <= maximum; j++ {

				tempMap := &GameInput{Coordinates: [2]int{GameMap[i].Coordinates[0], j}, Input: "#"}
				resultMap = append(resultMap, tempMap)

			}
		} else if GameMap[i].Coordinates[1] == GameMap[i+1].Coordinates[1] {
			minimum := min(GameMap[i].Coordinates[0], GameMap[i+1].Coordinates[0])
			maximum := max(GameMap[i].Coordinates[0], GameMap[i+1].Coordinates[0])
			for j := minimum; j <= maximum; j++ {
				tempMap := &GameInput{Coordinates: [2]int{j, GameMap[i].Coordinates[1]}, Input: "#"}
				resultMap = append(resultMap, tempMap)
			}
		}
	}

	return resultMap
}
