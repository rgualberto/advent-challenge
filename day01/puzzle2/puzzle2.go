package puzzle2

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rgualberto/advent-stuff/common"
	"github.com/rgualberto/advent-stuff/day01"
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

		fuel := day01.CalculateFuelRequirements(f)
		finalSum += int(fuel)

		fmt.Printf("\"%s\" -> %v\n", partMass, fuel)
	}

	fmt.Printf("\n\nfinal sum: %v\n\n", finalSum)
}
