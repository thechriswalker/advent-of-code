package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	"./aoc"
)

type Problem struct {
	Year int
	Day  int
}

// generates a skeleton for the given day.
func main() {
	// input should be year day
	prob := Problem{}
	var runOnlyTests bool
	var fetchProgress bool
	flag.BoolVar(&runOnlyTests, "test-only", false, "run only tests")
	flag.IntVar(&prob.Year, "year", time.Now().Year(), "the year")
	flag.IntVar(&prob.Day, "day", time.Now().Day(), "the day of december")
	flag.BoolVar(&fetchProgress, "progress", false, "fetch/update/view progress")
	flag.Parse()

	if fetchProgress {
		//log.Println("Checking Progress...")
		aoc.PrintHeader(0, 0)
		checkProgress()
		return
	}

	basePath := fmt.Sprintf("%d/%02d", prob.Year, prob.Day)
	if err := os.MkdirAll(basePath, 0755); err != nil {
		log.Fatalf("could not make directorys: %s", basePath)
	}
	if err := os.Chdir(basePath); err != nil {
		// this should not happen if the previous call succeeded
		log.Fatalf("could not change to directorys: %s", basePath)
	}
	f, err := os.Open("main.go")
	if err != nil {
		if os.IsNotExist(err) {
			// create the files
			if err := createFiles(prob); err != nil {
				log.Fatalln("error creating problem templates:", err)
			}
			fmt.Println("Created problem template for", basePath)
			fmt.Println("------------------------------------")
		} else {
			log.Fatalf("error checking file: %s/main.go", basePath)
		}
	} else {
		f.Close()
	}
	// file exists run it!
	var arg3 string
	if runOnlyTests {
		arg3 = "-test-only=true"
	} else {
		arg3 = "-test-only=false"
	}
	run := exec.Command("go", "run", "main.go", arg3)
	run.Stderr = os.Stderr
	run.Stdout = os.Stdout
	run.Run()
}

func createFiles(p Problem) error {
	files := []struct {
		name     string
		template *template.Template
	}{
		{"README.md", readmeTpl},
		{"main.go", mainTpl},
		{"main_test.go", testTpl},
		{"input.txt", nil},
	}

	for _, f := range files {
		file, err := os.Create(f.name)
		if err != nil {
			return err
		}
		if f.template != nil {
			err = f.template.Execute(file, p)
		}
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

var mainTpl = template.Must(template.New("main").Parse(`package main

import (
	"../../aoc"
)

func main() {
	aoc.Run({{.Year}}, {{.Day}}, solve1, solve2)
}


// Implement Solution to Problem 1
func solve1(input string) string {
	return "<unsolved>"
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return "<unsolved>"
}
`))

var testTpl = template.Must(template.New("test").Parse(`package main

import (
	"testing"
)

// tests for the AdventOfCode {{.Year}} day {{.Day}} solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"", ""},
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
	{"", ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
`))

var readmeTpl = template.Must(template.New("test").Parse(`# Advent of Code {{.Year}} day {{.Day}}

## Problem 1

...

## Problem 2

...
`))

func checkProgress() {
	// we need a cookie
	cookie, err := ioutil.ReadFile(".aoc-cookie")
	if err != nil {
		log.Fatalln("Could not read cookie from `./.aoc-cookie`, please make sure it is present. Error:", err)
	}
	//log.Println("found cookie", string(cookie))
	// now for each directory of problems, see if the progress file exists.
	dir, _ := os.Open(".")
	names, _ := dir.Readdirnames(-1)
	thisYear := time.Now().Year()
	years := []int{}
	for _, name := range names {
		year, err := strconv.Atoi(name)
		if err == nil && year >= 2014 && year <= thisYear {
			years = append(years, year)
		}
	}
	sort.Sort(sort.IntSlice(years))
	//log.Println("found years", years)
	for _, year := range years {
		// see if we have a progress file.
		file := fmt.Sprintf("./%d/.progress.json", year)
		f, err := os.Open(file)
		progress := &YearProgress{}
		if err != nil {
			// that's OK, just load it new
			progress = loadAndSaveProgress(year, file, strings.TrimSpace(string(cookie)))
		} else {
			json.NewDecoder(f).Decode(progress)
			f.Close()
		}
		// now just need a pretty way to display the data in a table.
		fmt.Printf("%s", progress)
	}
}

func loadAndSaveProgress(year int, file, cookie string) *YearProgress {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/leaderboard/self", year), nil)
	if err != nil {
		panic(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: cookie,
	})

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	progress := &YearProgress{
		Year: year,
	}
	err = parseStatsHTML(string(body), progress)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(progress)
	return progress
}

func parseStatsHTML(body string, progress *YearProgress) error {
	// we should test this one! so as not to hammer the AoC.
	// let's find the start, the second match of this.
	idx := strings.Index(body, " Score</span>\n")
	if idx < 0 {
		return fmt.Errorf("stats section start not found")
	}
	body = body[idx+14:]
	// the end is the </pre>
	idx = strings.Index(body, `</pre>`)
	if idx < 0 {
		return fmt.Errorf("stats section end not found")
	}
	body = body[:idx]
	lines := strings.Split(body, "\n")
	progress.Updated = time.Now()
	progress.Days = make([]DayProgress, 0, len(lines))
	getTime := func(t string) string {
		if t == "&gt;24h" {
			return "over 24 hours"
		}
		return t
	}
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 7 {
			// good
			day, _ := strconv.Atoi(fields[0])
			time1 := getTime(fields[1])
			rank1, _ := strconv.Atoi(fields[2])
			score1, _ := strconv.Atoi(fields[3])
			time2 := getTime(fields[4])
			rank2, _ := strconv.Atoi(fields[5])
			score2, _ := strconv.Atoi(fields[6])
			dp := DayProgress{
				Day: day,
				Part1: &PartProgress{
					Time:  time1,
					Rank:  rank1,
					Score: score1,
				},
			}
			// time2 will be nil if we haven't done it
			if time2 != "-" {
				dp.Part2 = &PartProgress{
					Time:  time2,
					Rank:  rank2,
					Score: score2,
				}
			}
			progress.Days = append(progress.Days, dp)
		}
	}
	return nil
}

type YearProgress struct {
	Updated time.Time
	Year    int
	Days    []DayProgress
}
type DayProgress struct {
	Day   int
	Part1 *PartProgress
	Part2 *PartProgress
}

type PartProgress struct {
	Time  string
	Rank  int
	Score int
}

func (yp *YearProgress) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "\x1b[1;90m[%s]\x1b[0m %d ", yp.Updated.Format(time.Stamp), yp.Year)

	// first we'll get the days into a map.
	days := make(map[int]DayProgress, len(yp.Days))
	for _, d := range yp.Days {
		days[d.Day] = d
	}

	// lets count yellow and grey stars.
	for i := 1; i <= 25; i++ {
		d, ok := days[i]
		if !ok || d.Part1 == nil {
			// zero stars. grey (bright black) color.
			fmt.Fprint(&b, "\x1b[1;90m*")
		} else {
			if d.Part2 == nil {
				// one star
				fmt.Fprint(&b, "\x1b[1;97m*")
			} else {
				// two stars!
				fmt.Fprint(&b, "\x1b[1;93m*")
			}
		}
	}
	fmt.Fprintln(&b, "\x1b[0m")
	// let's use fixed sizes for the results
	// so we can just iterate!

	return b.String()
}
