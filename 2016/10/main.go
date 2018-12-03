package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 10, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	chips := BuildTree(input)
	// find the bot responsible for 61 and 17
	for _, c := range chips.Bots {
		if (c.Inputs[0] == 17 && c.Inputs[1] == 61) || (c.Inputs[0] == 61 && c.Inputs[1] == 17) {

			return fmt.Sprintf("%d", c.Id)
		}
	}
	return "<unsolved>"
}

// Implement Solution to Problem 2
func solve2(input string) string {
	chips := BuildTree(input)
	v := chips.Outputs[0].Value * chips.Outputs[1].Value * chips.Outputs[2].Value
	return fmt.Sprintf("%d", v)
}

type ChipThing struct {
	Bots    map[int]*Bot
	Outputs map[int]*Output
}

type Inputter interface {
	Input(v int)
}

type Outputter interface {
	Output(lo, hi Inputter)
}

type Bot struct {
	Id     int
	Inputs [2]int
	count  int
	Lo, Hi Inputter
}

func (b *Bot) output() {
	if b.count == 2 {
		if b.Inputs[0] > b.Inputs[1] {
			if b.Hi != nil {
				b.Hi.Input(b.Inputs[0])
			}
			if b.Lo != nil {
				b.Lo.Input(b.Inputs[1])
			}
		} else {
			if b.Hi != nil {
				b.Hi.Input(b.Inputs[1])
			}
			if b.Lo != nil {
				b.Lo.Input(b.Inputs[0])
			}
		}
	}
}

func (b *Bot) Input(v int) {
	if b.count == 2 {
		panic(fmt.Sprintf("Trying to add 3rd value to bot %d (value: %d, has %d,%d)", b.Id, v, b.Inputs[0], b.Inputs[1]))
	}
	b.Inputs[b.count] = v
	b.count++
	b.output()
}

func (b *Bot) Output(lo, hi Inputter) {
	b.Lo = lo
	b.Hi = hi
	b.output()
}

type Output struct {
	Id    int
	Value int
}

func (o *Output) Input(v int) {
	o.Value = v
}

func BuildTree(input string) *ChipThing {
	chips := &ChipThing{
		Bots:    map[int]*Bot{},
		Outputs: map[int]*Output{},
	}
	rd := strings.NewReader(input)
	var value1, value2, value3 int
	var type1, type2 string

	var c rune
	var err error
	for {
		// read first rune, 'v' or 'b'
		c, _, err = rd.ReadRune()
		if err != nil {
			break
		}
		switch c {
		case 'v':
			// value X goes to bot Y
			_, err = fmt.Fscanf(rd, "alue %d goes to bot %d\n", &value1, &value2)
			if err != nil {
				break
			}
			bot := chips.MakeOrGetBot(value2)
			bot.Input(value1)
		case 'b':
			// bot X gives low to (bot|output) X and high to (bot|output) Y
			_, err = fmt.Fscanf(rd, "ot %d gives low to %s %d and high to %s %d\n", &value3, &type1, &value1, &type2, &value2)
			if err != nil {
				break
			}
			bot := chips.MakeOrGetBot(value3)
			lo := chips.MakeOrGetBotOrOutput(value1, type1)
			hi := chips.MakeOrGetBotOrOutput(value2, type2)
			bot.Output(lo, hi)
		}
	}
	return chips
}

func (c *ChipThing) MakeOrGetBotOrOutput(id int, kind string) Inputter {
	if kind == "bot" {
		return c.MakeOrGetBot(id)
	}
	return c.MakeOrGetOutput(id)
}

func (c *ChipThing) MakeOrGetOutput(id int) *Output {
	out, ok := c.Outputs[id]
	if !ok {
		out = &Output{Id: id}
		c.Outputs[id] = out
	}
	return out
}

func (c *ChipThing) MakeOrGetBot(id int) *Bot {
	bot, ok := c.Bots[id]
	if !ok {
		bot = &Bot{Id: id, Inputs: [2]int{0, 0}}
		c.Bots[id] = bot
	}
	return bot
}
