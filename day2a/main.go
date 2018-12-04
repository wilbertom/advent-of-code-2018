package main

import (
	"aoc2018"
	"fmt"
)

func main() {
	box_ids := aoc2018.FirstArgLines()
	hits_2 := 0
	hits_3 := 0

	for _, box_id := range box_ids {
		letter_counts := aoc2018.LetterCount(box_id)

		if aoc2018.ContainsValue(letter_counts, 2) {
			hits_2 += 1
		}

		if aoc2018.ContainsValue(letter_counts, 3) {
			hits_3 += 1
		}
	}

	checksum := hits_2 * hits_3

	fmt.Printf("Checksum: %d\n", checksum)
}
