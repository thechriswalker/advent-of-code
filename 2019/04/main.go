package main

import (
	"fmt"
	"log"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 4, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var min, max int
	fmt.Sscanf(input, "%d-%d", &min, &max)
	// the first password needs to be split into digits...
	p := &Password{
		digits: digits(min),
		value:  min,
	}
	// we should only have max-min increments
	// and we can cut out whole swaths of tests by being clever,
	// if 121xxx can be discarded.
	log.Println("exhaustive is", max-min, "tests")
	valid := 0
	for {
		ok := p.meetsCriteria()
		//log.Println("Checking", p.value, " => ", p.digits, " => ", ok)
		if ok {
			valid++
		}
		if !p.Inc(max) {
			break
		}
	}

	return fmt.Sprintf("%d", valid)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	var min, max int
	fmt.Sscanf(input, "%d-%d", &min, &max)
	// the first password needs to be split into digits...
	p := &Password{
		digits: digits(min),
		value:  min,
	}
	// we should only have max-min increments
	// and we can cut out whole swaths of tests by being a little smart
	//log.Println("exhaustive is", max-min, "tests")
	valid := 0
	for {
		ok := p.meetsCriteria2()
		//log.Println("Checking", p.value, " => ", p.digits, " => ", ok)
		if ok {
			valid++
		}
		if !p.Inc(max) {
			break
		}
	}

	return fmt.Sprintf("%d", valid)
}

func digits(n int) [6]uint8 {
	d := [6]uint8{}
	for i := 5; i >= 0; i-- {
		d[i] = uint8(n % 10)
		n /= 10
	}
	return d
}
func value(d [6]uint8) int {
	m := 1
	sum := 0
	for i := 5; i >= 0; i-- {
		sum += int(d[i]) * m
		m *= 10
	}
	return sum
}

type Password struct {
	digits [6]uint8
	value  int
}

func (p *Password) meetsCriteria() bool {
	// need a double, must never decrease
	foundDouble := false
	for i := 1; i < 6; i++ {
		if p.digits[i] < p.digits[i-1] {
			return false
		}
		if p.digits[i] == p.digits[i-1] {
			foundDouble = true
		}
	}
	return foundDouble
}

func (p *Password) meetsCriteria2() bool {
	// need a double, must never decrease
	// but double must be only 2 of the same.
	foundDouble := false
	for i := 1; i < 6; i++ {
		if p.digits[i] < p.digits[i-1] {
			return false
		}
		if p.digits[i] == p.digits[i-1] {
			var okFwd, okBck bool
			if i == 5 {
				// the end
				okFwd = true
			} else if p.digits[i] != p.digits[i+1] {
				// not the same as the next one
				okFwd = true
			}
			if i == 1 {
				okBck = true
			} else if p.digits[i] != p.digits[i-2] {
				okBck = true
			}
			if okBck && okFwd {
				foundDouble = true
			}
		}
	}
	return foundDouble
}

func (p *Password) Inc(max int) bool {
	// let make this function cleverer.
	// if the bump overflows, then we only
	// need to set the
	//bump the final digit by one.
	var bump func(idx int) uint8
	bump = func(idx int) uint8 {
		if p.digits[idx] == 9 {
			if idx == 0 {
				panic("password out of range!")
			}
			p.digits[idx] = bump(idx - 1)
		} else {
			p.digits[idx] = p.digits[idx] + 1
		}
		return p.digits[idx]
	}
	bump(5)
	p.value = value(p.digits)
	return p.value <= max
}
