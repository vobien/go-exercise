package main

import (
	"bank-app/utility"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	const balanceFile = "balance.txt"

	balance, err := utility.ReadFloatFromFile(balanceFile, 0.0)
	if err != nil {
		fmt.Println(err)
		fmt.Println("------------------")
	}

	fmt.Println("Welcome to Go bank")
	fmt.Printf("Contact us via hotline %s or email %s", randomdata.PhoneNumber(), randomdata.Email())
	for {
		printOptions()

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
			utility.WriteFloatToFile(balance, balanceFile)
			fmt.Println("Your new balance is", balance)
		case 3:
			var withdraw float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdraw)

			if withdraw > balance {
				fmt.Println("Insufficient balance, your balance is", balance)
			} else {
				balance -= withdraw
				utility.WriteFloatToFile(balance, balanceFile)
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
