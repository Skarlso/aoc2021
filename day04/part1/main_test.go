package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasWinner(t *testing.T) {
	// row winner
	boards := [][][]dot{
		{
			{
				dot{}, dot{}, dot{}, dot{},
			},
			{
				dot{}, dot{}, dot{}, dot{},
			},
			{
				dot{marked: true}, dot{marked: true}, dot{marked: true}, dot{marked: true},
			},
			{
				dot{}, dot{}, dot{}, dot{},
			},
		},
	}
	_, ok := hasWinner(boards)
	assert.True(t, ok)
	// column winner
	boards = [][][]dot{
		{
			{
				dot{marked: true}, dot{}, dot{}, dot{},
			},
			{
				dot{marked: true}, dot{}, dot{}, dot{},
			},
			{
				dot{marked: true}, dot{}, dot{}, dot{},
			},
			{
				dot{marked: true}, dot{}, dot{}, dot{},
			},
		},
	}
	_, ok = hasWinner(boards)
	assert.True(t, ok)
	boards = [][][]dot{
		{
			{
				dot{}, dot{}, dot{}, dot{marked: true},
			},
			{
				dot{}, dot{}, dot{}, dot{marked: true},
			},
			{
				dot{}, dot{}, dot{}, dot{marked: true},
			},
			{
				dot{}, dot{}, dot{}, dot{marked: true},
			},
		},
	}
	_, ok = hasWinner(boards)
	assert.True(t, ok)
}
