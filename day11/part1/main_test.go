package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestDoTheThing(t *testing.T) {
	content, _ := ioutil.ReadFile(filepath.Join("..", "test.txt"))
	input := strings.Split(string(content), "\n")
	grid := [10][10]*octopus{}
	for i, line := range input {
		for j, c := range line {
			n := c - '0'
			grid[i][j] = &octopus{
				energy: int(n),
			}
		}
	}
	doTheThing(grid)
}
