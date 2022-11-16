package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 18, solve1, solve2)
}

// 1 + (2 * 3) + (4 * (5 + 6))

func eval(rd *strings.Reader) int {
	value := 0
	var nextValue int
	op := byte('+')
	for {
		// next value is either
		// - a number
		// - open paren
		b, err := rd.ReadByte()
		//fmt.Println("next value hint:", b, err)
		if err != nil {
			panic(err)
		}
		if b == '(' {
			//	fmt.Println("open paren, eval sub-expression")
			nextValue = eval(rd)
		} else {
			//fmt.Println("non-paren, scanning for int")
			rd.UnreadByte() // rewind
			_, err = fmt.Fscanf(rd, "%d", &nextValue)
			if err != nil {
				panic(err)
			}
		}
		//fmt.Println("next evaluation is:", value, string(op), nextValue)
		switch op {
		case '+':
			value = value + nextValue
		case '*':
			value = value * nextValue
		}
		//fmt.Println("value:", value)
		// see what is next.
		// - an operator
		// - a close paren
		b, err = rd.ReadByte()
		if err == io.EOF || b == ')' {
			//	fmt.Println("EOL or paren", b)
			return value
		}
		switch b {
		case '+', '*':
			op = b
		default:
			panic("b is not op: " + string(b))
		}
		//fmt.Println("next op:", string(op))
	}
}

// Implement Solution to Problem 1
func solve1(input string) string {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		if line != "" {
			line = strings.ReplaceAll(line, " ", "")
			sum += eval(strings.NewReader(line))
		}
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
// 15341515275523405 is too high...
// 129770152447927
// 92260260094336    is too low...
func solve2(input string) string {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		if line != "" {
			line = strings.ReplaceAll(line, " ", "")
			//	fmt.Println("Expr:", line)
			sum += eval2(strings.NewReader(line), "#", breakOnEOF)
		}
	}
	return fmt.Sprintf("%d", sum)
}

type breaker uint8

const (
	breakOnEOF breaker = iota
	breakOnParen
	breakOnMultiply
)

func eval2(rd *strings.Reader, depth string, brk breaker) int {
	value := 0
	var nextValue int
	op := byte('+')
	for {
		// loop start
		//	fmt.Printf("%sloop start:value=%d, last nextValue=%d, op=%c\n", depth, value, nextValue, op)
		// next value is either
		// - a number
		// - open paren
		b, err := rd.ReadByte()
		//	fmt.Printf("%snext value hint: %c %v\n", depth, b, err)
		if err != nil {
			panic(err)
		}
		if b == '(' {

			if op == '*' {
				rd.UnreadByte()
				nextValue = eval2(rd, depth+" >", breakOnMultiply)
			} else {
				nextValue = eval2(rd, depth+" >", breakOnParen)
			}

		} else {
			rd.UnreadByte() // rewind
			if op == '+' {
				// if a plus, directly apply
				//			fmt.Printf("%snon-paren, op = +, scanning for int\n", depth)
				_, err = fmt.Fscanf(rd, "%d", &nextValue)
				if err != nil {
					panic(err)
				}
				//			fmt.Printf("%sFound %d, so adding: %d + %d = %d\n", depth, nextValue, value, nextValue, value+nextValue)
			} else {
				//			fmt.Printf("%snon-paren, op = *, eval RHS\n", depth)
				// multiply, we need to eval the right hand side first.
				// this is "stop on multiply..."
				nextValue = eval2(rd, depth+" >", breakOnMultiply)
				//			fmt.Printf("%srhs done: eval(%d * %d) = %d\n", depth, value, nextValue, value*nextValue)
			}
		}

		// see what is next.
		// - an operator
		// - a close paren
		b, err = rd.ReadByte()

		switch op {
		case '+':
			// if value != 0 {
			// 	fmt.Printf("%sEVAL(%d + %d) = %d\n", depth, value, nextValue, value+nextValue)
			// }
			value += nextValue

		case '*':
			// if value != 0 {
			// 	fmt.Printf("%sEVAL(%d * %d) = %d\n", depth, value, nextValue, value*nextValue)
			// }
			value *= nextValue
		default:
			panic("bad op")
		}
		if err != nil {
			if err == io.EOF {
				// fmt.Printf("%sEXPR: %d\n", depth, value)
				return value
			}
			panic(err)
		}
		if b == ')' {
			if brk != breakOnParen {
				rd.UnreadByte()
			}
			// fmt.Printf("%sEXPR: %d\n", depth, value)
			return value
		}

		switch b {
		case '+', '*':
			op = b
		default:
			panic("b is not op: " + string(b))
		}
		if op == '*' && brk == breakOnMultiply {
			rd.UnreadByte()
			//	fmt.Printf("%sEXPR: %d\n", depth, value)
			return value
		}
		//fmt.Printf("%snext op=%c, current value=%d\n", depth, op, value)
	}
}
