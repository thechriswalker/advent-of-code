package main

import (
	"testing"
)

// tests for the AdventOfCode 2021 day 16 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"D2FE28\n", "6"},
	{"38006F45291200\n", "9"},
	{"8A004A801A8002F478\n", "16"},
	{"620080001611562C8802118E34\n", "12"},
	{"C0015000016115A2E0802F182340\n", "23"},
	{"A0016C880162017C3686B18A3D4780\n", "31"},
}

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	{"C200B40A82", "3"},
	{"04005AC33890", "54"},
	{"880086C3E88112", "7"},
	{"CE00C43D881120", "9"},
	{"D8005AC2A8F0", "1"},
	{"F600BC2D8F", "0"},
	{"9C005AC2F8F0", "0"},
	{"9C0141080250320F1802104A08", "1"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
