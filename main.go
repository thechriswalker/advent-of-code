package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
	"time"
)

type Problem struct {
	Year int
	Day  int
}

// generates a skeleton for the given day.
func main() {
	// input should be year day
	prob := Problem{}
	flag.IntVar(&prob.Year, "year", time.Now().Year(), "the year")
	flag.IntVar(&prob.Day, "day", time.Now().Day(), "the day of december")
	flag.Parse()

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
	run := exec.Command("go", "run", "main.go")
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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("##############################")
	fmt.Println("#                            #")
	//          "#  AdventOfCode YYYY Day DD  #
	fmt.Println("#  AdventOfCode {{.Year}} Day {{ printf "%02d" .Day}}  #")
	fmt.Println("#                            #")
	fmt.Println("##############################")
	fmt.Println()
	fmt.Print("Running tests: ")
	tests, err := exec.Command("go", "test").Output()
	if err != nil {
		fmt.Println("FAIL")
		os.Stdout.Write(tests)
		os.Exit(1)
	}
	fmt.Println("PASS")
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("Please create 'input.txt' with your problem input")
		}
		log.Fatalln("Error trying to read input file ('input.txt'):", err)
	}
	input := string(b)

	fmt.Print("Running problem 1: ")
	fmt.Println(solve1(input))
	fmt.Print("Running problem 2: ")
	fmt.Println(solve2(input))
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
