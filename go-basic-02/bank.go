package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountFile = "balance.txt"

func writeToFile(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(accountFile, []byte(balanceText), 0644)
}

func readFromFile() (float64, error) {
	data, err := os.ReadFile(accountFile)
	if err != nil {
		fmt.Println("Error reading balance file:", err)
		return 0.0, errors.New("Error reading balance file or no balance file found, starting with balance of 0.0")
	}
	balance, err := strconv.ParseFloat(string(data), 64)
	fmt.Sscanf(string(data), "Your balance is: %f", &balance)
	if err != nil {
		fmt.Println("Error reading balance file:", err)
		return 0.0, errors.New("Error reading balance file, starting with balance of 0.0")
	}
	return balance, nil
}

func main() {
	// for i := 0; i < 20; i++ {

	// }
	accountBalance, err := readFromFile()
	if err != nil {
		fmt.Println("Error reading balance file")
		fmt.Println(err)
		fmt.Println("________________")
	}
	fmt.Println("Welcome to the go bank!")
mainLoop:
	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scanln(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount!")
				// return
				continue
			}
			accountBalance += depositAmount
			writeToFile(accountBalance)
			fmt.Println("Your new balance is: ", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scanln(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds!")
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Your new balance is: ", accountBalance)
			}
			writeToFile(accountBalance)
		default:
			// return
			break mainLoop
		}
	}
	fmt.Println("Thank you for using the go bank. Goodbye!")
}
