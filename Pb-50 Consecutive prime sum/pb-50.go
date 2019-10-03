package main

import (
	"fmt"
)

func ConsecutivePrimes(limit int) int {
	sum := 0
	for i := 2; i < limit; i++ {
		if sum+i > limit {
			return sum
		}
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}

func main() {
	fmt.Println(ConsecutivePrimes(41))
	fmt.Println(ConsecutivePrimes(1000))
	fmt.Println(ConsecutivePrimes(1000000))
}
