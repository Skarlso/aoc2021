package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	max = 256
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	input := strings.Split(string(content), ",")
	fishes := make([]int, 0)
	for _, i := range input {
		n, _ := strconv.Atoi(i)
		fishes = append(fishes, n)
	}

	days := make([]int, 9)
	for _, f := range fishes {
		days[f]++
	}
	for i := 0; i < max; i++ {
		first := days[0]
		for i := 0; i < 8; i++ {
			days[i] = days[i+1]
		}
		days[6] += first
		days[8] = first
	}

	sum := 0
	for _, f := range days {
		sum += f
	}

	fmt.Println("number of fish: ", sum)
}
