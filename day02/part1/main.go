package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := ioutil.ReadFile(file)

	split := strings.Split(string(content), "\n")
	x, y := 0, 0
	for _, l := range split {
		var (
			op string
			v  int
		)
		fmt.Sscanf(string(l), "%s %d", &op, &v)
		switch op {
		case "forward":
			x += v
		case "up":
			y -= v
		case "down":
			y += v
		}
	}
	fmt.Println("mult: ", x*y)
}
