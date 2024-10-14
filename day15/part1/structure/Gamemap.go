package structure

type Gamemap struct {
	Width    int
	Height   int
	Type     string
	Distance int
}

func GetMinWidth(Gamemap []Gamemap) int {
	minWidth := Gamemap[0].Width - Gamemap[0].Distance
	for i := 1; i < len(Gamemap); i++ {
		if (Gamemap[i].Width - Gamemap[i].Distance) < minWidth {
			minWidth = Gamemap[i].Width - Gamemap[0].Distance
		}
	}
	return minWidth
}

func GetMaxWidth(Gamemap []Gamemap) int {
	maxWidth := Gamemap[0].Width + Gamemap[0].Distance
	for i := 1; i < len(Gamemap); i++ {
		if Gamemap[i].Width+Gamemap[0].Distance > maxWidth {
			maxWidth = Gamemap[i].Width + Gamemap[0].Distance
		}
	}
	return maxWidth
}

func Find(Gamemap []Gamemap, x int, y int) bool {
	for i := 0; i < len(Gamemap); i++ {
		if Gamemap[i].Width == x && Gamemap[i].Height == y && Gamemap[i].Type == "B" {
			return true
		}
	}
	return false
}
