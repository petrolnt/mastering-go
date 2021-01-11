package main

import (
	"fmt"
)

func main() {
	arr := [8]string{"one", "two", "three", "four", "five", "six", "seven", "eight"}
	someMap := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		strValue := arr[i]
		someMap[strValue] = i + 1
	}

	fmt.Println(someMap)
}
