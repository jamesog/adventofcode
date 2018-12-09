package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var (
	reShift = regexp.MustCompile(`\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})\] Guard #(\d+) begins shift`)
	reSleep = regexp.MustCompile(`\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})\] falls asleep`)
	reWake  = regexp.MustCompile(`\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})\] wakes up`)
)

const timeFmt = "2006-01-02 15:04"

func parseTime(s string) time.Time {
	t, err := time.Parse(timeFmt, s)
	if err != nil {
		panic(err)
	}
	return t
}

type guard struct {
	ID             int
	LastSleep      time.Time
	AsleepDuration int
	Mins           map[int]int // Count each minute
}

type guards map[int]guard

func (g guard) MostCommonMin() (int, int) {
	var min, high int
	for m, c := range g.Mins {
		if c > high {
			high = c
			min = m
		}
	}
	return min, high
}

func findGuards(input []string) guards {
	// It's assumed that the input is pre-sorted
	// The input data is in a format that's easily sortable as it contains YYYY-mm-dd HH:MM
	// so it can be sorted with the Unix sort tool
	guardIDs := make(guards)
	var id int
	for _, log := range input {
		// var g guard
		// var sleep, wake time.Time
		if reShift.MatchString(log) {
			match := reShift.FindStringSubmatch(log)
			id, _ = strconv.Atoi(match[2])
			// Create a new guard entry if one doesn't exist
			if _, ok := guardIDs[id]; !ok {
				guardIDs[id] = guard{ID: id}
			}
			// fmt.Printf("Guard %d started\n", id)
		}
		if reSleep.MatchString(log) {
			match := reSleep.FindStringSubmatch(log)
			// Copy the guard entry for modification
			g := guardIDs[id]
			g.LastSleep = parseTime(match[1])
			// Put it back
			guardIDs[id] = g
		}
		if reWake.MatchString(log) {
			match := reWake.FindStringSubmatch(log)
			g := guardIDs[id]

			// Current sleep duration
			duration := int(parseTime(match[1]).Sub(g.LastSleep).Minutes())
			// Update total duration
			g.AsleepDuration += duration
			// fmt.Printf("Guard %d slept at %s, woke after %s\n", id, g.LastSleep.String(), parseTime(match[1]).Sub(g.LastSleep).String())
			if g.Mins == nil {
				g.Mins = make(map[int]int)
			}
			min := g.LastSleep.Minute()
			for i := 0; i < duration; i++ {
				g.Mins[min+i]++
			}
			guardIDs[id] = g
		}
	}

	return guardIDs
}

func laziest(gs guards) guard {
	// Map guard ID to mins slept
	mins := make(map[int]int)
	for _, g := range gs {
		mins[g.ID] = g.AsleepDuration
	}
	var guard, high int
	for id, dur := range mins {
		if dur >= high {
			high = dur
			guard = id
		}
	}
	return gs[guard]
}

func regularSleeper(gs guards) (int, int) {
	// Map guard to most common min count
	mins := make(map[int]int)
	// Map guard IDs to guards for lookup later
	guardMap := make(map[int]guard)
	for _, g := range gs {
		guardMap[g.ID] = g
		_, count := g.MostCommonMin()
		mins[g.ID] = count
	}
	var sleepy, high int
	for g, c := range mins {
		if c > high {
			high = c
			sleepy = g
		}
	}
	min, _ := guardMap[sleepy].MostCommonMin()
	return guardMap[sleepy].ID, min
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	guards := findGuards(input)
	sleepy := laziest(guards)
	mins, _ := sleepy.MostCommonMin()
	fmt.Printf("Part one: sleepy=%d, mins=%d, result=%d\n", sleepy.ID, mins, sleepy.ID*mins)
	guard, min := regularSleeper(guards)
	fmt.Printf("Part two: sleepy=%d, min=%d, result=%d\n", guard, min, guard*min)
}
