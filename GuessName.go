package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secreteNumber := rand.Intn(maxNum)
	//fmt.Println(secreteNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("gg", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("you fked", err)
			continue
		}
		fmt.Println("Your guess is", guess)

		if guess > secreteNumber {
			fmt.Println("Your guess is bigger")
		} else if guess < secreteNumber {
			fmt.Println("Your guess is smaller")
		} else {
			fmt.Println("Correct!")
			break
		}
	}

}
