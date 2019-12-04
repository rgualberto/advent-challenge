package day03

import (
	"strconv"
	"strings"
)

// TraceWirePath steps through a coord instruction set and returns a slice of x/y coord pairings
func TraceWirePath(wirepath string) ([][]int, error) {
	directions := strings.Split(wirepath, ",")
	x, y := 0, 0

	var wirePathCoordinates [][]int
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

			wirePathCoordinates = append(wirePathCoordinates, []int{x, y})
		}
	}

	return wirePathCoordinates, nil
}

func FindIntersectionPoints(wireCoords1 [][]int, wireCoords2 [][]int) ([][]int, error) {
	return [][]int{}, nil
}

func CalculateManhattanDistance(point []int) int {
	return 0
}

func CalculateClosestManhattanDistance(coordinates [][]int) int {
	return 0
}
