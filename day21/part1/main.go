package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type player struct {
	pos   int
	score int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	var (
		ps1, ps2 int
	)
	fmt.Sscanf(string(content), `Player 1 starting position: %d
	Player 2 starting position: %d`, &ps1, &ps2)
	fmt.Println(ps1, ps2)
	p1 := player{pos: ps1 - 1} // because 0-9
	p2 := player{pos: ps2 - 1}
	die := 0
	// game ends when one of the players has a score of 1000
	fmt.Println(p1, p2)
	for {
		// player 1
		steps := 0
		// roll the die
		for i := 0; i < 3; i++ {
			die++
			steps += die
		}
		p1.pos = (p1.pos + steps) % 10
		p1.score += p1.pos + 1
		if p1.score >= 1000 {
			fmt.Printf("player 2 score %d, times dies: %d; player 2 score: %d\n", p2.score, p2.score*die, die)
			break
		}

		// player 2
		steps = 0
		// roll the die
		for i := 0; i < 3; i++ {
			die++
			steps += die
		}
		p2.pos = (p2.pos + steps) % 10
		p2.score += p2.pos + 1
		if p2.score >= 1000 {
			fmt.Printf("player 1 score %d, times dies: %d; player 1 score: %d\n", p1.score, p1.score*die, die)
			break
		}
	}

}
