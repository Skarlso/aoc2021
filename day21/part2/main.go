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

// keep track of state of the game.
type game struct {
	p1, p2 player
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
	// keep track of each games current score in a memoization map
	win := play(game{p1: p1, p2: p2})
	fmt.Println(win)
}

// Keeping the game state.. Lot's of reading. Looked into:
// https://en.wikipedia.org/wiki/Dynamic_programming
// Further readings: Algorithms by Jeff Erickson ( http://jeffe.cs.illinois.edu/teaching/algorithms/ )
// Competitive Programming handbook https://cses.fi/book/book.pdf
// Grokking Algorithms https://www.amazon.com/Grokking-Algorithms-illustrated-programmers-curious/dp/1617292230
// Still not a 100% sure how this works yet... still reading and taking notes.
var dp = make(map[game][]int64)

func play(g game) []int64 {
	if g.p1.score >= 21 {
		return []int64{1, 0} // lots of wikipedia reading on this one
	}
	if g.p2.score >= 21 {
		return []int64{0, 1} // memoization to keep track and count of who won a game
	}
	if v, ok := dp[g]; ok {
		return v
	}
	win := []int64{0, 0}
	for d1 := 1; d1 < 4; d1++ {
		for d2 := 1; d2 < 4; d2++ {
			for d3 := 1; d3 < 4; d3++ {
				p1 := (g.p1.pos + d1 + d2 + d3) % 10
				s1 := g.p1.score + p1 + 1

				// do the switch
				w := play(game{
					p1: player{pos: g.p2.pos, score: g.p2.score},
					p2: player{pos: p1, score: s1},
				})
				win[0] += w[1]
				win[1] += w[0]
			}
		}
	}
	dp[g] = win
	return win
}
