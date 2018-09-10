package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	t := time.Now()
	count := 0
	for i := 0; i < 1000000; i++ {
		judgePrimeNumber(i, &count)
	}
	fmt.Println("*****************")
	fmt.Println(count)
	elapsed := time.Since(t)
	fmt.Println("all time: ", elapsed)
}

func judgePrimeNumber(number int, count *int) {
	if number == 1 ||
		number%6 != 1 && number%6 != 5 {
		return
	}
	if number == 2 || number == 3 {
		*count++
		fmt.Println(number)
		return
	}
	tmp := int(math.Sqrt(float64(number)))
	for i := 5; i <= tmp; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return
		}
		//排除所有，剩余的是质数
	}
	*count++
	fmt.Println(number)
}
