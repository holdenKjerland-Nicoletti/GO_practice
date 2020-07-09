package main

import (
	"fmt"
	"strconv"
	"os"
	"math/rand"
	"time"
)


// constant variable for max possible + 1
const Max_guess int = 100

// Incorrect _guesses should be global so both functions can see
var Incorrect_guesses int = 0

func Guess_feedback(guess int, correct_number int){
	if guess == correct_number{
		fmt.Println("Correct guess! Nice job! It took you ", Incorrect_guesses+1, " guesses.")
		os.Exit(0) // If the guess was correct we can exit program
	}else if guess > correct_number{
		fmt.Println("Too high")
		return
	}else{
		fmt.Println("Too low")
		return
	}
}

func main() {

	// We need a random seed, or else our random correct number will be the same

	rand.Seed(time.Now().UnixNano()) //https://flaviocopes.com/go-random/

	// Using the math/rand package to generate random int in range [0,Max_guess)

	correct_number := rand.Intn(Max_guess)

	// fmt.Println(correct_number)

	var guess string

	for Incorrect_guesses < 5{

		// Get next guess
		fmt.Println("Enter number to guess ( < 100 ): ")
		fmt.Scanln(&guess)

		// Convert guess from string to int
		var int_guess, err = strconv.Atoi(guess)


		// If user did not type an integer
		if err != nil{

			// fmt.Println(err)

			fmt.Println("Please type in an integer")

		}else{

			// Print high, low, or correct, and also incriment number of guesses
			Guess_feedback(int_guess, correct_number)

			// Incriment number of guesses
			Incorrect_guesses++
		}

	}

	// If they have gotten to this point, they ran out of guesses and lost

	fmt.Println("You have run out of guesses, and did not succesfully guess the number. Better luck next time!")
}