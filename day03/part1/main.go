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

	nums := strings.Split(string(content), "\n")

	var (
		gamma   string
		epsilon string
	)
	for i := 0; i < len(nums[0]); i++ {
		zeros := 0
		ones := 0

		for j := 0; j < len(nums); j++ {
			if nums[j][i] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(g * e)
}
