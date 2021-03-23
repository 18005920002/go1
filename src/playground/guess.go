package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Let's guess the number between 1 and 100, you have 10 times")
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, " guesses left.")
		fmt.Print("Make a guess:")
		input, error := reader.ReadString('\n')
		if error != nil {
			log.Fatal(error)
		}
		input = strings.TrimSpace(input)
		guess, error := strconv.Atoi(input)
		if error != nil {
			log.Fatal(error)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job！ You Guess it")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
