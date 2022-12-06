package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Rock     = "A" // +1 points
	Paper    = "B" // +2 points
	Scissors = "C" // +3 points

	Loss = "X" // +0 points
	Draw = "Y" // +3 points
	Win  = "Z" // +6 points
)

func scorePlayer(player string) {
	switch player {
	case Lose:
		round_score += 0
	case Draw:
		round_score += 3
	case Win:
		round_score += 6
	}
}

func main() {

	var score int

	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	index := 0
	for scanner.Scan() {
		index += 1
		round_score := 0

		throws := strings.Split(scanner.Text(), " ")

		opponent := throws[0]
		player := throws[1]

		switch player {
		case Loss:
			round_score += 0

			switch opponent {
			case Rock:
				round_score += 3
			case Paper:
				round_score += 1
			case Scissors:
				round_score += 2
			}

		case Draw:
			round_score += 3

			switch opponent {
			case Rock:
				round_score += 1
			case Paper:
				round_score += 2
			case Scissors:
				round_score += 3
			}

		case Win:
			round_score += 6

			switch opponent {
			case Rock:
				round_score += 2
			case Paper:
				round_score += 3
			case Scissors:
				round_score += 1
			}
		}

		fmt.Println("Round:", index, "- Score:", round_score)

		score = score + round_score
	}

	fmt.Println("-------------------")
	fmt.Println("Total Score: ", score)
}
