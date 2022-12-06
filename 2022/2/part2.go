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
	Loss     = "X" // +0 points
	Draw     = "Y" // +3 points
	Win      = "Z" // +6 points
)

func main() {

	var score int

	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {

		round_score := 0

		throws := strings.Split(scanner.Text(), " ")

		opponent := throws[0]
		player := throws[1]

		fmt.Println("You throw: ", player, "! Your opponent throws: ", opponent, "!")

		if player == Loss {
			if opponent == Rock {
				// Play Scissor
				round_score = round_score + 3
			} else if opponent == Paper {
				// Play Rock
				round_score = round_score + 1
			} else if opponent == Scissors {
				// Play Paper
				round_score = round_score + 2
			}
		} else if player == Draw {
			round_score = round_score + 3

			if opponent == Rock {
				// Play Rock
				round_score = round_score + 1
			} else if opponent == Paper {
				// Play Paper
				round_score = round_score + 2
			} else if opponent == Scissors {
				// Play Scissor
				round_score = round_score + 3
			}
		} else if player == Win {
			round_score = round_score + 6

			if opponent == Rock {
				// Play Paper
				round_score = round_score + 2
			} else if opponent == Paper {
				// Play Scissor
				round_score = round_score + 3
			} else if opponent == Scissors {
				// Play Rock
				round_score = round_score + 1
			}
		}

		fmt.Println("Round ", "1", " Score: ", round_score)

		score = score + round_score
	}

	fmt.Println("-------------------")
	fmt.Println("Total Score: ", score)
}
