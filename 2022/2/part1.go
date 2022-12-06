package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
		//fmt.Println("Opponent: ", opponent, " Player: ", player)

		// A or X = Rock
		// B or Y = Paper
		// C or Z = Scissors

		if player == "X" {
			round_score = round_score + 1

			if opponent == "A" {
				round_score = round_score + 3
			} else if opponent == "C" {
				round_score = round_score + 6
			}
		} else if player == "Y" {
			round_score = round_score + 2

			if opponent == "A" {
				round_score = round_score + 6
			} else if opponent == "B" {
				round_score = round_score + 3
			}
		} else if player == "Z" {
			round_score = round_score + 3

			if opponent == "B" {
				round_score = round_score + 6
			} else if opponent == "C" {
				round_score = round_score + 3
			}
		}

		fmt.Println("Round ", "1", " Score: ", round_score)

		score = score + round_score
	}

	fmt.Println("-------------------")
	fmt.Println("Total Score: ", score)
}
