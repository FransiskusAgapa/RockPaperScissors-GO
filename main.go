package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	/* TODO
	* prompt use to enter their name
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

	//clearScreen()

	// prompt action option
	fmt.Print("\n> Pick an action:\n  1. Rock\n  2. Paper\n  3. Scissors\n\n> What's your move:")

	// catch user input
	var player_move int

	// read error
	var err error

	// validate ''user_move''
	_, err = fmt.Scan(&player_move)

	for err != nil || player_move > len(actions) {
		if err != nil {
			fmt.Print("\n ! C'mon, you gotta pick a number\n\n> Try input your move again: ")
			_, err = fmt.Scan(&player_move)
		}

		if err == nil && player_move > len(actions) {
			fmt.Print("\n ! C'mon, you gotta pick a number between 1 to 3\n\n> Try input your move again: ")
			_, err = fmt.Scan(&player_move)
		}

		if err == nil && player_move <= len(actions) {
			break
		}
	}

	// generate random action for computer to simulate an opponent
	rand.Seed(time.Now().UnixNano())

	// generate random number
	computer_move := rand.Intn(3) + 1

	fmt.Printf("\n> Your move is: %s\n", actions[player_move])
	fmt.Printf("  Computer move is: %s\n", actions[computer_move])

	// find out the winner
	result := findWinner(actions[player_move], actions[computer_move])
	fmt.Printf("\n> Result: %s", result)

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

func findWinner(player_move string, computer_move string) string {
	result := ""
	// fmt.Printf("player: %s - opponent: %s", player_move, computer_move)
	if player_move == computer_move {
		result = "It is a Draw"
	} else if player_move == "Rock" && computer_move == "Scissors" {
		result = "You won!"
	} else if player_move == "Paper" && computer_move == "Rock" {
		result = "You won!"
	} else if player_move == "Scissors" && computer_move == "Paper" {
		result = "You won!"
	} else {
		result = "Computer won! "
	}
	return result
}
