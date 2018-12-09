package main

import (
  "fmt"
  "strings"
  "aoc2018"
)

func asciiRemoveChars(s string, offset, count int) (string) {
  ns := s[:offset]
  ns += s[offset + count:]
  return ns
}

func trigger(polymer *string) *string {
  if react(polymer) {
    return trigger(polymer)
  }

  return polymer
}

func destroy(x, y string) bool {
  return x == strings.ToUpper(y) && strings.ToLower(x) == y ||
    strings.ToUpper(x) == y && x == strings.ToLower(y)
}

func react(polymer *string) bool {
  for i := 0; i < (len(*polymer) - 1); i++ {
    j := i + 1

    if destroy(string((*polymer)[i]), string((*polymer)[j])) {
      *polymer = asciiRemoveChars(*polymer, i, 2)
      return true
    }
  }

  return false
}

func dropUnit(unit string, s *string) *string {
  unitInPolarity1 := strings.ToLower(unit)
  unitInPolarity2 := strings.ToUpper(unit)
  var output string

  for i := 0; i < len(*s); i++ {
    c := string((*s)[i])
    if c != unitInPolarity1 && c != unitInPolarity2 {
      output += c
    }
  }

  return &output
}

func main() {
  lines := aoc2018.FirstArgLines()
  polymer := lines[0]

  trigger(&polymer)
  fmt.Printf("%d units left on the polymer\n", len(polymer))

  units := "abcdefghijklmnopqrstuvwxyz"

  for i := 0; i < len(units); i++ {
    unit := string(units[i])
    tweaked_polymer := dropUnit(unit, &polymer)
    trigger(tweaked_polymer)
    fmt.Printf("%d units left on the polymer after dropping %s\n", len(*tweaked_polymer), unit)
  }
}
