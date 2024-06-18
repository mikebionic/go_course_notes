package main

import "fmt"

func PlayGame(start, takeTurn func(), haveWinner func() bool, winningPlayer func() int) {
	start()
	for !haveWinner() {
		takeTurn()
	}
	fmt.Printf("Player %d wins \n", winningPlayer())
}

func main() {
	turn, maxTurns, currentPlayer := 1, 10, 0
	start := func() {
		fmt.Println("Starting a game of chess")
	}
	takeTurn := func() {
		turn++
		fmt.Printf("turn %d taken by player %d \n", turn, currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
	}
	haveWinner := func() bool {
		return turn == maxTurns
	}
	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGame(start, takeTurn, haveWinner, winningPlayer)
}

// Starting a game of chess
// turn 2 taken by player 0
// turn 3 taken by player 1
// turn 4 taken by player 0
// turn 5 taken by player 1
// turn 6 taken by player 0
// turn 7 taken by player 1
// turn 8 taken by player 0
// turn 9 taken by player 1
// turn 10 taken by player 0
// Player 1 wins
