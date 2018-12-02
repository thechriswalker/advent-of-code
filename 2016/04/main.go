package main

import (
	"fmt"
	"sort"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 4, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	rooms := parseRooms(input)
	sum := 0
	for _, r := range rooms {
		if !r.IsDecoy() {
			sum += r.Sector
		}
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	rooms := parseRooms(input)
	for _, r := range rooms {
		if !r.IsDecoy() {
			name := r.DecryptName()
			// What is the sector ID of the room where North Pole objects are stored?
			if name == "northpole object storage" {
				return fmt.Sprintf("%d", r.Sector)
			}
		}
	}
	return "<unsolved>"
}

func parseRooms(s string) []Room {
	rooms := []Room{}
	r := strings.NewReader(s)
	var line string
	for {
		n, _ := fmt.Fscanf(r, "%s\n", &line)
		if n != 1 {
			break
		}
		rooms = append(rooms, NewRoom(line))
	}
	return rooms

}

type Room struct {
	Name     string
	Checksum string
	Sector   int
}

type Runes struct {
	List []rune
	Map  map[rune]int
}

func (r Runes) Len() int      { return len(r.List) }
func (r Runes) Swap(i, j int) { r.List[i], r.List[j] = r.List[j], r.List[i] }
func (r Runes) Less(i, j int) bool {
	ri, rj := r.List[i], r.List[j]
	ci, cj := r.Map[ri], r.Map[rj]

	if ci > cj {
		return true
	}
	if ci < cj {
		return false
	}
	return ri < rj
}

// A room is real (not a decoy) if the checksum is the five most common letters in the encrypted
// name, in order, with ties broken by alphabetization
func (r *Room) IsDecoy() bool {
	runes := &Runes{
		List: []rune{},
		Map:  map[rune]int{},
	}
	for _, c := range r.Name {
		if c == '-' {
			continue
		}
		n, ok := runes.Map[c]
		runes.Map[c] = n + 1
		if !ok {
			runes.List = append(runes.List, c)
		}
	}
	sort.Sort(runes)
	checksum := fmt.Sprintf("[%c%c%c%c%c]", runes.List[0], runes.List[1], runes.List[2], runes.List[3], runes.List[4])
	return checksum != r.Checksum
}

func (r *Room) DecryptName() string {
	builder := strings.Builder{}
	for _, c := range r.Name {
		if c == '-' {
			builder.WriteByte(' ')
		} else {
			builder.WriteByte(shift(c, r.Sector))
		}
	}
	return builder.String()
}

// a-z wrapped around.
// a == 97
func shift(b rune, n int) byte {
	x := (int(b) - 97 + n) % 26
	return byte(x + 97)
}

func NewRoom(s string) Room {
	r := Room{}
	i := strings.LastIndexByte(s, '-')
	r.Name = s[:i]
	fmt.Sscanf(s[i:], "-%d%s\n", &r.Sector, &r.Checksum)
	return r
}
