package main

import "fmt"

func main() {
	// 채널
	// 예제 1
	c := receiveOnly1(100) // 채널 반환
	output := total(c)

	fmt.Println("ex 1 : ", <-output)

}

func receiveOnly1(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 777
		close(tot)
	}()
	return tot
}

func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a = a + i
		}
		tot <- a
	}()
	return tot
}
