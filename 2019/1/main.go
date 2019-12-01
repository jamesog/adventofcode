package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// One option would be to use float64 as input and use math.Floor(), then
// convert back to int, however Go will do the right thing with int.
func fuel(mass int) int {
	return mass/3 - 2
}

func fuelForFuel(mass int) int {
	mf := fuel(mass)
	if mf < 0 {
		return 0
	}
	sum := mf
	for mf > 0 {
		mf = fuel(mf)
		if mf > 0 {
			sum += mf
		}
	}
	return sum
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		modFuel := fuel(mass)
		fuelFuel := fuelForFuel(modFuel)
		sum += modFuel + fuelFuel
	}
	fmt.Printf("Fuel needed: %d\n", sum)
}
