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

func ParseInt(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	exitIfError(err)
	return n
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
	exitIfError(err)
	return strings.Split(string(contents), "\n")
}

func Contains(collection []int64, element int64) bool {
	for _, c := range collection {
		if c == element {
			return true
		}
	}

	return false
}

func exitIfError(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(20)
	}
}
