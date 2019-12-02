package puzzle2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// Run starts the program
func Run() {
	file, err := os.Open("./day01/puzzle2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	finalSum := 0

	fmt.Printf("\n****** Scanning and calculating fuel for compounded fuel mass ********\n")
	for scanner.Scan() {
		in := scanner.Text()
		f, err := strconv.ParseFloat(in, 64)
		if err != nil {
			log.Fatal(err)
		}

		fuel := calculateFuelRequirements(f)
		finalSum += int(fuel)

		fmt.Printf("\"%s\" -> %v\n", in, fuel)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
