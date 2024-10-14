package visibility

func VisibleTrees(forest [][]int) int {
	visibleTrees := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			if isTreeVisible(forest, i, j) {
				visibleTrees++
			}
		}
	}
	return visibleTrees
}

func isTreeVisible(forest [][]int, i, j int) bool {
	if i == 0 || i == len(forest)-1 || j == 0 || j == len(forest[i])-1 {
		return true
	}
	if isVisibleFromLeft(forest, i, j) || isVisibleFromRight(forest, i, j) || isVisibleFromTop(forest, i, j) || isVisibleFromBottom(forest, i, j) {
		return true
	}
	return false
}

func isVisibleFromLeft(forest [][]int, i, j int) bool {
	for k := j - 1; k >= 0; k-- {
		if forest[i][k] >= forest[i][j] {
			return false
		}
	}
	return true
}

func isVisibleFromRight(forest [][]int, i, j int) bool {
	for k := j + 1; k < len(forest[i]); k++ {
		if forest[i][k] >= forest[i][j] {
			return false
		}
	}
	return true
}

func isVisibleFromTop(forest [][]int, i, j int) bool {
	for k := i - 1; k >= 0; k-- {
		if forest[k][j] >= forest[i][j] {
			return false
		}
	}
	return true
}

func isVisibleFromBottom(forest [][]int, i, j int) bool {
	for k := i + 1; k < len(forest); k++ {
		if forest[k][j] >= forest[i][j] {
			return false
		}
	}
	return true
}
