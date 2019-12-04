package puzzle2

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/rgualberto/advent-stuff/common"
)

// Run starts the program
func Run() {
	partsMass, err := common.Scan("./inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}

	finalSum := 0

	fmt.Printf("\n****** Calculating fuel for compounded fuel mass ********\n")
	for _, partMass := range partsMass {
		f, err := strconv.ParseFloat(partMass, 64)
		if err != nil {
			log.Fatal(err)
		}

		fuel := calculateFuelRequirements(f)
		finalSum += int(fuel)

		fmt.Printf("\"%s\" -> %v\n", partMass, fuel)
	}

	fmt.Printf("\n\nfinal sum: %v\n\n", finalSum)
}

func calculateFuelRequirements(mass float64) float64 {
	floatFuel := math.Floor(mass/3) - 2

	if floatFuel <= 0 {
		return 0
	}

	return floatFuel + calculateFuelRequirements(floatFuel)
}
