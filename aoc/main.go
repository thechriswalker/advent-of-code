package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func Run(YEAR, DAY int, solve1, solve2 func(string) string) {
	fmt.Println("##############################")
	fmt.Println("#                            #")
	//          "#  AdventOfCode YYYY Day DD  #
	fmt.Printf("#  AdventOfCode %d Day %02d  #\n", YEAR, DAY)
	fmt.Println("#                            #")
	fmt.Println("##############################")
	fmt.Println()
	runTest(1)
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("Please create 'input.txt' with your problem input")
		}
		log.Fatalln("Error trying to read input file ('input.txt'):", err)
	}
	input := string(b)

	fmt.Print("Solving problem 1: ")
	fmt.Println(solve1(input))
	runTest(2)
	fmt.Print("Solving problem 2: ")
	fmt.Println(solve2(input))
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
