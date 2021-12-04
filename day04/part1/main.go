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

	var (
		boards [][][]int
		board  [][]int
	)
	for i, l := range split {
		if i < 2 {
			continue
		}
		if l == "" {
			boards = append(boards, board)
			board = make([][]int, 0)
			continue
		}

		nums := strings.Fields(l)
		n := make([]int, 0)
		for _, number := range nums {
			i, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			n = append(n, i)
		}
		board = append(board, n)
	}
	// Add the last one.
	boards = append(boards, board)

	nums := make([]int, 0)
	for _, n := range strings.Split(numbers, ",") {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}
	for _, n := range nums {
		markBoards(n, boards)
		if winner, ok := hasWinner(boards); ok {
			fmt.Println("winner score: ", calculateScore(n, winner))
			fmt.Println("winner board: ", winner)
			os.Exit(0)
		}
	}
	fmt.Println("ops, no winner")
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

func hasWinner(boards [][][]int) ([][]int, bool) {
	for _, board := range boards {
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
				return board, true
			}
		}

		// these are 5x5 and not varring in size.
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
				return board, true
			}
		}
	}

	return nil, false
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
