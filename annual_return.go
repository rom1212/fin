package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	sp_dec := [][]float64{
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
	fmt.Println("sp_dec:", sp_dec)
	fmt.Println("sp_inc:", sp_inc)
}

func normal(sp [][]float64, start float64) {
	value := start
	for _, rate := range sp {
		fmt.Println("rate:", rate)
		value = (1 + rate[1]/100.0) * value
	}
	fmt.Println("normal value:", value)
}
