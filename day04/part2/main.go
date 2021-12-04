package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")
	numbers := split[0]

	boards := make([][][]int, 0)
	board := make([][]int, 0)
	for i, l := range split {
		if i < 2 {
			continue
		}
		if l == "" {
			boards = append(boards, board)
			board = make([][]int, 0)
			continue
		}

		numbers := strings.Fields(l)
		n := make([]int, 0)
		for _, number := range numbers {
			i, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			n = append(n, i)
		}
		board = append(board, n)
	}
	boards = append(boards, board)

	nums := make([]int, 0)
	for _, n := range strings.Split(numbers, ",") {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	for len(boards) > 0 {
		for _, n := range nums {
			markBoards(n, boards)
			if loc, winner, ok := hasWinner(boards); ok {
				fmt.Printf("winner score: number: %d, score: %d, loc: %d\n", n, calculateScore(n, winner), loc)
				boards = append(boards[:loc], boards[loc+1:]...)
				break
			}
		}
	}
}

func markBoards(n int, boards [][][]int) {
	for _, board := range boards {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == n {
					board[i][j] = -1
				}
			}
		}
	}
}

func hasWinner(boards [][][]int) (int, [][]int, bool) {
	// fmt.Println("board in hasWinner: ", boards)
	for loc, board := range boards {
		for i := 0; i < len(board); i++ {
			// check vertically
			rowWon := true
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] != -1 {
					rowWon = false
					break
				}
			}
			if rowWon {
				return loc, board, true
			}
		}

		// these are 5x5 and not varring in size.
		// check horizontally
		l := len(board[0])
		for col := 0; col < l; col++ {
			colWon := true
			for i := 0; i < len(board); i++ {
				if board[i][col] != -1 {
					colWon = false
					break
				}
			}
			if colWon {
				return loc, board, true
			}
		}
	}

	return -1, nil, false
}

func calculateScore(n int, board [][]int) int {
	score := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != -1 {
				score += board[i][j]
			}
		}
	}
	return score * n
}
