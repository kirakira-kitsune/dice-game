package main

import (
	"fmt"
	"math/rand"
	"time"
)

func chohan() {

	fmt.Println("Welcome to a round of 丁半, Aniki. Please place your bet. Your bet can be odd, even, Sakura or Dragon ")

	var myBet string

	fmt.Scan(&myBet)

	myMap := make(map[string]bool)

	myMap["even"] = true
	myMap["odd"] = true
	myMap["Sakura"] = true
	myMap["Dragon"] = true

	_, ok := myMap[myBet]

	if !ok {
		fmt.Println("Error! Invalid input!")
		return
	}

	fmt.Println("Your bet is: ", myBet)

	seed := time.Now().UnixNano()
	rand.Seed(seed)

	myDice1 := rand.Intn(5) + 1
	myDice2 := rand.Intn(5) + 1

	myDiceResult := myDice1 + myDice2

	fmt.Println("Dice 1 is: ", myDice1)
	fmt.Println("Dice 2 is: ", myDice2)

	fmt.Println("The sum of the dice is: ", myDiceResult)

	if myDice1 > 4 || myDice2 > 4 {
		fmt.Println("You have a high dice!")
	}

	var mySakuraDiceResult bool
	if (myDice1 == 5 && myDice2 == 3) || (myDice1 == 3 && myDice2 == 5) {
		mySakuraDiceResult = true
		fmt.Println("Sakura Panic!!!!")
	}

	var myDragonDiceResult bool
	if myDice1 == myDice2 {
		myDragonDiceResult = true
		fmt.Println("Amazing Dragon Equality!")
	}

	if myDiceResult%2 == 0 {
		fmt.Println("The result is even")
	} else {
		fmt.Println("The result is odd")
	}

	if (myDiceResult%2 == 0 && myBet == "even") ||
		(myDiceResult%2 == 1 && myBet == "odd") ||
		(mySakuraDiceResult && myBet == "Sakura") ||
		(myDragonDiceResult && myBet == "Dragon") {
		fmt.Println("You win!")

	} else {
		fmt.Println("You lose :-(")
	}
}

func main() {

	for {
		chohan()

		fmt.Println("Do you want to continue? [y/n]")

		var myAnswer string

		fmt.Scan(&myAnswer)

		if myAnswer != "y" {
			return
		}

	}

}
