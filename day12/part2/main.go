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
	for _, line := range input {
		var (
			a, b string
		)
		split := strings.Split(line, "-")
		a = split[0]
		b = split[1]
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

	path := make([]string, 0)
	seen := make(map[string]struct{})
	dfs(caves["start"], path, seen)
	fmt.Println(len(paths))
}

var paths = make([][]string, 0)

func dfs(curr *node, path []string, seen map[string]struct{}) {
	if isLower(curr.value) {
		seen[curr.value] = struct{}{}
	}
	path = append(path, curr.value)
	if curr.value == "end" {
		paths = append(paths, path)
	} else {
		for _, next := range curr.connections {
			if _, ok := seen[next.value]; !ok {
				dfs(next, path, seen)
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
