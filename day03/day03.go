package day03

import (
	"strconv"
	"strings"
)

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

	for _, wc1 := range wireCoords1 {
		for _, wc2 := range wireCoords2 {
			if wc1 == wc2 {
				intersectionPoints = append(intersectionPoints, wc2)
			}
		}
	}

	return intersectionPoints, nil
}

func CalculateManhattanDistance(point Coordinates) int {
	return 0
}

func CalculateClosestManhattanDistance(coordinates []Coordinates) int {
	return 0
}
