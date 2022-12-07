package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	count := 0
	for scanner.Scan() {
		signal := scanner.Text()

		var buffer []string

	Exit:
		for i := 0; i < len(signal); i++ {
			count++
			buffer = append(buffer, string(signal[i]))
			if len(buffer) == 14 {
				test := strings.Join(buffer, "")

				testCount := 0
				for x := 0; x < len(buffer); x++ {
					if strings.Count(test, buffer[x]) > 1 {
						testCount++
					}
				}

				if testCount == 0 {
					break Exit
				}

				buffer = buffer[1:]
			}
		}
	}

	fmt.Println("Total characters:", count)
}
