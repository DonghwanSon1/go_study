package main

import "fmt"

func main() {
	// 채널
	// 채널 또한 함수의 반환 값으로 사용 가능

	// 예제 1
	c := sum(100)

	fmt.Println("ex 1 : ", <-c)

}

func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
	}()
	return tot
}
