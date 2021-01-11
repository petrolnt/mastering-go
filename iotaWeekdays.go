package main

import (
	"fmt"
	"time"
)

func main() {

	const (
		sunday int = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)
	fmt.Println("Day number: ", saturday)
	weekday := time.Weekday(saturday)
	fmt.Println(weekday)
}
