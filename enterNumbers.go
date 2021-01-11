package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var numbers []float64
	fmt.Print("Enter a number: ")
	for reader.Scan() {
		fmt.Print("Enter a number: ")
		str := reader.Text()
		if str == "END" {
			fmt.Println("End of input")
			break
		}
		operand, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Operand", str, "is not a number")
		} else {
			numbers = append(numbers, operand)
		}

	}
	fmt.Println("Entered", len(numbers), "numbers")
	fmt.Println(numbers)
}
