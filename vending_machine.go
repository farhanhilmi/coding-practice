package main

import (
	"fmt"
	"sort"
)

type Drink struct {
	Name  string
	Price int
}

func totalMoney(money []int) int {
	total := 0

	for _, m := range money {
		total += m
	}

	return total
}

func vendingMachine(money []int) []string {
	result := []string{}

	drinkList := []Drink{
		{
			Name:  "Teh botol",
			Price: 5000,
		},
		{
			Name:  "Coca cola",
			Price: 7000,
		},
		{
			Name:  "Aqua",
			Price: 2000,
		},
	}

	totalMoney := totalMoney(money)

	sort.Slice(drinkList, func(i, j int) bool {
		return drinkList[i].Price > drinkList[j].Price
	})

	for _, drink := range drinkList {
		for totalMoney >= drink.Price {
			result = append(result, drink.Name)
			totalMoney -= drink.Price
		}
	}

	return result
}

func main() {
	money := []int{5000, 2000, 2000, 3000}
	purchasedDrinks := vendingMachine(money)
	fmt.Println(purchasedDrinks)
}
