package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"../../aoc"
)

func main() {
	aoc.Run(2018, 4, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	duty := parseInput(input)
	var id, max int
	for gid, sleep := range duty {
		v := sleep.MinutesAsleep()
		if v > max {
			max = v
			id = gid
		}
	}
	// now we have the id, find the most often asleep minute.
	min := duty[id].MinuteMostAsleep()
	return fmt.Sprintf("%d", min*id)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	duty := parseInput(input)
	var min, max, id int
	for gid, sleep := range duty {
		m := sleep.MinuteMostAsleep()
		if sleep[m] > max {
			max = sleep[m]
			id = gid
			min = m
		}
	}
	return fmt.Sprintf("%d", min*id)
}

// map minute => AWAKE or ASLEEP
type sleepMap map[int]int

func (s sleepMap) MinutesAsleep() int {
	sum := 0
	for _, c := range s {
		sum += c
	}
	return sum
}

func (s sleepMap) MinuteMostAsleep() int {
	var min, max int
	for m, c := range s {
		if c > max {
			min = m
			max = c
		}
	}
	return min
}

//Mon Jan 2 15:04:05 -0700 MST 2006
const timeformat = "[2006-01-02 15:04]"

func parseInput(input string) map[int]sleepMap {
	// sort lines first.
	lines := strings.Split(input, "\n")
	sort.Strings(lines)

	guards := map[int]sleepMap{}
	var sleep int
	var date time.Time
	var id int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		// 01234567890123456789012345678
		// [1518-03-07 00:04] Guard #1823 begins shift
		date, _ = time.Parse(timeformat, line[:18])
		switch line[19] {
		case 'G': // guard begins shift
			fmt.Sscanf(line[25:len(line)], "#%d", &id)
			if _, ok := guards[id]; !ok {
				guards[id] = sleepMap{}
			}
		case 'f': // falls asleep
			sleep = date.Minute()
		case 'w': // wakes up
			for i := sleep; i < date.Minute(); i++ {
				guards[id][i] = guards[id][i] + 1
			}
		}
	}
	return guards
}
