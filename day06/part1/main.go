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
	max = 80
)

type fish struct {
	timer int
}

func (f *fish) tick() *fish {
	if f.timer == 0 {
		f.timer = 6
		return newFish()
	}
	f.timer--
	return nil
}

func newFish() *fish {
	return &fish{
		timer: 8,
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	input := strings.Split(string(content), ",")
	fishes := make([]*fish, 0)
	for _, i := range input {
		n, _ := strconv.Atoi(i)
		fishes = append(fishes, &fish{
			timer: n,
		})
	}
	days := 0
	for days < max {
		for _, f := range fishes {
			if fish := f.tick(); fish != nil {
				fishes = append(fishes, fish)
			}
		}
		days++
	}

	fmt.Println("number of fish: ", len(fishes))
}
