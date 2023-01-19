package day2

import (
	"fmt"
	"testing"

	"advent.com/2022/data"
)

func TestFollowStrategyScore(t *testing.T) {
	result := CalcFollowStrategyScore(data.RPSgameInput)
	fmt.Printf("total score for follow up is %v\n", result)
}

func TestCalcEncryptedStrategyScore(t *testing.T) {
	result := CalcEncryptedStrategyScore(data.RPSgameInput)
	fmt.Printf("total score for encrypted is %v\n", result)
}
