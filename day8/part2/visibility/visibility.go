package visibility

func Score(forest [][]int) int {
	actualPoints, point := 0, 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			point = points(forest, i, j)
			if point > actualPoints {
				actualPoints = point
			}
		}
	}
	return actualPoints
}

func points(forest [][]int, i, j int) int {

	return visibleFromLeft(forest, i, j) * visibleFromRight(forest, i, j) * visibleFromTop(forest, i, j) * visibleFromBottom(forest, i, j)
}

func visibleFromLeft(forest [][]int, i, j int) int {
	counter := 0
	for k := j - 1; k >= 0; k-- {
		if forest[i][k] < forest[i][j] {
			counter++
		} else if forest[i][k] >= forest[i][j] {
			counter++
			return counter
		}
	}
	return counter
}

func visibleFromRight(forest [][]int, i, j int) int {
	counter := 0
	for k := j + 1; k < len(forest[i]); k++ {
		if forest[i][k] < forest[i][j] {
			counter++
		} else if forest[i][k] >= forest[i][j] {
			counter++
			return counter
		}
	}
	return counter
}

func visibleFromTop(forest [][]int, i, j int) int {
	counter := 0
	for k := i - 1; k >= 0; k-- {
		if forest[k][j] < forest[i][j] {
			counter++
		} else if forest[k][j] >= forest[i][j] {
			counter++
			return counter
		}
	}
	return counter
}

func visibleFromBottom(forest [][]int, i, j int) int {
	counter := 0
	for k := i + 1; k < len(forest); k++ {
		if forest[k][j] < forest[i][j] {
			counter++
		} else if forest[k][j] >= forest[i][j] {
			counter++
			return counter
		}
	}
	return counter
}
