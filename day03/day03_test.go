package day03

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDay03(t *testing.T) {
	Convey("CalculateClosestManhattanDistance", t, func() {
		Convey("inputs should have expected output", func() {
			wirepaths := []string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			}
			expectedResult := 159
			closestPoint := ExecutePuzzle1(wirepaths)

			So(closestPoint, ShouldEqual, expectedResult)

			wirepaths = []string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			}
			expectedResult = 135
			closestPoint = ExecutePuzzle1(wirepaths)

			So(closestPoint, ShouldEqual, expectedResult)
		})
	})

	Convey("CalculateSmallestStepCount", t, func() {
		Convey("inputs should have expected output", func() {
			wirepaths := []string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			}
			expectedResult := 610
			steps := ExecutePuzzle2(wirepaths)

			So(steps, ShouldEqual, expectedResult)

			wirepaths = []string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			}
			expectedResult = 410
			steps = ExecutePuzzle2(wirepaths)

			So(steps, ShouldEqual, expectedResult)
		})
	})
}
