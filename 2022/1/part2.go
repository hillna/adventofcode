package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	results := []int{}

	content, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()

	// fmt.Println(string(content))

	scanner := bufio.NewScanner(content)

	total := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			results = append(results, total)
			total = 0
		} else {
			count, _ := strconv.Atoi(scanner.Text())
			total = total + count
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(results)))

	top3 := results[0] + results[1] + results[2]

	fmt.Println("Top Elf's Calories: ", results[0])
	fmt.Println("Top 3 Elves' Calories: ", top3)
}
