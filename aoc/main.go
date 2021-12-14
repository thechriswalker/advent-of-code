package aoc

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

func PrintHeader(year, day int) {
	date := "       "
	url := "https://adventofcode.com/"
	if year != 0 {
		url += fmt.Sprintf("%d", year)
		if day != 0 {
			url += fmt.Sprintf("/day/%d", day)
			date = fmt.Sprintf("%d/%02d", year, day)
		} else {
			date = fmt.Sprintf(" %d  ", year)
		}
	}

	fmt.Println()
	fmt.Println("         ┏━━━━━━━━━━━━━━━━━\x1b[1;93m*\x1b[0m━━━┓")
	fmt.Println("         ┃                 \x1b[1;32m#\x1b[0m   ┃")
	fmt.Println("         ┃  \x1b[1;37mAdventOfCode  \x1b[1;32m###\x1b[0m  ┃")
	fmt.Printf("         ┃    \x1b[1;93m%s\x1b[0m    \x1b[1;32m#####\x1b[0m ┃\n", date)
	fmt.Println("         ┗━━━━━━━━━━━━━━━━━\x1b[1;31m#\x1b[0m━━━┛")
	fmt.Printf("%s\n", url)
	fmt.Println()
}

func timeAndPrint(fn func(in string) string, input string) {
	t := time.Now()
	s := fn(input)
	d := time.Since(t)
	// default red
	c := 31
	if d < time.Second {
		// less than a second, green
		c = 32
	} else if d < 15*time.Second {
		// less than 15 seconds, ok... yellow
		c = 93
	}
	fmt.Printf("\x1b[1;37m%s\x1b[0m \x1b[%dm%v\x1b[0m\n", s, c, d)
}

func Run(YEAR, DAY int, solve1, solve2 func(string) string) {
	var testsOnly = flag.Bool("test-only", false, "Only run the tests")
	flag.Parse()
	PrintHeader(YEAR, DAY)
	runTest(1)
	var input string
	if !*testsOnly {
		b, err := ioutil.ReadFile("input.txt")
		if err != nil {
			if os.IsNotExist(err) {
				log.Fatalln("Please create 'input.txt' with your problem input")
			}
			log.Fatalln("Error trying to read input file ('input.txt'):", err)
		}
		if len(b) == 0 {
			log.Fatalln("Please add your problem input to 'input.txt'")
		}
		input = string(b)
		fmt.Print("Solving problem 1: ")
		timeAndPrint(solve1, input)

	}
	runTest(2)
	if !*testsOnly {
		fmt.Print("Solving problem 2: ")
		timeAndPrint(solve2, input)
	}
}

func runTest(n int) {
	fmt.Printf("Testing problem %d: ", n)
	output, err := exec.Command("go", "test", "-run", fmt.Sprintf("^TestProblem%d$", n)).CombinedOutput()
	if err != nil {
		fmt.Println("\x1b[1;31mFAIL\x1b[0m")
		os.Stdout.Write(output)
		os.Exit(1)
	}
	fmt.Println("\x1b[1;32mPASS\x1b[0m")
}
