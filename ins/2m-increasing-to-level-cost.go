package main

import "fmt"

func main() {
	data := []int{
		24923,
		25045,
		25150,
		25324,
		25327,
		25390,
		25465,
		25582,
		25732,
		20302,
		14929,
		15205,
		15480,
		15800,
		16115,
		7815,
		7921,
		7892,
		7738,
		7393,
		6945,
		6307,
		5311,
		5920,
		6796,
		7236,
		7267,
		7055,
		6510,
		5517,
		6528,
		7722,
		9132,
		10777,
		12748,
		15215,
		18127,
		21571,
		25772,
		30912,
		36998,
		44325,
		53275,
		63890,
		77959,
		77221,
		71348,
		57988,
		34834,
		90,
		90,
		90,
		90,
		90}
	total := sum(data)
	fmt.Println("total: ", total)
}

func sum(data []int) int {
	result := 0
	for _, v := range data {
		result += v
	}
	return result
}

// total:  1106184
// https://go.dev/play/p/5Jq4dXgA7kt
//
// This also works: awk '{print $5}' 2m-increasing-to-level-illus.txt | sed 's/\$//g' | sed 's/,//g' | awk '{ sum += $1 } END { print sum }'
// Age, End of policy year, Premium outlay, Net distributions, Total charges, Accumulation value, Cash value, Death benefit
