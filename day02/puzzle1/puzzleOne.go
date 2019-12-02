package puzzle1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Run starts the program
func Run() {
	file, err := os.Open("./day02/puzzle1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("\n****** Running Intcode program - adjusting positions 1 and 2 with 12 and 2 ********\n")

	scanner := bufio.NewScanner(file)

	// only care about first line
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	splitVals := strings.Split(scanner.Text(), ",")

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

	program, err = CalculateIntcodeOutput(program)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("final program: %v\n", program)
	fmt.Printf("value at position 0: %v\n", program[0])
}

// CalculateIntcodeOutput runs the intcode program
func CalculateIntcodeOutput(program []int) ([]int, error) {
	index := 0
	var terminateProgram bool
	var err error
	for {
		program, terminateProgram, err = runProgramBlock(program, index)
		if err != nil {
			return []int{}, err
		}

		if terminateProgram {
			break
		}

		index += 4
	}

	return program, nil
}

func runProgramBlock(p []int, index int) ([]int, bool, error) {
	opCode := p[index]

	if opCode == 99 {
		return p, true, nil
	}

	firstIndex := p[index+1]
	secondIndex := p[index+2]
	destination := p[index+3]

	switch opCode {
	case 1:
		v := p[firstIndex] + p[secondIndex]
		p[destination] = v
		break
	case 2:
		v := p[firstIndex] * p[secondIndex]
		p[destination] = v
		break
	default:
		return p, true, fmt.Errorf("invalid opcode provided: %v at index: %v", opCode, index)
	}

	return p, false, nil
}