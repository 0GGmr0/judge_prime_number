package main

import "math"

const NUMBER_NUM = 1000000

func main() {
	
}

func generate


func judgePrimeNumber(in <-chan int) <-chan int {
	out := make(chan int)
	number := <- in
	go func() {
		if number == 2 || number == 3 {
			out <- number
		}
		if number % 6 == 1 || number % 6 == 5 {
			tmp := int(math.Sqrt(float64(number)))
			for i := 5; i <= tmp; i++ {
				if number % i == 0 || number % (i + 2) == 0 {
					break
				}
			}
			out <- number
		}
	}()
	return out
}