package puzzle2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/rgualberto/advent-stuff/common"
	"github.com/rgualberto/advent-stuff/day02"
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
	noun, verb, err := day02.FindIntcodeNounAndVerb(program, expectedOutput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Input Found! noun: %v, verb: %v\n", noun, verb)
	outputCode := 100*noun + verb
	fmt.Printf("Output Code: %v\n", outputCode)
}
