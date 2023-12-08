package main

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 7, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	games := []*Game{}
	aoc.MapLines(input, func(line string) error {
		games = append(games, parseGame(line, false))
		return nil
	})

	slices.SortFunc(games, gameSort)

	maxScore := len(games)

	total := 0
	for i, g := range games {
		//fmt.Printf("Game %d: %#v\n", i, g)
		total += (maxScore - i) * g.bid
	}

	return fmt.Sprint(total)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	games := []*Game{}
	aoc.MapLines(input, func(line string) error {
		games = append(games, parseGame(line, true))
		return nil
	})

	slices.SortFunc(games, gameSort)

	maxScore := len(games)

	total := 0
	for i, g := range games {
		// if g.hasJokers {
		// 	fmt.Printf("%s\n", g)
		// }
		total += (maxScore - i) * g.bid
	}

	//249851558 too high - I missed that jokers still count in the card-by-card ordering
	//249776650 - just right
	//249645808 too low - I missed that while they still count, they count as weaker than 2...
	return fmt.Sprint(total)
}

type Game struct {
	line           string
	cards          []byte
	cardWithJokers []byte
	rank           map[byte]int
	maxRank        int
	bid            int
	hasJokers      bool
}

var byteRank = []byte("AKQJT98765432j")

var byteRankMap = map[byte]byte{}

func init() {
	for r, b := range byteRank {
		byteRankMap[b] = byte(r)
		byteRankMap[byte(r)] = b
	}
}

func parseGame(line string, withJokers bool) *Game {
	g := &Game{
		line:           line,
		cards:          []byte(line[0:5]),
		cardWithJokers: make([]byte, 5),
		rank:           map[byte]int{},
	}
	g.bid, _ = strconv.Atoi(line[6:])

	jokers := []int{}

	// translate and add up
	for i, b := range g.cards {
		if withJokers && b == 'J' {
			// handle joker later
			jokers = append(jokers, i)
			// but the joker gets the weakest possible value - weaker than 2
			g.cards[i] = byteRankMap['j']
			continue
		}
		g.cards[i] = byteRankMap[b]
		g.rank[g.cards[i]]++
		if g.rank[g.cards[i]] > g.maxRank {
			g.maxRank = g.rank[g.cards[i]]
		}
	}
	copy(g.cardWithJokers, g.cards)

	if withJokers && len(jokers) > 0 {
		g.hasJokers = true
		// handle jokers - increment "best" position
		// basically, find cards we have most of or highest if tied.
		bestCardAtMaxRank := byteRankMap['j']
		//	fmt.Printf("%#v\n", g)
		for card, r := range g.rank {
			if r == g.maxRank && card < bestCardAtMaxRank {
				bestCardAtMaxRank = card
			}
		}
		g.rank[bestCardAtMaxRank] += len(jokers)

		if g.rank[bestCardAtMaxRank] > g.maxRank {
			g.maxRank = g.rank[bestCardAtMaxRank]
		}

		for _, i := range jokers {
			g.cardWithJokers[i] = bestCardAtMaxRank
		}
	}

	return g
}

func (g *Game) String() string {
	cards := make([]byte, 5)
	withj := make([]byte, 5)
	for i, b := range g.cards {
		cards[i] = byteRankMap[b]
		withj[i] = byteRankMap[g.cardWithJokers[i]]
	}
	return fmt.Sprintf("Cards:%s (%s), Rank:%d, MaxGroup:%d, Line:%s", string(cards), string(withj), len(g.rank), g.maxRank, g.line)
}

func gameSort(a, b *Game) int {
	// hand rank first
	if len(a.rank) != len(b.rank) {
		return len(a.rank) - len(b.rank)
	}
	// rank equal, could be still a better hand
	// better hand will have higher max same cards
	// i.e. 2pair has 3 rank, but only 2 highest group
	// 3 of king has 3 rank, but 3 in highest group
	if a.maxRank != b.maxRank {
		return b.maxRank - a.maxRank
	}
	// equal hands, compare actual cards.
	return bytes.Compare(a.cards, b.cards)
}
