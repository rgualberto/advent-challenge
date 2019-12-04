package day03

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Coordinates is a slice of 2 integers (x, y)
type Coordinates [2]int

// TraceWirePath steps through a coord instruction set and returns a slice of x/y coord pairings
func TraceWirePath(wirepath string) ([]Coordinates, error) {
	directions := strings.Split(wirepath, ",")
	x, y := 0, 0

	var wirePathCoordinates []Coordinates
	for _, d := range directions {
		dir := d[0:1]
		steps, err := strconv.Atoi(d[1:])
		if err != nil {
			return wirePathCoordinates, err
		}

		for i := 0; i < steps; i++ {
			switch dir {
			case "U":
				y++
				break
			case "D":
				y--
				break
			case "R":
				x++
				break
			case "L":
				x--
				break
			}

			wirePathCoordinates = append(wirePathCoordinates, Coordinates{x, y})
		}
	}

	return wirePathCoordinates, nil
}

// FindIntersectionPoints returns all intersection coordinates between two provided coordinates
func FindIntersectionPoints(wireCoords1 []Coordinates, wireCoords2 []Coordinates) ([]Coordinates, error) {
	var intersectionPoints []Coordinates

	// fmt.Printf("\nwc1 len: %v, wc2 len: %v\n", len(wireCoords1), len(wireCoords2))
	fmt.Println("finding intersection points")
	for index, wc1 := range wireCoords1 {
		if index%1000 == 0 {
			fmt.Printf(".")
		}

		for _, wc2 := range wireCoords2 {
			if wc1 == wc2 {
				intersectionPoints = append(intersectionPoints, wc2)
			}
		}
	}

	fmt.Println("intersection points: ", intersectionPoints)

	return intersectionPoints, nil
}

// CalculateManhattanDistance returns the distance from the central point (0, 0)
func CalculateManhattanDistance(point Coordinates) int {
	// distance := |a - c| + |b - d|
	// since we are assuming the central point to be 0, we can simplify
	// distance := |x| + |y|

	distance := math.Abs(float64(point[0])) + math.Abs(float64(point[1]))

	return int(distance)
}

// CalculateClosestManhattanDistance returns the closest distance to the central point (0, 0)
func CalculateClosestManhattanDistance(coordinates []Coordinates) (int, error) {
	var distance int // 0 is an invalid distance so lets start there

	fmt.Printf("\n distances: ")
	for _, point := range coordinates {
		dist := CalculateManhattanDistance(point)

		fmt.Printf("%v,", dist)

		if distance == 0 || dist < distance {
			distance = dist
		}
	}

	if distance == 0 {
		return distance, errors.New("Distance cannot be 0 (implies intersection at central port)")
	}

	return distance, nil
}
