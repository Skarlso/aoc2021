package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Item struct {
	point    point
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

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
	grid = expandGrid(grid)

	start := point{x: 0, y: 0}
	goal := point{x: len(grid[0]) - 1, y: len(grid) - 1}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		point:    start,
		priority: grid[0][0],
	})
	cost := map[point]int{
		start: grid[0][0],
	}
	cameFrom := map[point]point{
		start: start,
	}
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)

		if current.point == goal {
			break
		}
		for _, next := range neighbors(current.point, grid) {
			newCost := cost[current.point] + grid[next.y][next.x]
			if v, ok := cost[next]; !ok || newCost < v {
				cameFrom[next] = current.point
				cost[next] = newCost
				heap.Push(&pq, &Item{
					point:    next,
					priority: newCost,
				})
			}
		}
	}

	var sum int
	current := goal
	for current != start {
		sum += grid[current.y][current.x]
		current = cameFrom[current]
	}
	fmt.Println("sum: ", sum)
}

func expandGrid(grid [][]int) [][]int {
	expandedGrid := make([][]int, len(grid)*5)
	for i := 0; i < len(expandedGrid); i++ {
		expandedGrid[i] = make([]int, len(grid[0])*5)
	}

	// Prime the first tile
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			expandedGrid[y][x] = grid[y][x]
		}
	}
	// fill out first line for all of the original grid sizes.
	for y := 0; y < len(grid); y++ {
		for x := len(grid[y]); x < len(expandedGrid[y]); x++ {
			// x - the length of the initial grid
			newValue := (expandedGrid[y][x-len(grid[y])] + 1) % 10
			if newValue == 0 {
				newValue = 1
			}
			expandedGrid[y][x] = newValue
		}
	}

	// fill out downwards
	for y := len(grid); y < len(expandedGrid); y++ {
		for x := 0; x < len(expandedGrid[y]); x++ {
			// x - the length of the initial grid
			newValue := (expandedGrid[y-len(grid)][x] + 1) % 10
			if newValue == 0 {
				newValue = 1
			}
			expandedGrid[y][x] = newValue
		}
	}
	return expandedGrid
}

func neighbors(p point, grid [][]int) []point {
	var result []point
	for _, d := range directions {
		np := point{x: p.x + d.x, y: p.y + d.y}
		if np.x >= 0 && np.x < len(grid[p.y]) && np.y >= 0 && np.y < len(grid) {
			result = append(result, np)
		}
	}
	return result
}
