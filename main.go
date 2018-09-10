package main

import (
	"fmt"
	"math"
	"time"
)

const GenRoutineNum = 1000
const MaxNumber = 100000
func main() {
	t := time.Now()
	var channels [GenRoutineNum]<-chan int
	for i := 0; i < GenRoutineNum; i++ {
		channels[i] = createWorker(i)
	}
	count := 0
	for i := 0; i < GenRoutineNum; i++ {
		for v := range channels[i] {
			judgePrimeNumber(v , &count)
		}

	}
	fmt.Println("*******")
	fmt.Println(count)
	elapsed := time.Since(t)
	fmt.Println("all time: ", elapsed)
}


//创建channel
func createWorker(i int) <-chan int {
	c := make(chan int, MaxNumber)
	go func() {
		for j := i * MaxNumber; j < (i+1)*MaxNumber; j++ {
			c <- j
		}
		close(c)
	}()
	return c
}


//判断素数
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
	}
	fmt.Println(number)
	*count++
}
