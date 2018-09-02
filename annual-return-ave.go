/*
total normal value:  12229674.189605, for 48 years
total boxed  value:   8382690.412815, for 48 years
i:  0, normal value: 4743879.082111, boxed  value: 1895368.010901
i:  1, normal value: 4145934.127141, boxed  value: 1835050.078816
i:  2, normal value: 3195680.657356, boxed  value: 1616564.980639
i:  3, normal value: 2092314.029316, boxed  value: 1415548.639568
i:  4, normal value: 3154897.695013, boxed  value: 1616564.980639
i:  5, normal value: 4757446.707780, boxed  value: 1779987.339158
i:  6, normal value: 3637782.318609, boxed  value: 1623812.797835
i:  7, normal value: 3401314.717957, boxed  value: 1623812.797835
i:  8, normal value: 3865596.742052, boxed  value: 1701052.751178
i:  9, normal value: 2285403.479254, boxed  value: 1607507.620530
i: 10, normal value: 2440156.399750, boxed  value: 1607507.620530
i: 11, normal value: 2118976.568719, boxed  value: 1607507.620530
i: 12, normal value: 2275648.900209, boxed  value: 1630015.919884
i: 13, normal value: 2171742.265934, boxed  value: 1630015.919884
i: 14, normal value: 2345928.186904, boxed  value: 1630015.919884
i: 15, normal value: 2509725.939298, boxed  value: 1743827.137777
i: 16, normal value: 1931496.361694, boxed  value: 1537297.349807
i: 17, normal value: 1822283.076222, boxed  value: 1496659.228560
i: 18, normal value: 2109346.766519, boxed  value: 1635304.620280
normal average value: 2895029.159044, over 30 years
boxed  average value: 1643864.280749, over 30 years
*/
package main

import (
	"fmt"
)

func main() {
	sp_inc := [][]float64{
		{1970, 4.01},
		{1971, 14.31},
		{1972, 18.98},
		{1973, -14.66},
		{1974, -26.47},
		{1975, 37.20},
		{1976, 23.84},
		{1977, -7.18},
		{1978, 6.56},
		{1979, 18.44},
		{1980, 32.50},
		{1981, -4.92},
		{1982, 21.55},
		{1983, 22.56},
		{1984, 6.27},
		{1985, 31.73},
		{1986, 18.67},
		{1987, 5.25},
		{1988, 16.61},
		{1989, 31.69},
		{1990, -3.10},
		{1991, 30.47},
		{1992, 7.62},
		{1993, 10.08},
		{1994, 1.32},
		{1995, 37.58},
		{1996, 22.96},
		{1997, 33.36},
		{1998, 28.58},
		{1999, 21.04},
		{2000, -9.10},
		{2001, -11.89},
		{2002, -22.10},
		{2003, 28.68},
		{2004, 10.88},
		{2005, 4.91},
		{2006, 15.79},
		{2007, 5.49},
		{2008, -37.00},
		{2009, 26.46},
		{2010, 15.06},
		{2011, 2.11},
		{2012, 16.00},
		{2013, 32.39},
		{2014, 13.69},
		{2015, 1.38},
		{2016, 11.96},
		{2017, 21.83},
	}

	fmt.Printf("number of years: %v, sp_inc: %v\n", len(sp_inc), sp_inc)

	start := 100000.0
	// normal value
	normal_value := normal(sp_inc, start, false)

	// boxed value
	boxed_value := box(sp_inc, start, false)

	fmt.Printf("total normal value: %16f, for %v years\n", normal_value, len(sp_inc))
	fmt.Printf("total boxed  value: %16f, for %v years\n", boxed_value, len(sp_inc))

	ave_years(sp_inc, start, 30, true)
}

func normal(sp [][]float64, start float64, verbose bool) float64 {
	value := start
	for _, r := range sp {
		if verbose {
			fmt.Println("year rate:", r)
		}
		rate := r[1] / 100.00
		value = (1 + rate) * value
	}
	if verbose {
		fmt.Printf("normal value: %f\n", value)
	}
	return value
}

func box(sp [][]float64, start float64, verbose bool) float64 {
	cap := 0.15
	bot := 0.007
	value := start
	for _, r := range sp {
		rate := r[1] / 100.00
		if rate > cap {
			rate = cap
		} else if rate < 0 {
			rate = bot
		}
		value = (1 + rate) * value
		if verbose {
			fmt.Printf("box value: %f, box rate: %f, year rate: %v\n", value, rate, r)
		}
	}
	if verbose {
		fmt.Printf("box value: %f\n", value)
	}
	return value
}

func ave_years(sp [][]float64, start float64, num_years int, verbose bool) {
	normal_ave_value := 0.0
	boxed_ave_value := 0.0

	count := len(sp) - num_years + 1
	for i := 0; i < count; i++ {
		range_sp := sp[i : i+num_years]

		normal_value := normal(range_sp, start, false)
		boxed_value := box(range_sp, start, false)

		if verbose {
			fmt.Printf("i: %2d, normal value: %f, boxed  value: %f\n", i, normal_value, boxed_value)
		}

		normal_ave_value += normal_value
		boxed_ave_value += boxed_value

	}

	fmt.Printf("normal average value: %f, over %v years\n", normal_ave_value/float64(count), num_years)
	fmt.Printf("boxed  average value: %f, over %v years\n", boxed_ave_value/float64(count), num_years)
}
