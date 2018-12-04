package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 11, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	init := GetInitialPuzzleState(input)
	return fmt.Sprintf("%d", solveInt(init))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	init := GetInitialPuzzleState(input)
	// add the 2 gen/chips on the first floor
	// i.e. 4 entries all 0
	init.Items = append(init.Items, 0, 0, 0, 0)
	return fmt.Sprintf("%d", solveInt(init))
}

func solveInt(init *State) int64 {
	//	fmt.Println("initial state", init)
	cache := make(StateCache)
	cache.Add(init)
	countMoves := int64(0)
	availableNextMoves := []*State{init}

	for {
		countMoves++
		nextSteps := []*State{}
		//	fmt.Printf("Depth: %d, moves: %d\n", countMoves, len(availableNextMoves))
		for _, move := range availableNextMoves {
			//	fmt.Print(move.String())
			for _, option := range move.PossibleMoves(cache) {
				if option.IsDone() {
					//	fmt.Print(move.String())
					return countMoves
				}
				nextSteps = append(nextSteps, option)
			}
		}
		availableNextMoves = nextSteps
		if len(availableNextMoves) == 0 {
			panic("No more moves!")
		}
	}
}

const (
	FLOORS = 4
)

var matcher = regexp.MustCompile(`\b([a-z]+)( generator|-compatible)\b`)

func GetInitialPuzzleState(input string) *State {
	lines := strings.Split(input, "\n")
	pairMap := map[string][2]uint8{}

	for i, line := range lines {
		matches := matcher.FindAllStringSubmatch(line, -1)
		if matches == nil {
			continue
		}
		for _, match := range matches {
			pair, ok := pairMap[match[1]]
			if !ok {
				pair = [2]uint8{0, 0}
			}
			if match[2][0] == '-' {
				// chip
				pair[CHP] = uint8(i)
			} else {
				// generator
				pair[GEN] = uint8(i)
			}
			pairMap[match[1]] = pair
		}
	}

	state := make([]uint8, 2*len(pairMap))
	i := 0
	for _, pair := range pairMap {
		state[i] = pair[GEN]
		state[i+1] = pair[CHP]
		i += 2
	}

	return &State{Items: state}
}

type StateCache map[string]struct{}

func (sc StateCache) Add(s *State) (wasNew bool) {
	h := hash(s)
	_, old := sc[h]
	if old {
		return false
	}
	sc[h] = struct{}{}
	return true
}

// State is the slice of pairs of generators and chips
// the 0 index is the generator, the 1 index is the chip
type State struct {
	Lift   uint8
	Items  []uint8
	sorted bool
}

func (s *State) String() string {
	// pretty print
	floors := [FLOORS][]rune{}
	chars := []rune{'G', 'M'}
	for i := uint8(0); i < FLOORS; i++ {
		floors[i] = make([]rune, len(s.Items))
		for j, f := range s.Items {
			if f == i {
				floors[i][j] = chars[j%2]
			} else {
				floors[i][j] = ' '
			}
		}
	}
	b := &strings.Builder{}
	var writeFloor = func(f uint8) {
		fmt.Fprintf(b, "F%d |", f+1)
		if f == s.Lift {
			b.WriteRune('E')
		} else {
			b.WriteRune(' ')
		}
		b.WriteRune('|')

		for i, c := range floors[f] {
			if i%2 == 0 {
				b.WriteRune(' ')
			}
			b.WriteRune(c)
		}
		b.WriteRune('\n')
	}
	b.WriteRune('\n')
	for i := uint8(FLOORS); i > 0; i-- {
		writeFloor(i - 1)
	}
	return b.String()
}

func (s *State) Clone() *State {
	return &State{
		Lift:  s.Lift,
		Items: append(make([]uint8, 0, len(s.Items)), s.Items...),
	}
}

var hasher = &strings.Builder{}

// assumes the state is sorted by floor
func hash(s *State) string {
	s.Sort()
	hasher.Reset()
	fmt.Fprintf(hasher, "%d", s.Lift)
	for i := 0; i < len(s.Items); i++ {
		fmt.Fprintf(hasher, "%d", s.Items[i])
	}
	return hasher.String()
}

const (
	GEN = 0
	CHP = 1
)

func (s *State) Sort() {
	if !s.sorted {
		sort.Sort(s)
		s.sorted = true
	}
}

func (s *State) Len() int { return len(s.Items) / 2 }
func (s *State) Swap(ii, jj int) {
	i, j := ii*2, jj*2
	s.Items[i+GEN], s.Items[j+GEN] = s.Items[j+GEN], s.Items[i+GEN]
	s.Items[i+CHP], s.Items[j+CHP] = s.Items[j+CHP], s.Items[i+CHP]
}
func (s *State) Less(ii, jj int) bool {
	i, j := ii*2, jj*2
	//if GEN on same floor, sort by chip Floor
	// but we are sorting one array so we need to keep gens and chips together.
	if s.Items[i+GEN] == s.Items[j+GEN] {
		return s.Items[i+CHP] < s.Items[j+CHP]
	}
	return s.Items[i+GEN] < s.Items[j+GEN]
}

func (s *State) IsValid() bool {
	for i := uint8(0); i < FLOORS; i++ {
		gen, chip := false, false
		for j := 0; j < len(s.Items); j += 2 {
			if s.Items[j+GEN] != i && s.Items[j+CHP] != i {
				// these aren't on this floor
				continue
			}

			if s.Items[j+GEN] == i {
				// gen on this floor, if we have an unpaired chip, bad
				if chip {
					return false
				}
				// remember there is a generator
				gen = true
			}

			if s.Items[j+GEN] != i && s.Items[j+CHP] == i {
				// we have an unpaired chip. If there is a gen on the floor, bad
				if gen {
					return false
				}
				chip = true
			}
		}
	}
	return true
}

func (s *State) IsDone() bool {
	return s.FloorsBelowCleared(FLOORS - 1)
}

func (s *State) FloorsBelowCleared(floor uint8) bool {
	for _, p := range s.Items {
		if p < floor {
			return false
		}
	}
	return true
}

func (s *State) PossibleMoves(moveCache StateCache) []*State {
	moves := []*State{}
	/// work out moves
	/// discard seen moves
	for i := 0; i < len(s.Items); i++ {
		// we work for each item individually
		if s.Items[i] != s.Lift {
			// not on this floor, we can't move it
			continue
		}
		canMoveDown := s.Lift > 0
		canMoveUp := s.Lift < FLOORS-1

		if canMoveDown {
			// we could move down
			nxt := s.Clone()
			nxt.Lift--
			nxt.Items[i]--
			if nxt.IsValid() && moveCache.Add(nxt) {
				moves = append(moves, nxt)
			}
		}

		// check if we could move 2 up

		// we can only move the matching GEN/CHP so
		// if i%2 == GEN  then check the CHP next
		if i%2 == GEN && s.Items[i+CHP] == s.Lift {
			// chip is on the same floor, we could move up both
			if canMoveUp {
				nxt := s.Clone()
				nxt.Lift++
				nxt.Items[i]++
				nxt.Items[i+CHP]++
				if nxt.IsValid() && moveCache.Add(nxt) {
					moves = append(moves, nxt)
				}
			}
			// sometimes we need to move two down...
			if canMoveDown {
				nxt := s.Clone()
				nxt.Lift--
				nxt.Items[i]--
				nxt.Items[i+CHP]--
				if nxt.IsValid() && moveCache.Add(nxt) {
					moves = append(moves, nxt)
				}
			}
		}

		// now check the rest of the items, which we can only move if they
		// match the same type, so iterate in 2s
		//	didMoveTwo := false
		for j := i + 2; j < len(s.Items); j += 2 {
			if s.Items[j] != s.Lift {
				//not on this floor
				continue
			}
			if canMoveUp {
				nxt := s.Clone()
				nxt.Lift++
				nxt.Items[i]++
				nxt.Items[j]++
				if nxt.IsValid() && moveCache.Add(nxt) {
					moves = append(moves, nxt)
					//	didMoveTwo = true
				}
			}
			if canMoveDown {
				nxt := s.Clone()
				nxt.Lift--
				nxt.Items[i]--
				nxt.Items[j]--
				if nxt.IsValid() && moveCache.Add(nxt) {
					moves = append(moves, nxt)
					//	didMoveTwo = true
				}
			}
		}

		// don't bother moving one up, if we moved two already
		//if !didMoveTwo {
		if canMoveUp {
			nxt := s.Clone()
			nxt.Lift++
			nxt.Items[i]++
			if nxt.IsValid() && moveCache.Add(nxt) {
				moves = append(moves, nxt)
			}
		}
	}
	return moves
}
