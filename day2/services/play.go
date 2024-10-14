package services

type play struct {
    hand string
    points int
    winsAgainst string
    losesAgainst string
}

func (p *play) equals(s string) bool {
    return p.hand == s
}

func ConstructHands() []play {
    return []play{
        {
            hand: "ROCK",
            points: 1,
            winsAgainst: "PAPER",
            losesAgainst: "SCISSORS",
        },
        {
            hand: "PAPER",
            points: 2,
            winsAgainst: "SCISSORS",
            losesAgainst: "ROCK",
        },
        {
            hand: "SCISSORS",
            points: 3,
            winsAgainst: "ROCK",
            losesAgainst: "PAPER",
        },
    }
}
