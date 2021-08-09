package main

import "fmt"

const freezingF = 32.0
const boilingF = 212.0

func main() {
	fmt.Printf("freezeing point = %g˚F or %g˚C\n", freezingF, fToC(freezingF))
	fmt.Printf("boiling point = %g˚F or %g˚C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
