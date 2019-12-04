package day02

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDay02(t *testing.T) {
	Convey("CalculateIntcodeOutput", t, func() {
		Convey("inputs should have expected output", func() {
			intcode := []int{1, 0, 0, 0, 99}
			expectedResult := []int{2, 0, 0, 0, 99}
			resultIntcode, _ := CalculateIntcodeOutput(intcode)

			So(resultIntcode, ShouldResemble, expectedResult)

			intcode = []int{2, 3, 0, 3, 99}
			expectedResult = []int{2, 3, 0, 6, 99}
			resultIntcode, _ = CalculateIntcodeOutput(intcode)

			So(resultIntcode, ShouldResemble, expectedResult)

			intcode = []int{2, 4, 4, 5, 99, 0}
			expectedResult = []int{2, 4, 4, 5, 99, 9801}
			resultIntcode, _ = CalculateIntcodeOutput(intcode)

			So(resultIntcode, ShouldResemble, expectedResult)

			intcode = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
			expectedResult = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
			resultIntcode, _ = CalculateIntcodeOutput(intcode)

			So(resultIntcode, ShouldResemble, expectedResult)
		})
	})

	Convey("FindIntcodeNounAndVerb", t, func() {
		Convey("program input should find expected noun & verb", func() {
			intcode := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
			expectedOutput := 30

			noun, verb, _ := FindIntcodeNounAndVerb(intcode, expectedOutput)

			So(noun, ShouldEqual, 0)
			So(verb, ShouldEqual, 0)
		})
	})
}
