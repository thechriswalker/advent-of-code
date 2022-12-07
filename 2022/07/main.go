package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 7, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	_, list := parseTree(input)
	// find all dirs that are less than 100000
	limit := 100000

	total := 0

	for _, d := range list {
		s := d.Size()
		if s <= limit {
			total += s
		}
	}

	return fmt.Sprint(total)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	root, list := parseTree(input)
	used := root.Size()
	free := 70000000 - used
	target := 30000000 - free

	smallest := used

	for _, d := range list {
		s := d.Size()
		if s < smallest && s >= target {
			smallest = s
		}
	}

	return fmt.Sprint(smallest)
}

// returns root and flat list of dirs
func parseTree(input string) (*Dir, []*Dir) {
	root := &Dir{
		Name:     "/",
		Parent:   nil,
		Children: map[string]*Dir{},
		Files:    map[string]*File{},
	}
	all := []*Dir{root}
	current := root
	aoc.MapLines(input, func(line string) error {
		l := strings.Split(line, " ")
		switch l[0] {
		case "$":
			// command ls / cd
			// we should know the dir...
			// and ignore ls
			switch l[1] {
			case "cd":
				// handle cd.
				switch l[2] {
				case "/":
					// to root
					current = root

				case "..":
					// up
					current = current.Parent
				default:
					// dir in current directory.
					d, ok := current.Children[l[2]]
					if !ok {
						d = &Dir{
							Name:     l[2],
							Parent:   current,
							Children: map[string]*Dir{},
							Files:    map[string]*File{},
						}
						current.Children[l[2]] = d
						all = append(all, d)
					}
					current = d
				}
			}
		case "dir":
			// dir in current dir
			d := &Dir{
				Name:     l[1],
				Parent:   current,
				Children: map[string]*Dir{},
				Files:    map[string]*File{},
			}
			current.Children[l[1]] = d
			all = append(all, d)
		default:
			// size file
			s, _ := strconv.Atoi(l[0])
			current.Files[l[1]] = &File{
				Name: l[1],
				Dir:  current,
				Size: s,
			}
		}
		return nil
	})

	return root, all
}

type Dir struct {
	Name     string
	Parent   *Dir
	Children map[string]*Dir
	Files    map[string]*File
}

type File struct {
	Name string
	Dir  *Dir
	Size int
}

func (d *Dir) Size() int {
	s := 0
	for _, f := range d.Files {
		s += f.Size
	}
	for _, c := range d.Children {
		s += c.Size()
	}
	return s
}
