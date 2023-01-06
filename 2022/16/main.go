package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 16, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	//return solve1_take1(input)
	//return solve1_take2(input)
	return solve1_take3(input)
}

// this was my first attempt
func solve1_take1(input string) string {
	// we need to make a tree of the input.
	valves := parseInput(input)

	// we will populate this with all 0 flow valves,
	// as we will consider them open.
	initialOpenValves := map[*Valve]int{}

	// perhaps we should pre-calculate the shortest routes
	// from any valve to any other, and then only allow travel
	// along those routes while the valve is still open.
	// we need to check those routes against the given "state"s
	// open valves.
	//routeMap := RouteMap{}
	// for each valve, and each direction find the minimal route
	// to each other valve.
	for _, va := range valves {
		// here we do a sneaky check of 0 flow before the meat.
		if va.Flow == 0 {
			//initialOpenValves[va] = 0
		}

		// routeMap[va] = map[*Valve][]*Valve{}
		// // now in each direction find a minimal route to each other valve
		// for _, vb := range valves {
		// 	// don't do this for ourselves.
		// 	if va == vb {
		// 		continue
		// 	}
		// 	dir := va.FindShortestDirection(vb)
		// 	routeMap[va][dir] = append(routeMap[va][dir], vb)
		// }

		// fmt.Println("Route from", va, "\n> ", routeMap[va])

		fmt.Printf("Valve[%s] flow:%d, neighbours: %v\n", va.Name, va.Flow, va.Next)

	}

	current := []*State{
		{At: valves["AA"], Released: 0, Open: initialOpenValves},
	}

	// testing from a given state.
	// initialOpenValves[valves["DD"]] = 2
	// initialOpenValves[valves["BB"]] = 5
	// initialOpenValves[valves["JJ"]] = 9
	// current := []*State{
	// 	{At: valves["JJ"], Released: 0, Open: initialOpenValves},
	// }

	current[0].regenOpenKey()

	cache := map[string]int{
		current[0].Key(): 0,
	}

	var next []*State

	testOptimal := []string{
		"AA>",               // 0 no time = 0
		"DD>",               // 1 t = 1 move to dd
		"DD>DD",             // 2 open DD at 2
		"CC>DD",             // 3 move CC
		"BB>DD",             // 4 move to BB
		"BB>BB:DD",          // open b at 5
		"AA>BB:DD",          //  move to AA
		"II>BB:DD",          // move to II
		"JJ>BB:DD",          // move to JJ
		"JJ>BB:DD:JJ",       // open JJ at 9
		"II>BB:DD:JJ",       // move to II
		"AA>BB:DD:JJ",       // move to AA
		"DD>BB:DD:JJ",       // move to DD
		"EE>BB:DD:JJ",       // move to ee
		"FF>BB:DD:JJ",       // move to ff
		"GG>BB:DD:JJ",       // move to gg
		"HH>BB:DD:JJ",       // move to hh
		"HH>BB:DD:HH:JJ",    // open hh at 17
		"GG>BB:DD:HH:JJ",    // to GG
		"FF>BB:DD:HH:JJ",    // to FF
		"EE>BB:DD:HH:JJ",    // to EE
		"EE>BB:DD:EE:HH:JJ", // open EE at 21
		"DD>BB:DD:EE:HH:JJ", // to DD
		"CC>BB:DD:EE:HH:JJ", // to CC
		"DD>BB:DD:EE:HH:JJ", // open CC at 24
		"DD>BB:DD:EE:HH:JJ", // stay
		"DD>BB:DD:EE:HH:JJ", // stay
		"DD>BB:DD:EE:HH:JJ", // stay
		"DD>BB:DD:EE:HH:JJ", // stay
		"DD>BB:DD:EE:HH:JJ", // stay
		"DD>BB:DD:EE:HH:JJ", // stay
	}

	// 30 minutes.
	for time := 1; time <= 30; time++ {
		fmt.Println("======= TIME", time, "================")
		for _, state := range current {

			// first release the pressure in this state.
			state.Release()

			// all valves open?
			if len(state.Open) == len(valves) {
				// perpetuate the current state, there is nothing more worth doing.
				next = append(next, state)
				continue
			}

			newState := false

			// now what are our options
			// if this valve is not open and has > 0 flow
			if state.ValveIsWorthOpening() {
				// we could open it.
				n := state.Clone()
				n.OpenCurrentValve(time)
				// we cannot have seen this state
				cache[n.Key()] = time
				if n.Key() == testOptimal[time] {
					fmt.Println("Found optimal move: open valve. from", state, "to", n)
				}
				next = append(next, n)
				newState = true
			}
			// otherwise, check all possible movements
			for _, v := range state.At.Next {
				// if there is a minimal route to another open valve
				// in this direction, take it.
				//if routeMap.HasValves(state.At, v, state.Open) {
				n := state.Clone()
				n.MoveTo(v)
				//
				if n.Key() == testOptimal[time] {
					fmt.Println("Found optimal move: move tunnel. from", state, "to", n)
				}
				// if we have seen this state, discard
				if _, seen := cache[n.Key()]; !seen {
					cache[n.Key()] = time
					next = append(next, n)
					newState = true
				}
			}
			if !newState {
				// just perpetuate the old.
				next = append(next, state)
			}
		}
		fmt.Println("After Time:", time, "Current States:", len(current), "Next States:", len(next), "cache:", len(cache))
		//fmt.Println("Cache:", cache)
		found := false
		for _, n := range next {
			if n.Key() == testOptimal[time] {
				found = true
			}
		}
		if !found {
			fmt.Println("Could not got from current:", current, "to expected next value:", testOptimal[time])
			panic("ABORT")
		}
		current = next
		next = nil
	}
	// find the one with the most Released.
	sort.Slice(current, func(i, j int) bool {
		return current[i].Released > current[j].Released
	})

	fmt.Println("All states:")
	for _, s := range current {
		fmt.Println(s)
	}

	return fmt.Sprint(current[0].Released)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return "<unsolved>"
}

type Valve struct {
	Name string // not really important.
	Flow int
	Next map[string]*Valve
}

func (v *Valve) String() string {
	return v.Name
}

type Walk struct {
	Start   *Valve
	Current *Valve
}

func (w Walk) String() string {
	return fmt.Sprintf("(%s via %s)", w.Current, w.Start)
}

// returns a valve in our ".Next" map that represent the shortest distance
// to the target valve.
func (v *Valve) FindShortestDirection(other *Valve) *Valve {
	// this is a breath first search with shortcut.
	current := []Walk{}

	// we should ignore loops, as if we have seen a node, we already got there quicker.
	cache := map[*Valve]struct{}{v: {}}

	// populate the current, with the v.Next map.
	for _, n := range v.Next {
		// might as well shorcut for adjacent valves.
		if n == other {
			return n
		}
		// treat this as visited
		cache[n] = struct{}{}

		// and add to list
		current = append(current, Walk{
			Start:   n,
			Current: n,
		})
	}

	debug := false

	if debug {
		fmt.Println(">>>>>>>>>>>>>>: from", v.Name, "to", other.Name)
	}
	for {
		if debug {
			fmt.Println("Route Finding: current", current)
		}
		next := []Walk{}
		for _, c := range current {
			// look at this node's Next
			for _, n := range c.Current.Next {
				// if this is the target, then we are done.
				if debug {
					fmt.Println("Route Finding: at", c, "checking:", n.Name)
				}
				if n == other {
					// return the valve in the right direction.
					return c.Start
				}
				// if we have been not reached this node yet, use this path
				if _, seen := cache[n]; !seen {
					if debug {
						fmt.Println("Route Finding: at", c, "checking:", n.Name, "not in cache, adding to next")
					}
					cache[n] = struct{}{}
					// try again with this node
					next = append(next, Walk{
						Start:   c.Start,
						Current: n,
					})
				} else {
					if debug {
						fmt.Println("Route Finding: at", c, "checking:", n.Name, "ALREADY SEEN, skip this route")
					}
				}
			}
		}
		if debug {
			fmt.Println("Route Finding:    next", next)
			// we should expect to see the optimal path in one of our options.
		}
		if len(next) == 0 {
			panic(fmt.Sprintf("cannot find route from %q to %q", v.Name, other.Name))
		}
		current = next
	}
}

// This map is a list of valves, and their "next" directions
// and the valves reachable via those direction that are minimal.
type RouteMap map[*Valve]map[*Valve][]*Valve

func (rm RouteMap) HasValves(src, dir *Valve, ignore map[string]*Valve) bool {
	valves := rm[src][dir]
	// all valves in that direction.
	for _, valve := range valves {
		// if this valve is not in our ignore list (already opened, or zero flow)
		// then return true.
		if _, ignored := ignore[valve.Name]; !ignored {
			return true
		}
	}
	return false
}

type State struct {
	At       *Valve
	Open     map[*Valve]int // valve an minute it was opened!
	Released int
	openkey  string
}

func (s *State) String() string {
	sb := strings.Builder{}
	fmt.Fprintf(&sb, "At: %v, Released Flow: %d (key=%s)\n", s.At, s.Released, s.Key())
	// best to sort the valves by open time.
	opened := make([]*Valve, 0, len(s.Open))
	for v, m := range s.Open {
		if m > 0 {
			opened = append(opened, v)
		}
	}
	sort.Slice(opened, func(i, j int) bool { return s.Open[opened[i]] < s.Open[opened[j]] })

	for _, v := range opened {
		m := s.Open[v]
		fmt.Fprintf(&sb, "Opened %s at minute %d\n", v, m)
	}
	return sb.String()
}

// a unique key for this state.
// allows us to drop "duplicate" state achieved "later"
// as they cannot be optimal.
func (s *State) Key() string {
	return s.At.Name + ">" + s.openkey
}

func (s *State) OpenCurrentValve(minute int) {
	s.Open[s.At] = minute
	// re generate openkey
	s.regenOpenKey()
}

func (s *State) regenOpenKey() {
	type vm struct {
		v *Valve
		m int
	}
	valves := []vm{}
	keys := make([]string, 0, len(s.Open))
	for k, m := range s.Open {
		valves = append(valves, vm{v: k, m: m})
	}
	// sort the keys by ORDER OPENED
	sort.Slice(valves, func(i, j int) bool { return valves[i].m < valves[j].m })
	for _, x := range valves {
		keys = append(keys, x.v.Name)
	}
	s.openkey = strings.Join(keys, ":")

}

func (s *State) MoveTo(v *Valve) {
	s.At = v
}

func (s *State) ValveIsWorthOpening() bool {
	if _, open := s.Open[s.At]; open {
		return false // already open
	}
	return s.At.Flow > 0
}

func (s *State) Release() {
	for v := range s.Open {
		s.Released += v.Flow
	}
}

func (s *State) Clone() *State {
	clone := &State{
		At:       s.At,
		Open:     map[*Valve]int{},
		Released: s.Released,
		openkey:  s.openkey,
	}
	for k, v := range s.Open {
		clone.Open[k] = v
	}

	return clone
}

func parseInput(input string) map[string]*Valve {
	valves := map[string]*Valve{}

	getOrCreateValve := func(name string) *Valve {
		v, ok := valves[name]
		if !ok {
			v = &Valve{Name: name, Next: map[string]*Valve{}}
			valves[name] = v
		}
		return v
	}

	var name string
	var flow int
	aoc.MapLines(input, func(line string) error {
		fmt.Sscanf(line, "Valve %s has flow rate=%d", &name, &flow)
		v := getOrCreateValve(name)
		v.Flow = flow
		// Valve JJ has flow rate=21; tunnel leads to valve // 48 characters
		// if rate < 10 the we have one less character, but that will be a space.
		links := strings.Split(strings.TrimSpace(line[49:]), ", ")

		for _, link := range links {
			lv := getOrCreateValve(link)
			v.Next[link] = lv
		}
		return nil
	})

	return valves
}

// I couldn't work out how to search arouter := &Router{}ll the possibilities.
// the space gets too big, and I couldn't work out a reasonable
// way to trim "equivalent paths", either trimmed too many, and
// missed the optimal solution, or not enough and it would take
// too long (and too much memory and crash).
//
// So this is attempt 2, a more focussed approach.
// At each time step I will find all the possible valves to open
// and how much pressure they would release until the end of the
// 30 minutes.
//
// if I take the greatest path here, then move repeat, until
// I couldn't find anything new before the 30 minutes runs out.
//
// I can see this not working either as each step may be optimal
// in isolation, but as a path they might not be. But if it works
// for the test data, then it'll probably work for the real input.
//
// No, but instead of taking the "best" route, if we use all of them
// and branch that way, we have cut the problem space down, to a reasonable level
// and perhaps we can solve it!
//
// that will be take 3
func solve1_take2(input string) string {
	valves := parseInput(input)

	vslice := make([]*Valve, 0, len(valves))
	for _, v := range valves {
		vslice = append(vslice, v)
	}

	// valve and minute opened
	opened := map[*Valve]int{}
	current := valves["AA"]
	router := &Router{}

	// t = 0.
	for t := 0; t < 30; {
		t, current = router.FindBestNextValve(t, current, vslice, opened)
	}

	// now add up the flows
	pressure := 0
	for v, m := range opened {
		pressure += v.Flow * (30 - m)
	}

	return fmt.Sprint(pressure)
}

type State2 struct {
	Time   int
	At     *Valve
	Opened map[*Valve]int
}

func (s *State2) Released() int {
	pressure := 0
	for v, m := range s.Opened {
		pressure += v.Flow * (30 - m)
	}
	return pressure
}

func solve1_take3(input string) string {
	valves := parseInput(input)

	vslice := make([]*Valve, 0, len(valves))
	for _, v := range valves {
		vslice = append(vslice, v)
	}
	router := &Router{max: 30} // 30 mins

	// damn it, we need states here as well.
	// valve and minute opened

	current := []State2{{
		Time:   0,
		At:     valves["AA"],
		Opened: map[*Valve]int{},
	}}

	completed := []State2{}

	for {
		next := []State2{}
		for _, s := range current {
			// find available next choices.
			routes := router.NextRoutes(s.Time, s.At, vslice, s.Opened)
			if len(routes) == 0 {
				// nothing to do!
				completed = append(completed, s)
				continue
			}
			// explore the routes.
			for _, r := range routes {
				t := s.Time + r.Distance + 1
				opened := openValveClone(r.Destination, s.Opened, t)
				next = append(next, State2{
					Time:   t,
					At:     r.Destination,
					Opened: opened,
				})
			}
		}
		if len(next) == 0 {
			// done
			break
		}
		current = next
	}

	// sort the completed states by pressure
	sort.Slice(completed, func(i, j int) bool {
		return completed[i].Released() > completed[j].Released()
	})

	opened := completed[0].Opened

	// now add up the flows
	pressure := 0
	for v, m := range opened {
		pressure += v.Flow * (30 - m)
	}

	return fmt.Sprint(pressure)
}

func openValveClone(v *Valve, m map[*Valve]int, t int) map[*Valve]int {
	open := make(map[*Valve]int, len(m)+1)
	for k, v := range m {
		open[k] = v
	}
	open[v] = t
	return open
}

type Route struct {
	Distance    int
	Destination *Valve
}

type Router struct {
	max int
	m   map[*Valve]map[*Valve]int
}

func (r Route) PressureReleased(t, max int) int {
	if r.Destination.Flow == 0 {
		return 0
	}
	// it will take r.Distance to get there and 1 to open the valve
	t += r.Distance + 1
	// and then we will get r.Destination.Flow units until
	// t = 30
	// if t is already > 30, then nothing.
	if t >= max {
		return 0
	}
	return r.Destination.Flow * (max - t)
}

func (r *Router) NextRoutes(t int, current *Valve, all []*Valve, opened map[*Valve]int) []Route {
	options := make([]Route, 0, len(all)-len(opened))
	for _, v := range all {
		if _, ok := opened[v]; ok || v.Flow == 0 {
			// no flow, or already opened, skip
			continue
		}
		// worth checking
		route := r.getShortestRoute(current, v)
		//	fmt.Println("Route from ", current, "to", v, "is", route, "at", t, "pressure:", route.PressureReleased(t))
		if route.PressureReleased(t, r.max) > 0 {
			options = append(options, route)
		}
	}
	return options
}

func (r *Router) FindBestNextValve(t int, current *Valve, all []*Valve, opened map[*Valve]int) (int, *Valve) {
	// for all valves, what is the max pressure we would get out of moving to and
	// opening all other valves. pick the greatest
	options := r.NextRoutes(t, current, all, opened)
	if len(options) == 0 {
		// nothing to do.
		// bump time and return
		return math.MaxInt, nil
	}

	// else sort and use the biggest output.
	sort.Slice(options, func(i, j int) bool {
		return options[i].PressureReleased(t, r.max) > options[j].PressureReleased(t, r.max)
	})

	winner := options[0]
	fmt.Println("Best option:", winner)

	return t + winner.Distance + 1, winner.Destination
}

func (r *Router) getShortestRoute(src, dst *Valve) Route {
	if r.m == nil {
		r.m = map[*Valve]map[*Valve]int{}
	}
	// if we have it, use it, else work it out and cache it.
	d, ok := r.m[src][dst]
	if !ok {
		// we must find the route.
		d = getDistance(src, dst)

		if _, ok := r.m[src]; !ok {
			r.m[src] = map[*Valve]int{}
		}
		if _, ok := r.m[dst]; !ok {
			r.m[dst] = map[*Valve]int{}
		}
		r.m[src][dst] = d
		r.m[dst][src] = d
	}
	return Route{Distance: d, Destination: dst}
}

func getDistance(src, dst *Valve) int {
	current := []*Valve{src}

	// we should ignore loops, as if we have seen a node, we already got there quicker.
	cache := map[*Valve]struct{}{src: {}}

	d := 0
	for {
		next := []*Valve{}
		for _, c := range current {
			if dst == c {
				return d
			}
			// look at this node's Next
			for _, n := range c.Next {
				if n == dst {
					// return the valve in the right direction.
					return d + 1
				}
				// if we have been not reached this node yet, use this path
				if _, seen := cache[n]; !seen {
					cache[n] = struct{}{}
					// try again with this node
					next = append(next, n)
				}
			}
		}
		if len(next) == 0 {
			panic(fmt.Sprintf("cannot find route from %q to %q", src.Name, dst.Name))
		}
		current = next
		d++
	}
}
