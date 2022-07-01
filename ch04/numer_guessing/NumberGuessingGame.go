package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/* This is number guessing game played against the computer. The computer produces a four-digit number
that is tried to be correctly guessed by the player. Four-digit number can't start with 0 and have repeating digits.
*/
// Ini1234tial Guess object created by the computer.
var initialGuess Guess

// All guesses the player enters
var allGuesses []*Guess
var numberOfGuesses int

// Guess represents an entry of four-digit number and its result according to how it matches with initial guess.
type Guess struct {
	// Number represented in a four-digit array
	entry [4]int
	// In the form of +X-Y where X is number of right positions and Y is number of wrong positions.
	result string
}

// createGuess creates a Guess instance.
func createGuess(entry [4]int) *Guess {
	guess := new(Guess)
	guess.entry = entry
	return guess
}

// evaluate checks entry of its Guess instance in terms of position matches.
func (guess *Guess) evaluate() {
	rightPositionCount := 0
	wrongPositionCount := 0

	for i := 0; i < 4; i++ {
		// First check all right positions
		if guess.entry[i] == initialGuess.entry[i] {
			rightPositionCount++
		} else { // Then if it is not in right position check if in wrong positions
			if arrayContains(initialGuess.entry, guess.entry[i]) {
				wrongPositionCount++
			}
		}
	}

	var result string = fmt.Sprintf("+%d-%d", rightPositionCount, wrongPositionCount)
	guess.result = result
}

// print prints out its Guess instance in a format.
func (guess *Guess) print() {
	fmt.Printf("Guess: %d Result: %s\n", guess.entry, guess.result)
}

// saveGuess saves given Guess instance for further use.
func saveGuess(guess *Guess) {
	allGuesses = append(allGuesses, guess)
	numberOfGuesses++
}

// printGuesses prints all guesses so far
func printGuesses() {
	println("Your guesses so far:")
	//var guess *Guess
	for i, guess := range allGuesses {
		//guessPtr = &guess
		fmt.Printf("%d ", i)
		guess.print()
	}
}

// arrayContains checks whether given number is in given array.
func arrayContains(array [4]int, number int) bool {
	for _, element := range array {
		if element == number {
			return true
		}
	}

	return false
}

// convertStringToIntArray  converts given string with 4 digits into an array of int of size 4.
func convertStringToIntArray(str string) [4]int {
	var intArray [4]int
	for i := 0; i < 4; i++ {
		intArray[i], _ = strconv.Atoi(str[i : i+1])
	}
	return intArray
}

// printArray prints given array.
func printArray(array [4]int) {
	print("Your guess: ")
	for i := range array {
		print(array[i])
	}
}

// createInitialGuess creates a Guess instance for the computer to be guessed by the player.
func createInitialGuess() {
	var numbers [4]int
	for i := 0; i < 4; i++ {
		number := rand.Intn(10)
		if i == 0 && number == 0 {
			i--
			continue
		}
		if !arrayContains(numbers, number) {
			numbers[i] = number
		} else {
			i--
			continue
		}
	}

	initialGuess = Guess{numbers, "+4"}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	createInitialGuess()
}

func main() {
	//initialGuess.print()
	println("\n   *   *   *   Welcome to Number Guessing Game   *   *   *")
	println("Enter a four-digit-number guess to match what the computer has in its mind.")
	println("Your guess shouldn't start with zero and include repeating digits.")
	flag := true
	// Continue until either player wins or quits
	for flag {
		println("Please enter your guess as a four-digit number:")
		var guessEnteredAsString string
		fmt.Scanf("%s", &guessEnteredAsString)
		if len(guessEnteredAsString) != 4 {
			println("Please enter your guess as a four-digit number:")
		} else {
			guessEnteredAsIntArray := convertStringToIntArray(guessEnteredAsString)
			guessEntered := createGuess(guessEnteredAsIntArray)
			saveGuess(guessEntered)
			guessEntered.evaluate()
			if strings.Contains(guessEntered.result, "+4") {
				fmt.Printf("You win at %dth attemp!", numberOfGuesses)
				flag = false
			} else {
				guessEntered.print()
				println("Keep trying!\n")
				printGuesses()
			}
		}
	}
}
