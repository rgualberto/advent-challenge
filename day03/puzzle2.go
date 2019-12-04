package day03

import (
	"fmt"
	"log"

	"github.com/rgualberto/advent-stuff/common"
)

// RunPuzzle2 starts the program
func RunPuzzle2() {
	fmt.Printf("\n****** Calculating intersection point with smallest steps to reach from central port ********\n")

	wirepaths, err := common.Scan("./inputs/day03.txt")
	if err != nil {
		log.Fatal(err)
	}

	ExecutePuzzle2(wirepaths)
}

// ExecutePuzzle2 executes duh
func ExecutePuzzle2(wirepaths []string) int {
	var wireCoordinates [][]Coordinates
	for _, wp := range wirepaths {
		wc, err := TraceWirePath(wp)
		if err != nil {
			log.Fatal(err)
		}

		wireCoordinates = append(wireCoordinates, wc)
	}

	_, intersectionPoints, err := FindIntersectionPoints(wireCoordinates[0], wireCoordinates[1])
	if err != nil {
		log.Fatal(err)
	}

	smallestStepCount, err := CalculateSmallestStepCount(intersectionPoints)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\nSmallest step count to central port: %v\n", smallestStepCount)

	return smallestStepCount
}
