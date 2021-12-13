package main

import (
	"fmt"
	"math"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 22, solve1, solve2)
}

type Fight struct {
	pHP, bHP  int
	bDamage   int
	pArmor    int
	manaSpent int
	manaLeft  int
	effects   map[string]*Effect
}

type Action func(f *Fight)
type Effect struct {
	Timer  int
	During Action
	After  Action
}

type Spell struct {
	ID     string
	Cost   int
	Action Action
	Effect *Effect
}

var spells = []Spell{
	{
		ID:     "missile",
		Cost:   53,
		Action: func(f *Fight) { f.bHP -= 4 },
	},
	{
		ID:   "drain",
		Cost: 73,
		Action: func(f *Fight) {
			f.bHP -= 2
			f.pHP += 2
		},
	},
	{
		ID:     "shield",
		Cost:   113,
		Action: func(f *Fight) { f.pArmor = 7 },
		Effect: &Effect{
			Timer: 6,
			After: func(f *Fight) { f.pArmor = 0 },
		},
	},
	{
		ID:   "poison",
		Cost: 173,
		Effect: &Effect{
			Timer:  6,
			During: func(f *Fight) { f.bHP -= 3 },
		},
	},
	{
		ID:   "recharge",
		Cost: 229,
		Effect: &Effect{
			Timer:  5,
			During: func(f *Fight) { f.manaLeft += 101 },
		},
	},
}

func (f *Fight) Clone() *Fight {
	c := &Fight{
		pHP:       f.pHP,
		bHP:       f.bHP,
		bDamage:   f.bDamage,
		pArmor:    f.pArmor,
		manaSpent: f.manaSpent,
		manaLeft:  f.manaLeft,
		effects:   make(map[string]*Effect, len(f.effects)),
	}
	for id, ef := range f.effects {
		c.effects[id] = &Effect{
			Timer:  ef.Timer,
			After:  ef.After,
			During: ef.During,
		}
	}
	return c
}

func (f *Fight) TickEffects() {
	for id, ef := range f.effects {
		if ef.During != nil {
			ef.During(f)
		}
		ef.Timer--
		if ef.Timer == 0 {
			if ef.After != nil {
				ef.After(f)
			}
			delete(f.effects, id)
		}
	}
}

func (f *Fight) Done() bool {
	return f.bHP <= 0 || f.pHP <= 0
}
func (f *Fight) Won() bool {
	return f.bHP <= 0
}

// find all the next options where we don't die.
// note that this is AFTER the effect for our turn have
// happened. So we choose an option and then we advance the game by
// result of our action, boss turn: effects, boss damage, our turn: effect
func (f *Fight) NextOptions(hardMode bool) []*Fight {
	// we can do one of 5 things.
	options := []*Fight{}
	// let's try all of them.
	for _, spell := range spells {
		// first check if we have enough mana to cast it.
		if spell.Cost > f.manaLeft {
			continue //nope
		}
		// We cannot cast it if we already have the effect running.
		if spell.Effect != nil {
			if _, ok := f.effects[spell.ID]; ok {
				// already cast... boo
				continue
			}
		}
		// OK, we can cast it!
		next := f.Clone()
		next.manaLeft -= spell.Cost
		next.manaSpent += spell.Cost
		options = append(options, next)
		if spell.Action != nil {
			spell.Action(next)
			if next.Done() {
				// don't need to go any further
				continue
			}
		}
		if spell.Effect != nil {
			next.effects[spell.ID] = &Effect{
				Timer:  spell.Effect.Timer,
				During: spell.Effect.During,
				After:  spell.Effect.After,
			}
		}
		// now we have the boss turn.
		// first tick all effects.
		next.TickEffects()
		if next.Done() {
			continue
		}
		// otherwise have the boss's go
		damage := next.bDamage - next.pArmor
		next.pHP -= damage
		if next.Done() {
			continue
		}
		// start of another turn. tick the effects.
		if hardMode {
			// in hard mode the play looses an HP at the start of their turn
			next.pHP--
			if next.Done() {
				continue
			}
		}
		next.TickEffects()
	}
	return options
}

// Implement Solution to Problem 1
func solve1(input string) string {
	fight := &Fight{
		// from input
		bHP:     55,
		bDamage: 8,
		// from description
		pHP:      50,
		manaLeft: 500,
		// init effects
		effects: map[string]*Effect{},
	}
	return fmt.Sprint(solveForFight(fight, false))
}
func solveForFight(f *Fight, hardMode bool) int {
	// depth first search for lowest winning option
	min := math.MaxInt64
	curr := []*Fight{f}
	var next []*Fight
	for {
		next = []*Fight{}
		for _, opt := range curr {
			if opt.Done() {
				// we stop here.
				if opt.Won() && min > opt.manaSpent {
					// we won! (and with a better spend!)
					min = opt.manaSpent
				}
			} else if opt.manaSpent < min {
				// we still have a chance to do better.
				next = append(next, opt.NextOptions(hardMode)...)
			}
		}
		if len(next) == 0 {
			break
		}
		curr = next
	}
	return min
}

// Implement Solution to Problem 2
func solve2(input string) string {
	fight := &Fight{
		// from input
		bHP:     55,
		bDamage: 8,
		// from description
		pHP:      49, // I think I need to do this to simulate the player first turn in hard mode
		manaLeft: 500,
		// init effects
		effects: map[string]*Effect{},
	}
	return fmt.Sprint(solveForFight(fight, true))
}
