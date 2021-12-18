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
	// THANK YOU KHAN ACADEMY!!!
	// https://www.khanacademy.org/science/physics/one-dimensional-motion/old-projectile-motion/v/projectile-motion-part-1
	// and: https://openstax.org/books/university-physics-volume-1/pages/4-3-projectile-motion
	// Especially, this sentence: Note also that the maximum height depends only on the vertical component of the initial velocity,
	yv := -y1 - 1
	y := 0
	for yv > 0 {
		y += yv
		yv--
	}
	fmt.Println("max hight:", y)
}
