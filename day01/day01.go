package day01

import "math"

// CalculateFuel calculates fuel requirement for mass input
func CalculateFuel(mass float64) int {
	floatFuel := math.Floor(mass/3) - 2
	return int(floatFuel)
}

// CalculateFuelRequirements calculates fuel requirement for mass input and mass of fuel added
func CalculateFuelRequirements(mass float64) float64 {
	floatFuel := math.Floor(mass/3) - 2

	if floatFuel <= 0 {
		return 0
	}

	return floatFuel + CalculateFuelRequirements(floatFuel)
}
