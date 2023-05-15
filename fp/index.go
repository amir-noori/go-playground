package fp

import (
	"errors"
	"fmt"
)

func Run() {

	//fmt.Println("imperative fib: ", imperativeFibonacci(100))
	//fmt.Println("tail recursive fib: ", tailRecFibonacci(10, 0, 1, 1, 0))
	//fmt.Println("recursive fib: ", RecursiveFibonacci(10))
	tryClosure()

}

type Car struct {
	Model string
}

type Cars []Car

func (cars *Cars) find(model string) (*Car, error) {
	for _, car := range *cars {
		if model == car.Model {
			return &car, nil
		}
	}
	return nil, errors.New("no Car Found")
}

func imperativeFind(model string) {
	accord := Car{"Accord"}
	is250 := Car{"IS250"}
	blazer := Car{"Blazer"}
	cars := Cars{accord, is250, blazer}

	for _, car := range cars {
		if car.Model == model {
			fmt.Printf("Found %s", model)
		}
	}
}

func functionalFind(targetCar string) {
	accord := Car{"Accord"}
	is250 := Car{"IS250"}
	blazer := Car{"Blazer"}
	cars := Cars{accord, is250, blazer}
	cars.find(targetCar)
}
