package main

import "fmt"

func factorial(x uint64) uint64 {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(50))
}
