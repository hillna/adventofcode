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

	Exit:
		for x := elf1Start; x <= elf1End; x++ {
			for y := elf2Start; y <= elf2End; y++ {
				if x == y {
					total++
					break Exit
				}
			}

		}
	}

	fmt.Println("Total:", total)
}
