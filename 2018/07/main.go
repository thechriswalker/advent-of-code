package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2018, 7, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	ins := parseInput(input)
	for ins.DoNextTask() {
	}
	return ins.Path.String()
}

// Implement Solution to Problem 2
func solve2(input string) string {
	ins := parseInput(input)
	return fmt.Sprintf("%d", ins.TimeWorkers())
}

type Step struct {
	Id           rune
	Prerequistes []*Step
	Done         bool
	Started      int
	InProgress   bool
	available    bool // memoize the available state
}

// available if !done and all prequistes done
func (s *Step) Available() bool {
	if s.Done || s.InProgress {
		return false
	}
	if s.available {
		return true
	}
	for _, p := range s.Prerequistes {
		if !p.Done {
			return false
		}
	}
	s.available = true
	return true
}

func parseInput(input string) *Instructions {
	ins := &Instructions{
		Steps: make([]*Step, 0, 26), // won't be bigger than the alphabet!
		temp:  make([]*Step, 0, 26),
		Path:  &strings.Builder{},
	}
	rd := strings.NewReader(input)
	var err error
	var a, b rune
	var dependency, dependent *Step
	var found bool
	byId := map[rune]*Step{}
	for {
		if _, err = fmt.Fscanf(rd, "Step %c must be finished before step %c can begin.\n", &a, &b); err != nil {
			break
		}
		// `a` is a prerequiste of `b`
		dependency, found = byId[a]
		if !found {
			dependency = &Step{Id: a, Prerequistes: []*Step{}}
			byId[a] = dependency
			ins.Steps = append(ins.Steps, dependency)
		}
		dependent, found = byId[b]
		if !found {
			dependent = &Step{Id: b, Prerequistes: []*Step{}}
			byId[b] = dependent
			ins.Steps = append(ins.Steps, dependent)
		}
		dependent.Prerequistes = append(dependent.Prerequistes, dependency)
	}
	return ins
}

type Instructions struct {
	Steps []*Step
	Path  *strings.Builder
	temp  []*Step // to save re-allocations
}

func (i *Instructions) Complete() bool {
	for _, s := range i.Steps {
		if !s.Done {
			return false
		}
	}
	return true
}
func (i *Instructions) FindNextSteps() []*Step {
	if i.temp == nil {
		i.temp = make([]*Step, 0, len(i.Steps))
	}
	next := i.temp[0:0]

	// this is probably going to be easiest by
	// simply iteration All steps and then sorting
	// rather than worrying about what is already in the slice.

	for _, s := range i.Steps {
		if s.Available() {
			next = append(next, s)
		}
	}
	sort.Sort(StepSlice(next))
	return next
}

// returns true until all tasks complete
func (i *Instructions) DoNextTask() bool {
	next := i.FindNextSteps()
	if len(next) == 0 {
		return false
	}
	// take the first task and "do" it
	next[0].Done = true
	i.Path.WriteRune(next[0].Id)
	return true
}

// for sorting
type StepSlice []*Step

func (ss StepSlice) Len() int           { return len(ss) }
func (ss StepSlice) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }
func (ss StepSlice) Less(i, j int) bool { return ss[i].Id < ss[j].Id }

// part 2
var BaseTaskTime = 60
var Workers = 5

// lets do part 2 the awesome way, with goroutines for workers!
// probably more complex and less efficient and prone to race conditions
// and I couldn't get it to work right, the co-ordination required so much
// overhead I decided to do it the easy way.
// maybe I'll come back to the goroutine version some time.
func (ins *Instructions) TimeWorkers() int {
	workers := make([]*Worker, Workers)
	for i := 0; i < Workers; i++ {
		workers[i] = &Worker{Id: i}
	}
	clock := 0
	tasks := len(ins.Steps)
	refreshTasks := false
	availableTasks := ins.FindNextSteps()
	for {
		// loop through workers to see if any have finished
		for _, w := range workers {
			// has the task finished
			if w.Step != nil && w.Finish == clock {
				//		fmt.Printf("Worker %d finished task %c at %d\n", w.Id, w.Step.Id, clock)
				tasks--
				w.Step.Done = true
				refreshTasks = true
				w.Step = nil
			}
		}
		// if some finished, we need to refresh our list
		if refreshTasks {
			availableTasks = ins.FindNextSteps()
		}
		// now try and allocate work to any free workers
		for _, w := range workers {
			if w.Step == nil && len(availableTasks) > 0 {
				w.Step = availableTasks[0]
				//	fmt.Printf("Worker %d started task %c at %d\n", w.Id, w.Step.Id, clock)
				availableTasks = availableTasks[1:]
				w.Step.InProgress = true
				w.Step.Started = clock
				w.Finish = clock + BaseTaskTime + int(w.Step.Id) - 64
			} else if w.Step == nil {
				//				fmt.Printf("Worker %d idle, but no tasks\n", w.Id)
			} else {
				//			fmt.Printf("Worker %d busy\n", w.Id)
			}
		}
		// if no more tasks we are done!
		if tasks == 0 {
			return clock
		}
		clock++
	}
}

type Worker struct {
	Id     int
	Step   *Step
	Finish int
}
