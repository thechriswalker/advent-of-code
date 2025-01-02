package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 20, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	modules := map[string]*Module{}

	aoc.MapLines(input, func(line string) error {
		parseModule(line, modules)
		return nil
	})

	// now link all the inputs!
	for _, m := range modules {
		for _, o := range m.outputs {
			if mod, ok := modules[o]; ok {
				mod.inputs = append(mod.inputs, m.name)
			} else {
				// we dont care? these are dummy modules
				// they receive input, but don't do anything.
				// fmt.Println("no module:", o)
				// panic("no module!")
			}
		}
	}

	// now start with the broadcaster
	hi, lo := int64(0), int64(0)

	for i := 0; i < 1000; i++ {
		// one press for the button, then the broadcaster recieves lo
		h, l := pressButton(modules)
		//fmt.Printf("run %03d: hi=%d, lo=%d\n", i, h, l)
		hi += h
		lo += l
	}

	return fmt.Sprint(hi * lo)
}

func pressButton(mods map[string]*Module) (hi, lo int64) {

	curr := []Signal{{
		src:   nil,
		dst:   mods["broadcaster"],
		pulse: LO,
	}}
	var next []Signal

	for {
		next = []Signal{}
		for _, c := range curr {
			// emit the pulse!
			if c.pulse == HI {
				hi++
			} else {
				lo++
			}
			// if no destination, just stop
			if c.dst == nil {
				continue
			}
			if p := c.dst.logic(c.dst, c.src, c.pulse); p != NO {

				for _, out := range c.dst.outputs {
					next = append(next, Signal{
						src:   c.dst,
						dst:   mods[out],
						pulse: p,
					})
				}
			}
		}

		if len(next) == 0 {
			break
		}
		curr = next
	}

	return
}

// Implement Solution to Problem 2
func solve2(input string) string {
	modules := map[string]*Module{}

	aoc.MapLines(input, func(line string) error {
		parseModule(line, modules)
		return nil
	})
	// let's add a module to act as RX

	// rxReached := false
	//
	// modules["rx"] = &Module{
	// 	modules: modules,
	// 	name:    "rx",
	// 	inputs:  []string{},
	// 	outputs: []string{},
	// 	logic: func(self, src *Module, p Pulse) Pulse {
	// 		if p == LO {
	// 			rxReached = true
	// 		}
	// 		return NO
	// 	},
	// }

	// boo this will not happen any time soon....
	// instead we have to inspect our input manually
	// as we have a bunch of conjunctions wired into the
	// conjunction that feeds rx.
	// so we want to know when all those conjunctions will be
	// triggered together.
	// another LCM on the period...
	// so we need to work out the period of all the feeding
	// conjunctions.
	//

	// now link all the inputs!
	var rxInput string

	for _, m := range modules {
		for _, o := range m.outputs {
			if mod, ok := modules[o]; ok {
				mod.inputs = append(mod.inputs, m.name)
			} else {
				// we dont care? these are dummy modules
				// they receive input, but don't do anything.
				// fmt.Println("no module:", o)
				// panic("no module!")
				if o == "rx" {
					// this module outputs to rx (the rxInput)
					rxInput = m.name
				}
			}
		}
	}
	presses := 0

	// the names of the modules are those that are connected
	// to the module that rx is connected to
	//
	// note that I don't valid the assumption, so it probably won't hold for other inputs...
	// i.e. I don't check that those final 2 layers are a many-to-one conjunction group...
	nInput := 0
	nFound := 0
	nValue := 1
	for _, m := range modules[rxInput].inputs {
		// so we monkey-patch them
		m := m
		found := false
		nInput++
		monkeyPatch(modules[m], func() {
			if !found {
				found = true
				nFound++
				nValue *= presses // all co-prime, so no problem just multiplying
				// fmt.Println(m, "is high at press", presses)
			}
		})
	}

	for {
		presses++
		pressButton(modules)
		if nFound == nInput {
			break
		}
	}

	return fmt.Sprint(nValue)
}

func monkeyPatch(m *Module, fn func()) {
	// fmt.Println("patching:", m.name)
	logic := m.logic
	m.logic = func(self, src *Module, p Pulse) Pulse {
		out := logic(self, src, p)
		if out == HI {
			fn()
		}
		return out
	}
}

func parseModule(line string, all map[string]*Module) *Module {
	name, outputList, _ := strings.Cut(line, " -> ")
	m := &Module{
		modules: all,
		inputs:  []string{},
	}
	if outputList == "" {
		m.outputs = []string{}
	} else {
		m.outputs = strings.Split(outputList, ", ")
	}
	switch name[0] {
	case '%':
		// flip flop
		m.logic = flipFlop()
		m.name = name[1:]
	case '&':
		// conjunction
		m.logic = conjunction()
		m.name = name[1:]
	default:
		// broadcaster
		if name == "broadcaster" {
			m.logic = broadcaster()
			m.name = name
		}
	}
	all[m.name] = m
	return m

}

type Pulse uint8

const (
	NO Pulse = iota
	LO
	HI
)

type Logic func(self, src *Module, p Pulse) Pulse

type Module struct {
	name    string
	modules map[string]*Module
	inputs  []string
	outputs []string
	logic   Logic
}

type Signal struct {
	src, dst *Module
	pulse    Pulse
}

func broadcaster() Logic {
	return func(_, src *Module, p Pulse) Pulse { return p }
}

func flipFlop() Logic {
	var on bool
	return func(_, src *Module, p Pulse) Pulse {
		if p == HI {
			return NO
		}
		on = !on // flip
		if on {
			return HI
		} else {
			return LO
		}
	}
}

func conjunction() Logic {
	var l int
	var state map[string]struct{}
	return func(self, src *Module, p Pulse) Pulse {
		// on first run, we initialise
		if l == 0 {
			l = len(self.inputs)
			state = make(map[string]struct{}, l)
		}

		// first update state
		if p == LO {
			delete(state, src.name)
		} else {
			state[src.name] = struct{}{}
		}
		if len(state) == l {
			// all hi
			return LO
		}
		return HI
	}

}
