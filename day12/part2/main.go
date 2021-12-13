package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

type node struct {
	connections []*node
	value       string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	input := strings.Split(string(content), "\n")
	caves := make(map[string]*node)
	smallCaves := make(map[string]struct{})
	for _, line := range input {
		var (
			a, b string
		)
		split := strings.Split(line, "-")
		a = split[0]
		b = split[1]
		if isLower(a) {
			smallCaves[a] = struct{}{}
		}
		if isLower(b) {
			smallCaves[b] = struct{}{}
		}
		if _, ok := caves[a]; ok {
			if _, ok := caves[b]; ok {
				caves[a].connections = append(caves[a].connections, caves[b])
				caves[b].connections = append(caves[b].connections, caves[a])
			} else {
				n := &node{value: b, connections: []*node{caves[a]}}
				caves[a].connections = append(caves[a].connections, n)
				caves[b] = n
			}
		} else {
			an := &node{
				value: a,
			}
			if _, ok := caves[b]; ok {
				an.connections = append(an.connections, caves[b])
				caves[b].connections = append(caves[b].connections, an)
			} else {
				bn := &node{
					value: b,
				}
				an.connections = append(an.connections, bn)
				bn.connections = append(bn.connections, an)
				caves[b] = bn
			}
			caves[a] = an
		}
	}

	// fmt.Println(smallCaves)
	for k := range smallCaves {
		if k != "start" && k != "end" {
			path := make([]string, 0)
			seen := make(map[string]struct{})
			seenTwice := map[string]int{k: 0}
			dfs(caves["start"], path, seenTwice, seen)
		}
	}
	fmt.Println("len: ", len(paths))
	// fmt.Println("len: ", paths)
	for v := range paths {
		fmt.Println(v)
	}
}

// use a map of joined strings
var paths = make(map[string]struct{})

// var paths = make([][]string, 0)

func dfs(curr *node, path []string, seenTwice map[string]int, seen map[string]struct{}) {
	if isLower(curr.value) {
		if v, ok := seenTwice[curr.value]; ok && v > 0 {
			seen[curr.value] = struct{}{}
		} else if !ok {
			seen[curr.value] = struct{}{}
		} else if ok {
			seenTwice[curr.value]++
		}
	}
	path = append(path, curr.value)
	if curr.value == "end" {
		paths[strings.Join(path, ",")] = struct{}{}
		// paths = append(paths, path)
	} else {
		for _, next := range curr.connections {
			if _, ok := seen[next.value]; !ok {
				dfs(next, path, seenTwice, seen)
			}
		}
	}
	delete(seen, curr.value)
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
