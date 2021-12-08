package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	digitsToSegments = map[int]string{
		2: "1",
		4: "4",
		3: "7",
		7: "8",
	}
	numberToSegment = make(map[string]string)
	segmentToNumber = make(map[string]string)
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)

	sum := 0
	for _, line := range strings.Split(string(content), "\n") {
		input := strings.Split(string(line), " | ")
		signalPatterns := input[0]
		outputValue := input[1]

		var (
			sortedPatterns []string
			sortedOutput   []string
		)

		for _, v := range strings.Split(signalPatterns, " ") {
			s := []rune(v)
			sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
			sortedPatterns = append(sortedPatterns, string(s))
		}

		for _, v := range strings.Split(outputValue, " ") {
			s := []rune(v)
			sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
			sortedOutput = append(sortedOutput, string(s))
		}

		for _, v := range sortedPatterns {
			if num, ok := digitsToSegments[len(v)]; ok {
				numberToSegment[num] = v
				segmentToNumber[v] = num
			}
		}
		findNumberMappings(sortedPatterns)

		num := gatherOutput(sortedOutput)
		sum += num
	}

	fmt.Println("sum: ", sum)
}

func gatherOutput(s []string) int {
	var result string
	for _, v := range s {
		result += segmentToNumber[v]
	}
	i, _ := strconv.Atoi(result)
	return i
}

func findNumberMappings(s []string) {
	// find 6 and 9
	var (
		ninezero                          []string
		six, three, five, two, nine, zero string
		topOne, bottomOne                 string
	)

	// find out which one of 1's "ab" is in 6 and 9
	// determine which of ones lines is top and bottom
	// find 2, 3, 5 and determine which is 2 and which is 5
	// by checking which one it contains.
	// the rest is 3.
	for _, v := range s {
		// 9, 6, 0
		if len(v) == 6 && strings.Contains(v, string(numberToSegment["1"][0])) && strings.Contains(v, string(numberToSegment["1"][1])) {
			ninezero = append(ninezero, v)
		} else if len(v) == 6 && (strings.Contains(v, string(numberToSegment["1"][0])) || strings.Contains(v, string(numberToSegment["1"][1]))) {
			six = v
		}
		// 3
		if len(v) == 5 && strings.Contains(v, string(numberToSegment["1"][0])) && strings.Contains(v, string(numberToSegment["1"][1])) {
			three = v
		}
	}

	if strings.Contains(six, string(numberToSegment["1"][0])) {
		topOne = string(numberToSegment["1"][1])
		bottomOne = string(numberToSegment["1"][0])
	} else {
		topOne = string(numberToSegment["1"][0])
		bottomOne = string(numberToSegment["1"][1])
	}

	// find 2, 5
	for _, v := range s {
		if len(v) == 5 && v != three {
			if strings.Contains(v, topOne) {
				two = v
			} else if strings.Contains(v, bottomOne) {
				five = v
			}
		}
	}

	// 9 contains all segments of 4
	nine = ninezero[0]
	zero = ninezero[1]
	isNine := true
	for _, c := range numberToSegment["4"] {
		if !strings.Contains(nine, string(c)) {
			isNine = false
			break
		}
	}
	if !isNine {
		nine, zero = zero, nine
	}

	segmentToNumber[nine] = "9"
	segmentToNumber[six] = "6"
	segmentToNumber[two] = "2"
	segmentToNumber[three] = "3"
	segmentToNumber[five] = "5"
	segmentToNumber[zero] = "0"
}
