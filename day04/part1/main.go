package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type dot struct {
	value  int
	marked bool
}

var (
	numbers = "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1"

// numbers = "92,12,94,64,14,4,99,71,47,59,37,73,29,7,16,32,40,53,30,76,74,39,70,88,55,45,17,0,24,65,35,20,63,68,89,84,33,66,18,50,38,10,83,75,67,42,3,56,82,34,90,46,87,52,49,2,21,62,93,86,25,78,19,57,77,26,81,15,23,31,54,48,98,11,91,85,60,72,8,69,6,22,97,96,80,95,58,36,44,1,51,43,9,61,41,79,5,27,28,13"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing file name argument.")
	}
	name := os.Args[1]
	content, _ := ioutil.ReadFile(name)
	split := strings.Split(string(content), "\n")

	boards := make([][][]dot, 0)
	board := make([][]dot, 0)
	for _, l := range split {
		if l == "" {
			boards = append(boards, board)
			board = make([][]dot, 0)
			continue
		}

		numbers := strings.Fields(l)
		n := make([]dot, 0)
		for _, number := range numbers {
			i, _ := strconv.Atoi(number)
			n = append(n, dot{value: i, marked: false})
		}
		board = append(board, n)
	}

	nums := make([]int, 0)
	for _, n := range strings.Split(numbers, ",") {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}
	for _, n := range nums {
		markBoards(n, boards)
		if ok, winner := hasWinner(boards); ok {
			fmt.Println("winner score: ", calculateScore(n, winner))
			break
		}
	}
}

func markBoards(n int, boards [][][]dot) {

}

func hasWinner(boards [][][]dot) (bool, [][]dot) {
	return false, nil
}

func calculateScore(n int, board [][]dot) int {
	return 0
}
