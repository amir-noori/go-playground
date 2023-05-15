package race

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func calculateIncome() {

	totalIncome := 0
	wg := sync.WaitGroup{}
	incomeMutex := sync.Mutex{}

	weeklyIncomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "part time job", Amount: 500},
		{Source: "investment", Amount: 500},
	}

	wg.Add(len(weeklyIncomes))

	for i, income := range weeklyIncomes {
		go func(i int, income Income) {
			defer wg.Done()
			incomeMutex.Lock()
			for week := 0; week < 52; week++ {
				totalIncome += income.Amount
			}
			incomeMutex.Unlock()
			fmt.Printf("income for %s in week $%d\n", income.Source, i)
		}(i, income)
	}

}
