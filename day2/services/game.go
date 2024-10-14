package services

const (
	ROCK     = 'A'
	PAPER    = 'B'
	SCISSORS = 'C'
	DRAW     = 'Y'
	WIN      = 'Z'
	LOSE     = 'X'
)

func PlayGame(s string) int {
	player1 := decryptPlay(s[0])
	player2 := s[2]
	play := getPlay(player1)
	return getplayerPoints(player2, ConstructHands(), play)
}

func getPlay(player string) play {
	plays := ConstructHands()
	for _, p := range plays {
		if p.equals(player) {
			return p
		}
	}
	return play{}
}

func getplayerPoints(s byte, plays []play, play play) int {
	matchPoints := map[byte]int{
		DRAW: 3,
		WIN:  6,
		LOSE: 0,
	}
	if DRAW == s {
		return matchPoints[s] + play.points
	} else if WIN == s {
		winplay := getPlay(play.winsAgainst)
		return matchPoints[s] + winplay.points
	} else {
		loseplay := getPlay(play.losesAgainst)
		return matchPoints[s] + loseplay.points
	}
}

func decryptPlay(s byte) string {
	hands := map[byte]string{
		ROCK:     "ROCK",
		PAPER:    "PAPER",
		SCISSORS: "SCISSORS",
	}

	return hands[s]
}
