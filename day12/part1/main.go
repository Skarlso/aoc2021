package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type cave struct {
	pos       string
	pathSoFar []string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	input := strings.Split(string(content), "\n")
	caves := make(map[string][]string)
	for _, line := range input {
		var (
			a, b string
		)
		split := strings.Split(line, "-")
		a = split[0]
		b = split[1]
		caves[a] = append(caves[a], b)
		caves[b] = append(caves[b], a)
	}

	seen := func(a string, b []string) bool {
		for _, v := range b {
			if v == a {
				return true
			}
		}
		return false
	}
	count := 0
	start := cave{
		pos:       "start",
		pathSoFar: []string{"start"},
	}
	queue := []cave{start}
	var current cave
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if current.pos == "end" {
			count++
			continue
		}
		for _, next := range caves[current.pos] {
			if !seen(next, current.pathSoFar) {
				path := make([]string, 0)
				path = append(path, current.pathSoFar...)
				if strings.ToLower(next) == next {
					path = append(path, next)
				}
				queue = append(queue, cave{
					pos:       next,
					pathSoFar: path,
				})
			}
		}
	}
	fmt.Println(count)
}
