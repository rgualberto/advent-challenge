package puzzle1

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

	fmt.Printf("\n****** Calculating fuel ********\n")
	for _, partMass := range partsMass {
		f, err := strconv.ParseFloat(partMass, 64)
		if err != nil {
			log.Fatal(err)
		}

		fuel := calculateFuel(f)
		finalSum += fuel

		fmt.Printf("\"%s\" -> %v\n", partMass, fuel)
	}

	fmt.Printf("\n\nfinal sum: %v\n\n", finalSum)
}

func calculateFuel(mass float64) int {
	floatFuel := math.Floor(mass/3) - 2
	return int(floatFuel)
}
