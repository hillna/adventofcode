package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	stacks := [][]string{
		{"C", "Z", "N", "B", "M", "W", "Q", "V"},
		{"H", "Z", "R", "W", "C", "B"},
		{"F", "Q", "R", "J"},
		{"Z", "S", "W", "H", "F", "N", "M", "T"},
		{"G", "F", "W", "L", "N", "Q", "P"},
		{"L", "P", "W"},
		{"V", "B", "D", "R", "G", "C", "Q", "J"},
		{"Z", "Q", "N", "B", "W"},
		{"H", "L", "F", "C", "G", "T", "J"},
	}

	fmt.Println("Stacks (original):", stacks)

	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var positions []string
	for scanner.Scan() {
		instructions := strings.Split(scanner.Text(), " ")

		count, _ := strconv.Atoi(instructions[1])

		fromStack, _ := strconv.Atoi(instructions[3])
		fromStack--

		toStack, _ := strconv.Atoi(instructions[5])
		toStack--

		for i := 0; i < count; i++ {
			stacks[toStack] = append(stacks[toStack], stacks[fromStack][len(stacks[fromStack])-1])

			stacks[fromStack] = stacks[fromStack][:len(stacks[fromStack])-1]
		}
	}

	for i := 0; i < len(stacks); i++ {
		positions = append(positions, stacks[i][len(stacks[i])-1])
	}

	fmt.Println("Stacks (updated):", stacks)
	fmt.Println("Positions:", positions)
}
