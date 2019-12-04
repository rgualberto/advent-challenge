package day03

import (
	"fmt"
	"log"

	"github.com/rgualberto/advent-stuff/common"
)

// RunPuzzle1 starts the program
func RunPuzzle1() {
	fmt.Printf("\n****** Calculating closest distance (manhattan) to central port ********\n")

	wirepaths, err := common.Scan("./inputs/day03.txt")
	if err != nil {
		log.Fatal(err)
	}

	ExecutePuzzle1(wirepaths)
}

// ExecutePuzzle1 executes duh
func ExecutePuzzle1(wirepaths []string) int {
	var wireCoordinates [][]Coordinates
	for _, wp := range wirepaths {
		wc, err := TraceWirePath(wp)
		if err != nil {
			log.Fatal(err)
		}

		wireCoordinates = append(wireCoordinates, wc)
	}

	intersectionPoints, _, err := FindIntersectionPoints(wireCoordinates[0], wireCoordinates[1])
	if err != nil {
		log.Fatal(err)
	}

	closestPoint, err := CalculateClosestManhattanDistance(intersectionPoints)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\nClosest Manhatten Distance to central port: %v\n", closestPoint)

	return closestPoint
}
