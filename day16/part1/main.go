package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	parseBits(binary, 0, 0)
	fmt.Println("sum: ", sum)
}

var sum int

// Does the parsing. Gets an index and the depth of the current package.
// Returns the accumulated version number.
func parseBits(binary string, i, depth int) int {
	version := binary[i : i+3]
	output, err := strconv.ParseInt(version, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	sum += int(output)
	typeID := binary[i+3 : i+6]
	t, err := strconv.ParseInt(typeID, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	// parse literal value
	if t == 4 {
		i += 6
		for {
			i += 5
			if binary[i-5] == '0' {
				return i
			}
		}
	} else {
		ltid := binary[i+6]
		// 15 bits
		if ltid == '0' {
			lenBits, err := strconv.ParseInt(binary[i+7:i+7+15], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			startIndex := i + 7 + 15
			i = startIndex
			for {
				nextIndex := parseBits(binary, i, depth)
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
				nextIndex := parseBits(binary, i, depth)
				i = nextIndex
			}
		}
	}

	return i
}
