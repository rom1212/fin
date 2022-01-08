package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	sp_inc := sp_rate()

	yearly := 140000.0
	yearsOfContrib := 10
	totalYears := 15

	// Value with fixed rate
	fixedRate := 0.0615
	value_with_fixed_rate := fixed_rate_grow(totalYears, yearly, yearsOfContrib, fixedRate)

	// Value with S&P rate
	value_with_sp_rate := sp_rate_grow(sp_inc, totalYears, yearly, yearsOfContrib)

	p := message.NewPrinter(language.English)
	p.Printf("value_with_fixed_rate: %.2f\n", value_with_fixed_rate)
	p.Printf("value_with_sp_rate   : %.2f\n", value_with_sp_rate)
}

// For illustration
func fixed_rate_grow(totalYears int, yearly float64, yearsOfContrib int, rate float64) float64 {
	value := 0.0
	for i := 0; i < totalYears; i++ {
		if i < yearsOfContrib {
			value += yearly
		}
		value *= (1 + rate)
		fmt.Printf("fixed_rate_grow year: %02d, value: %.2f\n", i+1, value)
	}
	return value
}

func sp_rate_grow(sp [][]float64, totalYears int, yearly float64, yearsOfContrib int) float64 {
	value := 0.0
	for i, r := range sp {
		if i >= totalYears {
			break
		}
		if i < yearsOfContrib {
			value += yearly
		}
		rate := r[1] / 100.00
		value *= (1 + rate)
		fmt.Printf("sp_rate_grow year: %02d, value: %.2f\n", i+1, value)
	}
	return value
}

func sp_rate() [][]float64 {
	sp_dec := [][]float64{
		{2021, 28.71},
		{2020, 18.40},
		{2019, 31.49},
		{2018, -4.38},
		{2017, 21.83},
		{2016, 11.96},
		{2015, 1.38},
		{2014, 13.69},
		{2013, 32.39},
		{2012, 16.00},
		{2011, 2.11},
		{2010, 15.06},
		{2009, 26.46},
		{2008, -37.00},
		{2007, 5.49},
		{2006, 15.79},
		{2005, 4.91},
		{2004, 10.88},
		{2003, 28.68},
		{2002, -22.10},
		{2001, -11.89},
		{2000, -9.10},
		{1999, 21.04},
		{1998, 28.58},
		{1997, 33.36},
		{1996, 22.96},
		{1995, 37.58},
		{1994, 1.32},
		{1993, 10.08},
		{1992, 7.62},
		{1991, 30.47},
		{1990, -3.10},
		{1989, 31.69},
		{1988, 16.61}}
	sp_inc := [][]float64{}
	for i, _ := range sp_dec {
		sp_inc = append(sp_inc, sp_dec[len(sp_dec)-i-1])
	}
	return sp_inc
}

// https://go.dev/play/p/Gm_ck3HTS96
// value_with_fixed_rate: 2,658,583.48
// value_with_sp_rate   : 4,005,126.04
