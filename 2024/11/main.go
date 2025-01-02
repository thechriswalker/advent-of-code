package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 11, solve1, solve2)
}

func solve1(input string) string {
	_ = solve1_slow
	// return solve1_slow(input)
	_ = solve1_fast
	//return solve1_fast(input)
	return solve1_map(input)
}

// Implement Solution to Problem 1
func solve1_slow(input string) string {
	head := parse(input)
	head = blinkN(head, 25)
	return fmt.Sprint(head.Len())
}

func solve1_fast(input string) string {
	ints := aoc.ToIntSlice(input, ' ')
	// nope the naive approach is too slow.
	//head = blinkN(head, 75)
	// we need to be smarter. The order is unimportant, then the split can be an append.
	// create a massive list to allocate the memory up front.
	out := make([]int, 0, 250_000)
	out = append(out, ints...)

	for range 25 {
		out = unorderedBlink(out)
	}

	return fmt.Sprint(len(out))
}

func solve1_map(input string) string {
	ints := aoc.ToIntSlice(input, ' ')
	return fmt.Sprint(mapSolve(ints, 25))
}

func parse(input string) *Stone {
	// obviously a linked list is the way to go.
	// stones will link to each other, in the first problem we don't care about order, but I guarrantee we will in the second.
	// well I was wrong, and this is the slowest possible implementation....
	var head *Stone
	var curr *Stone
	ints := aoc.ToUint64Slice(input, ' ')
	for _, n := range ints {
		if head == nil {
			head = &Stone{Value: n}
			curr = head
		} else {
			n := &Stone{Value: n, Prev: curr}
			curr.Next = n
			curr = n
		}
	}
	return head
}

func blinkN(head *Stone, n int) *Stone {
	curr := head
	for i := range n {
		_ = i
		// blink all the stones.
		// check our head is still the head (a stone may have been added to the left)
		for head.Prev != nil {
			head = head.Prev
		}
		curr = head
		inc := 0
		for curr != nil {
			inc += curr.Blink()
			curr = curr.Next
		}
		// if i < 5 {
		// 	fmt.Println(head.String())
		// }
		fmt.Println("Blinked", i, "times, added", inc, "stones, num stones", head.Len())
	}

	// ensure the head is rewound.
	for head.Prev != nil {
		head = head.Prev
	}
	return head
}

// Implement Solution to Problem 2
func solve2(input string) string {
	ints := aoc.ToIntSlice(input, ' ')
	// nope the naive approach is too slow.
	//head = blinkN(head, 75)
	// we need to be smarter. The order is unimportant, then the split can be an append.
	// create a massive list to allocate the memory up front.
	// out := make([]int, 0, 1_000_000)
	// out = append(out, ints...)
	// for range 75 {
	_ = unorderedBlink

	// 	out = unorderedBlink(out)
	// }

	// still too slow. There must be a pattern.
	// i.e. we must have the "same" number" a number of times.
	// i.e. we can compact our ints, to a map of int -> count
	// and increment the count each time.
	// then we can just "blink" each number once.

	// 5600371258628357520 is too high
	// 276661131175807 - is correct
	return fmt.Sprint(mapSolve(ints, 75))
}

func mapSolve(ints []int, blinks int) int {
	curr := make(map[int]int, 1_000_000)
	next := make(map[int]int, 1_000_000)
	for _, n := range ints {
		curr[n] = curr[n] + 1
	}
	for range blinks {
		mapBlink(curr, next)
		next, curr = curr, next
	}
	// now count
	count := 0
	for _, c := range curr {
		count += c
	}
	return count
}

func mapBlink(curr, next map[int]int) {
	// need to do it into a new map, or intermediate values will be wrong.
	for n, count := range curr {
		if count == 0 {
			continue
		}
		if n == 0 {
			next[1] = next[1] + count
		} else {
			digits := int(math.Floor(math.Log10(float64(n))) + 1)
			if digits%2 == 0 {
				// even, split them in half.
				p := int(math.Pow10(digits / 2))
				right := n % p
				left := (n - right) / p
				next[right] = next[right] + count
				next[left] = next[left] + count
			} else {
				// odd, just multiply by 2024
				next[n*2024] = next[n*2024] + count
			}
		}
	}
	// now empty the current map.
	for k := range curr {
		delete(curr, k)
	}
}

// this still does a lot of work when the numbers are the same.
// should have spotted that and used a map first time round.
// but I "assumed" order would be important...
func unorderedBlink(ints []int) []int {
	l := len(ints)
	for i := 0; i < l; i++ {
		n := ints[i]
		if n == 0 {
			// update to 1
			ints[i] = 1
		} else {
			digits := int(math.Floor(math.Log10(float64(n))) + 1)
			if digits%2 == 0 {
				// even, split them in half.
				p := int(math.Pow10(digits / 2))
				right := n % p
				left := (n - right) / p
				ints[i] = right
				ints = append(ints, left)
			} else {
				// odd, just multiply by 2024
				ints[i] *= 2024
			}
		}
	}
	return ints
}

type Stone struct {
	Value uint64
	Next  *Stone
	Prev  *Stone
}

func (s *Stone) Len() int {
	for s.Prev != nil {
		s = s.Prev
	}
	count := 0
	for curr := s; curr != nil; curr = curr.Next {
		count++
	}
	return count
}

func (s *Stone) String() string {
	// rewind to the start, then print all the values.
	for s.Prev != nil {
		s = s.Prev
	}
	sb := strings.Builder{}
	for curr := s; curr != nil; curr = curr.Next {
		sb.WriteString(fmt.Sprint(curr.Value))
		if curr.Next != nil {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func (s *Stone) Blink() int {
	if s.Value == 0 {
		s.Value = 1
		return 0
	}
	// how many digits are there?
	digits := int(math.Floor(math.Log10(float64(s.Value))) + 1)
	if digits%2 == 0 {
		// this is wrong
		p := uint64(math.Pow10(digits / 2))
		// even, split them in half.
		right := s.Value % p
		left := (s.Value - right) / p
		//fmt.Println("Splitting", s.Value, "into", left, "and", right, "num digits", digits, "p", p)
		// split into 2 new stones.
		// we will add the left stone and update the current as right
		// this has the effect of meaning that the current stone's next is still the next stone
		leftStone := &Stone{Value: left, Prev: s.Prev, Next: s}
		if s.Prev != nil {
			s.Prev.Next = leftStone
		}
		s.Prev = leftStone
		s.Value = right
		return 1
	} else {
		// odd, just multiply by 2024
		s.Value *= 2024
		return 0
	}
}
