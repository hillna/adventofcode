package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {

	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		index := len(text) / 2

		container1 := text[:index]
		container2 := text[index:]

		for x := 0; x < len(container1); x++ {
			item := string(container1[x])

			if strings.Index(container2, item) > -1 {
				total += strings.Index(priorities, item) + 1
				break
			}
		}
	}

	fmt.Println("Total:", total)
}
