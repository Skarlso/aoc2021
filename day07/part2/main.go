package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	input := strings.Split(string(content), ",")

	// find the maximum and the minimum
	crabs := make([]int, 0)
	for _, in := range input {
		n, _ := strconv.Atoi(in)
		crabs = append(crabs, n)
	}

	max, min := 0, math.MaxInt

	for _, crab := range crabs {
		if crab < min {
			min = crab
		}
		if crab > max {
			max = crab
		}
	}

	sum := func(n int) int {
		return (n * (1 + n)) / 2
	}
	minFuel := math.MaxInt
	for i := min; i <= max; i++ {
		currFuel := 0
		for _, crab := range crabs {
			v := sum(abs(i - crab))
			currFuel += v
		}
		if currFuel < minFuel {
			minFuel = currFuel
		}
	}

	fmt.Println("min fuel: ", minFuel)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
