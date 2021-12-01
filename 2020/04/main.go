package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 4, solve1, solve2)
}

type passport map[string]string

func (p passport) HasFields(required []string) bool {
	for _, s := range required {
		if _, ok := p[s]; !ok {
			return false
		}
	}
	return true
}

func (p passport) HasValidFields(v map[string]func(string) bool) bool {
	for key, fn := range v {
		val, ok := p[key]
		if !ok || !fn(val) {
			return false
		}
	}
	return true
}

func parsePassports(input string) []passport {
	stack := make([]passport, 0, 500) // I don't know how many there will be...
	// we kinda want to iterate over the file line by line.
	lines := strings.Split(input, "\n")
	current := make(passport)
	for _, line := range lines {
		if line == "" {
			stack = append(stack, current)
			current = make(passport)
		}
		for _, kv := range strings.Split(line, " ") {
			x := strings.Split(kv, ":")
			if len(x) == 2 {
				current[x[0]] = x[1]
			}
		}
	}
	if len(current) != 0 {
		stack = append(stack, current)
	}

	return stack
}

// Implement Solution to Problem 1
func solve1(input string) string {
	passports := parsePassports(input)
	sum := 0
	req := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, p := range passports {
		if p.HasFields(req) {
			sum++
		}
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	passports := parsePassports(input)
	sum := 0
	for _, p := range passports {
		if p.HasValidFields(validators) {
			sum++
		}
	}
	return fmt.Sprintf("%d", sum)
	return "<unsolved>"
}

var validators = map[string]func(string) bool{
	"byr": func(s string) bool {
		if len(s) != 4 {
			return false
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return i >= 1920 && i <= 2002
	},
	"iyr": func(s string) bool {
		if len(s) != 4 {
			return false
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return i >= 2010 && i <= 2020
	},
	"eyr": func(s string) bool {
		if len(s) != 4 {
			return false
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return i >= 2020 && i <= 2030
	},
	"hgt": func(s string) bool {
		if len(s) < 3 {
			return false
		}
		u := s[len(s)-2:]
		v, err := strconv.Atoi(s[:len(s)-2])
		if err != nil {
			return false
		}
		switch u {
		case "cm":
			return v >= 150 && v <= 193
		case "in":
			return v >= 59 && v <= 76
		default:
			return false
		}
	},
	"hcl": func(s string) bool {
		if len(s) != 7 {
			return false
		}
		if s[0] != '#' {
			return false
		}
		for _, c := range s[1:] {
			if c >= 'a' || c <= 'f' {
				continue
			}
			if c >= '0' || c <= '9' {
				continue
			}
			return false
		}
		return true
	},
	"ecl": func(s string) bool {
		switch s {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		default:
			return false
		}
	},
	"pid": func(s string) bool {
		if len(s) != 9 {
			return false
		}
		for _, c := range s {
			if c >= '0' || c <= '9' {
				continue
			}
			return false
		}
		return true
	},
}
