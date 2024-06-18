package main

import "fmt"

// skeleton algorithm

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

func PlayGame(g Game) {
	g.Start()
	// abstract implementation to use as High level algorithm
	for !g.HaveWinner() {
		g.TakeTurn()
	}
	fmt.Printf("Player %d wins \n", g.WinningPlayer())
}

type chess struct {
	turn, maxTurns, currentPlayer int
}

func (c *chess) Start() {
	fmt.Println("Starting a new game of chess.")
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d taken by player %d \n",
		c.turn, c.currentPlayer)
	c.currentPlayer = 1 - c.currentPlayer
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}

func NewGameOfChess() Game {
	return &chess{1, 10, 0}
}

func main() {
	chess := NewGameOfChess()
	PlayGame(chess)
}

// Starting a new game of chess.
// Turn 2 taken by player 0
// Turn 3 taken by player 1
// Turn 4 taken by player 0
// Turn 5 taken by player 1
// Turn 6 taken by player 0
// Turn 7 taken by player 1
// Turn 8 taken by player 0
// Turn 9 taken by player 1
// Turn 10 taken by player 0
// Player 1 wins
