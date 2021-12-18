package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	var (
		x1, x2, y1, y2 int
	)
	fmt.Sscanf(string(content), "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	fmt.Println(x1, x2, y1, y2)
}
