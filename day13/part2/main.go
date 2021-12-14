package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	input := strings.Split(string(content), "\n")
	numbers := make([]string, 0)
	folders := make([]string, 0)

	readingFolds := false
	for _, line := range input {
		// gather numbers
		if line == "" {
			readingFolds = true
			continue
		}

		if !readingFolds {
			numbers = append(numbers, line)
		} else {
			folders = append(folders, line)
		}
	}

	points := make([]point, 0)
	maxx, maxy := 0, 0
	for _, v := range numbers {
		var (
			x, y int
		)
		fmt.Sscanf(v, "%d,%d", &x, &y)
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
		points = append(points, point{x: x, y: y})
	}

	grid := make([][]string, maxy+1)
	for i := 0; i < maxy+1; i++ {
		grid[i] = make([]string, maxx+1)
	}

	for _, p := range points {
		grid[p.y][p.x] = "#"
	}

	// display(grid) // parsed
	// do one fold

	for _, f := range folders {
		split := strings.Split(f, " ")
		fold := split[len(split)-1]
		if fold[0] == 'x' {
			// fmt.Println("folding left")
			var x int
			fmt.Sscanf(fold, "x=%d", &x)
			for y := 0; y < len(grid); y++ {
				for j := x; j < len(grid[y]); j++ {
					if grid[y][j] == "#" {
						grid[y][x-(j-x)] = "#"
					}
				}
			}
			for i, v := range grid {
				v = v[:x]
				grid[i] = v
			}
		} else if fold[0] == 'y' {
			// fmt.Println("folding up")
			var y int
			fmt.Sscanf(fold, "y=%d", &y)
			// scan downwards
			// offset is i - y 7-3 -> 4 and we subtract that from level (7)
			for i := y; i < len(grid); i++ {
				for x := 0; x < len(grid[i]); x++ {
					if grid[i][x] == "#" {
						grid[y-(i-y)][x] = "#"
					}
				}
			}
			grid = grid[:y]
		}
	}
	display(grid) // parsed
}

func display(grid [][]string) {
	for _, v := range grid {
		for _, s := range v {
			if s != "#" {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
