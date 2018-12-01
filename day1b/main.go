package main

import (
	"aoc2018"
	"fmt"
	"os"
)

func main() {
	var frequency int64 = 0
	lines := aoc2018.FirstArgLines()
	frequencies := make(map[int64]bool)
	frequencies[frequency] = true

	for {
		for _, line := range lines {
			if aoc2018.Blank(line) {
				continue
			}

			frequency_delta := aoc2018.ParseInt(line)
			frequency += frequency_delta

			if frequencies[frequency] {
				fmt.Printf("Duplicate Frequency: %d\n", frequency)
				os.Exit(0)
			}

			frequencies[frequency] = true
		}
	}
}
