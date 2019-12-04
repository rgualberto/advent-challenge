package puzzle1

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
	fmt.Printf("\n****** Running Intcode program - adjusting positions 1 and 2 with 12 and 2 ********\n")

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

	program[1] = 12
	program[2] = 2

	program, err = day02.CalculateIntcodeOutput(program)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("final program: %v\n", program)
	fmt.Printf("value at position 0: %v\n", program[0])
}
