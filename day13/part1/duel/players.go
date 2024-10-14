package duel

import "strconv"

type player interface {
	Match(player) int
	ToString() string
}

type PlayerInt int

type PlayerSlice []player

func (p *PlayerInt) Match(p2 player) int {

	if playerVal, ok := p2.(*PlayerInt); ok {
		if *p == *playerVal {
			return 0
		} else if *p < *playerVal {
			return -1
		} else {
			return 1
		}
	}

	if playersSlice, ok := p2.(*PlayerSlice); ok {
		if len(*playersSlice) == 1 {
			return p.Match((*playersSlice)[0])
		} else if len(*playersSlice) == 0 {
			return 1
		} else {
			p1 := &PlayerSlice{p}
			return p1.Match(playersSlice)
		}
	}

	return 0
}

func (p *PlayerSlice) Match(p2 player) int {
	if playerSlice, ok := p2.(*PlayerSlice); ok {
		for i := range *p {
			if len(*playerSlice) < i+1 {
				return 1
			} else if (*p)[i].Match((*playerSlice)[i]) != 0 {
				return (*p)[i].Match((*playerSlice)[i])
			} else if i == len(*p)-1 && i < len(*playerSlice)-1 {
				return -1
			}
		}
		if len(*p) == 0 && len(*playerSlice) > 0 {
			return -1
		}
	} else if playerInt, ok := p2.(*PlayerInt); ok {
		playerSlice := &PlayerSlice{playerInt}
		return p.Match(playerSlice)
	}
	return 0
}

func (p PlayerSlice) ToString() string {
	result := ""
	for _, slice := range p {
		result += "," + slice.ToString()
	}
	if result != "" {
		result = "[" + result[1:] + "]"
	} else {
		result = "[]"
	}
	return result
}

func (p *PlayerInt) ToString() string {
	return strconv.Itoa(int(*p))
}

type Duel struct {
	Left  *PlayerSlice
	Right *PlayerSlice
}
