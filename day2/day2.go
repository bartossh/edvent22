package day2

import "strings"

const (
	A = 'A' // rock
	B = 'B' // paper
	C = 'C' // scissors

	X = 'X' // rock, lose
	Y = 'Y' // paper, draw
	Z = 'Z' // scissors, win
)

var (
	AColl = []rune{Z, X, Y} // lose, draw, win
	BColl = []rune{X, Y, Z}
	CColl = []rune{Y, Z, X}
)

func CalcFollowStrategyScore(data string) int {
	score := 0
	rows := strings.Split(data, "\n")
	for _, r := range rows {
		game := strings.Split(r, " ")
		if len(game) != 2 {
			continue
		}

		score += handValue(rune(game[1][0])) + calcSingleGameOutcome(rune(game[0][0]), rune(game[1][0]))
	}
	return score
}

func CalcEncryptedStrategyScore(data string) int {
	score := 0
	rows := strings.Split(data, "\n")
	for _, r := range rows {
		game := strings.Split(r, " ")
		if len(game) != 2 {
			continue
		}
		playHand := calcPlayedHand(rune(game[0][0]), rune(game[1][0]))
		score += handValue(playHand) + calcSingleGameOutcome(rune(game[0][0]), playHand)
	}
	return score
}

func handValue(hand rune) int {
	switch hand {
	case X:
		return 1
	case Y:
		return 2
	case Z:
		return 3
	default:
		return 0
	}
}

func calcPlayedHand(opponent, effect rune) rune {
	idx := getCollIndex(effect)
	switch opponent {
	case A:
		return AColl[idx]
	case B:
		return BColl[idx]
	default:
		return CColl[idx]
	}
}

func getCollIndex(r rune) int {
	switch r {
	case X:
		return 0
	case Y:
		return 1
	default:
		return 2
	}
}

func calcSingleGameOutcome(opponent, player rune) int {
	switch {
	case opponent == A && player == Y:
		return 6
	case opponent == B && player == Z:
		return 6
	case opponent == C && player == X:
		return 6
	case opponent == A && player == X:
		return 3
	case opponent == B && player == Y:
		return 3
	case opponent == C && player == Z:
		return 3
	default:
		return 0
	}
}
