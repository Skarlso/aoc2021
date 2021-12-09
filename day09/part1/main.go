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

var (
	directions = []point{
		{x: -1, y: 0},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
	}
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	input := strings.Split(string(content), "\n")

	grid := make([][]int, 0)
	for _, i := range input {
		row := make([]int, 0)
		for _, v := range i {
			c := v - '0'
			row = append(row, int(c))
		}
		grid = append(grid, row)
	}
	sum := 0
	// display(grid)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			lowest := true
			for _, d := range directions {
				np := point{x: x + d.x, y: y + d.y}
				if np.x >= 0 && np.x < len(grid[y]) && np.y >= 0 && np.y < len(grid) {
					if grid[np.y][np.x] <= grid[y][x] {
						lowest = false
						break
					}
				}
			}
			if lowest {
				fmt.Println("lowest: ", grid[y][x])
				sum += grid[y][x] + 1
			}
		}
	}
	fmt.Println("sum: ", sum)
}

func display(grid [][]int) {
	for _, v := range grid {
		for _, row := range v {
			fmt.Print(".", row)
		}
		fmt.Println()
	}
}
