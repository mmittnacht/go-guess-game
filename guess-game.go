package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	const min = 1
	const max = 100
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Generating random number between 1-100 to guess...")

	guessNumber := rand.IntN(max) + min
	fmt.Println("Done")

	for {
		fmt.Printf("Please enter your tip: ")
		text, _ := reader.ReadString('\n')

		playerGuess, err := strconv.Atoi(strings.Trim(text, "\r\n"))

		if err != nil {
			fmt.Println("Error with user input, please enter a number!")
			continue
		}

		if guessNumber == playerGuess {
			fmt.Printf("Congrats, you guess it, it's %d\n", guessNumber)
			break
		} else if guessNumber > playerGuess {
			fmt.Printf("Guessed number: %d is smaller than the random number between 1-100\n", playerGuess)
			continue
		} else {
			fmt.Printf("Guessed number: %d is bigger than the random number between 1-100\n", playerGuess)
			continue
		}
	}
}
