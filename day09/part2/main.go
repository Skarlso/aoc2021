package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
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
	// display(grid)
	basins := make([]point, 0)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 9 {
				continue
			}
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
				basins = append(basins, point{x: x, y: y})
			}
		}
	}
	sizes := make([]int, 0)
	mark := make(map[point]struct{})
	for _, p := range basins {
		points := calculateBasinSize(p, grid)
		for p := range points {
			mark[p] = struct{}{}
		}
		sizes = append(sizes, len(points))
	}
	sort.Ints(sizes)
	l := len(sizes)
	display(grid, mark)
	fmt.Println("sum: ", sizes[l-1]*sizes[l-2]*sizes[l-3])
}

func neighbors(p point, grid [][]int) []point {
	var result []point

	for _, d := range directions {
		np := point{x: p.x + d.x, y: p.y + d.y}
		if np.x >= 0 && np.x < len(grid[p.y]) && np.y >= 0 && np.y < len(grid) {
			if grid[np.y][np.x] > grid[p.y][p.x] && grid[np.y][np.x] != 9 {
				result = append(result, np)
			}
		}
	}

	return result
}

func calculateBasinSize(p point, grid [][]int) map[point]struct{} {
	start := p
	seen := map[point]struct{}{
		start: {},
	}
	queue := []point{start}
	var current point
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		for _, next := range neighbors(current, grid) {
			if _, ok := seen[next]; !ok {
				queue = append(queue, next)
				seen[next] = struct{}{}
			}
		}
	}
	// fmt.Println(seen)
	// fmt.Println("values: ")
	// for k := range seen {
	// fmt.Printf("%+v, ", grid[k.y][k.x])
	// }
	// fmt.Println()
	return seen
}

func display(grid [][]int, mark map[point]struct{}) {
	red := color.New(color.FgRed).PrintfFunc()
	for y, v := range grid {
		for x, row := range v {
			if _, ok := mark[point{x: x, y: y}]; ok {
				red("%d", row)
			} else {
				fmt.Print(row)
			}
		}
		fmt.Println()
	}
}
