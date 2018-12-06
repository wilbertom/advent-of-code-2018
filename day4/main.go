package main

import (
	"aoc2018"
	"fmt"
	"sort"
	"strings"
	"time"
)

type ShiftLine struct {
	timestamp                         time.Time
	line                              *string
	fallsAsleep, wakesUp, startsShift bool
	guardID                           int
}

type ShiftsParser struct{}

func NewShiftsParser() *ShiftsParser {
	return &ShiftsParser{}
}

func (self *ShiftsParser) isShiftStart(line *string) bool {
	return strings.HasSuffix(*line, "begins shift")
}

func (self *ShiftsParser) isWakesUp(line *string) bool {
	return strings.HasSuffix(*line, "wakes up")
}

func (self *ShiftsParser) isFallsAsleep(line *string) bool {
	return strings.HasSuffix(*line, "falls asleep")
}

func parseLineTime(line *string) time.Time {
	var year, month, day, hour, minute int

	_, err := fmt.Sscanf(*line, "[%d-%d-%d %d:%d] ", &year, &month, &day, &hour, &minute)
	aoc2018.ExitIfError(err)

	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
}

func (self *ShiftsParser) parserGuardID(line *string) int {
	if !self.isShiftStart(line) {
		panic("Trying to parse a guard id from a invalid line")
	}

	return aoc2018.ParseInt(strings.Split(*line, " ")[3][1:])
}

func (self *ShiftsParser) Parse(line *string) *ShiftLine {
	shiftLine := &ShiftLine{
		line:        line,
		timestamp:   parseLineTime(line),
		fallsAsleep: self.isFallsAsleep(line),
		wakesUp:     self.isWakesUp(line),
		startsShift: self.isShiftStart(line),
	}

	if shiftLine.startsShift {
		shiftLine.guardID = self.parserGuardID(line)
	} else {
		shiftLine.guardID = -1
	}

	return shiftLine
}

type Guard struct {
	id            int
	minutesAsleep map[int]int
}

func NewGuard(id int) *Guard {
	return &Guard{
		id:            id,
		minutesAsleep: make(map[int]int),
	}
}

func (self *Guard) MinutesAsleep() int {
	minutesAsleep := 0

	for _, timesAsleep := range self.minutesAsleep {
		minutesAsleep += timesAsleep
	}

	return minutesAsleep
}

func (self *Guard) MinuteMostAsleep() int {
	var minuteMostAsleep int
	var minuteMostAsleepTimes int

	for minute, timesAsleep := range self.minutesAsleep {
		if timesAsleep > minuteMostAsleepTimes {
			minuteMostAsleep = minute
			minuteMostAsleepTimes = timesAsleep
		}
	}

	return minuteMostAsleep
}

func (self *Guard) MinuteMostAsleepTimes() int {
	return self.minutesAsleep[self.MinuteMostAsleep()]
}

func LoadGuards(shiftLines *[]*ShiftLine) *map[int]*Guard {
	var guards = map[int]*Guard{}
	var guard *Guard
	var startSleepMinute int
	var wakesUpMinute int

	for _, line := range *shiftLines {
		if line.startsShift {
			if guard = guards[line.guardID]; guard == nil {
				guards[line.guardID] = NewGuard(line.guardID)
				guard = guards[line.guardID]
			}
		} else if line.fallsAsleep {
			startSleepMinute = line.timestamp.Minute()
		} else if line.wakesUp {
			wakesUpMinute = line.timestamp.Minute()

			for i := startSleepMinute; i < wakesUpMinute; i++ {
				guard.minutesAsleep[i]++
			}
		}
	}

	return &guards
}

func findMostGuard(guards *map[int]*Guard, p func(current, found *Guard) bool) *Guard {
	var found *Guard

	for _, current := range *guards {
		if found == nil {
			found = current
		}

		if p(current, found) {
			found = current
		}
	}

	return found
}

func mostSleepyGuard(guards *map[int]*Guard) *Guard {
	return findMostGuard(guards, func(current, found *Guard) bool {
		return current.MinutesAsleep() > found.MinutesAsleep()
	})
}

func mostFrequentlySleepingGuard(guards *map[int]*Guard) *Guard {
	return findMostGuard(guards, func(current, found *Guard) bool {
		return current.MinuteMostAsleepTimes() > found.MinuteMostAsleepTimes()
	})
}

func main() {
	lines := aoc2018.FirstArgLines()
	sort.Strings(lines)

	shiftLines := make([]*ShiftLine, len(lines))
	shiftsParser := NewShiftsParser()

	for i := range lines {
		shiftLines[i] = shiftsParser.Parse(&lines[i])
	}

	guards := LoadGuards(&shiftLines)
	sleepiestGuard := mostSleepyGuard(guards)
	frequentlySleepingGuard := mostFrequentlySleepingGuard(guards)

	fmt.Printf(
		"Guard #%d slept the most: %d minutes\n%d\n",
		sleepiestGuard.id,
		sleepiestGuard.MinutesAsleep(),
		sleepiestGuard.id*sleepiestGuard.MinuteMostAsleep(),
	)

	fmt.Printf(
		"Guard #%d was most frequently asleep at minute %d, %d times\n%d\n",
		frequentlySleepingGuard.id,
		frequentlySleepingGuard.MinuteMostAsleep(),
		frequentlySleepingGuard.MinuteMostAsleepTimes(),
		frequentlySleepingGuard.id*frequentlySleepingGuard.MinuteMostAsleep(),
	)
}
