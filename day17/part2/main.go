package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := os.ReadFile(name)
	var (
		x1, x2, y1, y2 int
	)
	fmt.Sscanf(string(content), "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	fmt.Println(x1, x2, y1, y2)

	count := 0
	for vx := 1; vx < 1000; vx++ {
		for vy := -1000; vy < 1000; vy++ { // it's surely within a 1000...
			var (
				x, y int
			)

			currentVX := vx
			currentVY := vy

			for {
				x += currentVX
				y += currentVY
				if currentVX > 0 {
					currentVX--
				}
				currentVY--

				if x1 <= x && x <= x2 && y1 <= y && y <= y2 {
					// hit
					count++
					break
				}

				if x2 < x || y < y1 {
					// miss
					break
				}
			}
		}
	}

	fmt.Println("found number of trajectories: ", count)
}
