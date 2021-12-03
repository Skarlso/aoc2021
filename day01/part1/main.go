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
		fmt.Println("Usage: go run part1/main.go [file]")
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
	for i := 0; i < len(n)-1; i++ {
		if n[i+1] > n[i] {
			count++
		}
	}
	fmt.Println("increases: ", count)
}
