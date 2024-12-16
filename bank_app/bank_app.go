package main

import (
	"fmt"
	"os"
	"strconv"
)


func writeBalanceToFile(balance float64, balanceFile string) {
	balanceValue := fmt.Sprintf("%f", balance)
	os.WriteFile(balanceFile, []byte(balanceValue), 0644)
}

func readBalanceFromFile(balanceFile string) float64 {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		fmt.Println("Error reading balance from file:", err)
		return 0
	}

	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		fmt.Println("Error parsing balance from file:", err)
		return 0
	}

	return balance
}

func main() {
	const balanceFile = "balance.txt"

	var balance float64 = readBalanceFromFile(balanceFile)

	fmt.Println("Welcome to Go bank")

	for {
		fmt.Printf("\nWhat do you want to do?\n")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Printf("4. Exit\n\n")

		var userChoose int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&userChoose)

		switch userChoose {
		case 1:
			fmt.Println("Your balance is", balance)
		case 2:
			var deposit float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&deposit)

			balance += deposit
			writeBalanceToFile(balance, balanceFile)
			fmt.Println("Your new balance is", balance)
		case 3:
			var withdraw float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdraw)

			if withdraw > balance {
				fmt.Println("Insufficient balance, your balance is", balance)
			} else {
				balance -= withdraw
				writeBalanceToFile(balance, balanceFile)
				fmt.Println("Your new balance is", balance)
			}
		case 4:
			fmt.Println("Thank you for using Go bank")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
