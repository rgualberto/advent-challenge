package day01

import "math"

// CalculateFuel calculates fuel requirement for mass input
func CalculateFuel(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

// CalculateFuelRequirements calculates fuel requirement for mass input and mass of fuel added
func CalculateFuelRequirements(mass float64) float64 {
	floatFuel := CalculateFuel(mass)

	if floatFuel <= 0 {
		return 0
	}

	return floatFuel + CalculateFuelRequirements(floatFuel)
}
