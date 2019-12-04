package day02

import (
	"errors"
	"fmt"
)

// CalculateIntcodeOutput runs the intcode program
func CalculateIntcodeOutput(program []int) ([]int, error) {
	index := 0
	var terminateProgram bool
	var err error
	for {
		program, terminateProgram, err = ExecuteInstruction(program, index)
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

// ExecuteInstruction executes an intcode program instruction
func ExecuteInstruction(ins []int, index int) ([]int, bool, error) {
	opCode := ins[index]

	if opCode == 99 {
		return ins, true, nil
	}

	firstIndex := ins[index+1]
	secondIndex := ins[index+2]
	destination := ins[index+3]

	switch opCode {
	case 1:
		v := ins[firstIndex] + ins[secondIndex]
		ins[destination] = v
		break
	case 2:
		v := ins[firstIndex] * ins[secondIndex]
		ins[destination] = v
		break
	default:
		return ins, true, fmt.Errorf("invalid opcode provided: %v at index: %v", opCode, index)
	}

	return ins, false, nil
}

// FindIntcodeNounAndVerb calculates possible outputs from a given intcode program to find the noun/verb that calculates it
func FindIntcodeNounAndVerb(p []int, expectedOutput int) (int, int, error) {
	maxMemAddresses := len(p) - 1

	var noun int
	var verb int

	for noun = 0; noun <= maxMemAddresses; noun++ {
		for verb = 0; verb <= maxMemAddresses; verb++ {
			// clone slice into new mem addr
			program := append(p[:0:0], p...)
			program[1] = noun
			program[2] = verb
			programResult, err := CalculateIntcodeOutput(program)
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
