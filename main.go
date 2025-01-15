package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

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
	// map of all possible action
	actions := map[int]string{
		1: "Rock",
		2: "Paper",
		3: "Scissors",
	}

	clearScreen()

	// prompt action option
	fmt.Print("\n> Pick an action:\n  1. Rock\n  2. Paper\n  3. Scissors\n\n> What's your move:")

	// catch user input
	var user_move int

	// read error
	var err error

	// validate ''user_move''
	_, err = fmt.Scan(&user_move)

	for err != nil || user_move > len(actions) {
		if err != nil{
			fmt.Print("\n ! Cmon, you gotta pick a number\n\n> Try input your move again: ")
			_, err = fmt.Scan(&user_move)
		}
		
		if err == nil && user_move > len(actions){
			fmt.Print("\n ! Cmon, you gotta pick a number between 1 to 3\n\n> Try input your move again: ")
			_, err = fmt.Scan(&user_move)
		}
		
		if err == nil && user_move <= len(actions) {
			break
		}
	}

	if user_move > len(actions){
		fmt.Println("\n ! only choose between 1 to 3")
	}

	fmt.Println("\n> Your move is ", actions[user_move])

	fmt.Println("\n============================\n")
}

// GO Primitive Data Type
// int, string, bool, float, rune

// 'clearScreen' every time new game is played : )
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	} else {
		cmd = exec.Command("clear") // Linux/macOS
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
