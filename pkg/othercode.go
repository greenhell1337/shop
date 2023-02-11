package pkg

import (
	"io/ioutil"
	"os"
)

func SumCost(card map[string]int) int {
	var sum int
	for _, v := range card {
		if v != 0 {
			sum += v
		}
	}
	return sum
}

func Discount(amount int) int {
	if amount <= 51 {
		amount = amount - (amount / 10)
	}
	return amount
}

func CreateFile(name string) (bool, error) {
	_, err := os.Create("./pkg/" + name)
	if err != nil {
		return false, err
	}
	return true, nil
}

func OpenFile(name string, Order any) (bool, error){
	file, err := os.Open("./pkg/" + name)
	if err != nil {
		return false, err
	}
	for {
		os.E
	}
}