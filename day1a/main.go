package main

import (
	"aoc2018"
	"fmt"
)

func main() {
	var frequency int64 = 0
	lines := aoc2018.FirstArgLines()

	for _, line := range lines {
		if aoc2018.Blank(line) {
			continue
		}

		frequency_delta := aoc2018.ParseInt(line)
		frequency += frequency_delta
	}

	fmt.Printf("Frequency: %d\n", frequency)
}
