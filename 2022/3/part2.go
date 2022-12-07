package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {

	input, _ := os.Open("input.txt")
	defer input.Close()
	reader := bufio.NewReader(input)

	var lines []string

	for {
		line, _ := reader.ReadString('\n')
		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		lines = append(lines, line)
	}

	total := 0
	for i := 0; i < len(lines); i += 3 {

		itemsList := lines[i]
		for x := 0; x < len(itemsList); x++ {
			item := string(itemsList[x])

			if strings.Index(lines[i+1], item) > -1 {
				if strings.Index(lines[i+2], item) > -1 {
					total += strings.Index(priorities, item) + 1
					break
				}
			}
		}
	}

	fmt.Println("Total:", total)
}
