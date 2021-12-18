package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	binary := ""
	for _, c := range string(content) {
		ui, _ := strconv.ParseUint(string(c), 16, 64)
		b := fmt.Sprintf("%04b", ui)
		binary += b
	}

	v, _ := parseBits(binary, 0, 0)
	fmt.Println("value: ", v)
}

// Does the parsing. Gets an index and the depth of the current package.
// Returns the accumulated version number.
func parseBits(binary string, i, depth int) (int, int) {
	typeID := binary[i+3 : i+6]
	t, err := strconv.ParseInt(typeID, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	// parse literal value
	if t == 4 {
		i += 6
		value := 0
		for {
			n, _ := strconv.ParseInt(binary[i+1:i+5], 2, 64)
			value = value*16 + int(n)
			i += 5
			if binary[i-5] == '0' {
				return value, i
			}
		}
	} else {
		ltid := binary[i+6]
		values := make([]int, 0)
		// 15 bits
		if ltid == '0' {
			lenBits, err := strconv.ParseInt(binary[i+7:i+7+15], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			startIndex := i + 7 + 15
			i = startIndex
			for {
				val, nextIndex := parseBits(binary, i, depth)
				values = append(values, val)
				i = nextIndex
				if nextIndex-startIndex == int(lenBits) {
					break
				}
			}
		} else {
			// 11 bits
			numberOfPackets, err := strconv.ParseInt(binary[i+7:i+7+11], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			i += (7 + 11)
			for n := 0; n < int(numberOfPackets); n++ {
				val, nextIndex := parseBits(binary, i, depth)
				values = append(values, val)
				i = nextIndex
			}
		}
		switch t {
		case 0:
			return sum(values), i
		case 1:
			s := 1
			for _, v := range values {
				s *= v
			}
			return s, i
		case 2:
			min := math.MaxInt
			for _, v := range values {
				if v < min {
					min = v
				}
			}
			return min, i
		case 3:
			max := 0
			for _, v := range values {
				if v > max {
					max = v
				}
			}
			return max, i
		case 5:
			if values[0] > values[1] {
				return 1, i
			} else {
				return 0, i
			}
		case 6:
			if values[0] < values[1] {
				return 1, i
			} else {
				return 0, i
			}
		case 7:
			if values[0] == values[1] {
				return 1, i
			} else {
				return 0, i
			}
		}
	}

	return 0, i
}

func sum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}
