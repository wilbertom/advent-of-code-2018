package main

import (
	"aoc2018"
	"fmt"
)

func main() {
	box_ids := aoc2018.FirstArgLines()

	for _, box_id_x := range box_ids {
		for _, box_id_y := range box_ids {
			if box_id_x == box_id_y {
				continue
			}

			differences := 0
			common := ""

			for index, x_letter := range box_id_x {
				y_letter := box_id_y[index]

				if string(x_letter) != string(y_letter) {
					differences += 1
				} else {
					common += string(x_letter)
				}
			}

			if differences == 1 {
				fmt.Printf("Box %s and box %s are the same product with %s in common\n", box_id_x, box_id_y, common)
			}
		}
	}
}
