package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	openingForClosed = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	openingDelimiters = map[rune]struct{}{
		'(': {},
		'[': {},
		'{': {},
		'<': {},
	}
	closingDelimiters = map[rune]struct{}{
		')': {},
		']': {},
		'}': {},
		'>': {},
	}
	points = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	input := strings.Split(string(content), "\n")

	score := 0
	for _, line := range input {
		stack := make([]rune, 0)
		var last rune
		for _, r := range line {
			// if it's an opening thing, push it in stack
			if _, ok := openingDelimiters[r]; ok {
				stack = append(stack, r)
			}
			// if it's a closing one, pop one if it's the opening of the previous one
			// we are good and we popped it.
			if _, ok := closingDelimiters[r]; ok {
				last, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if last != openingForClosed[r] {
					score += points[r]
					break
				}
			}
		}
	}
	fmt.Println("Number of corrupter lines: ", score)
}
