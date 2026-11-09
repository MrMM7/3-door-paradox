package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	var stayWins int
	var switchWins int
	var iters int

	fmt.Print("Enter number of iterations: ")
	_, err := fmt.Scanln(&iters)

	if err != nil {
		fmt.Println(err)
		panic("fmt threw err")
	}

	start := time.Now()
	for i := 0; i < iters; i++ {
		winningDoor := rand.IntN(3)
		playerGuess := rand.IntN(3)

		hostReveal := rand.IntN(3)
		for hostReveal == winningDoor || hostReveal == playerGuess {
			hostReveal = rand.IntN(3)
		}

		switchGuess := 3 - playerGuess - hostReveal

		if playerGuess == winningDoor {
			stayWins++
		}
		if switchGuess == winningDoor {
			switchWins++
		}
	}
	end := time.Now()

	fmt.Printf("Stay Wins: %.2f%%\n", float64(stayWins)/float64(iters)*100)
	fmt.Printf("Switch Wins: %.2f%%\n", float64(switchWins)/float64(iters)*100)
	fmt.Printf("Time taken: %s\n", end.Sub(start))
}
