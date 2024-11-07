package main

import "fmt"

// y = 5*x+3
func mathAbolfazl(zahl int) int {
	return 5*zahl + 3
}

func main() {
	fmt.Println("Hello Abolfazl")
	var sum = 0
	for index := range 11 {
		if index%2 == 0 {
			sum = sum + index
		}
		fmt.Println("index = ", index)
	}
	fmt.Println("sum = ", sum)

	fmt.Println("Abolfazl function = ", mathAbolfazl(10))
}
