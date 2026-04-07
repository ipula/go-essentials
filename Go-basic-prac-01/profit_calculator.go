package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const profitsFile = "profits.txt"

func writeToFile(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(profitsFile, []byte(balanceText), 0644)
}

func readFromFile() (float64, error) {
	data, err := os.ReadFile(profitsFile)
	if err != nil {
		fmt.Println("Error reading profits file:", err)
		return 0.0, errors.New("Error reading profits file or no profits file found, starting with profits of 0.0")
	}
	balance, err := strconv.ParseFloat(string(data), 64)
	fmt.Sscanf(string(data), "Your profits are: %f", &balance)
	if err != nil {
		fmt.Println("Error reading profits file:", err)
		return 0.0, errors.New("Error reading profits file, starting with profits of 0.0")
	}
	return balance, nil
}

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue, err := getUserInput("Enter total revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err = getUserInput("Enter total expenses: ")
	if err != nil {
		fmt.Println(err)
	}
	taxRate, err = getUserInput("Enter tax rate (as a decimal): ")
	if err != nil {
		fmt.Println(err)
	}

	currentProfits, err := readFromFile()
	if err != nil {
		fmt.Println("Error reading profits file")
		fmt.Println(err)
		fmt.Println("________________")
	}
	revenue = revenue + currentProfits

	ebt, netProfit, ratio := calculateProfit(revenue, expenses, taxRate)
	writeToFile(netProfit)

	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Net Profit: %.2f\n", netProfit)
	fmt.Printf("EBT to Net Profit Ratio: %.2f\n", ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scanln(&userInput)
	if userInput == 0 {
		return 0, errors.New("Please enter a valid number.")
	} else if userInput < 0 {
		return 0, errors.New("Please enter a positive number.")
	}
	return userInput, nil
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	netProfit := ebt * (1 - taxRate/100)
	ratio := ebt / netProfit
	return ebt, netProfit, ratio
}
