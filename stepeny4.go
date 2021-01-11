package main

import (
	"fmt"
)

type stepen int

func main() {
	const (
		zero stepen = 1 << (iota * 2)
		one
		two
		three
		four
		five
		six
	)
	fmt.Println("4^0 =", zero)
	fmt.Println("4^1 =", one)
	fmt.Println("4^2 =", two)
	fmt.Println("4^3 =", three)
	fmt.Println("4^4 =", four)
	fmt.Println("4^5 =", five)
}
