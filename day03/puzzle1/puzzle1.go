package puzzle1

import (
	"fmt"
	"log"

	"github.com/rgualberto/advent-stuff/common"
	"github.com/rgualberto/advent-stuff/day03"
)

// Run starts the program
func Run() {
	fmt.Printf("\n****** Calculating closest distance (manhattan) to central port ********\n")

	wirepaths, err := common.Scan("./inputs/day03.txt")
	if err != nil {
		log.Fatal(err)
	}

	var wireCoordinates [][][]int
	for _, wp := range wirepaths {
		wc, err := day03.TraceWirePath(wp)
		if err != nil {
			log.Fatal(err)
		}

		wireCoordinates = append(wireCoordinates, wc)
	}

	intersectionPoints, err := day03.FindIntersectionPoints(wireCoordinates[0], wireCoordinates[1])
	if err != nil {
		log.Fatal(err)
	}

	closestPoint := day03.CalculateClosestManhattanDistance(intersectionPoints)

	fmt.Printf("Closest Manhatten Distance to central port: %v\n", closestPoint)
}
