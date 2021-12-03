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

	oxygens := strings.Split(string(content), "\n")
	co2s := strings.Split(string(content), "\n")

	// var (
	// oxygenGeneratorRating string
	// co2ScrubberRating     string
	// )
	oxygen := filter(oxygens, func(zeros, ones int) bool {
		return zeros > ones
	})
	co2 := filter(co2s, func(zeros, ones int) bool {
		return zeros < ones || zeros == ones
	})

	c, _ := strconv.ParseInt(co2, 2, 64)
	o, _ := strconv.ParseInt(oxygen, 2, 64)
	fmt.Println(c * o)
}

func filter(list []string, pred func(zeros, ones int) bool) string {
	bitPosition := 0
	for len(list) != 1 {
		zeros := 0
		ones := 0

		for _, o := range list {
			if o[bitPosition] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		var bit byte
		if pred(zeros, ones) {
			bit = '1'
		} else {
			bit = '0'
		}
		for i := 0; i < len(list); i++ {
			if list[i][bitPosition] == bit {
				list = append(list[:i], list[i+1:]...)
				i--
			}
		}
		bitPosition++
	}
	return list[0]
}
