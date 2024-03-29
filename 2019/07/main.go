package main

import (
	"fmt"
	"sync"

	"github.com/thechriswalker/advent-of-code/2019/intcode"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 7, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	p := intcode.New(input)
	// find the best order of the 5 number,
	// this is combinations. super naive implementation.
	max := int64(0)
	// 0,1,2,3,4
	// for a := 0; a < 5; a++ {
	// 	for b := 0; b < 5; b++ {
	// 		if a == b {
	// 			continue
	// 		}
	// 		for c := 0; c < 5; c++ {
	// 			if c == a || c == b {
	// 				continue
	// 			}
	// 			for d := 0; d < 5; d++ {
	// 				if d == a || d == b || d == c {
	// 					continue
	// 				}
	// 				for e := 0; e < 5; e++ {
	// 					if e == a || e == b || e == c || e == d {
	// 						continue
	// 					}
	// 					// test
	// 					if output := amplify(p, [5]int{a, b, c, d, e}); output > max {
	// 						max = output
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	//}
	for _, s := range combinations([]int64{4, 3, 2, 1, 0}) {
		//
		output := amplify(p, s)
		//log.Printf("order: %v: %d\n", s, output)

		if output > max {
			max = output

		}
		//return fmt.Sprintf("%d", max)
	}
	return fmt.Sprintf("%d", max)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	p := intcode.New(input)
	max := int64(0)
	for _, s := range combinations([]int64{5, 6, 7, 8, 9}) {
		if output := amplify2(p, s); output > max {
			max = output
		}
	}
	return fmt.Sprintf("%d", max)
}

// produce all combinations of a given set of ints.
// needed for problem 2, as the thing is bigger
func combinations(options []int64) [][]int64 {
	all := [][]int64{}
	var recur func(a, b []int64)
	recur = func(base, rem []int64) {
		if len(rem) == 0 {
			all = append(all, base)
		}
		//log.Printf("base: %v, remaining: %v\n", base, rem)
		for i := range rem {
			nextBase := make([]int64, len(base)+1)
			copy(nextBase, base)
			// pick i
			nextBase[len(base)] = rem[i]
			// create the next remainder.
			if len(rem) == 1 {
				recur(nextBase, []int64{})
			} else {
				nextRem := make([]int64, len(rem)-1)
				copy(nextRem, rem[:i])
				copy(nextRem[i:], rem[i+1:])
				recur(nextBase, nextRem)
			}
		}
	}
	recur([]int64{}, options)
	return all
}

func amplify(p *intcode.Program, order []int64) int64 {
	// clone before each run.
	output := int64(0)
	for _, i := range order {
		pi := p.Copy()
		//pi.Debug = true
		pi.ID = fmt.Sprintf("amp:%d", i)
		done := pi.RunAsync()
		pi.Input <- func() int64 { return i }
		pi.Input <- func() int64 { return output }
		output = pi.GetOutput()
		<-done
	}
	return output
}

// in this version we have to initialise 5 machines and keep feeding them input
// without resetting their memory...
func amplify2(p *intcode.Program, order []int64) int64 {
	amps := make([]*intcode.Program, len(order))
	var wg sync.WaitGroup

	output := int64(0)

	for i, v := range order {
		amp := p.Copy()
		amp.ID = fmt.Sprintf("amp:%d", v)
		amps[i] = amp
		amp.RunAsync()
		amp.Input <- func() int64 { return v }
		wg.Add(1)
		go func(idx int) {
			for {
				select {
				case <-amps[idx].Halted:
					wg.Done()
					return
				case out := <-amps[idx].Output:
					// capture last output
					output = out
					// each amp feeds into the next amp, unless that one halts first
					next := amps[(idx+1)%len(order)]
					// either wait for halt OR send output on.
					// we cannot do this now...
					// how we see if this is waiting for input.
					// I need a block until input used.
					select {
					case <-next.Halted:
					case next.Input <- func() int64 { return out }:
					}
				}
			}
		}(i)
	}
	amps[0].Input <- func() int64 { return 0 } // start the first amp.
	wg.Wait()
	return output
}
