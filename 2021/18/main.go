package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 18, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	list := parseNumbers(input)
	//fmt.Println(list)
	// add them all.
	sum := list[0]
	for i := 1; i < len(list); i++ {
		sum = sum.Add(list[i])
	}
	//fmt.Println("sum", sum)
	m := sum.Magnitude()
	return fmt.Sprintf("%d", m)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	list := parseNumbers(input)
	//fmt.Println(list)
	// we need to sum them in both directions and
	// find the largest magnitude.
	m := 0
	var a, b *Snailfish
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			a, b = list[i].Clone(nil), list[j].Clone(nil)
			sm := a.Add(b).Magnitude()
			if sm > m {
				m = sm
			}
			a, b = list[i].Clone(nil), list[j].Clone(nil)
			sm = b.Add(a).Magnitude()
			if sm > m {
				m = sm
			}
		}
	}

	return fmt.Sprint(m)
}

func parseNumbers(input string) []*Snailfish {
	list := []*Snailfish{}
	aoc.MapLines(input, func(line string) error {
		n, _ := parseSnailFish(line, 1)
		list = append(list, n)
		return nil
	})
	return list
}

func parseSnailFish(s string, offset int) (*Snailfish, int) {
	// the numbers are single digits so we can parse character by character.
	// this is recursive to allow keeping state more easily, each returns
	// a single snailfish number.
	sf := &Snailfish{}
	isLHS := true
	i := offset
	l := len(s)
	for {
		if i >= l {
			break
		}
		c := s[i]
		i++
		switch c {
		case '[':
			// start of a new pair.
			inner, n := parseSnailFish(s, i)
			i = n
			inner.Parent = sf
			if isLHS {
				sf.LeftFish = inner
			} else {
				sf.RightFish = inner
			}
		case ']':
			// end of a pair.
			return sf, i
		case ',':
			// switch to right hand of pair
			isLHS = false
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// a regular number
			if isLHS {
				sf.LeftVal = int(c - '0')
			} else {
				sf.RightVal = int(c - '0')
			}
		default:
			// ignore...
			panic("unexpected character! => " + string([]byte{c}))
		}
	}
	panic("bad parse")
}

// the number is represented by a left/right pair.
// each left/right is also a snailfish number or possibly
// a regular number.
// a `nil` *Fish means a regular number
// we also have a parent ref for backtracking to find
// adjacent regular numbers.
type Snailfish struct {
	Parent *Snailfish

	LeftFish *Snailfish
	LeftVal  int

	RightFish *Snailfish
	RightVal  int
}

func (sf *Snailfish) Clone(parent *Snailfish) *Snailfish {
	if sf == nil {
		return nil
	}
	clone := &Snailfish{
		Parent:   parent,
		LeftVal:  sf.LeftVal,
		RightVal: sf.RightVal,
	}
	clone.LeftFish = sf.LeftFish.Clone(clone)
	clone.RightFish = sf.RightFish.Clone(clone)
	return clone
}

func (sf *Snailfish) String() string {
	sb := strings.Builder{}
	sf.AppendTo(&sb)
	return sb.String()
}

func (sf *Snailfish) AppendTo(sb *strings.Builder) {
	sb.WriteByte('[')
	if sf.LeftFish == nil {
		sb.WriteByte(byte(int('0') + sf.LeftVal))
	} else {
		sf.LeftFish.AppendTo(sb)
	}
	sb.WriteByte(',')
	if sf.RightFish == nil {
		sb.WriteByte(byte(int('0') + sf.RightVal))
	} else {
		sf.RightFish.AppendTo(sb)
	}
	sb.WriteByte(']')
}

func (sf *Snailfish) IsLeft() bool {
	if sf.Parent == nil {
		panic("tried to call IsLeft() on top level pair")
	}
	return sf.Parent.LeftFish == sf
}

func (sf *Snailfish) IsRight() bool {
	if sf.Parent == nil {
		panic("tried to call IsRight() on top level pair")
	}
	return sf.Parent.RightFish == sf
}

func (sf *Snailfish) Magnitude() int {
	l, r := 0, 0
	if sf.LeftFish == nil {
		l = sf.LeftVal
	} else {
		l = sf.LeftFish.Magnitude()
	}
	if sf.RightFish == nil {
		r = sf.RightVal
	} else {
		r = sf.RightFish.Magnitude()
	}
	return 3*l + 2*r
}

func (sf *Snailfish) Add(n *Snailfish) *Snailfish {
	//fmt.Printf("  %s\n+ %s\n", sf, n)
	result := &Snailfish{LeftFish: sf, RightFish: n}
	// keep the parent refs intact
	sf.Parent, n.Parent = result, result
	// reduce as needed
	result.Reduce()

	//fmt.Printf("= %s\n", result)

	return result
}

func (sf *Snailfish) Reduce() {
	i := 0
	for {
		i++
		if i > 10000 {
			panic("reduction probably fubar'd")
		}

		// if we can find a pair to explode, do that.
		if sf.tryExplode(1) {
			// we exploded a pair, return to top of loop
			continue
		}
		// if we can find a pair to split, do that
		if sf.trySplit() {
			// we made a split, return to top of loop
			continue
		}
		// nothing happened, we are done.
		return
	}
}

// attempt to recursively find a split (using a depth first
// left biased search)
func (sf *Snailfish) trySplit() bool {
	// attempt a left split first
	if sf.LeftFish == nil {
		if sf.LeftVal > 9 {
			x := sf.LeftVal / 2
			// this puts the "odd" value on the right
			sf.LeftFish = &Snailfish{Parent: sf, LeftVal: x, RightVal: x + sf.LeftVal%2}
			return true
		}
		// otherwise keep trying
	} else {
		// recurse to the left
		if sf.LeftFish.trySplit() {
			return true
		}
		// otherwise keep trying
	}
	// now the same on the right.
	if sf.RightFish == nil {
		if sf.RightVal > 9 {
			x := sf.RightVal / 2
			// this puts the "odd" value on the right
			sf.RightFish = &Snailfish{Parent: sf, LeftVal: x, RightVal: x + sf.RightVal%2}
			return true
		}
		// otherwise keep trying
	} else {
		// recurse to the right
		if sf.RightFish.trySplit() {
			return true
		}
		// otherwise keep trying
	}
	// nothing doing, no splits available
	return false
}

func (sf *Snailfish) tryExplode(depth int) bool {
	if depth > 4 {
		return false
	}
	// is our left pair a pair?
	if sf.LeftFish != nil {
		// are we deep enough?
		if depth == 4 {
			// explode, assume both inners are regular.
			sf.LeftFish.addToNextLeft()
			sf.LeftFish.addToNextRight()
			// replace pair with 0.
			sf.LeftFish = nil
			sf.LeftVal = 0
		}
		// not deep enough, recurse.
		if sf.LeftFish.tryExplode(depth + 1) {
			return true
		}
	}
	// now try to the right
	if sf.RightFish != nil {
		if depth == 4 {
			// explode, assume both inners are regular.
			sf.RightFish.addToNextLeft()
			sf.RightFish.addToNextRight()
			// replace pair with 0.
			sf.RightFish = nil
			sf.RightVal = 0
		}
		// not deep enough, recurse.
		if sf.RightFish.tryExplode(depth + 1) {
			return true
		}
	}
	return false
}

func (sf *Snailfish) addToNextLeft() {
	// take the LeftVal
	n := sf.LeftVal
	// and add it to the next regular value to the left.
	// this means traversing up until we are a right hand element,
	// then going left and if needed coming back down the right
	// hand from that left point.
	curr := sf
	for {
		if curr.Parent == nil {
			// fine, we hit the top
			return
		}
		if curr.IsRight() {
			if curr.Parent.LeftFish == nil {
				// the left is regular.
				curr.Parent.LeftVal += n
				return
			} else {
				// we need to go back down the right...
				curr = curr.Parent.LeftFish
				for {
					if curr.RightFish == nil {
						// this is it
						curr.RightVal += n
						return
					}
					curr = curr.RightFish
				}
			}
		}
		curr = curr.Parent
	}
}

func (sf *Snailfish) addToNextRight() {
	// take the RightVal
	n := sf.RightVal
	// and add it to the next regular value to the right.
	// this means traversing up until we are a left hand element,
	// then going right and if needed coming back down the left
	// hand from that right point.
	curr := sf
	for {
		if curr.Parent == nil {
			// fine, we hit the top
			return
		}
		if curr.IsLeft() {
			if curr.Parent.RightFish == nil {
				// the right is regular.
				curr.Parent.RightVal += n
				return
			} else {
				// we need to go back down the left...
				curr = curr.Parent.RightFish
				for {
					if curr.LeftFish == nil {
						// this is it
						curr.LeftVal += n
						return
					}
					curr = curr.LeftFish
				}
			}
		}
		curr = curr.Parent
	}
}
