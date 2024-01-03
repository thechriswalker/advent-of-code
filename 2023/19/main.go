package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 19, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	metals := []*Metal{}

	workflows := map[string]*Workflow{}

	aoc.MapLines(input, func(line string) error {
		if line == "" {
			return nil
		}
		// if it starts with { it is a metal
		if line[0] == '{' {
			metals = append(metals, parseMetal(line))
		} else {
			w := parseWorkflow(line)
			workflows[w.name] = w
		}
		return nil
	})

	sum := 0

	in := workflows["in"]

	for _, m := range metals {
		// from "in" to R or A
		flow := in
		stop := false
		for {
			for _, c := range flow.cond {
				if c.check(m) {
					switch c.target {
					case "R":
						//rejected
						stop = true
					case "A":
						// accepted
						sum += m.x + m.m + m.a + m.s
						stop = true
					default:
						flow = workflows[c.target]
					}
					break
				}
			}
			if stop {
				break
			}
		}

	}

	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	workflows := map[string]*Workflow{}

	aoc.MapLines(input, func(line string) error {
		if line == "" {
			return nil
		}
		// if it starts with { it is a metal
		if line[0] == '{' {
			// metals = append(metals, parseMetal(line))
		} else {
			w := parseWorkflow(line)
			workflows[w.name] = w
		}
		return nil
	})

	// let's go recursively not bfs
	var runFlow func(metal *MetalRange, flow *Workflow) int64
	runFlow = func(metal *MetalRange, flow *Workflow) int64 {
		var sum int64
		m := metal
		for _, cond := range flow.cond {
			pass, fail := cond.split(m)
			if pass != nil {
				if cond.target == "R" {
					// nothing
				} else if cond.target == "A" {
					sum += pass.Combinations()
				} else {
					sum += runFlow(pass, workflows[cond.target])
				}
			}
			if fail == nil {
				break
			}
			m = fail
		}
		return sum
	}

	sum := runFlow(&MetalRange{
		x: &Range{lo: 1, hi: 4000},
		m: &Range{lo: 1, hi: 4000},
		a: &Range{lo: 1, hi: 4000},
		s: &Range{lo: 1, hi: 4000},
	},
		workflows["in"],
	)

	return fmt.Sprint(sum)

}

type MetalRange struct {
	x, m, a, s *Range
}

func (m *MetalRange) Combinations() int64 {
	return m.x.Size() * m.m.Size() * m.a.Size() * m.s.Size()

}

func (m *MetalRange) CopyX(r *Range) *MetalRange {
	if r == nil {
		return nil
	}
	if m.x == r {
		return m
	}
	mr := *m
	mr.x = r
	return &mr
}
func (m *MetalRange) CopyM(r *Range) *MetalRange {
	if r == nil {
		return nil
	}
	if m.m == r {
		return m
	}
	mr := *m
	mr.m = r
	return &mr
}

func (m *MetalRange) CopyA(r *Range) *MetalRange {
	if r == nil {
		return nil
	}
	if m.a == r {
		return m
	}
	mr := *m
	mr.a = r
	return &mr
}

func (m *MetalRange) CopyS(r *Range) *MetalRange {
	if r == nil {
		return nil
	}
	if m.s == r {
		return m
	}
	mr := *m
	mr.s = r
	return &mr
}

// 2 objects a "metal"
type Metal struct {
	x, m, a, s int
}

// and a Workflow
type Workflow struct {
	name string
	cond []Condition
}
type Condition struct {
	target string
	cond   string
	check  func(m *Metal) bool
	split  func(r *MetalRange) (pass, fail *MetalRange)
}

func parseMetal(line string) *Metal {
	var x, m, a, s int
	fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
	return &Metal{
		x: x, m: m, a: a, s: s,
	}
}

func parseWorkflow(line string) *Workflow {
	// ltm{a>3281:qd,m>2147:dbd,pc}
	name, rest, _ := strings.Cut(line, "{")
	ins := strings.Split(rest[:len(rest)-1], ",")
	w := &Workflow{
		name: name,
		cond: make([]Condition, 0, len(ins)),
	}
	for _, s := range ins {
		a, b, ok := strings.Cut(s, ":")
		if !ok {
			// this is a always condition
			w.cond = append(w.cond, Condition{
				target: a,
				cond:   s,
				check:  func(m *Metal) bool { return true },
				split: func(r *MetalRange) (pass, fail *MetalRange) {
					return r, nil
				},
			})
		} else {
			n, _ := strconv.Atoi(a[2:])
			op := a[1]
			field := a[0]
			w.cond = append(w.cond, Condition{
				target: b,
				cond:   s,
				check: func(m *Metal) bool {
					var v int
					switch field {
					case 'x':
						v = m.x
					case 'm':
						v = m.m
					case 'a':
						v = m.a
					case 's':
						v = m.s
					}
					if op == '<' && v < n {
						return true
					}
					if op == '>' && v > n {
						return true
					}
					return false
				},
				split: func(m *MetalRange) (pass, fail *MetalRange) {
					var r *Range
					switch field {
					case 'x':
						r = m.x
					case 'm':
						r = m.m
					case 'a':
						r = m.a
					case 's':
						r = m.s
					}
					var p, f *Range
					if op == '<' {
						p, f = splitLessThan(r, n)
					} else {
						p, f = splitGreaterThan(r, n)
					}
					switch field {
					case 'x':
						return m.CopyX(p), m.CopyX(f)
					case 'm':
						return m.CopyM(p), m.CopyM(f)
					case 'a':
						return m.CopyA(p), m.CopyA(f)
					case 's':
						return m.CopyS(p), m.CopyS(f)
					}
					panic("unreachable")
				},
			})
		}
	}
	return w
}

type Range struct{ lo, hi int }

func (r *Range) Size() int64 {
	return 1 + int64(r.hi-r.lo)
}

func splitLessThan(r *Range, n int) (pass, fail *Range) {
	if r.lo > n {
		// lo value is greater than the check, everything fails
		return nil, r
	}
	if r.hi < n {
		// if the hi value is less than, then everything passes
		return r, nil
	}
	// half and half.
	return &Range{lo: r.lo, hi: n - 1}, &Range{lo: n, hi: r.hi}
}

func splitGreaterThan(r *Range, n int) (pass, fail *Range) {
	if r.hi < n {
		// all fail if hi is less that n
		return nil, r
	}
	if r.lo > n {
		// low is higher, so all pass
		return r, nil
	}
	return &Range{lo: n + 1, hi: r.hi}, &Range{lo: r.lo, hi: n}
}
