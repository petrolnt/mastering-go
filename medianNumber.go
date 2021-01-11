package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	argsCount := len(arguments) - 1
	if argsCount == 0 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	summ := 0.0
	for i := 1; i <= argsCount; i++ {
		operand, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			log.Fatal(err)
		}
		summ += operand
	}
	median := summ / float64(argsCount)
	fmt.Println("Median:", median)
}
