package main

import (
	"fmt"

	"../../aoc"

	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	aoc.Run(2020, 7, solve1, solve2)
}

// The bags can contain other bags and we should probably
// only have a single value per bag. so we are looking at
// tree structure, then we can search into the tree for
// `shiny gold` bags, keeping track of our path and
// then count unique colors for the number.
type Bag struct {
	color       string
	contains    []*BagCount
	containedBy []*Bag
}

type BagCount struct {
	bag   *Bag
	count int
}

var contentsRegex = regexp.MustCompile(`\s*([0-9]+) (.*) bags?`)

func parseInput(in string) map[string]*Bag {
	// every bag will be at the top level.
	// we will always look to the top level to
	// prevent making duplicates
	lines := strings.Split(in, "\n")
	out := make(map[string]*Bag, len(lines))

	getBag := func(c string) *Bag {
		bag, ok := out[c]
		if !ok {
			bag = &Bag{
				color:       c,
				contains:    []*BagCount{},
				containedBy: []*Bag{},
			}
			out[c] = bag
		}
		return bag
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		s := strings.Split(line, " bags contain ")
		if len(s) != 2 {
			log.Fatalln("Bad line!", line)
		}
		color, inner := s[0], s[1]
		bag := getBag(color)
		if inner != "no other bags." {
			for _, innerBagCount := range strings.Split(inner, ",") {
				m := contentsRegex.FindStringSubmatch(innerBagCount)
				//log.Println(innerBagCount, m)
				innerBag := getBag(m[2])
				innerBag.containedBy = append(innerBag.containedBy, bag)
				count, err := strconv.Atoi(m[1])
				if err != nil {
					log.Fatalln("bad inner bag declaration:", line, m)
				}
				bag.contains = append(bag.contains, &BagCount{
					count: count,
					bag:   innerBag,
				})
			}
		}
	}
	return out
}

// Implement Solution to Problem 1
func solve1(input string) string {
	bags := parseInput(input)

	// find how many can contain a Shiny - Gold.
	// we mapped the "contained by", so we iterate
	// outwards from the shiny, and count without
	// duplicates.
	found := map[*Bag]int{}
	shinyGold := bags["shiny gold"]
	// let's assume this is non-nil...
	var contained func(*Bag)
	contained = func(b *Bag) {
		for _, bb := range b.containedBy {
			found[bb] = 1
			contained(bb)
		}
	}
	contained(shinyGold)

	return fmt.Sprintf("%d", len(found))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	bags := parseInput(input)

	// we will cache every time we make a count to make this easier
	cache := map[*Bag]int{}

	var countInnerBags func(*Bag) int
	countInnerBags = func(b *Bag) int {
		if cached, ok := cache[b]; ok {
			return cached
		}
		count := 1 // the bag itself!
		for _, bc := range b.contains {
			//log.Printf("Counting inner bags: outer:%s, count:%d, inner:%s\n", b.color, bc.count, bc.bag.color)
			count += bc.count * countInnerBags(bc.bag)
		}
		cache[b] = count
		return count
	}

	answer := countInnerBags(bags["shiny gold"]) - 1 // we count the bag itself so we are "off-by-one"!

	return fmt.Sprintf("%d", answer)
}
