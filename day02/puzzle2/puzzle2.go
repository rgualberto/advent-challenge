package puzzle2

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/rgualberto/advent-stuff/common"
	d2puz1 "github.com/rgualberto/advent-stuff/day02/puzzle1"
)

// Run starts the program
func Run() {
	fmt.Printf("\n****** Running Intcode program - finding noun & verb to output 19690720 ********\n")

	intcode, err := common.Scan("./inputs/day02.txt")
	if err != nil {
		log.Fatal(err)
	}

	// only care about first line
	splitVals := strings.Split(intcode[0], ",")

	var program []int
	for _, v := range splitVals {
		posVal, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		program = append(program, posVal)
	}

	fmt.Printf("program input: %v\n", program)
	fmt.Printf("=== running program ===\n")

	expectedOutput := 19690720
	noun, verb, err := findInput(program, expectedOutput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Input Found! noun: %v, verb: %v\n", noun, verb)
	outputCode := 100*noun + verb
	fmt.Printf("Output Code: %v\n", outputCode)
}

func findInput(p []int, expectedOutput int) (int, int, error) {
	maxMemAddresses := len(p) - 1

	var noun int
	var verb int

	for noun = 0; noun <= maxMemAddresses; noun++ {
		for verb = 0; verb <= maxMemAddresses; verb++ {
			// clone slice into new mem addr
			program := append(p[:0:0], p...)
			program[1] = noun
			program[2] = verb
			programResult, err := d2puz1.CalculateIntcodeOutput(program)
			if err != nil {
				// invalid input -- terminate
				fmt.Println(err)
				continue
			}

			if programResult[0] == expectedOutput {
				fmt.Println(programResult[0])
				return noun, verb, nil
			}
		}
	}

	return 0, 0, errors.New("input not found -- output not possible")
}
