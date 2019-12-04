package day01

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDay01(t *testing.T) {
	Convey("CalculateFuel", t, func() {
		Convey("mass of 12 should return 2", func() {
			fuel := CalculateFuel(12)
			expectedFuel := 2

			So(fuel, ShouldEqual, expectedFuel)
		})

		Convey("mass of 14 should return 2", func() {
			fuel := CalculateFuel(14)
			expectedFuel := 2

			So(fuel, ShouldEqual, expectedFuel)
		})

		Convey("mass of 1969 should return 654", func() {
			fuel := CalculateFuel(1969)
			expectedFuel := 654

			So(fuel, ShouldEqual, expectedFuel)
		})

		Convey("mass of 100756 should return 33583", func() {
			fuel := CalculateFuel(100756)
			expectedFuel := 33583

			So(fuel, ShouldEqual, expectedFuel)
		})
	})

	Convey("CalculateFuelRequirements", t, func() {
		Convey("mass of 12 should return 2", func() {
			fuel := CalculateFuelRequirements(12)
			expectedFuel := 2

			So(fuel, ShouldEqual, expectedFuel)
		})

		Convey("mass of 1969 should return 966", func() {
			fuel := CalculateFuelRequirements(1969)
			expectedFuel := 966

			So(fuel, ShouldEqual, expectedFuel)
		})

		Convey("mass of 100756 should return 50346", func() {
			fuel := CalculateFuelRequirements(100756)
			expectedFuel := 50346

			So(fuel, ShouldEqual, expectedFuel)
		})
	})
}
