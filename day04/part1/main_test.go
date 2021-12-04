package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasWinner(t *testing.T) {
	// row winner
	boards := [][][]int{
		{
			{
				1, 1, 1, 1,
			},
			{
				1, 1, 1, 1,
			},
			{
				-1, -1, -1, -1,
			},
			{
				1, 1, 1, 1,
			},
		},
	}
	_, ok := hasWinner(boards)
	assert.True(t, ok)
	// column winner
	boards = [][][]int{
		{
			{
				-1, 1, 1, 1,
			},
			{
				-1, 1, 1, 1,
			},
			{
				-1, 1, 1, 1,
			},
			{
				-1, 1, 1, 1,
			},
		},
	}
	_, ok = hasWinner(boards)
	assert.True(t, ok)
	boards = [][][]int{
		{
			{
				1, 1, 1, -1,
			},
			{
				1, 1, 1, -1,
			},
			{
				1, 1, 1, -1,
			},
			{
				1, 1, 1, -1,
			},
		},
	}
	_, ok = hasWinner(boards)
	assert.True(t, ok)
}
