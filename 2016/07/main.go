package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 7, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	count := 0
	for _, ip := range parseIPv7(input) {
		if ip.SupportsTLS() {
			count++
		}
	}
	return fmt.Sprintf("%d", count)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	count := 0
	for _, ip := range parseIPv7(input) {
		if ip.SupportsSSL() {
			count++
		}
	}
	return fmt.Sprintf("%d", count)
}

func parseIPv7(in string) []*IPv7 {
	out := []*IPv7{}
	curr := &IPv7{
		Net: []string{},
		Hyp: []string{},
	}
	s := strings.Builder{}
	for _, c := range in {
		switch c {
		case '\n', '[':
			// end of net section
			net := s.String()
			if len(net) > 0 {
				curr.Net = append(curr.Net, net)
			}
			s.Reset()
			if c == '\n' {
				// also EOL
				out = append(out, curr)
				curr = &IPv7{
					Net: []string{},
					Hyp: []string{},
				}
			}
		case ']':
			hyp := s.String()
			curr.Hyp = append(curr.Hyp, hyp)
			s.Reset()
		default:
			s.WriteRune(c)
		}
	}

	net := s.String()
	if len(net) > 0 {
		curr.Net = append(curr.Net, net)
	}

	if len(curr.Net) > 0 || len(curr.Hyp) > 0 {
		out = append(out, curr)
	}
	return out
}

type IPv7 struct {
	Net []string
	Hyp []string
}

func (ip *IPv7) SupportsTLS() bool {
	// any four-character sequence which consists of a pair of two different characters
	//followed by the reverse of that pair, such as xyyx or abba. However, the IP also
	//must not have an ABBA within any hypernet sequences, which are contained by square brackets.

	// rule out those with ABBAs in hypernet sequences
	for _, h := range ip.Hyp {
		if hasABBA(h) {
			return false
		}
	}
	// find those with
	for _, n := range ip.Net {
		if hasABBA(n) {
			return true
		}
	}
	return false
}

func hasABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}

func (ip *IPv7) SupportsSSL() bool {
	// An IP supports SSL if it has an Area-Broadcast Accessor, or ABA, anywhere in
	// the supernet sequences (outside any square bracketed sections), and a corresponding
	// Byte Allocation Block, or BAB, anywhere in the hypernet sequences.
	for _, n := range ip.Net {
		for _, bab := range findABAs(n) {
			for _, h := range ip.Hyp {
				if strings.Contains(h, bab) {
					return true
				}
			}
		}
	}
	return false
}

func findABAs(s string) []string {
	found := []string{}
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i] == s[i+2] {
			found = append(found, string([]byte{s[i+1], s[i], s[i+1]}))
		}
	}
	return found
}
