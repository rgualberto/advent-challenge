package puzzle1

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
	file, err := os.Open("./day01/puzzle1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	finalSum := 0

	fmt.Printf("\n****** Scanning and calculating fuel ********\n")
	for scanner.Scan() {
		in := scanner.Text()
		f, err := strconv.ParseFloat(in, 64)
		if err != nil {
			log.Fatal(err)
		}

		fuel := calculateFuel(f)
		finalSum += fuel

		fmt.Printf("\"%s\" -> %v\n", in, fuel)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\nfinal sum: %v\n\n", finalSum)
}

func calculateFuel(mass float64) int {
	floatFuel := math.Floor(mass/3) - 2
	return int(floatFuel)
}
