package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func main() {
	var wg sync.WaitGroup
	var balance sync.Mutex

	// variable for bank balance
	var bankBalance int

	// print out starting values
	fmt.Printf("Initial account balance $%d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Investements", Amount: 1000},
		{Source: "Part time Job", Amount: 50},
	}

	wg.Add(len(incomes))
	// loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week < 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On the week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()

	// print out final balance
	fmt.Printf("Final bank balance $%d.00", bankBalance)
	fmt.Println()
}
