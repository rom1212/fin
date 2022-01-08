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
		7838,
		8617,
		9442,
		10368,
		11353,
		12634,
		14280,
		16272,
		18659,
		21457,
		24743,
		27554,
		30677,
		34127,
		37925,
		42234,
		47021,
		52325,
		58115,
		64696,
		72691,
		81532,
		91366,
		102816,
		116194,
		131078,
		148065,
		167878,
		190019,
		218982,
		256152,
		298029,
		343016,
		389371,
		438507,
		490172,
		548444,
		626321,
		714100}
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

// total:  6300839
// https://go.dev/play/p/KUSgKs2VkfT
//
// This also works: awk '{print $5}' 2m-increasing.txt | sed 's/\$//g' | sed 's/,//g' | awk '{ sum += $1 } END { print sum }'
// Age, End of policy year, Premium outlay, Net distributions, Total charges, Accumulation value, Cash value, Death benefit
