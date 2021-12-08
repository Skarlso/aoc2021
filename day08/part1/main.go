package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var digitsToSegments = map[int]string{
	2: "1",
	4: "4",
	3: "7",
	7: "8",
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)

	out := 0
	for _, line := range strings.Split(string(content), "\n") {
		input := strings.Split(string(line), " | ")
		// signalPatterns := input[0]
		outputValue := input[1]
		for _, v := range strings.Split(outputValue, " ") {
			if _, ok := digitsToSegments[len(v)]; ok {
				out++
			}
		}
	}

	fmt.Println("out: ", out)
}
