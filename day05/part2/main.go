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

type line struct {
	begin point
	end   point
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")
	seaFloor := make(map[point]int)
	lines := make([]line, 0)
	for _, l := range split {
		var (
			x1, x2 int
			y1, y2 int
		)
		fmt.Sscanf(l, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		lines = append(lines, line{
			begin: point{x: x1, y: y1},
			end:   point{x: x2, y: y2},
		})
	}
	for _, l := range lines {
		addX := 0
		addY := 0
		if l.begin.x > l.end.x {
			addX = -1
		}
		if l.begin.x < l.end.x {
			addX = 1
		}
		if l.begin.y > l.end.y {
			addY = -1
		}
		if l.begin.y < l.end.y {
			addY = 1
		}

		startX := l.begin.x
		startY := l.begin.y
		targetX := l.end.x
		targetY := l.end.y

		for startX != targetX || startY != targetY {
			seaFloor[point{x: startX, y: startY}]++

			startX += addX
			startY += addY
		}
		seaFloor[point{x: startX, y: startY}]++
	}

	overlaps := 0
	for _, v := range seaFloor {
		if v > 1 {
			overlaps++
		}
	}
	fmt.Println("overlaps: ", overlaps)
}
