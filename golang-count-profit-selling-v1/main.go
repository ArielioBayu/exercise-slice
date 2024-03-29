package main

import "fmt"

func CountProfit(data [][][2]int) []int {

	resultMap := make(map[int]int)
	for _, cabang := range data {
		fmt.Println("cabang", cabang)
		for bulanKe, month := range cabang {
			fmt.Println("Bulan", month)

			income := month[0]
			expense := month[1]

			fmt.Println("income", income)
			fmt.Println("expense", expense)

			profit := income - expense
			resultMap[bulanKe+1] += profit
		}
		fmt.Println()
	}
	numOfMonth := 0
	for k := range resultMap {
		if k > numOfMonth {
			numOfMonth = k
		}
	}
	result := make([]int, numOfMonth)

	for k, v := range resultMap {
		result[k-1] = v
	}

	return result // TODO: replace this
}

func main() {
	fmt.Println(CountProfit([][][2]int{
		{{1000, 500}, {500, 200}},
		{{1200, 200}, {1000, 800}},
		{{500, 100}, {700, 100}},
	}))
}
