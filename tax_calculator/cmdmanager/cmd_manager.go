package cmdmanager

import "fmt"

type CmdManager struct{}

func New() CmdManager {
	return CmdManager{}
}

func (cmd CmdManager) ReadLines() ([]string, error) {
	var prices []string
	for {
		var price string
		fmt.Print("Enter a price: ")
		_, err := fmt.Scan(&price)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid price.")
			continue
		}

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CmdManager) Write(data interface{}) error {
	fmt.Println(data)
	return nil
}
