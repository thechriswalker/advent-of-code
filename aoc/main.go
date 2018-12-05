package aoc

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var testsOnly = flag.Bool("test-only", false, "Only run the tests")

func Run(YEAR, DAY int, solve1, solve2 func(string) string) {
	flag.Parse()
	fmt.Println("##############################")
	fmt.Println("#                            #")
	//          "#  AdventOfCode YYYY Day DD  #
	fmt.Printf("#  AdventOfCode %d Day %02d  #\n", YEAR, DAY)
	fmt.Println("#                            #")
	fmt.Println("##############################")
	fmt.Println()
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
		fmt.Println(solve1(input))
	}
	runTest(2)
	if !*testsOnly {
		fmt.Print("Solving problem 2: ")
		fmt.Println(solve2(input))
	}
}

func runTest(n int) {
	fmt.Printf("Testing problem %d: ", n)
	output, err := exec.Command("go", "test", "-run", fmt.Sprintf("^TestProblem%d$", n)).CombinedOutput()
	if err != nil {
		fmt.Println("FAIL")
		os.Stdout.Write(output)
		os.Exit(1)
	}
	fmt.Println("PASS")
}
