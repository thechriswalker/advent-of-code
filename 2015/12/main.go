package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 12, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// we don't reallyneed to JSON parse,
	//but go has a nice streaming parser
	// so we might as well us it.
	sum := int64(0)
	dec := json.NewDecoder(strings.NewReader(input))
	dec.UseNumber()
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if n, ok := t.(json.Number); ok {
			v, _ := n.Int64()
			sum += v
		}
	}
	return fmt.Sprintf("%d", sum)
}

func sumArray(dec *json.Decoder) int64 {
	sum := int64(0)
	for {
		t, err := dec.Token()
		if err == io.EOF {
			return sum
		}
		if err != nil {
			log.Fatal(err)
		}
		switch x := t.(type) {
		case json.Delim:
			switch x {
			case '[':
				sum += sumArray(dec)
			case ']':
				return sum
			case '{':
				sum += sumObject(dec)
			}
		case json.Number:
			v, _ := x.Int64()
			sum += v
		}
	}
}

func sumObject(dec *json.Decoder) int64 {
	sum := int64(0)
	// we assume that this is the _start_ of an object.
	// if any value is "red", then return 0;
	foundRed := false
	for {
		// key first
		t, err := dec.Token()
		if err == io.EOF {
			return sum
		}
		if err != nil {
			log.Fatal(err)
		}
		// this is a KEY or close object
		switch x := t.(type) {
		case string:
			// expected
			//fmt.Println("KEY:", x)
		case json.Delim:
			if x != '}' {
				panic("expecting object end!")
			}
			if foundRed {
				return 0
			}
			return sum
		default:
			fmt.Printf("BAD KEY: %T: %v\n", t, t)
			panic("expecting string key or object end")
		}
		// now grab another for the value:
		t, err = dec.Token()
		if err == io.EOF {
			return sum
		}
		if err != nil {
			log.Fatal(err)
		}
		// this is VALUE
		switch x := t.(type) {
		case string:
			if x == "red" {
				foundRed = true
			}
		case json.Delim:
			switch x {
			case '[':
				sum += sumArray(dec)
			case '{':
				sum += sumObject(dec)
			default:
				panic("unexpected delim in object")
			}
		case json.Number:
			v, _ := x.Int64()
			sum += v
		default:
			// some other json value.
		}
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	sum := int64(0)
	dec := json.NewDecoder(strings.NewReader(input))

	dec.UseNumber()
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if d, ok := t.(json.Delim); ok {
			if d == '{' {
				sum += sumObject(dec)
			}
		} else {
			// if this is the start of a number, we should break out
			if n, ok := t.(json.Number); ok {
				v, _ := n.Int64()
				sum += v
			}
		}
	}
	return fmt.Sprintf("%d", sum)
}
