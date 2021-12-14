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

	pairs := make(map[string]string)
	for _, line := range input {
		split := strings.Split(line, " -> ")
		pair := split[0]
		character := split[1]
		pairs[pair] = character
	}

	steps := 10
	// insert then skip because pairs overlap.
	for step := 0; step < steps; step++ {
		newTemplate := " " // cheeky
		for i := 0; i < len(template)-1; i++ {
			pair := string(template[i]) + string(template[i+1])
			pair = string(pair[0]) + pairs[pair] + string(pair[1])
			newTemplate = newTemplate[:len(newTemplate)-1]
			newTemplate += pair
		}
		template = newTemplate
	}

	min := math.MaxInt
	max := 0

	for _, c := range template {
		count := strings.Count(template, string(c))
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	fmt.Println(max - min)
}
