package aoc2018

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func Blank(s string) bool {
	return strings.TrimSpace(s) == ""
}

func ParseInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	ExitIfError(err)
	return n
}

func ParseInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	ExitIfError(err)
	return int(n)
}

func FirstArgContents() ([]byte, error) {
	if len(os.Args) <= 1 {
		return nil, errors.New("Missing command line argument")
	}

	filename := os.Args[1]
	return ioutil.ReadFile(filename)
}

func FirstArgLines() []string {
	contents, err := FirstArgContents()
	ExitIfError(err)

	all_lines := strings.Split(string(contents), "\n")
	lines := make([]string, 0)

	for _, line := range all_lines {
		if Blank(line) {
			continue
		}

		lines = append(lines, line)
	}

	return lines
}

func Contains(collection []int64, element int64) bool {
	for _, c := range collection {
		if c == element {
			return true
		}
	}

	return false
}

func ContainsValue(collection map[rune]int, value int) bool {
	for _, element := range collection {
		if element == value {
			return true
		}
	}

	return false
}

func LetterCount(word string) map[rune]int {
	letter_counts := make(map[rune]int)

	for _, letter := range word {
		letter_counts[letter]++
	}

	return letter_counts
}

func ExitIfError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(20)
	}
}
