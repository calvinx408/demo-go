package main

import "fmt"

func sum(values []int, result chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	result <- sum
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7}
	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan
	fmt.Println("Result:", sum1, sum2, sum1+sum2)
}
