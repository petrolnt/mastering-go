package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	arguments := os.Args
	result := 0.0
	for i := 1; i < len(arguments); i++ {
		operand, _ := strconv.ParseFloat(arguments[i], 64)
		result += operand
	}
	fmt.Println("Summ:", result)
}
