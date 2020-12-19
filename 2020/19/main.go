package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2020, 19, solve1, solve2)
}

// the question is to build a regex engine, or just the regex....
// easier to build the regex. Also it is only "a" and "b" to match.

// so I will build an intermediate representation we can recursively
// collapse into an regexp string.

type rule struct {
	fragment string
	subrules [][]int
}

func (r *rule) toRegexp(all ruleset) string {
	if r.fragment == "" {
		// create from subrules.
		subs := make([]string, len(r.subrules))
		for i, sub := range r.subrules {
			stack := make([]string, len(sub))
			for j, idx := range sub {
				stack[j] = all[idx].toRegexp(all)
			}
			// this stack is all required, so we simply join them
			subs[i] = strings.Join(stack, "")
		}
		// this stack is "one of" so if > 1 wrap in non-capturing parens and |
		if len(subs) == 1 {
			r.fragment = subs[0]
		} else {
			r.fragment = `(?:` + strings.Join(subs, "|") + ")"
		}
	}
	return r.fragment
}

type ruleset map[int]*rule

func parseRuleset(input string) ruleset {
	lines := strings.Split(input, "\n")
	set := map[int]*rule{}
	var idx int
	var literal string
	for _, line := range lines {
		n, _ := fmt.Sscanf(line, `%d: "%1s"`, &idx, &literal)
		if n == 0 {
			continue
		}
		if n == 2 {
			// we have a literal
			set[idx] = &rule{fragment: literal, subrules: [][]int{}}
			continue
		}
		// we have sub-options
		options := strings.Split(strings.Split(line, ":")[1], "|")
		r := &rule{
			subrules: [][]int{},
		}
		for _, option := range options {
			// a list of ints.
			stack := []int{}
			for _, num := range strings.Split(option, " ") {
				if n, err := strconv.Atoi(strings.TrimSpace(num)); err == nil {
					stack = append(stack, n)
				}
			}
			r.subrules = append(r.subrules, stack)
		}
		set[idx] = r
	}
	return set
}

// Implement Solution to Problem 1
func solve1(input string) string {
	parts := strings.Split(input, "\n\n")
	rules := parseRuleset(parts[0])
	re := `^` + rules[0].toRegexp(rules) + `$`
	//fmt.Println(re)
	regex := regexp.MustCompile(re)
	sum := 0
	for _, msg := range strings.Split(parts[1], "\n") {
		if regex.MatchString(strings.TrimSpace(msg)) {
			sum++
		}
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	parts := strings.Split(input, "\n\n")
	rules := parseRuleset(parts[0])

	// we need to update 8 and 11 to have a new option, which is recursive
	//  8: 42     =>  8: 42 | 42 8
	// 11: 42 31  => 11: 42 31 | 42 11 31
	// note that both these rules "only" affect themselves.
	// so we cheat and work out the required fragments first!
	// process 42 31 first
	fortyTwo := rules[42].toRegexp(rules)  // this will pre-process 42
	thirtyOne := rules[31].toRegexp(rules) // this will pre-process 31
	// now we have a string for these 2 we manually construct the fragments for
	// 8 and 11.
	// `8: 42 | 42 8` means one or more 42, this is easy
	rules[8].fragment = `(?:` + fortyTwo + `)+`
	// 11: 42 31 | 42 11 31
	// this means "the same number of 42s as 31s" which I am not sure regexp can do.
	// so we will cheat and assume that it is never going to be more than 10.
	options := make([]string, 10)
	for i := 1; i <= 10; i++ {
		options[i-1] = fmt.Sprintf(`(?:(%s){%d}(%s){%d})`, fortyTwo, i, thirtyOne, i)
	}
	rules[11].fragment = `(?:` + strings.Join(options, "|") + `)`

	re := `^` + rules[0].toRegexp(rules) + `$`
	regex := regexp.MustCompile(re)
	sum := 0
	for _, msg := range strings.Split(parts[1], "\n") {
		if regex.MatchString(strings.TrimSpace(msg)) {
			sum++
		}
	}
	return fmt.Sprintf("%d", sum)
}
