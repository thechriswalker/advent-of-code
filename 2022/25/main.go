package main

import (
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 25, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	sum := Snafu(0)
	aoc.MapLines(input, func(s string) error {
		sum += ParseSnafu(s)
		return nil
	})

	return sum.String()
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return "<unsolved>"
}

type Snafu int

func (s Snafu) String() string {
	// finding the string representation of a snafu is harder...
	// find the value mod 5.
	// if > 2 we need to add 5 to the next digit and make this a -1 or -2
	// the again with a multiply by 5 and continue...
	digits := []byte{}
	n := s
	for {
		d := int(n % 5)
		switch d {
		case 0:
			// fine, no 5s.
			digits = append(digits, '0')
		case 1:
			digits = append(digits, '1')
			// now subtract that one from the n
			n--
		case 2:
			digits = append(digits, '2')
			// now subtract that one from the n
			n -= 2
		case 3:
			// this is a -2 from the next digit
			digits = append(digits, '=')
			// so to get to a multiple of 5 we need to add 2 to the next digit
			n += 2
		case 4:
			// this is a -1 from the next digit
			digits = append(digits, '-')
			// so to get to a multiple of 5 we need to add 1 to the next digit
			n++
		}
		if n == 0 {
			break
		}
		n /= 5
	}

	sb := strings.Builder{}
	for i := len(digits) - 1; i >= 0; i-- {
		sb.WriteByte(digits[i])
	}

	return sb.String()
}

func ParseSnafu(input string) (s Snafu) {
	// parsing is easy.
	m := 1 // 5^0
	for d := len(input) - 1; d >= 0; d-- {
		c := int(input[d] - '0')
		switch input[d] {
		case '2', '1', '0':
			// easy already done
		case '-':
			c = -1
		case '=':
			c = -2
		default:
			panic("bad input")
		}
		s += Snafu(c * m)
		m *= 5
	}

	return
}
