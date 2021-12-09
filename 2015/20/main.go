package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 20, solve1, solve2)
}

// thought I might need this
//var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53}

func numPresents(house int) int {
	sqrt := int(math.Ceil(math.Sqrt(float64(house))))
	cache := map[int]struct{}{}
	cache[house] = struct{}{}
	cache[1] = struct{}{}
	for i := 2; i <= sqrt; i++ {
		if house%i == 0 {
			cache[i] = struct{}{}
			cache[house/i] = struct{}{}
		}
	}
	//fmt.Println("house:", house, "divisors:", cache)
	sum := 0
	for v := range cache {
		sum += 10 * v
	}
	return sum
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// basically each house i get 10 * (sum of numbers < i where n % i == 0)
	// the numbers are the factors and combinations of the factors
	target, _ := strconv.Atoi(strings.TrimSpace(input))

	for house := 100000; house < 2000000; house++ {
		p := numPresents(house)
		if p > target {
			return fmt.Sprintf("%d", house)
		}
		// if house < 20 {
		// 	fmt.Println("house:", house, "presents:", p)
		// } else {
		// 	panic("")
		// }
	}
	return "nope"
}

// Implement Solution to Problem 2
func solve2(input string) string {
	target, _ := strconv.Atoi(strings.TrimSpace(input))
	// my first try gave 4989600m which I now know is "too big"
	for house := 786240; house < 4989600; house++ {
		p := numPresents2(house)
		if p > target {
			return fmt.Sprintf("%d", house)
		}
	}
	return "nope"
}

func numPresents2(house int) int {
	sqrt := int(math.Ceil(math.Sqrt(float64(house))))
	cache := map[int]struct{}{}
	cache[house] = struct{}{}
	cache[1] = struct{}{}
	for i := 2; i <= sqrt; i++ {
		if house%i == 0 {
			cache[i] = struct{}{}
			cache[house/i] = struct{}{}
		}
	}
	//fmt.Println("house:", house, "divisors:", cache)
	sum := 0
	for v := range cache {
		// only sum if this elf has delivered to <= 50 houses.
		if house/v <= 50 {
			// now 11 presents
			sum += 11 * v
		}
	}
	return sum
}
