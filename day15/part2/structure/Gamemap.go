package structure

import (
	"errors"
	"parser"
)

type SBMap struct {
    Sensor  []Gamemap
    Beacon  []Gamemap
}
type Gamemap struct {
	Width    int
	Height   int
	Distance int
}
type Rangex struct {
    Min int
    Max int
}

func (g Gamemap) Minmaxin(y int) (Rangex, error) {
    distancex := g.Distance - parser.Abs(g.Height-y)
    if distancex < 0 {
        return Rangex{}, errors.New("distance is less than zero")
    }
    return Rangex{Min: g.Width - distancex, Max: g.Width + distancex}, nil
}

func GetMinMaxWidth(Gamemap []Gamemap) (int, int) {
    minWidth := Gamemap[0].Width - Gamemap[0].Distance
    maxWidth := Gamemap[0].Width + Gamemap[0].Distance

    for i := 1; i < len(Gamemap); i++ {
        currentMinWidth := Gamemap[i].Width - Gamemap[i].Distance
        currentMaxWidth := Gamemap[i].Width + Gamemap[i].Distance

        if currentMinWidth < minWidth {
            minWidth = currentMinWidth
        }
        if currentMaxWidth > maxWidth {
            maxWidth = currentMaxWidth
        }
    }

    return minWidth, maxWidth
}

func Find(Gamemap []Gamemap, x int, y int) bool {
    for _, mapElement := range Gamemap {
        if mapElement.Width == x && mapElement.Height == y {
            return true
        }
    }
    return false
}
