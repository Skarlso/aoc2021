package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Errorf("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := ioutil.ReadFile(file)

	split := strings.Split(string(content), "\n")
	n := make([]int, 0)
	for _, l := range split {
		i, _ := strconv.Atoi(strings.Trim(l, "\n"))
		n = append(n, i)
	}
	count := 0
	prev := 9999999
	for i := 0; i < len(n)-2; i++ {
		sum := n[i] + n[i+1] + n[i+2]
		if sum > prev {
			count++
		}
		prev = sum
	}
	fmt.Println("increases: ", count)
}
