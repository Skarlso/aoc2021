package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type point struct {
	x, y, z int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	lines := strings.Split(string(content), "\n")
	grid := make(map[point]string)
	for _, line := range lines {
		var (
			action                 string
			x1, x2, y1, y2, z1, z2 int
		)
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &action, &x1, &x2, &y1, &y2, &z1, &z2)
		// fmt.Printf("action: %s, x1=%d, x2=%d, y1=%d, y2=%d, z1=%d, z2=%d\n", action, x1, x2, y1, y2, z1, z2)

		if !inRegion(x1) || !inRegion(x2) || !inRegion(y1) || !inRegion(y2) || !inRegion(z1) || !inRegion(z2) {
			continue
		}
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				for z := z1; z <= z2; z++ {
					grid[point{x: x, y: y, z: z}] = action
				}
			}
		}
	}

	count := 0
	for _, v := range grid {
		if v == "on" {
			count++
		}
	}
	fmt.Println("count: ", count)
}

func inRegion(a int) bool {
	return a <= 50 && a >= -50
}
