package main

import (
	"aoc2018"
	"errors"
	"fmt"
)

type Claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Height int
}

func NewClaim(id, x, y, width, height int) *Claim {
	return &Claim{
		ID:     id,
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func ParseClaim(line string) (*Claim, error) {
	var id, x, y, width, height int

	n, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x, &y, &width, &height)

	if err != nil {
		return nil, err
	}

	if n != 5 {
		return nil, errors.New("Only read " + string(n) + " fields")
	}

	return NewClaim(id, x, y, width, height), nil
}

type Fabric struct {
	Width  int
	Height int
	grid   [][]int
}

func NewFabric() Fabric {
	return Fabric{
		Width:  0,
		Height: 0,
		grid:   make([][]int, 0),
	}
}

func (self *Fabric) ClaimPiece(claim *Claim) {
	padding := 1
	self.Grow(claim.X+claim.Width+padding, claim.Y+claim.Height+padding)

	for row := claim.Y; row < claim.Y+claim.Height; row++ {
		for column := claim.X; column < claim.X+claim.Width; column++ {
			self.grid[row][column]++
		}
	}
}

func (self *Fabric) Grow(width, height int) {
	if self.Width < width {
		grew_width := width - self.Width
		self.Width = width

		for i, _ := range self.grid {
			for column := 0; column < grew_width; column++ {
				self.grid[i] = append(self.grid[i], 0)
			}
		}
	}

	if self.Height < height {
		grew_height := height - self.Height
		self.Height = height

		for row := 0; row < grew_height; row++ {
			self.grid = append(self.grid, make([]int, self.Width))
		}
	}
}

func (self *Fabric) ClaimedXTimesOrMore(x int) int {
	claimed_x_times_or_more := 0

	for _, row := range self.grid {
		for _, column := range row {
			if column >= x {
				claimed_x_times_or_more++
			}
		}
	}

	return claimed_x_times_or_more
}

func (self *Fabric) Intact(claim *Claim) bool {
	for row := claim.Y; row < claim.Y+claim.Height; row++ {
		for column := claim.X; column < claim.X+claim.Width; column++ {
			if self.grid[row][column] != 1 {
				return false
			}
		}
	}

	return true
}

func (self *Fabric) Print() {
	for i, row := range self.grid {
		fmt.Printf("%4d: %v\n", i, row)
	}
}

func main() {
	lines := aoc2018.FirstArgLines()

	fabric := NewFabric()
	claims := make([]*Claim, 0)

	for _, line := range lines {
		claim, err := ParseClaim(line)
		aoc2018.ExitIfError(err)

		fabric.ClaimPiece(claim)
		claims = append(claims, claim)
	}

	x := 2
	claimed_x_times_or_more := fabric.ClaimedXTimesOrMore(x)

	fmt.Printf("%d square inches of fabric were claimed %d or more times.\n", claimed_x_times_or_more, x)

	for _, claim := range claims {
		if fabric.Intact(claim) {
			fmt.Printf("Claim %d is intact\n", claim.ID)
		}
	}
}
