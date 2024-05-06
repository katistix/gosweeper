package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var BOMB_COUNT int
var x, y int

func printBanner() {
	// Empty the console
	fmt.Print("\033[H\033[2J")
	const banner = `
     ______                                             
    / ____/___  ______      _____  ___  ____  ___  _____
   / / __/ __ \/ ___/ | /| / / _ \/ _ \/ __ \/ _ \/ ___/
  / /_/ / /_/ (__  )| |/ |/ /  __/  __/ /_/ /  __/ /    
  \____/\____/____/ |__/|__/\___/\___/ .___/\___/_/     
                                    /_/
`
	fmt.Println(banner)
}

func placeBombs(count int, bombArray *[][]int) {
	// While loop to place bombs
	for count > 0 {
		// Generate random x and y coordinates
		x := rand.Intn(len(*bombArray))
		y := rand.Intn(len((*bombArray)[0]))

		// Check if the cell is already a bomb
		if (*bombArray)[x][y] == 0 {
			(*bombArray)[x][y] = -1
			count--
		}
	}

	// After placing bombs, calculate the number of bombs around each cell
	for i := 0; i < len(*bombArray); i++ {
		for j := 0; j < len((*bombArray)[0]); j++ {

			// If the current cell is not a bomb, check all 8 directions
			if (*bombArray)[i][j] == -1 {
				// Check all 8 directions
				for k := -1; k <= 1; k++ {
					for l := -1; l <= 1; l++ {
						if i+k >= 0 && i+k < len(*bombArray) && j+l >= 0 && j+l < len((*bombArray)[0]) {
							if (*bombArray)[i+k][j+l] != -1 {
								(*bombArray)[i+k][j+l]++
							}
						}
					}
				}
			}
		}
	}
}

func askSpot() (int, int) {
	var row, col int
	var input string
	// The first char of the input is the row, the second char is the column, the row is a letter and the column is a number
	fmt.Print("Enter spot: ")
	fmt.Scanf("%s", &input)
	input = strings.ToUpper(input) // convert to uppercase a1 -> A1

	// Convert the input to row and column
	row = int(input[0]) - 65
	col = int(input[1]) - 49

	// Check if the input is valid
	if row < 0 || row >= x || col < 0 || col >= y {
		fmt.Println("Invalid input")
		return askSpot() // ask again
	}

	return row, col
}

// Returns true if the spot is valid(or already tried), false if we hit a bomb
func trySpot(x, y int, inputArray *[][]int, bombArray [][]int) bool {

	if (*inputArray)[x][y] == 0 {
		(*inputArray)[x][y] = 1 // mark as tried
	} else {
		fmt.Println("Already tried this spot")
		return true
	}

	if bombArray[x][y] == -1 {
		fmt.Println("You hit a bomb!") // game over
		return false
	}

	return true // if we reach here, the spot is valid
}

func printBoard(inputArray *[][]int, bombArray [][]int) {
	// Print column numbers
	fmt.Print(" ")
	for i := 0; i < len((*inputArray)[0]); i++ {
		fmt.Print(i + 1)
	}
	fmt.Println()
	for i := 0; i < len(*inputArray); i++ {
		// Print row letters
		fmt.Print(string(rune(i + 65)))
		for j := 0; j < len((*inputArray)[0]); j++ {
			if (*inputArray)[i][j] == 0 {
				fmt.Print("#")
			} else if bombArray[i][j] == -1 {
				fmt.Print("B")
			} else {
				fmt.Print(bombArray[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	// Pretty welcome message
	printBanner()
	fmt.Println("A simple minesweeper game in Go by Paul Tal (@katistix)")

	fmt.Print("Enter the number of rows (1-9): ")
	fmt.Scanf("%d", &x)
	if x > 9 || x < 1 {
		fmt.Println("Rows should be between 1 and 9")
		return
	}
	fmt.Print("Enter the number of columns (1-9): ")
	fmt.Scanf("%d", &y)
	if y > 9 || y < 1 {
		fmt.Println("Columns should be between 1 and 9")
		return
	}
	fmt.Print("Enter the number of bombs (no more than rows*cols): ")
	fmt.Scanf("%d", &BOMB_COUNT)

	// Assert that the number of bombs is less than the number of cells
	if BOMB_COUNT >= x*y {
		fmt.Println("Too many bombs")
		return
	}

	// Create the input slice and the bomb slice
	inputArray := make([][]int, x)
	bombArray := make([][]int, x)

	for i := 0; i < x; i++ {
		inputArray[i] = make([]int, y)
		bombArray[i] = make([]int, y)
	}

	// Initialize both slices to 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			inputArray[i][j] = 0
			bombArray[i][j] = 0
		}
	}

	// Place bombs
	placeBombs(BOMB_COUNT, &bombArray)

	var spotsLeft = x*y - BOMB_COUNT

	for spotsLeft > 0 {
		printBanner()
		printBoard(&inputArray, bombArray)
		row, col := askSpot()
		if !trySpot(row, col, &inputArray, bombArray) {
			// If we hit a bomb, print the board and break the loop
			fmt.Print("\033[H\033[2J")
			printBoard(&inputArray, bombArray)
			break
		}
		spotsLeft--
	}
}
