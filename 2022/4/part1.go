package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	total := 0
	for scanner.Scan() {
		elves := strings.Split(scanner.Text(), ",")

		elf1Start, _ := strconv.Atoi(strings.Split(elves[0], "-")[0])
		elf1End, _ := strconv.Atoi(strings.Split(elves[0], "-")[1])

		elf2Start, _ := strconv.Atoi(strings.Split(elves[1], "-")[0])
		elf2End, _ := strconv.Atoi(strings.Split(elves[1], "-")[1])

		if (elf1Start <= elf2Start) && (elf1End >= elf2End) {
			total++
		} else if (elf2Start <= elf1Start) && (elf2End >= elf1End) {
			total++
		}
	}

	fmt.Println("Total:", total)
}
