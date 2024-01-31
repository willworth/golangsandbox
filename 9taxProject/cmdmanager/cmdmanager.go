package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm each with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		prices = append(prices, price)
		if price == "0" {
			break
		}
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
