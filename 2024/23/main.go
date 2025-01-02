package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 23, solve1, solve2)
}

type PC struct {
	name        string
	connections map[string]*PC
}

// Implement Solution to Problem 1
func solve1(input string) string {
	network := map[string]*PC{}

	aoc.MapLines(input, func(line string) error {
		a, b, _ := strings.Cut(line, "-")
		if _, ok := network[a]; !ok {
			network[a] = &PC{name: a, connections: map[string]*PC{}}
		}
		if _, ok := network[b]; !ok {
			network[b] = &PC{name: b, connections: map[string]*PC{}}
		}
		network[a].connections[b] = network[b]
		network[b].connections[a] = network[a]

		return nil
	})

	// we are looking for loops of 3. so aa -> bb -> cc --> aa
	// this is the same as aa -> bb and aa -> cc and bb -> cc
	// i.e. for machine aa, find connections whose connection are also connected to aa

	// we have to "walk" the network to find all connected machines.
	// assigning groups

	threes := map[string]struct{}{}

	for _, pc := range network {
		connections := make([]string, 0, len(pc.connections))
		for _, conn := range pc.connections {
			connections = append(connections, conn.name)
		}
		for i, a := range connections {
			for j, b := range connections {
				if i == j {
					continue
				}
				// are these 2 connections of pc, connected to each other??
				if _, ok := network[a].connections[b]; ok {
					// if any of pc.name, a, and b start with a 't' then we can add them
					if strings.HasPrefix(pc.name, "t") || strings.HasPrefix(a, "t") || strings.HasPrefix(b, "t") {

						key := []string{pc.name, a, b}
						sort.Strings(key)
						cache := strings.Join(key, ",")
						threes[cache] = struct{}{}
					}
				}
			}
		}
	}

	return fmt.Sprint(len(threes))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	network := map[string]*PC{}
	list := []*PC{}

	aoc.MapLines(input, func(line string) error {
		a, b, _ := strings.Cut(line, "-")
		if _, ok := network[a]; !ok {
			network[a] = &PC{name: a, connections: map[string]*PC{}}
			list = append(list, network[a])
		}
		if _, ok := network[b]; !ok {
			network[b] = &PC{name: b, connections: map[string]*PC{}}
			list = append(list, network[b])
		}
		network[a].connections[b] = network[b]
		network[b].connections[a] = network[a]

		return nil
	})

	// we are looking to sort the computers into groups.
	// so we can find the "largest" group.
	// a group is defined by interconnected computers.

	// so we find a group we need to start with a computer
	// and find all it's connections.
	// the for each of it's connections see if they are also connected to the group.

	// we will take the PC's in pairs, and for each pair, we will check if they are connected to each other,
	// if they are then we will check if any of their connections are also connected to the current members of the group.

	// let us find the sets of three.
	// then we can see if we can make any sets of four by adding a new computer to the group.
	// if so, thake the sets of four and try to make them 5, and so on, until we can't make any groups larger.

	groups := map[string][]string{}

	for _, pc := range network {
		connections := make([]string, 0, len(pc.connections))
		for _, conn := range pc.connections {
			connections = append(connections, conn.name)
		}
		for i, a := range connections {
			for j, b := range connections {
				if i == j {
					continue
				}
				// are these 2 connections of pc, connected to each other??
				if _, ok := network[a].connections[b]; ok {
					key := []string{pc.name, a, b}
					sort.Strings(key)
					cache := strings.Join(key, ",")
					groups[cache] = key

				}
			}
		}
	}

	// now we try and make the groups larger
	// this is massively inefficient, but it works. >3.5 seconds
	for {
		increased := false
		nextGroups := map[string][]string{}
		for _, group := range groups {
			size := len(group)
			// go through the connections of each member of the group
			// and see if there is another that is part of all of this
			for _, name := range group {
				pc := network[name]
				for _, conn := range pc.connections {
					if containsAllAndIsNew(conn.name, group, conn.connections) {
						newGroup := make([]string, size+1)
						copy(newGroup, group)
						newGroup[size] = conn.name
						sort.Strings(newGroup)
						cache := strings.Join(newGroup, ",")
						nextGroups[cache] = newGroup
						increased = true
					}
				}
			}
		}
		if !increased {
			break
		}
		groups = nextGroups
	}
	// there should be only one group left.
	for k := range groups {
		return k
	}
	return ""
}

func containsAllAndIsNew(name string, group []string, connections map[string]*PC) bool {
	for _, g := range group {
		if g == name {
			return false
		}
		if _, ok := connections[g]; !ok {
			return false
		}
	}
	return true
}
