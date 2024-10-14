package maze

type Maze [][]string

type Position struct {
	i, j int
}

func (m Maze) findStart() []Position {
	positions := []Position{}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == "S" || m[i][j] == "a" {
				positions = append(positions, Position{i, j})
			}
		}
	}
	return positions
}

func (m Maze) findEnd() Position {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == "E" {
				return Position{i, j}
			}
		}
	}
	return Position{-1, -1}
}

func (m Maze) scalable(from, to Position) bool {
	letterFrom := m[from.i][from.j]
	letterTo := m[to.i][to.j]
	if letterTo == "E" && letterFrom != "z" {
		return false
	}
	return (int(letterFrom[0])+1 >= int(letterTo[0]) || (letterFrom == "S" && letterTo == "a") || (letterTo == "E" && letterFrom == "z"))
}

func (m Maze) short(paths []int) int {
	shortest := paths[0]
	for _, path := range paths {
		if path < shortest {
			shortest = path
		}
	}
	return shortest
}

func (m Maze) FindShortestPath() int {
	starts := m.findStart()
	paths := []int{}
	for _, start := range starts {
		paths = append(paths, m.findShortestRoute(start))
	}
	path := m.short(paths)
	return path
}

type Node struct {
	position Position
	distance int
}

func (m Maze) findShortestRoute(start Position) int {
	directions := []Position{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	alreadyVisited := make([][]bool, len(m))
	for i := range alreadyVisited {
		alreadyVisited[i] = make([]bool, len(m[0]))
	}

	queue := []Node{{position: start, distance: 0}}
	alreadyVisited[start.i][start.j] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if m[current.position.i][current.position.j] == "E" {
			return current.distance
		}

		for _, direction := range directions {
			ni, nj := current.position.i+direction.i, current.position.j+direction.j
			if ni >= 0 && ni < len(m) && nj >= 0 && nj < len(m[0]) && !alreadyVisited[ni][nj] && m.scalable(current.position, Position{ni, nj}) {
				queue = append(queue, Node{position: Position{ni, nj}, distance: current.distance + 1})
				alreadyVisited[ni][nj] = true
			}
		}
	}

	return 99999999999
}
