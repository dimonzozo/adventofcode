package main

import (
	"github.com/dimonzozo/adventofcode/common"
	"github.com/sirupsen/logrus"
	"math"
	"strconv"
)

func main() {
	sumFuel := 0
	totalFuel := 0

	for _, mass := range common.Input() {
		if mass == "" {
			continue
		}

		massInt, err := strconv.Atoi(mass)
		if err != nil {
			logrus.Fatalf("Mass should be integer")
		}

		sumFuel += CalculateFuel(massInt)
		totalFuel += CalculateFuelTotal(massInt)
	}

	logrus.Printf("Sum of fuel requirements: %d", sumFuel)
	logrus.Printf("Sum of total fuel requirements: %d", totalFuel)
}

func CalculateFuel(mass int) int {
	return int(math.Floor(float64(mass)/3.0)) - 2
}

func CalculateFuelTotal(mass int) int {
	sum := 0

	for {
		mass = CalculateFuel(mass)
		if mass > 0 {
			sum += mass
		} else {
			break
		}
	}

	return sum
}
