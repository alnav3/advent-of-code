package player

import "day14/part1/parser"

type Droplet parser.GameInput

func (d Droplet) canMoveDown(GameMap []*parser.GameInput) bool {

	if findEmpty([2]int{d.Coordinates[0], d.Coordinates[1] + 1}, GameMap) ||
		findEmpty([2]int{d.Coordinates[0] - 1, d.Coordinates[1] + 1}, GameMap) ||
		findEmpty([2]int{d.Coordinates[0] + 1, d.Coordinates[1] + 1}, GameMap) {
		return true
	}
	return false
}

func (d Droplet) MoveDown(GameMap []*parser.GameInput) ([]*parser.GameInput, bool) {
	for d.canMoveDown(GameMap) && !d.isGameOver(GameMap) {
		if findEmpty([2]int{d.Coordinates[0], d.Coordinates[1] + 1}, GameMap) {
			d.Coordinates[1]++
		} else if findEmpty([2]int{d.Coordinates[0] - 1, d.Coordinates[1] + 1}, GameMap) {
			d.Coordinates[0]--
			d.Coordinates[1]++
		} else if findEmpty([2]int{d.Coordinates[0] + 1, d.Coordinates[1] + 1}, GameMap) {
			d.Coordinates[0]++
			d.Coordinates[1]++
		}
	}
	GameMap = append(GameMap, (*parser.GameInput)(&d))
	return GameMap, d.isGameOver(GameMap)
}

func (d Droplet) isGameOver(GameMap []*parser.GameInput) bool {
	if d.Coordinates[1] == 0 && d.Coordinates[0] == 500 && !d.canMoveDown(GameMap) {
		return true
	}
	return false

}
func maximumCoordinates(GameMap []*parser.GameInput) (int, int, int) {
	maximumCoordinates, left, right := 0, 100000000, 0
	for i := 0; i < len(GameMap); i++ {
		if GameMap[i].Coordinates[1] > maximumCoordinates {
			maximumCoordinates = GameMap[i].Coordinates[1]
		}
		if GameMap[i].Coordinates[0] < left {
			left = GameMap[i].Coordinates[0]
		}
		if GameMap[i].Coordinates[0] > right {
			right = GameMap[i].Coordinates[0]
		}
	}
	return maximumCoordinates, left, right
}

func findEmpty(coordinates [2]int, GameMap []*parser.GameInput) bool {
	for i := 0; i < len(GameMap); i++ {
		if GameMap[i].Coordinates[0] == coordinates[0] && GameMap[i].Coordinates[1] == coordinates[1] {
			return false
		}
	}
	return true
}

func CreateFloor(GameMap []*parser.GameInput) []*parser.GameInput {
	actualFloor, left, right := maximumCoordinates(GameMap)
	for i := left - 1000; i <= right+1000; i++ {
		tempMap := &parser.GameInput{Coordinates: [2]int{i, actualFloor + 2}, Input: "#"}
		GameMap = append(GameMap, tempMap)
	}

	return GameMap
}
