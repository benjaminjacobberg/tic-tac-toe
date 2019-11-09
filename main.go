package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	state := initializeBoard()

	player := "X"
	for true {
		if isGameOver(state) {
			println("NEW GAME? [y/n]")
			input := playAgain()
			if input == "y" {
				state = initializeBoard()
				player = "X"
			} else {
				break
			}
		}
		renderBoard(state)
		print("\nPLAYER " + player + ": ")
		updateBoard(state, player, input())
		switchPlayers(&player)
	}
}

func playAgain() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input != "y" && input != "n" {
			println("'" + input + "' is not correct you dumb dumb. Enter 'y' for YES or 'n' for NO.")
			continue
		}
		return input
	}

	return ""
}

func initializeBoard() [][]string {
	return [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
}

func switchPlayers(player *string) {
	if *player == "X" {
		*player = "O"
	} else {
		*player = "X"
	}
}

func input() int {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		position, e := strconv.Atoi(input)
		if e != nil {
			print("'" + input + "' is not an integer you dumb dumb. Please try again.")
			continue
		}
		return position

	}
	return 0
}

func renderBoard(state [][]string) {
	print("\033[H\033[2J") // clear the terminal
	println("TIC-TAC-TOE\n")
	for y := 0; y < len(state); y++ {
		for x := 0; x < len(state[y]); x++ {
			print("[" + state[y][x] + "]")
		}
		println()
	}
}

func updateBoard(state [][]string, player string, position int) {
	for y := 0; y < len(state); y++ {
		for x := 0; x < len(state[y]); x++ {
			if state[y][x] == strconv.Itoa(position) {
				state[y][x] = player
			}
		}
	}
}

func isGameOver(state [][]string) bool {
	if isWinningGame(state, "X") {
		return true
	}
	if isWinningGame(state, "O") {
		return true
	}
	if isTieGame(state) {
		return true
	}

	return false
}

func isTieGame(state [][]string) bool {
	for y := 0; y < len(state); y++ {
		for x := 0; x < len(state[y]); x++ {
			if _, err := strconv.Atoi(state[y][x]); err == nil {
				return false
			}
		}
	}

	print("\033[H\033[2J") // clear the termina
	println("IT'S A TIE")
	return true
}

func isWinningGame(state [][]string, player string) bool {
	if state[0][0] == player && state[1][0] == player && state[2][0] == player ||
		state[0][1] == player && state[1][1] == player && state[2][1] == player ||
		state[0][2] == player && state[1][2] == player && state[2][2] == player ||
		state[0][0] == player && state[0][1] == player && state[0][2] == player ||
		state[1][0] == player && state[1][1] == player && state[1][2] == player ||
		state[2][0] == player && state[2][1] == player && state[2][2] == player ||
		state[0][0] == player && state[1][1] == player && state[2][2] == player ||
		state[2][0] == player && state[1][1] == player && state[0][2] == player {
		print("\033[H\033[2J") // clear the terminal
		println("PLAYER " + player + " WINS!")

		return true
	}

	return false
}
