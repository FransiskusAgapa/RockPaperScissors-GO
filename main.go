package main

import "fmt"

func main() {
	/* TODO
	* prompt use to enter their name
	* prompt a set of action to take
	* randomize an action for computer that'll act as opponent
	* checking answer and announce the winner
	* ask if user want to play again
	* add scoreboard
	 */
	fmt.Println("\n==== Rock Paper Scissors ===")

	// prompt action option
	fmt.Print("\n> Pick an action:\n  1. Rock\n  2. Paper\n  3. Scissors\n\n> What's your move:")

	// catch user input
	var user_move int 

	// validate ''user_move''
	_, err := fmt.Scan(&user_move)
	if err != nil {
		fmt.Println("not a valid move")
	} else {
		fmt.Println("\n> Your move is ",user_move)
	}

	fmt.Println("\n============================\n")
}

// GO Primitive Data Type
// int, string, bool, float, rune
