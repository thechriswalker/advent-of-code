package aoc

import (
	"flag"
	"fmt"
	"io"
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

// Remember AoC says <15 seconds on 10 year old hardware.
// So we are probably looking at <2 seconds on current hardware.
// And a good time will be much quicker, we will give ourselves 10ms?
const (
	goodTiming = 10 * time.Millisecond
	badTiming  = time.Second
)

func timeAndPrint(fn func(in string) string, input string, answers io.Writer, timingOnly bool) {
	t := time.Now()
	s := fn(input)
	fmt.Fprintln(answers, s)
	d := time.Since(t)
	// default red
	c := 31
	if d < goodTiming {
		// less than a second, green
		c = 32
	} else if d < badTiming {
		// less than 15 seconds, ok... yellow
		c = 93
	}
	if !timingOnly {
		fmt.Printf("\x1b[1;37m%s\x1b[0m ", s)
	}

	fmt.Printf("\x1b[%dm%v\x1b[0m\n", c, d)

}

func SupressOutput() func() {
	stdout := os.Stdout
	stderr := os.Stdin
	devnull, _ := os.Open(os.DevNull)
	os.Stdout = devnull
	os.Stderr = devnull
	return func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}
}

func wrapSuppressOutput(f func(string) string) func(string) string {
	return func(s string) string {
		defer SupressOutput()()
		return f(s)
	}
}

// if we have the environment variable for check answers
// --check-answers=1:abc
func Run(YEAR, DAY int, solve1, solve2 func(string) string) {
	testsOnly := flag.Bool("test-only", false, "Only run the tests")
	timingOnly := flag.Bool("timing-only", false, "Show only timings")
	answersCheck := flag.Bool("check-answers", false, "Check Solutions against know answers")
	recordAnswers := flag.Bool("record-answers", false, "Save current answers as correct")
	flag.Parse()

	suppresStdoutDuringExecution := *answersCheck || *timingOnly || *recordAnswers || *testsOnly

	if suppresStdoutDuringExecution {
		solve1 = wrapSuppressOutput(solve1)
		solve2 = wrapSuppressOutput(solve2)
	}

	// from here, i.e. a relative path
	privateDataPath := fmt.Sprintf("../../not_public/%d/%02d/", YEAR, DAY)

	var input string
	getInput := func() string {
		if input == "" {
			b, err := os.ReadFile(privateDataPath + "input.txt")
			if err != nil {
				if os.IsNotExist(err) {
					log.Fatalln("Please create 'input.txt' in private path with your problem input")
				}
				log.Fatalln("Error trying to read input file ('input.txt') from private path:", err)
			}
			if len(b) == 0 {
				log.Fatalln("Please add your problem input to 'input.txt'")
			}
			input = string(b)
		}
		return input
	}

	if *answersCheck {
		fails := 100 //NB this is supposed to start at `100`
		in := getInput()
		answer1, answer2 := readAnswers(YEAR, DAY)
		result1 := solve1(in)
		result2 := solve2(in)
		fmt.Printf("%d-%02d Part 1: ", YEAR, DAY)
		if result1 == answer1 {
			fmt.Println("\x1b[1;32mPASS\x1b[0m")
		} else {
			fails++
			fmt.Printf("\x1b[1;31mFAIL\x1b[0m (expected %q, got %q)\n", answer1, result1)
		}
		fmt.Printf("%d-%02d Part 2: ", YEAR, DAY)
		if result2 == answer2 {
			fmt.Println("\x1b[1;32mPASS\x1b[0m")
		} else {
			fails++
			fmt.Printf("\x1b[1;31mFAIL\x1b[0m (expected %q, got %q)\n", answer2, result2)
		}
		os.Exit(fails)
		return
	}
	answerRecorder, _ := os.Open(os.DevNull)
	if *timingOnly {
		// noheader no tests
		fmt.Printf("%d-%02d Part 1: ", YEAR, DAY)
		timeAndPrint(solve1, getInput(), answerRecorder, true)
		fmt.Printf("%d-%02d Part 2: ", YEAR, DAY)
		timeAndPrint(solve2, getInput(), answerRecorder, true)
		return
	}

	PrintHeader(YEAR, DAY)
	runTest(0)
	runTest(1)
	if *recordAnswers && !*testsOnly {
		var err error
		answerRecorder, err = os.Create(privateDataPath + "/answers.txt")
		if err != nil {
			log.Fatalln("could not open answers file for writing")
		}
		defer answerRecorder.Close()
	}
	if !*testsOnly {
		fmt.Print("Solving problem 1: ")
		timeAndPrint(solve1, getInput(), answerRecorder, false)
	}
	runTest(2)
	if !*testsOnly {
		fmt.Print("Solving problem 2: ")
		timeAndPrint(solve2, getInput(), answerRecorder, false)
	}
}

func runTest(n int) {
	args := []string{"test"}
	if n == 0 {
		fmt.Printf("Any Extra Testing: ")
		args = append(args, "-skip", "^TestProblem[12]$")
	} else {
		fmt.Printf("Testing problem %d: ", n)
		args = append(args, "-run", fmt.Sprintf("^TestProblem%d$", n))
	}
	output, err := exec.Command("go", args...).CombinedOutput()
	if err != nil {
		fmt.Println("\x1b[1;31mFAIL\x1b[0m")
		os.Stdout.Write(output)
		os.Exit(1)
	}
	fmt.Println("\x1b[1;32mPASS\x1b[0m")
}

func readAnswers(year, day int) (a1, a2 string) {
	// assume they are in a file answers.txt
	privateDataPath := fmt.Sprintf("../../not_public/%d/%02d/", year, day)

	b, err := os.ReadFile(privateDataPath + "/answers.txt")
	if err != nil {
		log.Fatalf("Could not open answers file at %q\n", privateDataPath+"/answers.txt")
	}
	i := 0
	MapLines(string(b), func(line string) error {
		switch i {
		case 0:
			a1 = line
		case 1:
			a2 = line
		default:
			log.Fatalln("More than 2 lines in answers.txt")
		}
		i++
		return nil
	})
	return
}
