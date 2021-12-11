package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type octopus struct {
	energy      int
	currentStep int
	flashed     bool // how do I reset this?
}

type point struct {
	x, y int
}

var (
	directions = []point{
		{x: 0, y: -1},
		{x: -1, y: -1},
		{x: -1, y: 0},
		{x: 1, y: 1},
		{x: 0, y: 1},
		{x: -1, y: 1},
		{x: 1, y: 0},
		{x: 1, y: -1},
	}
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	input := strings.Split(string(content), "\n")
	grid := [10][10]*octopus{}
	for i, line := range input {
		for j, c := range line {
			n := c - '0'
			grid[i][j] = &octopus{
				currentStep: 0,
				energy:      int(n),
				flashed:     false,
			}
		}
	}

	doTheThing(grid)
}

var flashCount int

func doTheThing(grid [10][10]*octopus) {
	steps := 100

	for i := 0; i < steps; i++ {
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				// we only ++ if the octopus is on the same step.
				// if it reset, it will be on the next step anyways.
				// if it's 0 and not on the same step it's resting.
				if grid[y][x].currentStep == i {
					grid[y][x].energy++
					if grid[y][x].energy == 10 {
						flash(point{x: x, y: y}, i, grid)
					}
				}
			}
		}
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x].currentStep == i {
					grid[y][x].currentStep++
				}
			}
		}
		// fmt.Println("==========")
		// display(grid)
	}
	fmt.Println("total flash counts: ", flashCount)
}

func display(grid [10][10]*octopus) {
	for _, v := range grid {
		for _, c := range v {
			fmt.Print(c.energy)
		}
		fmt.Println()
	}
}

func flash(currentPoint point, currentStep int, grid [10][10]*octopus) {
	// flash the current octopus
	flashCount++
	grid[currentPoint.y][currentPoint.x].energy = 0
	grid[currentPoint.y][currentPoint.x].currentStep++
	// select neighbors and increase their energy
	for _, d := range directions {
		np := point{x: currentPoint.x + d.x, y: currentPoint.y + d.y}
		if np.x >= 0 && np.x < len(grid[currentPoint.y]) && np.y >= 0 && np.y < len(grid) {
			if grid[np.y][np.x].currentStep == currentStep {
				grid[np.y][np.x].energy++
				if grid[np.y][np.x].energy == 10 {
					flash(np, currentStep, grid)
				}
			}
		}
	}
}
