package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 12, solve1, solve2)
}

func solve1(input string) string {
	return solve1steps(input, 1000)
}

// Implement Solution to Problem 1
func solve1steps(input string, steps int) string {
	m := parseInput(input)
	for i := 0; i < steps; i++ {
		tick(m)
	}
	return fmt.Sprint(energy(m))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// we need a good way to match the data.
	// we have 6 values, and I reckon they don't go so
	// high. But we have to keep ALL the states that don't
	// repeat to test to repetition.
	// the example had 4686774924 steps.
	// to keep the outcome of all these steps with 6 bytes
	// per step is 28_120_649_544 (28 Gb)
	// clearly I can't do that.
	// if each value never goes higher than +-15 then there are
	// only 32 values so 5bits is enough, so 6 of those numbers
	// could fit in 30 bits. and we only need 4bytes per state.
	// but that is still only 2/3 of the 6 (so _only_ 16Gb)

	// actually it's worse because there are 4 moons, each with 6
	// values. so that's 24 pieces of information.

	// first thing to do is to estimate the bounds, the min/max value
	// for each co-ordinate and velocity, by running the simulation for
	// say, 1e6 iterations?

	// bounds are at least -2222, 2256.
	// so I think we need a trie to reduce duplication...
	// a 24 step trie?
	// no that is too deep.
	// how about a compromise, 1 per moon per attribute
	// i.e a 4 * 2 = 8 step trie and key it on the [3]int16
	// or we just keep the state of each moon separately
	// in 4 * 2 level tries
	// we could just use a single map for each moon

	// still takes too long (and eventually runs outof memory)

	// so there must be a pattern.
	// let's see if the moons cycle independantly with a given frequency
	// then we can extrapolate when those cycles will align.
	// it will be the least common multiple!
	// no, there is no common cycle pattern I can see.

	// and I run out of memory eventually.

	// shit this is getting tough the clue is in the problem for sure:

	// > Clearly, you might need to find a more efficient way to simulate the universe.

	// how do I represent this universe any differently? with the same rules?
	// clearly I cannot do 4,686,774,924 steps in <10secs on reasonable hardware
	// So there must be a different way of representing this whole thing
	// that I am missing.

	// x y and z ARE independant,so we should only care about x,and vx for each
	// moon at a time, so only 3 periods to match, then lowest common multiple!

	moons := parseInput(input)
	//fmt.Println("Number of moons", len(moons))
	if len(moons) != 4 {
		panic("expected 4 moons!")
	}
	// one for each

	// we will store the states in a map[[8]int16]uint64{}
	X := map[[8]int16]uint64{}
	Y := map[[8]int16]uint64{}
	Z := map[[8]int16]uint64{}

	// this is the last period
	//	var xLast, yLast, zLast uint64

	checkX := [8]int16{}
	checkY := [8]int16{}
	checkZ := [8]int16{}

	// the first repetition will have an offset and a period.
	// no offset needed, we are actually finding the first time the initial state is
	// seen again, that would have eliminated the memory bounds problem as we would
	// only have to keep the initial state (not everything), but would not have fixed
	// the fact that 4.7 billion ticks will take longer than I can wait.
	var dx, dy, dz uint64

	var fx, fy, fz bool

	ticks := uint64(0)
	for {
		for i, m := range moons {
			// check X, y , z values separately
			copy(checkX[i*2:(i*2)+2], []int16{m[x], m[vx]})
			copy(checkY[i*2:(i*2)+2], []int16{m[y], m[vy]})
			copy(checkZ[i*2:(i*2)+2], []int16{m[z], m[vz]})
		}
		// now check for collisions on each axis
		if !fx {
			xl, xok := X[checkX]
			if xok {
				// this is an X collision
				//fmt.Printf("X match: curr % 6d, last % 06d, diff % 6d\n", ticks, xl, ticks-xl)
				dx = ticks - xl // period
				fx = true
			} else {
				//set it
				X[checkX] = ticks
			}
		}
		if !fy {
			yl, yok := Y[checkY]
			if yok {
				// this is an Y collision
				//fmt.Printf("Y match: curr % 6d, last % 06d, diff % 6d\n", ticks, yl, ticks-yl)
				dy = ticks - yl // period
				fy = true
			} else {
				//set it
				Y[checkY] = ticks
			}
		}
		if !fz {
			zl, zok := Z[checkZ]
			if zok {
				// this is an Z collision
				//fmt.Printf("Z match: curr % 6d, last % 06d, diff % 6d\n", ticks, zl, ticks-zl)
				dz = ticks - zl // period
				fz = true
			} else {
				//set it
				Z[checkZ] = ticks
			}
		}
		if fx && fy && fz {
			// we have the offset an periods.
			//fmt.Printf("X offset=%d,period=%d\nY offset=%d,period=%d\nZ offset=%d,period=%d\n", ox, dx, oy, dy, oz, dz)
			return fmt.Sprint(lcm(dx, dy, dz))
		}
		tick(moons)
		ticks++
	}
}

var mem = runtime.MemStats{}

func dumpMemoryUse(i int) {
	runtime.ReadMemStats(&mem)
	fmt.Printf("Tick %09d: Mem %0.4fGB\n", i, float64(mem.HeapAlloc)/(1024*1024*1024))
}

// // is this how they work?
// type Trie map[[3]int16]Trie

// // make this one function to save iteration time
// // the boolean is whether it Has the value,
// // it always Sets.
// func (t Trie) HasOrSet(values ...[3]int16) bool {
// 	if len(values) == 1 {
// 		_, ok := t[values[0]]
// 		if ok {
// 			// we already had this value.
// 			return true
// 		}
// 		// we didn't already have it.
// 		next := Trie{}
// 		t[values[0]] = next
// 		return false
// 	}
// 	next, ok := t[values[0]]
// 	if !ok {
// 		next = Trie{}
// 		t[values[0]] = next
// 		next.HasOrSet(values[1:]...)
// 		return false
// 	}
// 	return next.HasOrSet(values[1:]...)
// }

type Moon [6]int16

func parseInput(input string) []*Moon {
	rd := strings.NewReader(input)
	moons := []*Moon{}
	var x, y, z, id int16
	for {
		if n, _ := fmt.Fscanf(rd, "<x=%d, y=%d, z=%d>", &x, &y, &z); n == 3 {
			moons = append(moons, &Moon{x, y, z})
			id++
			rd.ReadByte()
		} else {
			break
		}
	}
	return moons
}

func dump(moons []*Moon) {
	for _, m := range moons {
		fmt.Println(m)
	}
}

func (m Moon) String() string {
	return fmt.Sprintf("Pos<x=%d, y=%d, z=%d> Vel<x=%d, y=%d, z=%d>", m[x], m[y], m[z], m[vx], m[vy], m[vz])
}

func tick(moons []*Moon) {
	// for each pair of moons, calculate the gravity between them and ammend velocity
	for i := 0; i < len(moons)-1; i++ {
		for j := i + 1; j < len(moons); j++ {
			a, b := moons[i], moons[j]
			updateVforG(a, b, x)
			updateVforG(a, b, y)
			updateVforG(a, b, z)
		}
	}
	// now ammend position for velocity
	for _, m := range moons {
		m[x] += m[vx]
		m[y] += m[vy]
		m[z] += m[vz]
	}
}

func updateVforG(a, b *Moon, c int) {
	pa, pb := a[c], b[c]
	//fmt.Printf("testing moons %d and %d on axis %d => %d vs %d ", a.id, b.id, c, pa, pb)
	if pa > pb {
		// a towards b
		a[c+3]--
		b[c+3]++
		//fmt.Printf("=> %d > %d\n", pa, pb)
	}
	if pa < pb {
		// b towards a
		a[c+3]++
		b[c+3]--
		//fmt.Printf("=> %d < %d\n", pa, pb)
	}
	// if pa == pb {
	// 	// if equal, nothing
	// 	fmt.Printf("=> %d == %d\n", pa, pb)
	// }
}

const (
	x  = 0
	y  = 1
	z  = 2
	vx = 3
	vy = 4
	vz = 5
)

func energy(moons []*Moon) int {
	sum := 0
	for _, m := range moons {
		sum += (abs(m[x]) + abs(m[y]) + abs(m[z])) * (abs(m[vx]) + abs(m[vy]) + abs(m[vz]))
	}
	return sum
}

func abs(a int16) int {
	if a < 0 {
		return -1 * int(a)
	}
	return int(a)
}

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b uint64, integers ...uint64) uint64 {
	result := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}
