package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 12, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return solve(input, parseSpringsFolded)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return solve(input, parseSpringsUnfolded)
}

func solve(input string, parser func(l string) *Springs) string {
	all := []*Springs{}

	aoc.MapLines(input, func(line string) error {
		all = append(all, parser(line))
		return nil
	})

	sum := 0
	for i, s := range all {
		_ = i
		//fmt.Printf("[%03d] %s (%v) => ", i+1, s.layout, s.arrangements)
		x := s.CalculateArrangments4()
		//fmt.Printf("%d\n", x)
		sum += x
	}

	return fmt.Sprint(sum)
}

// in our examples, any runs of empties can be ignored
// but this optimisation means almost nothing in the latest verion...
//var manyDots = regexp.MustCompile(`[.]+`)

func parseSpringsFolded(line string) *Springs {
	s := strings.Split(line, " ")
	return parseSprings(s[0], s[1])
}

func x5(str string, joiner string) string {
	str += joiner
	str = strings.Repeat(str, 5)
	str = strings.TrimSuffix(str, joiner)
	return str
}

func parseSpringsUnfolded(line string) *Springs {
	s := strings.Split(line, " ")
	return parseSprings(x5(s[0], "?"), x5(s[1], ","))
}

func parseSprings(layout, arrangements string) *Springs {
	sp := &Springs{
		//		layout:       manyDots.ReplaceAllLiteralString(layout, "."),
		layout:       layout,
		arrangements: aoc.ToIntSlice(arrangements, ','),
		// springLeftCache:   map[int]int{},
		// possibleLeftCache: map[int]int{},
	}
	// sp.countOptions = strings.Count(sp.layout, "?")
	// sp.countSprings = strings.Count(sp.layout, "#")
	// for _, x := range sp.arrangements {
	// 	sp.targetSprings += x
	// }
	// left := sp.countSprings
	// possible := sp.countOptions
	// for i, x := range sp.layout {
	// 	if x == SPRING {
	// 		left--
	// 	}
	// 	if x == BROKE {
	// 		possible--
	// 	}
	// 	sp.springLeftCache[i] = left
	// 	sp.possibleLeftCache[i] = left + possible
	// }

	return sp
}

type Springs struct {
	layout       string
	arrangements []int
	// countSprings      int
	// countOptions      int
	// targetSprings     int
	// springLeftCache   map[int]int
	// possibleLeftCache map[int]int
}

// func (s *Springs) SpringsLeft(i int) int {
// 	return s.springLeftCache[i]
// }
// func (s *Springs) PossibleSpringsLeft(i int) int {
// 	return s.possibleLeftCache[i]
// }

// func (s *Springs) CalculateArrangments() int {
// 	curr := []PossibleState{{}}
// 	sum := 0
// 	fmt.Printf("%#v\n", s)
// 	for i, r := range s.layout {
// 		//fmt.Printf("Rune %c (%d) of %s\n", r, i, s.layout)
// 		if len(curr) == 0 {
// 			return sum
// 		}
// 		next := []PossibleState{}
// 		for _, c := range curr {
// 			for _, n := range c.Next(r, i, s, false) {
// 				fmt.Printf("%#v\n", n)
// 				if n.accepted {
// 					sum++
// 				} else {
// 					next = append(next, n)
// 				}
// 			}
// 		}
// 		curr = next
// 	}
// 	return sum
// }

const (
	EMTPY_S = "."
	EMPTY   = '.'
	SPRING  = '#'
	OPTION  = '?'
)

// the way I will solve this is to take the starting point
// and make every possible choice for the "next" state
// by choosing something in the layout. Then we will discard impossible
// states, and iterate on from the states we have. The number of valid
// states at the end will be our count.
//
// We track how far in to the string we are. This is "global" as we will
// tick one index at a time, so not stored here
//
// We will track which of the "arrangements" we are going for, and
// how long a run of "springs" we have found.
// we also need a "last was spring" boolean to
// indicate whether the piece before was a spring as we cannot
// put the runs together
// type PossibleState struct {
// 	output           string
// 	arrangementIndex int
// 	currentRun       int
// 	prevWasSpring    bool
// 	springsPlaced    int
// 	optionsUsed      int
// 	springsSeen      int
// 	accepted         bool // this means the state is done. good and no more branches
// }

// func (ps PossibleState) Complete(s *Springs) bool {
// 	return ps.arrangementIndex == len(s.arrangements)
// }

// var EmptyStates = []PossibleState{}

// func (ps PossibleState) Next(r rune, i int, s *Springs, fake bool) []PossibleState {
// 	// Now check the length of the arrangement we are going for
// 	// We should not hit this out of bounds.
// 	wanted := s.arrangements[ps.arrangementIndex]

// 	if fake {
// 		ps.optionsUsed++
// 	}

// 	switch r {
// 	case EMPTY:
// 		// nothing to do
// 		ps.output += "."
// 		ps.prevWasSpring = false
// 		ps.currentRun = 0
// 		return []PossibleState{ps}
// 	case SPRING:
// 		ps.output += "#"
// 		ps.currentRun++
// 		ps.prevWasSpring = true
// 		ps.springsPlaced++
// 		if !fake {
// 			ps.springsSeen++ // these are the actual springs
// 		}
// 		// a spring. We must add this one, so it could invalidate our run
// 		// does this "finish a run?"
// 		if ps.currentRun == wanted {
// 			fmt.Printf("(%v) current run looks good wanted=%d currentRun=%d => %s\n", s.arrangements, wanted, ps.currentRun, ps.output)
// 			// check the next is not a spring.
// 			if len(s.layout)-1 < i {
// 				if s.layout[i+1] == SPRING {
// 					// bad
// 					return EmptyStates
// 				}
// 			}
// 			//good
// 			ps.arrangementIndex++

// 			// do we have enough possible? springs left?
// 			remainingForcedSprings := s.countSprings - ps.springsSeen
// 			remainingOptionalSprings := s.countOptions - ps.optionsUsed
// 			remainingWantedSprings := s.targetSprings - ps.springsPlaced

// 			if remainingWantedSprings < remainingForcedSprings {
// 				// bad
// 				return EmptyStates
// 			}
// 			if remainingWantedSprings > remainingForcedSprings+remainingOptionalSprings {
// 				// also bad
// 				return EmptyStates
// 			}

// 			//fmt.Printf("is this the last arrangement satisfied? index=%d, len=%d => index == len => %v", ps.arrangementIndex, len(s.arrangements), ps.arrangementIndex == len(s.arrangements))
// 			if ps.arrangementIndex == len(s.arrangements) {
// 				// this should be the last one!
// 				fmt.Println("Have we place the right number of springs?", ps.springsPlaced == s.targetSprings)
// 				if ps.springsPlaced != s.targetSprings {
// 					// bad
// 					return EmptyStates
// 				}
// 				// place the right amount how many more "forced" springs do we have?
// 				fmt.Println("Are there any springs left we _must_ add?", s.countSprings-ps.springsSeen > 0)
// 				if s.countSprings-ps.springsSeen > 0 {
// 					// bad.
// 					return EmptyStates
// 				}
// 				fmt.Println("We have a _good_ solution!")
// 				// ok, so we are accepted now.
// 				ps.accepted = true
// 			}
// 			return []PossibleState{ps}
// 		}

// 		if ps.currentRun > wanted {
// 			// adding this spring has broken the run (it will be too long) // shoud not reach here
// 			//fmt.Printf("bad state for arrangments (%v, idx=%d): %s\n", arrangements, ps.arrangementIndex, string(ps.output))
// 			return EmptyStates
// 		}

// 		// // this can only be the start of a run, if the last was not a spring
// 		// if ps.currentRun == 0 && ps.prevWasSpring {
// 		// 	// nope, it was a spring
// 		// 	//fmt.Printf("bad state for arrangments (%v): %s\n", arrangements, string(ps.output))
// 		// 	return []PossibleState{}
// 		// }

// 		// otherwise we could start (or continue) a run
// 		return []PossibleState{ps}
// 	case BROKE:
// 		// a possible spring. here we consider both options (but we can cheat and re-use our code)
// 		possible := []PossibleState{}
// 		possible = append(possible, ps.Next(EMPTY, i, s, true)...)
// 		possible = append(possible, ps.Next(SPRING, i, s, true)...)
// 		return possible
// 	default:
// 		// something odd
// 		panic("bad rune in string")
// 	}
// }

// // I think I overthought that.
// // Let us try something simpler
// // instead of this complex ephemeral state machine which is too slow for the long strings, we will track the actual output
// type State struct {
// 	Buf         string
// 	Arrangement []int
// }

// func (s *State) Append(r rune) *State {
// 	// this will be EMPTY or SPRING, never an "option"
// 	next := &State{
// 		Buf:         s.Buf + string([]byte{byte(r)}),
// 		Arrangement: make([]int, len(s.Arrangement)),
// 	}
// 	copy(next.Arrangement, s.Arrangement)

// 	if len(s.Buf) == 0 {
// 		// first element, was it a spring?
// 		if r == SPRING {
// 			// we now have 1 spring
// 			next.Arrangement = append(next.Arrangement, 1)
// 		} else {
// 			// just one gap.
// 		}
// 	} else {
// 		// buf len > 0 so we need to look at the prev element
// 		prev := s.Buf[len(s.Buf)-1]
// 		if prev == SPRING {
// 			if r == EMPTY {
// 				// nothing
// 			} else {
// 				// another spring, increase the current arrangement value
// 				next.Arrangement[len(next.Arrangement)-1]++

// 			}
// 		} else {
// 			// prev was empty
// 			if r == EMPTY {
// 				// nothing
// 			} else {
// 				// a "new" this is a "new" arrangement
// 				next.Arrangement = append(next.Arrangement, 1)

// 			}
// 		}
// 	}
// 	//next.Value = next.value()
// 	return next
// }

// func (s *State) Validate(p *Springs) (valid bool, complete bool) {
// 	ls, lp := len(s.Arrangement), len(p.arrangements)
// 	if ls == 0 {
// 		// still OK, but not complete.
// 		return true, false
// 	}
// 	// too many springs
// 	if ls > lp {
// 		return false, false
// 	}

// 	// all of _our_ arrangements should be equal to the
// 	// wanted, except the LAST one, which can be less.
// 	ourSprings := 0
// 	finalMatch := false
// 	for i, sv := range s.Arrangement {
// 		ourSprings += sv
// 		sp := p.arrangements[i]
// 		if i == ls-1 {
// 			// the last one
// 			if sv > sp {
// 				return false, false
// 			}
// 			finalMatch = sv == sp
// 		} else if sv != sp {
// 			return false, false
// 		}
// 	}

// 	if ls == lp && finalMatch {
// 		// all matched. so, are there any springs left over?
// 		if strings.ContainsRune(p.layout[len(s.Buf):], SPRING) {
// 			// yes. so this is invalid
// 			return false, false
// 		}
// 		// no. perfect, valid and complete.
// 		return true, true
// 	} else {
// 		// we have not got everything yet,
// 		// we could check for "enough springs left" and discount
// 		// do we have enough length to get the remaining springs.
// 		return true, false
// 	}
// }

// func (s *Springs) CalculateArrangments2() int {
// 	curr := []*State{{}}
// 	var next []*State
// 	var valid, complete bool
// 	var c1, c2 *State
// 	found := 0
// 	for _, r := range s.layout {
// 		next = []*State{}
// 		for _, c := range curr {
// 			switch r {
// 			case EMPTY, SPRING:
// 				c = c.Append(r)
// 				valid, complete = c.Validate(s)
// 				if complete {
// 					found++
// 				} else if valid {
// 					next = append(next, c)
// 				}
// 			case BROKE:
// 				// do both
// 				c1 = c.Append(EMPTY)
// 				valid, complete = c1.Validate(s)
// 				if complete {
// 					found++
// 				} else if valid {
// 					next = append(next, c1)
// 				}
// 				c2 = c.Append(SPRING)
// 				valid, complete = c2.Validate(s)
// 				if complete {
// 					found++
// 				} else if valid {
// 					next = append(next, c2)
// 				}
// 			}
// 		}
// 		if len(next) == 0 {
// 			return found
// 		}
// 		curr = next
// 	}
// 	return found
// }

// // nope still too slow.
// // works correctly but the program blows up.
// // maybe go back to the first version, which didn't
// // keep history, but use the method from the
// // second, where the outer calculate function does
// // the "possibilities"

// func (s *Springs) CalculateArrangments3() int {
// 	curr := []*State3{{}}
// 	var next []*State3
// 	var valid, complete bool
// 	found := 0
// 	update := func(c *State3, r rune, i int) {
// 		//fmt.Printf("from state: %#v with rune: %c (index=%d)\n", c, r, i)
// 		c, valid, complete = c.Append(r, i, s)
// 		//fmt.Printf("to state: %#v (valid=%v, complete=%v)\n", c, valid, complete)
// 		if complete {
// 			found++
// 		} else if valid {
// 			next = append(next, c)
// 		}
// 	}
// 	for i, r := range s.layout {
// 		next = []*State3{}
// 		for _, c := range curr {
// 			switch r {
// 			case EMPTY, SPRING:
// 				update(c, r, i)
// 			case BROKE:
// 				// do both
// 				update(c, EMPTY, i)
// 				update(c, SPRING, i)
// 			}
// 		}
// 		if len(next) == 0 {
// 			return found
// 		}
// 		curr = next
// 	}
// 	return found
// }

// type State3 struct {
// 	arrIndex         int // index into the arrangement array
// 	runSprings       int // current consecutive run
// 	prevCompletedSet bool
// 	springsUsed      int
// }

// func (s *State3) Append(r rune, i int, p *Springs) (next *State3, valid, complete bool) {

// 	switch r {
// 	case SPRING:
// 		// a spring.
// 		if s.prevCompletedSet {
// 			// 2 consecutive sets. bad
// 			return nil, false, false
// 		}
// 		x := *s
// 		next = &x // copy of s
// 		next.prevCompletedSet = false
// 		next.runSprings++
// 		next.springsUsed++
// 		// does this complete a run?
// 		if next.runSprings == p.arrangements[next.arrIndex] {
// 			// yes.
// 			next.arrIndex++
// 			next.runSprings = 0
// 			next.prevCompletedSet = true
// 			// this means we completed a set just now.
// 			// was it the final one?
// 			if next.arrIndex == len(p.arrangements) {
// 				// yes.
// 				// are there any springs left?
// 				//fmt.Printf("any springs left? %s\n", p.layout[i+1:])
// 				if strings.ContainsRune(p.layout[i+1:], SPRING) {
// 					// bad
// 					return nil, false, false
// 				}
// 				// no all good!
// 				return next, true, true
// 			}
// 		}
// 	case EMPTY:
// 		// an empty space.
// 		if s.runSprings > 0 {
// 			// we didn't match the arrangement
// 			return nil, false, false
// 		}
// 		x := *s
// 		next = &x // copy of s
// 		next.prevCompletedSet = false
// 	}
// 	// we didn't complete the set. can we still do it? assume so
// 	// we need to strip the "not possibles" here.
// 	// perhaps a cache of "springs left from index i"?
// 	springsLeft := p.targetSprings - next.springsUsed
// 	if springsLeft > p.PossibleSpringsLeft(i) || springsLeft < p.SpringsLeft(i) {
// 		// springsLeft > possible - we cannot add that many springs
// 		// springsLeft < p.SpringsLeft(i) - we will have to add more springs than we want
// 		return nil, false, false
// 	}
// 	return next, true, false

// }

// nope. still doesn't work - we cannot do it brute force.
// perhaps there is another way
// this is way to use cached dynamic programming to
// iterate. because we use a massively recursive algorithm
// and cache all the intermediate lookups we don't have
// to explode the search space like in my other 3 methods.
//
// turns out this is the right way. it runs insanely fast (in comparison)
func (s *Springs) CalculateArrangments4() int {

	cache := make(map[[2]int]int)

	text := s.layout
	groups := s.arrangements

	var step func(strIdx, groupIdx int) int
	step = func(strIdx, groupIdx int) int {
		if strIdx >= len(text) {
			// we got to the end!
			// if we satisfied all groups, then return 1, else 0
			// all the other calls use this recursively, so they will add up
			if groupIdx < len(groups) {
				return 0
			}
			return 1
		}
		// check for a cached value at this point/group
		if c, ok := cache[[2]int{strIdx, groupIdx}]; ok {
			return c
		}
		res := 0
		if text[strIdx] == EMPTY {
			// move on (we can ignore an empty space) the value for this index will be the same as the value for the "next" index
			res = step(strIdx+1, groupIdx)
		} else {
			if text[strIdx] == OPTION {
				// this treats the ? as a '.'
				// the code _after_ this if block treats it as #
				res += step(strIdx+1, groupIdx)

			}
			if groupIdx < len(groups) {
				count := 0 // how many #'s (or possible ones) in a row?
				for k := strIdx; k < len(text); k++ {
					if count > groups[groupIdx] || // we have overrun run the group we are aiming for
						text[k] == EMPTY || // there is an empty space at this point (the other option is # or ? which are valid)
						count == groups[groupIdx] && text[k] == OPTION { // we have the right number for this group and the current place is a ? (becuase we checked for . and we don't want this to be a #)
						break
					}
					count++
				}
				if count == groups[groupIdx] {
					if count+strIdx < len(text) {
						// we got a group, so lets check for the next group
						res += step(count+strIdx+1, groupIdx+1)
					} else {
						// string is too long - move on
						res += step(count+strIdx, groupIdx+1)
					}
				}
			}
		}
		cache[[2]int{strIdx, groupIdx}] = res
		return res
	}

	return step(0, 0)
}
