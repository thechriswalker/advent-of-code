package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/thechriswalker/advent-of-code/aoc"
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
	var refresh string
	var checkAnswers bool
	var recordAnswers bool
	var timingOnly bool
	flag.BoolVar(&runOnlyTests, "test-only", false, "run only tests")
	flag.BoolVar(&checkAnswers, "check-answers", false, "run answer check")
	flag.BoolVar(&recordAnswers, "record-answers", false, "record current solutions as correct")
	flag.BoolVar(&timingOnly, "timing-only", false, "Show only timings")
	flag.IntVar(&prob.Year, "year", 0, "the year")
	flag.IntVar(&prob.Day, "day", 0, "the day of december")
	flag.BoolVar(&fetchProgress, "progress", false, "fetch/refresh/view progress")
	flag.StringVar(&refresh, "refresh", "", "comma-separated years to refresh")
	flag.Parse()

	if fetchProgress {
		//log.Println("Checking Progress...")
		aoc.PrintHeader(0, 0)
		checkProgress(refresh)
		return
	}
	// check all if no year.day given
	if checkAnswers && prob.Day == 0 {
		checkAllAnswers(prob.Year)
		return
	}

	// and now set the defaults
	// but only in December, and only if no year is set.
	if time.Now().Month() == time.December {
		if prob.Year == 0 {
			prob.Year = time.Now().Year()

			if prob.Day == 0 {
				prob.Day = time.Now().Day()
			}
		}
	}
	currentYear := time.Now().Year()
	// for the timing, we can cope with 0 day
	if timingOnly {
		runTimings(prob.Year, prob.Day)
		return
	}

	if prob.Year == 0 || prob.Year > currentYear || prob.Day == 0 || prob.Day > 25 {
		log.Fatalf("bad input year/day: year=%d, day=%d", prob.Year, prob.Day)
	}

	basePath := fmt.Sprintf("%d/%02d", prob.Year, prob.Day)
	if err := os.MkdirAll(basePath, 0755); err != nil {
		log.Fatalf("could not make directorys: %s", basePath)
	}
	f, err := os.Open(basePath + "/main.go")
	if err != nil {
		if os.IsNotExist(err) {
			// create the files
			if err := createFiles(prob, basePath); err != nil {
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
	if err := os.Chdir(basePath); err != nil {
		// this should not happen if the previous call succeeded
		log.Fatalf("could not change to directorys: %s", basePath)
	}
	// file exists run it!
	var arg3 string
	if recordAnswers {
		arg3 = "-record-answers=true"
	} else if checkAnswers {
		arg3 = "-check-answers=true"
	} else if runOnlyTests {
		arg3 = "-test-only=true"
	} else {
		arg3 = "-test-only=false"
	}
	run := exec.Command("go", "run", "main.go", arg3)
	run.Stdin = os.Stdin
	run.Stderr = os.Stderr
	run.Stdout = os.Stdout
	run.Run()
}

func createFiles(p Problem, basePath string) error {
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
		file, err := os.Create(basePath + "/" + f.name)
		if err != nil {
			return err
		}
		if f.template != nil {
			err = f.template.Execute(file, p)
		}
		if f.name == "input.txt" {
			// let's try to fetch from AOC.
			cookie, err := os.ReadFile(".aoc-cookie")
			if err != nil {
				log.Println("We can fetch the input data from AoC if we have the cookie!")
				log.Println("Could not read cookie from `./.aoc-cookie`, please make sure it is present. Error:", err)
			} else {
				req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", p.Year, p.Day), nil)
				if err == nil {
					req.AddCookie(&http.Cookie{
						Name:  "session",
						Value: string(cookie),
					})
					res, err := http.DefaultClient.Do(req)
					if err != nil {
						log.Printf("Failed to fetch input data. Bad cookie? %s\n", err)
					} else if res.StatusCode != http.StatusOK {
						log.Printf("Failed to fetch input data. Bad cookie? (status:%d)\n", res.StatusCode)
					} else {
						io.Copy(file, res.Body)
						res.Body.Close()
					}
				}
			}
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
	"github.com/thechriswalker/advent-of-code/aoc"
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

const example = ""

var problem1cases = []Case{
	// cases here
	{example, ""},
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
	{example, ""},
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

func checkProgress(refresh string) {
	yearsToRefresh := []int{}
	refreshAll := refresh == "all"
	if refresh != "" && !refreshAll {
		for _, y := range strings.Split(refresh, ",") {
			if n, err := strconv.Atoi(y); err == nil {
				yearsToRefresh = append(yearsToRefresh, n)
			} else {
				panic("Bad Year in -refresh arg: " + y)
			}

		}

	}
	// we need a cookie
	cookie, err := os.ReadFile(".aoc-cookie")
	if err != nil {
		log.Fatalln("Could not read cookie from `./.aoc-cookie`, please make sure it is present. Error:", err)
	}
	//log.Println("found cookie", string(cookie))
	// now for each directory of problems, see if the progress file exists.
	dir, _ := os.Open(".")
	names, _ := dir.Readdirnames(-1)
	thisYear := time.Now().Year()
	if time.Now().Month() < time.December {
		thisYear--
	}
	years := []int{}
	for _, name := range names {
		year, err := strconv.Atoi(name)
		if err == nil && year >= 2014 && year <= thisYear {
			years = append(years, year)
		}
	}
	sort.Ints(years)
	//log.Println("found years", years)
	for _, year := range years {
		// see if we have a progress file.
		file := fmt.Sprintf("./%d/.progress.json", year)
		f, err := os.Open(file)
		progress := &YearProgress{}
		// load new
		load := refreshAll || contains(year, yearsToRefresh) || err != nil
		if !load {
			// try and read.
			if err := json.NewDecoder(f).Decode(progress); err != nil {
				// failed to read, try and load again.
				load = true
			}
			f.Close()
		}
		if load {
			// that's OK, just load it new
			progress = loadAndSaveProgress(year, file, strings.TrimSpace(string(cookie)))
		}
		// now just need a pretty way to display the data in a table.
		fmt.Printf("%s", progress)
	}
}

func contains(needle int, haystack []int) bool {
	for _, i := range haystack {
		if i == needle {
			return true
		}
	}
	return false
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
	body, err := io.ReadAll(res.Body)
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
	fmt.Fprintf(&b, "\x1b[1;90m[%s]\x1b[0m %d ", yp.Updated.Format(time.ANSIC), yp.Year)

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

func checkAllAnswers(oneYearOnly int) {
	// loop through all dirs.
	// we start in 2015
	// and go to today.
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get working dir: %s", err)
	}
	fmt.Println("              1111111111222222")
	fmt.Println("     1234567890123456789012345")
	buf := &strings.Builder{}
	for year := 2015; year <= time.Now().Year(); year++ {
		if oneYearOnly > 0 && oneYearOnly != year {
			continue
		}
		fmt.Printf("%d ", year)
		for day := 1; day < 26; day++ {
			// check if we have a main.go in that folder.
			basePath := fmt.Sprintf("%s/%d/%02d", cwd, year, day)
			if err := os.Chdir(basePath); err != nil {
				// probably doesn't exist.
				continue
			}
			// otherwise try to run it!
			buf.Reset()
			run := exec.Command("go", "run", "main.go", "-check-answers=true")
			run.Stdin = os.Stdin
			run.Stderr = buf
			run.Stdout = io.Discard
			run.Run()
			// we have to parse the exit code from the last line of stderr.
			stderr := buf.String()
			lines := strings.Split(strings.TrimSpace(stderr), "\n")
			exit := 0

			if len(lines) > 0 {
				lastline := lines[len(lines)-1]
				fmt.Sscanf(lastline, "exit status %d", &exit)
			}

			switch exit {
			case 100: // both passed
				fmt.Print("\x1b[1;93m*")
			case 101: // one passed
				fmt.Print("\x1b[1;97m*")
			case 102: // none passed
				fmt.Print("\x1b[1;90m*")
			default:
				// probably an error...
				fmt.Print("\x1b[1;31me")
			}
		}
		fmt.Print("\x1b[0m\n")
	}
}

func runTimings(thisYearOnly, thisDayOnly int) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get working dir: %s", err)
	}
	for year := 2015; year <= time.Now().Year(); year++ {
		if thisYearOnly > 0 && thisYearOnly != year {
			continue
		}
		for day := 1; day < 26; day++ {
			if thisDayOnly > 0 && thisDayOnly != day {
				continue
			}
			printTiming(year, day, cwd)
		}
	}
}

func printTiming(year, day int, cwd string) {
	// does the code exist?
	basePath := fmt.Sprintf("%s/%d/%02d", cwd, year, day)
	if err := os.Chdir(basePath); err != nil {
		// probably doesn't exist.
		return
	}
	run := exec.Command("go", "run", "main.go", "-timing-only")
	run.Stdin = os.Stdin
	run.Stderr = io.Discard
	run.Stdout = os.Stdout
	run.Run()
}
