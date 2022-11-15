package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 21, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// parse positions of players.
	p1, p2 := parseInitialPlayers(input)
	d100 := &Deterministic{Sides: 100}

	winnerIdx := runGame(d100, 1000, []*Player{p1, p2})
	var losingScore int
	if winnerIdx == 0 {
		losingScore = p2.score
	} else {
		losingScore = p1.score
	}
	n := losingScore * d100.RollCount()

	return fmt.Sprint(n)
}

func parseInitialPlayers(input string) (p1, p2 *Player) {
	aoc.MapLines(input, func(line string) error {
		p := &Player{}
		idx := 0
		_, err := fmt.Sscanf(line, "Player %d starting position: %d", &idx, &(p.position))
		if err != nil {
			return err
		}
		//decrement position, as we use 0-based
		p.position--
		if idx == 1 {
			p1 = p
		} else {
			p2 = p
		}
		return nil
	})
	return
}

// Implement Solution to Problem 2
func solve2(input string) string {

	// this is a much more maths related question.
	// the dice rolls all values at once, which means
	// in 3 rolls, we roll 3^3 different values.
	// the first numbers are the rolls (with 1 to start)
	// the second number are the sums, with 1, then 2, then 3 to start
	// 1,1,1 - 3 4 5
	// 1,1,2 - 4 5 6
	// 1,1,3 - 5 6 7
	// 1,2,1 - 4 5 6
	// 1,2,2 - 5 6 7
	// 1,2,3 - 6 7 8
	// 1,3,1 - 5 6 7
	// 1,3,2 - 6 7 8
	// 1,3,3 - 7 8 9
	// all of those with 2 at the beginning
	// all of those with 3 at the beginning
	/* of the 27 possible throws we get the sums with frequency.
	3: 1
	4: 3
	5: 6
	6: 7
	7: 6
	8: 3
	9: 1
	*/
	// so we have 9 possibilities to keep track of for each roll.
	// we need to keep track of all the possible player positions for
	// each step until a win, with the number of universes in which those
	// possibilities happened, remembering we need to "multiply" the universes
	// on each split. it would be best to have a Value for game state and map
	// that to the  number of universes, we can then mutate that each tick, discarding
	// states that win.
	wins := &Wins{}
	p1, p2 := parseInitialPlayers(input)

	initialState := StateMap{
		GameState{
			Pos1: uint8(p1.position),
			Pos2: uint8(p2.position),
		}: 1,
	}
	isPlayer1 := true
	for state := initialState; len(state) > 0; isPlayer1 = !isPlayer1 {
		state = state.Tick(isPlayer1, wins)
	}
	var n uint64
	if wins.P1 > wins.P2 {
		n = wins.P1
	} else {
		n = wins.P2
	}

	return fmt.Sprint(n)
}

type GameState struct {
	Pos1, Score1 uint8 // keep these small!
	Pos2, Score2 uint8
}

type Wins struct {
	P1, P2 uint64
}

// map from game state to universe count.
type StateMap map[GameState]uint64

var tick1 = func(state GameState, universeCount uint64, value uint8, times uint64, m StateMap, w *Wins) {
	nextPos := (state.Pos1 + value) % 10
	score := state.Score1 + nextPos + 1
	if score >= 21 {
		w.P1 += universeCount * times
	} else {
		// add the new states (the number of times this happened, multiplied by the number
		// current number of universes we had reached this state.
		nextState := GameState{
			Pos1:   nextPos,
			Pos2:   state.Pos2,
			Score1: score,
			Score2: state.Score2,
		}
		// we might already have this state, reached another way, so we check
		existing := m[nextState]
		// add the new instances of this state to the existing (probably 0)
		m[nextState] = existing + (universeCount * times)
	}
}
var tick2 = func(state GameState, universeCount uint64, value uint8, times uint64, m StateMap, w *Wins) {
	nextPos := (state.Pos2 + value) % 10
	score := state.Score2 + nextPos + 1
	if score >= 21 {
		w.P2 += universeCount * times
	} else {
		// add the new states (the number of times this happened, multiplied by the number
		// current number of universes we had reached this state.
		nextState := GameState{
			Pos2:   nextPos,
			Pos1:   state.Pos1,
			Score2: score,
			Score1: state.Score1,
		}
		// we might already have this state, reached another way, so we check
		existing := m[nextState]
		// add the new instances of this state to the existing (probably 0)
		m[nextState] = existing + (universeCount * times)
	}
}

// do we need to copy the map? or can we work on it.
func (sm StateMap) Tick(isPlayer1 bool, wins *Wins) StateMap {

	// tick for player one
	var tick func(state GameState, universeCount uint64, value uint8, times uint64, m StateMap, w *Wins)
	if isPlayer1 {
		tick = tick1
	} else {
		tick = tick2
	}
	next := make(StateMap, len(sm)*9)
	for state, universeCount := range sm {
		// iterate the possible game ticks.
		tick(state, universeCount, 3, 1, next, wins)
		tick(state, universeCount, 4, 3, next, wins)
		tick(state, universeCount, 5, 6, next, wins)
		tick(state, universeCount, 6, 7, next, wins)
		tick(state, universeCount, 7, 6, next, wins)
		tick(state, universeCount, 8, 3, next, wins)
		tick(state, universeCount, 9, 1, next, wins)
	}

	return next
}

// returns index of winning player.
func runGame(d Dice, target int, players []*Player) int {
	for {
		for i, p := range players {
			a, b, c := Roll3(d)
			p.position = (p.position + a + b + c) % 10
			p.score += p.position + 1 // positions are zero-based here, but score 1-based
			if p.score >= target {
				// player wins.
				return i
			}
		}
	}
}

type Player struct {
	position int
	score    int
}

type Dice interface {
	Roll() int
	RollCount() int
}

type Deterministic struct {
	Sides     int
	rollcount int
	last      int
}

func Roll3(d Dice) (int, int, int) {
	a := d.Roll()
	b := d.Roll()
	c := d.Roll()
	return a, b, c
}

func (d *Deterministic) Roll() int {
	d.last++
	d.rollcount++
	if d.last > d.Sides {
		d.last = 1
	}
	return d.last
}

func (d *Deterministic) RollCount() int {
	return d.rollcount
}
