package pkg

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