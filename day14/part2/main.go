package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	input := strings.Split(string(content), "\n")
	template := input[0]
	input = input[2:]

	rules := make(map[string]string)
	for _, line := range input {
		split := strings.Split(line, " -> ")
		pair := split[0]
		character := split[1]
		rules[pair] = character
	}

	steps := 40
	trackPairs := make(map[string]int)

	for i := 0; i < len(template)-1; i++ {
		trackPairs[string(template[i])+string(template[i+1])]++
	}

	for i := 0; i < steps; i++ {
		update := make(map[string]int)
		for k, v := range trackPairs {
			update[string(k[0])+rules[k]] += v
			update[rules[k]+string(k[1])] += v
		}
		trackPairs = update
	}

	fmt.Println(trackPairs)

	counts := map[string]int{}
	for k, v := range trackPairs {
		counts[string(k[0])] += v
	}
	fmt.Println(counts)

	max, min := 0, math.MaxInt

	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Println(max - min + 1)
}
